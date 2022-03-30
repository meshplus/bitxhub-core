package cache

import (
	"sync"

	"github.com/meshplus/bitxhub-core/tss/message"
)

// LocalCacheItem used to cache the unconfirmed broadcast message
type LocalCacheItem struct {
	Msg  *message.WireMessage
	Hash string
	Lock *sync.Mutex
	// parytyID.id -> msgHash
	ConfirmedList map[string]string
}

func NewLocalCacheItem(msg *message.WireMessage, hash string) *LocalCacheItem {
	return &LocalCacheItem{
		Msg:           msg,
		Hash:          hash,
		Lock:          &sync.Mutex{},
		ConfirmedList: make(map[string]string),
	}
}

// UpdateConfirmList add the given party's hash into the confirm list
func (l *LocalCacheItem) UpdateConfirmList(partyID, hash string) {
	l.Lock.Lock()
	defer l.Lock.Unlock()
	l.ConfirmedList[partyID] = hash
}

// TotalConfirmParty counts the number of parties that already confirmed their hash
func (l *LocalCacheItem) TotalConfirmParty() int {
	l.Lock.Lock()
	defer l.Lock.Unlock()
	return len(l.ConfirmedList)
}

func (l *LocalCacheItem) GetPeers() []string {
	peers := make([]string, 0, len(l.ConfirmedList))
	l.Lock.Lock()
	defer l.Lock.Unlock()
	for peer := range l.ConfirmedList {
		peers = append(peers, peer)
	}
	return peers
}
