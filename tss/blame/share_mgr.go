package blame

import (
	"sync"
)

// ShareMgr is used to manage the share request.
// When the current participant does not receive a broadcast message (or when the received broadcast message
//   is inconsistent with the majority of participants), the current participant requests the share of the broadcast
//   message from the majority of participants in agreement.
type ShareMgr struct {
	// requested indicates whether the current participant has requested sharing of the broadcast message from other participants
	requested map[string]bool
	reqLocker *sync.Mutex
}

func NewTssShareMgr() *ShareMgr {
	return &ShareMgr{
		reqLocker: &sync.Mutex{},
		requested: make(map[string]bool),
	}
}

// Set Sets the 「requested」 to true when making requests to other participants
func (sm *ShareMgr) Set(key string) {
	sm.reqLocker.Lock()
	defer sm.reqLocker.Unlock()
	sm.requested[key] = true
}

// QueryAndDelete determines whether the share result is needed according to the 「requested」 when receiving the result from other
// participants, and it deletes the 「requested」at the same time.
func (sm *ShareMgr) QueryAndDelete(key string) bool {
	sm.reqLocker.Lock()
	defer sm.reqLocker.Unlock()
	ret := sm.requested[key]
	delete(sm.requested, key)
	return ret
}
