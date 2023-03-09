package blame

import (
	"sync"

	"github.com/meshplus/bitxhub-core/tss/message"
)

// RoundMgr is used to record the confirmation of broadcast messages of each round.
// In the round of keygen and keysign, a participant needs to broadcast the hash of a broadcast message after receiving it
// to ensure that the broadcast message received by the participant is consistent with that received by other participants.
type RoundMgr struct {
	storedMsg   map[string]*message.TaskMessage
	storeLocker *sync.Mutex
}

func NewTssRoundMgr() *RoundMgr {
	return &RoundMgr{
		storeLocker: &sync.Mutex{},
		storedMsg:   make(map[string]*message.TaskMessage),
	}
}

func (tr *RoundMgr) Get(key string) *message.TaskMessage {
	tr.storeLocker.Lock()
	defer tr.storeLocker.Unlock()
	ret, ok := tr.storedMsg[key]
	if !ok {
		return nil
	}
	return ret
}

func (tr *RoundMgr) Set(key string, msg *message.TaskMessage) {
	tr.storeLocker.Lock()
	defer tr.storeLocker.Unlock()
	tr.storedMsg[key] = msg
}

// GetByRound queries all senders of messages in the specified round.
// That is, query participants that normally send messages in the specified round.
func (tr *RoundMgr) GetByRound(roundInfo string) []string {
	var standbyPeers []string
	tr.storeLocker.Lock()
	defer tr.storeLocker.Unlock()
	for _, el := range tr.storedMsg {
		if el.RoundInfo == roundInfo {
			standbyPeers = append(standbyPeers, el.Routing.From.Id)
		}
	}
	return standbyPeers
}
