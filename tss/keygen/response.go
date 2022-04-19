package keygen

import (
	"crypto/ecdsa"

	"github.com/meshplus/bitxhub-core/tss/blame"
	"github.com/meshplus/bitxhub-core/tss/conversion"
)

// Response keygen response
type Response struct {
	PubKey      *ecdsa.PublicKey  `json:"pub_key"`
	PoolAddress string            `json:"pool_address"`
	Status      conversion.Status `json:"status"`
	Blame       *blame.Blame      `json:"blame"`
}

// NewResponse create a new instance of keygen.Response
func NewResponse(pk *ecdsa.PublicKey, addr string, status conversion.Status, blame *blame.Blame) *Response {
	return &Response{
		PubKey:      pk,
		PoolAddress: addr,
		Status:      status,
		Blame:       blame,
	}
}
