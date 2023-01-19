package tss

import (
	"fmt"
	"sync"

	bkg "github.com/binance-chain/tss-lib/ecdsa/keygen"
	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p-core/crypto"
	peer_mgr "github.com/meshplus/bitxhub-core/peer-mgr"
	"github.com/meshplus/bitxhub-core/tss/blame"
	"github.com/meshplus/bitxhub-core/tss/cache"
	"github.com/meshplus/bitxhub-core/tss/conversion"
	"github.com/meshplus/bitxhub-core/tss/message"
	"github.com/meshplus/bitxhub-core/tss/p2p"
	"github.com/meshplus/bitxhub-core/tss/storage"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
)

var _ Tss = (*TssInstance)(nil)

// TssInstance is the structure that can provide all Keygen and Keysign features
type TssInstance struct {
	// =================== tssmgr info
	localPrivK       crypto.PrivKey
	localPubK        crypto.PubKey
	threshold        int
	keygenPreParams  *bkg.LocalPreParams
	conf             TssConfig
	keygenLocalState *storage.KeygenLocalState

	// ==================== task info
	msgID string
	// the number of msgs contained in the current request (1 for Keygen request is and n for Keysign request)
	msgNum int

	localPartyID string
	partyInfo    *conversion.PartyInfo
	partyLock    *sync.Mutex

	// receive task Messages
	TssMsgChan chan *pb.Message
	// receive end signal of the coroutine for inbound message handling (abnormal abort, timeout abort, end of task)
	inMsgHandleStopChan chan struct{}
	// receive the task end signal
	taskDoneChan chan struct{}
	// receive TSS shutdown signal
	stopChan chan struct{}

	// msg cache list
	cachedWireBroadcastMsgLists *sync.Map
	cachedWireUnicastMsgLists   *sync.Map

	// record the confirmation of broadcast messages
	// partyID + round -> msgContent, msgHash
	unConfirmedMessages map[string]*cache.LocalCacheItem
	unConfirmedMsgLock  *sync.Mutex

	// record the participants who have completed the task
	finishedParties map[string]bool

	// malicious participation node
	culprits     []*btss.PartyID
	culpritsLock *sync.RWMutex

	blameMgr *blame.Manager
	p2pComm  *p2p.Communication
	logger   logrus.FieldLogger
}

// NewTss creates a new instance of Tss
func NewTssInstance(
	conf TssConfig,
	privKey crypto.PrivKey,
	preParams *bkg.LocalPreParams,
	keygenLocalState *storage.KeygenLocalState,
	peerMgr peer_mgr.OrderPeerManager,
	logger logrus.FieldLogger,
) (tssInstance *TssInstance, err error) {
	// keygen pre params
	// When using the keygen party it is recommended that you pre-compute the
	// "safe primes" and Paillier secret beforehand because this can take some
	// time.
	// This code will generate those parameters using a concurrency limit equal
	// to the number of available CPU cores.
	if preParams == nil || !preParams.Validate() {
		preParams, err = bkg.GeneratePreParams(conf.PreParamTimeout)
		if err != nil {
			return nil, fmt.Errorf("fail to generate pre parameters: %w", err)
		}
	}
	if !preParams.Validate() {
		return nil, fmt.Errorf("invalid preparams")
	}

	// network
	comm, err := p2p.NewCommunication(peerMgr, logger)
	if err != nil {
		return nil, fmt.Errorf("fail to create communication layer: %w", err)
	}
	comm.Start()

	tss := &TssInstance{
		keygenLocalState: keygenLocalState,
		keygenPreParams:  preParams,
		conf:             conf,
		localPrivK:       privKey,
		localPubK:        privKey.GetPublic(),
		partyLock:        &sync.Mutex{},
		partyInfo: &conversion.PartyInfo{
			PartyMap:   nil,
			PartyIDMap: make(map[string]*btss.PartyID),
		},
		inMsgHandleStopChan:         make(chan struct{}),
		taskDoneChan:                make(chan struct{}),
		stopChan:                    make(chan struct{}),
		cachedWireBroadcastMsgLists: &sync.Map{},
		cachedWireUnicastMsgLists:   &sync.Map{},
		unConfirmedMessages:         make(map[string]*cache.LocalCacheItem),
		unConfirmedMsgLock:          &sync.Mutex{},
		finishedParties:             make(map[string]bool),
		culpritsLock:                &sync.RWMutex{},
		p2pComm:                     comm,
		blameMgr:                    blame.NewBlameManager(logger),
		logger:                      logger,
	}
	return tss, nil
}

