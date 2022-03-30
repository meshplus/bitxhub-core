package blame

import (
	"sync"

	btss "github.com/binance-chain/tss-lib/tss"
	"github.com/meshplus/bitxhub-core/tss/conversion"
	"github.com/sirupsen/logrus"
)

type Manager struct {
	Blame    *Blame
	ShareMgr *ShareMgr
	RoundMgr *RoundMgr

	// partyInfo is same to the partyInfo of TssManager
	partyInfo *conversion.PartyInfo
	// localPartyID is the partyID.id of localParty
	localPartyID string

	// lastUnicastParty stores the mapping between the round information and the participant information of the last unicast message
	lastUnicastParty map[string][]string
	// lastMsg stores the last msg
	lastMsg       btss.Message
	lastMsgLocker *sync.RWMutex

	// acceptedShares stores participants that have accepted all shared information for the specified round
	acceptedShares    map[conversion.RoundInfo][]string
	acceptShareLocker *sync.Mutex

	logger logrus.FieldLogger
}

func NewBlameManager(logger logrus.FieldLogger) *Manager {
	blame := NewBlame("", nil)
	return &Manager{
		Blame:    &blame,
		ShareMgr: NewTssShareMgr(),
		RoundMgr: NewTssRoundMgr(),
		partyInfo: &conversion.PartyInfo{
			PartyMap:   nil,
			PartyIDMap: make(map[string]*btss.PartyID),
		},
		lastUnicastParty:  make(map[string][]string),
		lastMsgLocker:     &sync.RWMutex{},
		acceptedShares:    make(map[conversion.RoundInfo][]string),
		acceptShareLocker: &sync.Mutex{},
		logger:            logger,
	}
}

func (m *Manager) SetPartyInfo(partyInfo *conversion.PartyInfo) {
	m.lastMsgLocker.Lock()
	defer m.lastMsgLocker.Unlock()

	m.partyInfo = partyInfo

	var localParty btss.Party
	m.partyInfo.PartyMap.Range(func(key, value interface{}) bool {
		localParty = value.(btss.Party)
		return false
	})

	m.localPartyID = localParty.PartyID().Id
}

// UpdateAcceptShare adds a participant who have accepted all shared information for the specified round
func (m *Manager) UpdateAcceptShare(round conversion.RoundInfo, id string) {
	m.acceptShareLocker.Lock()
	defer m.acceptShareLocker.Unlock()
	partyList, ok := m.acceptedShares[round]
	if !ok {
		partyList := []string{id}
		m.acceptedShares[round] = partyList
		return
	}
	partyList = append(partyList, id)
	m.acceptedShares[round] = partyList
}

// CheckMsgDuplication checks whether a participant is already in the acceptedShares list for the specified round
func (m *Manager) CheckMsgDuplication(round conversion.RoundInfo, id string) bool {
	m.acceptShareLocker.Lock()
	defer m.acceptShareLocker.Unlock()
	partyList, ok := m.acceptedShares[round]
	if ok {
		for _, el := range partyList {
			if el == id {
				return true
			}
		}
	}
	return false
}

func (m *Manager) SetLastMsg(lastMsg btss.Message) {
	m.lastMsgLocker.Lock()
	defer m.lastMsgLocker.Unlock()
	m.lastMsg = lastMsg
}

func (m *Manager) GetLastMsg() btss.Message {
	m.lastMsgLocker.RLock()
	defer m.lastMsgLocker.RUnlock()
	return m.lastMsg
}

func (m *Manager) SetLastUnicastParty(partyID string, roundInfo string) {
	m.lastMsgLocker.Lock()
	defer m.lastMsgLocker.Unlock()

	l, ok := m.lastUnicastParty[roundInfo]
	if !ok {
		peerList := []string{partyID}
		m.lastUnicastParty[roundInfo] = peerList
	} else {
		l = append(l, partyID)
		m.lastUnicastParty[roundInfo] = l
	}
}
