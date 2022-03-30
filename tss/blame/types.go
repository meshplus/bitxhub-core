package blame

import "fmt"

const (
	HashCheckFail = "hash check failed"
	TssTimeout    = "Tss timeout"
	TssBrokenMsg  = "tss share verification failed"
	InternalError = "fail to start the join party "
)

var (
	ErrHashFromOwner     = fmt.Errorf(" hash sent from data owner")
	ErrNotEnoughPeer     = fmt.Errorf("not enough nodes to evaluate hash")
	ErrNotMajority       = fmt.Errorf("message we received does not match the majority")
	ErrTssTimeOut        = fmt.Errorf("error Tss Timeout")
	ErrHashCheck         = fmt.Errorf("error in processing hash check")
	ErrHashInconsistency = fmt.Errorf("fail to agree on the hash value")
)