// The current instance may be reused. Need to clear some information before using it
func (t *TssInstance) InitTssInfo(
	msgId string,
	msgNum int,
	privKey crypto.PrivKey,
	threshold uint64,
	conf TssConfig,
	preParams *bkg.LocalPreParams,
	keygenLocalState *storage.KeygenLocalState,
	peerMgr peer_mgr.OrderPeerManager,
	logger logrus.FieldLogger,
) (err error) {
	// keygen pre params
	// When using the keygen party it is recommended that you pre-compute the
	// "safe primes" and Paillier secret beforehand because this can take some
	// time.
	// This code will generate those parameters using a concurrency limit equal
	// to the number of available CPU cores.
	if preParams == nil || !preParams.Validate() {
		preParams, err = bkg.GeneratePreParams(conf.PreParamTimeout)
		if err != nil {
			return fmt.Errorf("fail to generate pre parameters: %w", err)
		}
	}
	if !preParams.Validate() {
		return fmt.Errorf("invalid preparams")
	}
	t.keygenPreParams = preParams

	t.localPrivK = privKey
	t.localPubK = privKey.GetPublic()
	t.threshold = int(threshold)
	t.conf = conf
	t.keygenLocalState = keygenLocalState

	t.msgID = msgId
	t.msgNum = msgNum
	t.localPartyID = ""
	t.partyInfo = &conversion.PartyInfo{
		PartyMap:   nil,
		PartyIDMap: make(map[string]*btss.PartyID),
	}
	t.partyLock = &sync.Mutex{}
	t.TssMsgChan = make(chan *pb.Message, msgNum)
	t.inMsgHandleStopChan = make(chan struct{})
	t.taskDoneChan = make(chan struct{})
	t.stopChan = make(chan struct{})
	t.cachedWireBroadcastMsgLists = &sync.Map{}
	t.cachedWireUnicastMsgLists = &sync.Map{}
	t.unConfirmedMessages = make(map[string]*cache.LocalCacheItem)
	t.unConfirmedMsgLock = &sync.Mutex{}
	t.finishedParties = make(map[string]bool)
	t.culprits = []*btss.PartyID{}
	t.culpritsLock = &sync.RWMutex{}
	t.blameMgr = blame.NewBlameManager(logger)

	comm, err := p2p.NewCommunication(peerMgr, logger)
	if err != nil {
		return fmt.Errorf("fail to create communication layer: %w", err)
	}
	t.p2pComm = comm
	t.p2pComm.Start()

	t.logger = logger
	t.logger.Debugf("..... init instance %s", msgId)
	return nil
}

func (t *TssInstance) setPartyInfo(partyInfo *conversion.PartyInfo) {
	t.partyLock.Lock()
	defer t.partyLock.Unlock()
	t.partyInfo = partyInfo
}

func (t *TssInstance) getPartyInfo() *conversion.PartyInfo {
	t.partyLock.Lock()
	defer t.partyLock.Unlock()
	return t.partyInfo
}

func (t *TssInstance) setLocalUnconfirmedMessages(key string, cacheItem *cache.LocalCacheItem) {
	t.unConfirmedMsgLock.Lock()
	defer t.unConfirmedMsgLock.Unlock()
	t.unConfirmedMessages[key] = cacheItem
}

func (t *TssInstance) getLocalCacheItem(key string) *cache.LocalCacheItem {
	t.unConfirmedMsgLock.Lock()
	defer t.unConfirmedMsgLock.Unlock()
	localCacheItem, ok := t.unConfirmedMessages[key]
	if !ok {
		return nil
	}
	return localCacheItem
}

func (t *TssInstance) renderToP2P(sendMsg *message.SendMsgChan) {
	if t.p2pComm.SendMsgChan == nil {
		t.logger.Warn("broadcast channel is not set")
		return
	}
	t.p2pComm.SendMsgChan <- sendMsg
}

func (t *TssInstance) PutTssMsg(msg *pb.Message) {
	t.TssMsgChan <- msg
	t.logger.WithFields(logrus.Fields{
		"msgID": t.msgID,
	}).Debugf("PutTssMsg")
}
