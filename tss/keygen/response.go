package keygen

import (
	"crypto/ecdsa"

	"github.com/meshplus/bitxhub-core/tss/blame"
	"github.com/meshplus/bitxhub-core/tss/storage"
)

// Response keygen response
type Response struct {
	PubKey      *ecdsa.PublicKey          `json:"pub_key"`
	PoolAddress string                    `json:"pool_address"`
	LocalState  *storage.KeygenLocalState `json:"local_state"`
	//Status      conversion.Status `json:"status"`
	Blame *blame.Blame `json:"blame"`
}

// NewResponse create a new instance of keygen.Response
func NewResponse(pk *ecdsa.PublicKey, addr string, state *storage.KeygenLocalState, blame *blame.Blame) *Response {
	return &Response{
		PubKey:      pk,
		PoolAddress: addr,
		LocalState:  state,
		//Status:      status,
		Blame: blame,
	}
}
