package conversion

import (
	"bytes"
	"encoding/base64"

	bcommon "github.com/binance-chain/tss-lib/common"
	"github.com/libp2p/go-libp2p-core/crypto"
)

type Signature struct {
	Msg         string `json:"signed_msg"`
	R           string `json:"r"`
	S           string `json:"s"`
	RecoveryID  string `json:"recovery_id"`
	SignEthData []byte `json:"sign_eth_data"`
}

func NewSignature(msg, r, s, recoveryID string, data []byte) Signature {
	return Signature{
		Msg:         msg,
		R:           r,
		S:           s,
		RecoveryID:  recoveryID,
		SignEthData: data,
	}
}

// BatchSignatures package the signature list and message list as message signature list
func BatchSignatures(sigs []*bcommon.ECSignature, msgsToSign [][]byte) []Signature {
	signatures := []Signature{}
	for i, sig := range sigs {
		msg := base64.StdEncoding.EncodeToString(msgsToSign[i])
		r := base64.StdEncoding.EncodeToString(sig.R)
		s := base64.StdEncoding.EncodeToString(sig.S)
		recovery := base64.StdEncoding.EncodeToString(sig.SignatureRecovery)

		signData := []byte{}
		signData = append(signData, sig.Signature...)
		signData = append(signData, sig.SignatureRecovery...)

		signature := NewSignature(msg, r, s, recovery, signData)
		signatures = append(signatures, signature)
	}

	return signatures
}

// Signing and validation before and after sending and receiving messages to ensure that you can specify who sent each message

func GenerateSignature(msg []byte, msgID string, privKey crypto.PrivKey) ([]byte, error) {
	var dataForSigning bytes.Buffer
	dataForSigning.Write(msg)
	dataForSigning.WriteString(msgID)
	return privKey.Sign(dataForSigning.Bytes())
}

func VerifySignature(pubKey crypto.PubKey, message, sig []byte, msgID string) (bool, error) {
	var dataForSign bytes.Buffer
	dataForSign.Write(message)
	dataForSign.WriteString(msgID)
	return pubKey.Verify(dataForSign.Bytes(), sig)
}
