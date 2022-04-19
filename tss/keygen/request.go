package keygen

import (
	"fmt"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/meshplus/bitxhub-core/tss/conversion"
)

// Request request to do keygen
type Request struct {
	// parties pubkey
	Pubkeys []crypto.PubKey `json:"pubkeys"`
}

// NewRequest creeate a new instance of keygen.Request
func NewRequest(pubkeys []crypto.PubKey) Request {
	return Request{
		Pubkeys: pubkeys,
	}
}

func (r *Request) RequestToMsgId() (string, error) {
	keys := r.Pubkeys
	conversion.SortPubKey(keys)

	keyAccumulation := ""
	for _, el := range keys {
		pid, err := conversion.GetPIDFromPubKey(el)
		if err != nil {
			return "", fmt.Errorf("fail to get pid from pubkey")
		}
		keyAccumulation += pid.String()
	}

	var dat []byte
	dat = append(dat, []byte(keyAccumulation)...)
	return conversion.MsgToHashString(dat)
}
