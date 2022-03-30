package keysign

import (
	"github.com/meshplus/bitxhub-core/tss/blame"
	"github.com/meshplus/bitxhub-core/tss/conversion"
)

// Response key sign response
type Response struct {
	Signatures []conversion.Signature `json:"signatures"`
	Status     conversion.Status      `json:"status"`
	Blame      *blame.Blame           `json:"blame"`
}

func NewResponse(signatures []conversion.Signature, status conversion.Status, blame *blame.Blame) *Response {
	return &Response{
		Signatures: signatures,
		Status:     status,
		Blame:      blame,
	}
}
