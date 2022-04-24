package tss

import (
	"crypto/ecdsa"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"

	"github.com/meshplus/bitxhub-core/tss/cache"

	bkg "github.com/binance-chain/tss-lib/ecdsa/keygen"
	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/libp2p/go-libp2p-core/crypto"
	peer_mgr "github.com/meshplus/bitxhub-core/peer-mgr"
	"github.com/meshplus/bitxhub-core/tss/blame"
	"github.com/meshplus/bitxhub-core/tss/conversion"
	"github.com/meshplus/bitxhub-core/tss/message"
	"github.com/meshplus/bitxhub-core/tss/p2p"
	"github.com/meshplus/bitxhub-core/tss/storage"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
)

var _ Tss = (*TssManager)(nil)

// TssManager is the structure that can provide all keysign and key gen features
type TssManager struct {
	msgID string
	// the number of msgs contained in the current request (1 for Keygen request is and n for Keysign request)
	msgNum          int
	threshold       int
	keygenPreParams *bkg.LocalPreParams
	conf            TssConfig
	repoPath        string

	// local party info
	localPrivK   crypto.PrivKey
	localPubK    crypto.PubKey
	localPartyID string
	// other parties info
	partyLock *sync.Mutex
	partyInfo *conversion.PartyInfo

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

	tssKeyGenLocker *sync.Mutex

	p2pComm  *p2p.Communication
	stateMgr storage.LocalStateManager
	blameMgr *blame.Manager
	logger   logrus.FieldLogger
}

// NewTss creates a new instance of Tss
func NewTss(
	repoPath string,
	peerMgr peer_mgr.OrderPeerManager,
	conf TssConfig,
	threshold int,
	privKey crypto.PrivKey,
	logger logrus.FieldLogger,
	baseFolder string,
	preParams *bkg.LocalPreParams,
) (*TssManager, error) {
	// Persistent storage of data
	stateManager, err := storage.NewFileStateMgr(baseFolder)
	if err != nil {
		return nil, fmt.Errorf("fail to create file state manager: %w", err)
	}

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

	tss := &TssManager{
		keygenPreParams: preParams,
		conf:            conf,
		repoPath:        repoPath,
		localPrivK:      privKey,
		localPubK:       privKey.GetPublic(),
		partyLock:       &sync.Mutex{},
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
		tssKeyGenLocker:             &sync.Mutex{},
		p2pComm:                     comm,
		stateMgr:                    stateManager,
		blameMgr:                    blame.NewBlameManager(logger),
		logger:                      logger,
	}
	return tss, nil
}

// Set message related information at the start of Keygen and Keysign tasks
func (t *TssManager) setTssMsgInfo(
	msgID string,
	msgNum int,
) {
	t.msgID = msgID
	t.msgNum = msgNum
	t.TssMsgChan = make(chan *pb.Message, msgNum)
	t.inMsgHandleStopChan = make(chan struct{})
	t.taskDoneChan = make(chan struct{})
}

func (t *TssManager) setPartyInfo(partyInfo *conversion.PartyInfo) {
	t.partyLock.Lock()
	defer t.partyLock.Unlock()
	t.partyInfo = partyInfo
}

func (t *TssManager) getPartyInfo() *conversion.PartyInfo {
	t.partyLock.Lock()
	defer t.partyLock.Unlock()
	return t.partyInfo
}

func (t *TssManager) setLocalUnconfirmedMessages(key string, cacheItem *cache.LocalCacheItem) {
	t.unConfirmedMsgLock.Lock()
	defer t.unConfirmedMsgLock.Unlock()
	t.unConfirmedMessages[key] = cacheItem
}

func (t *TssManager) getLocalCacheItem(key string) *cache.LocalCacheItem {
	t.unConfirmedMsgLock.Lock()
	defer t.unConfirmedMsgLock.Unlock()
	localCacheItem, ok := t.unConfirmedMessages[key]
	if !ok {
		return nil
	}
	return localCacheItem
}

func (t *TssManager) renderToP2P(sendMsg *message.SendMsgChan) {
	if t.p2pComm.SendMsgChan == nil {
		t.logger.Warn("broadcast channel is not set")
		return
	}
	t.p2pComm.SendMsgChan <- sendMsg
}

// External interface ==================================================================================================

func (t *TssManager) Start(threshold uint64) {
	t.threshold = int(threshold)
	t.logger.Infof("Starting the TSS Manager: n-%d, t-%d", len(t.partyInfo.PartyIDMap), threshold)
}

func (t *TssManager) Stop() {
	close(t.stopChan)
	err := t.p2pComm.Stop()
	if err != nil {
		t.logger.Error("error in shutdown the p2p server")
	}
	t.logger.Info("The Tss and p2p server has been stopped successfully")
}

func (t *TssManager) PutTssMsg(msg *pb.Message) {
	t.TssMsgChan <- msg
	return
}

func (t *TssManager) GetTssPubkey() (string, *ecdsa.PublicKey, error) {
	// 1. get pool addr from file
	filePath := filepath.Join(t.repoPath, t.conf.TssConfPath, storage.PoolPkAddrFileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", nil, err
	}

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", nil, fmt.Errorf("file to read from file(%s): %w", filePath, err)
	}

	// 2. get local state by pool addr
	state, err := t.stateMgr.GetLocalState(string(buf))
	if err != nil {
		return "", nil, fmt.Errorf("failed to get local state: %s,  %v", string(buf), err)
	}

	// 3. get tss pk from local state
	pk, err := conversion.GetECDSAPubKeyFromPubKeyData(state.PubKeyData)
	if err != nil {
		return "", nil, fmt.Errorf("failed to get ECDSA pubKey from pubkey data: %v", err)
	}

	return state.PubKeyAddr, pk, nil
}

func (t *TssManager) GetTssInfo() (*pb.TssInfo, error) {
	// 1. get pool addr from file
	filePath := filepath.Join(t.repoPath, t.conf.TssConfPath, storage.PoolPkAddrFileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, err
	}

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("file to read from file(%s): %w", filePath, err)
	}

	// 2. get local state by pool addr
	state, err := t.stateMgr.GetLocalState(string(buf))
	if err != nil {
		return nil, fmt.Errorf("failed to get local state: %s,  %v", string(buf), err)
	}

	// 3. get parties pks from local state
	return &pb.TssInfo{
		PartiesPkMap: state.ParticipantPksMap,
		Pubkey:       state.PubKeyData,
	}, nil
}

func (t *TssManager) DeleteCulpritsFromLocalState(culprits []string) error {
	// 1. get pool addr from file
	filePath := filepath.Join(t.repoPath, t.conf.TssConfPath, storage.PoolPkAddrFileName)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return err
	}

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("file to read from file(%s): %w", filePath, err)
	}

	// 2. get local state by pool addr
	state, err := t.stateMgr.GetLocalState(string(buf))
	if err != nil {
		return fmt.Errorf("failed to get local state: %s,  %v", string(buf), err)
	}

	// 3. delete culprits
	for _, id := range culprits {
		delete(state.ParticipantPksMap, id)
	}

	// 4. update local state
	return t.stateMgr.SaveLocalState(state)
}
