package validatorlib

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/asn1"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"

	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/hyperledger/fabric/protoutil"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/sm3"
	"github.com/tjfoc/gmsm/x509"
)

func fabric_sm2_validate(proof, validator []byte) (bool, error) {
	artifact, err := extractArtifacts(proof)
	if err != nil {
		return false, err
	}
	endorser := artifact.endorsements[0].Endorser
	signature := artifact.endorsements[0].Signature
	fmt.Println(signature)

	serializedID := &msp.SerializedIdentity{}
	err = serializedID.XXX_Unmarshal(endorser)
	if err != nil {
		return false, err
	}
	block, _ := pem.Decode(serializedID.IdBytes)
	if block == nil {
		return false, fmt.Errorf("block is nil")
	}
	rBlock, _ := pem.Decode(validator)
	if rBlock == nil {
		return false, fmt.Errorf("block is nil")
	}
	rCert, err := x509.ParseCertificate(rBlock.Bytes)
	if err != nil {
		return false, err
	}
	if !bytes.Equal(block.Bytes, rBlock.Bytes) {
		return false, fmt.Errorf("block not equal")
	}
	r, s, err := UnmarshalSM2Signature(signature)
	if err != nil {
		return false, err
	}

	PubKey := rCert.PublicKey.(*ecdsa.PublicKey)
	sm2PubKey := &sm2.PublicKey{
		Curve: PubKey.Curve,
		X:     PubKey.X,
		Y:     PubKey.Y,
	}
	h := sm3.New()
	h.Write(append(artifact.cap.Action.ProposalResponsePayload, artifact.cap.Action.Endorsements[0].Endorser...))
	out2 := h.Sum(nil)
	hDigest := HashMsgZa(out2, sm2PubKey)

	return sm2.Verify(sm2PubKey, hDigest, r, s), nil

}

func HashMsgZa(msg []byte, pub *sm2.PublicKey) []byte {
	za, _ := sm2.ZA(pub, []byte{0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38})

	var buffer bytes.Buffer
	buffer.Write(za)
	buffer.Write(msg)
	m := buffer.Bytes()

	h := sm3.New()
	h.Write(m)
	digest := h.Sum(nil)

	return digest
}

func extractArtifacts(proof []byte) (*valiadationArtifacts, error) {
	cap, err := protoutil.UnmarshalChaincodeActionPayload(proof)
	if err != nil {
		return nil, err
	}

	pRespPayload, err := protoutil.UnmarshalProposalResponsePayload(cap.Action.ProposalResponsePayload)
	if err != nil {
		err = fmt.Errorf("GetProposalResponsePayload error %s", err)
		return nil, err
	}
	if pRespPayload.Extension == nil {
		err = fmt.Errorf("nil pRespPayload.Extension")
		return nil, err
	}
	respPayload, err := protoutil.UnmarshalChaincodeAction(pRespPayload.Extension)
	if err != nil {
		err = fmt.Errorf("GetChaincodeAction error %s", err)
		return nil, err
	}

	return &valiadationArtifacts{
		rwset:        respPayload.Results,
		prp:          cap.Action.ProposalResponsePayload,
		endorsements: cap.Action.Endorsements,
		cap:          cap,
	}, nil
}

type SM2Signature struct {
	R, S *big.Int
}

func UnmarshalSM2Signature(raw []byte) (*big.Int, *big.Int, error) {
	// Unmarshal
	sig := new(SM2Signature)
	_, err := asn1.Unmarshal(raw, sig)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed unmashalling signature [%s]", err)
	}

	// Validate sig
	if sig.R == nil {
		return nil, nil, errors.New("Invalid signature. R must be different from nil.")
	}
	if sig.S == nil {
		return nil, nil, errors.New("Invalid signature. S must be different from nil.")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, errors.New("Invalid signature. R must be larger than zero")
	}
	if sig.S.Sign() != 1 {
		return nil, nil, errors.New("Invalid signature. S must be larger than zero")
	}

	return sig.R, sig.S, nil
}

func intToBytes(x int) []byte {
	var b = make([]byte, 4)
	b[0] = byte(x >> 24)
	b[1] = byte(x >> 16)
	b[2] = byte(x >> 8)
	b[3] = byte(x)
	return b
}
