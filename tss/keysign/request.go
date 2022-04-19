package keysign

import (
	"crypto/ecdsa"
	"fmt"
	"sort"
	"strings"

	"github.com/libp2p/go-libp2p-core/crypto"
	"github.com/meshplus/bitxhub-core/tss/conversion"
)

// Request request to sign a message
type Request struct {
	PoolPubKey    *ecdsa.PublicKey `json:"pool_pub_key"` // pub key of the pool that we would like to send this message from
	Messages      []string         `json:"messages"`     // base64 encoded message to be signed
	SignerPubKeys []crypto.PubKey  `json:"signer_pub_keys"`
}

func NewRequest(pk *ecdsa.PublicKey, msgs []string, signers []crypto.PubKey) Request {
	return Request{
		PoolPubKey:    pk,
		Messages:      msgs,
		SignerPubKeys: signers,
	}
}

func (r *Request) RequestToMsgId() (string, error) {
	var dat []byte

	sort.Strings(r.Messages)
	dat = []byte(strings.Join(r.Messages, ","))

	keys := r.SignerPubKeys
	conversion.SortPubKey(keys)
	keyAccumulation := ""
	for _, el := range keys {
		pid, err := conversion.GetPIDFromPubKey(el)
		if err != nil {
			return "", fmt.Errorf("fail to get pid from pubkey")
		}
		keyAccumulation += pid.String()
	}
	dat = append(dat, []byte(keyAccumulation)...)
	return conversion.MsgToHashString(dat)
}
