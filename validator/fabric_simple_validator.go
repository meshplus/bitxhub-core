package validator

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type FabSimValidator struct {
	logger logrus.FieldLogger
	pkMap  map[string]*ecdsa.PublicKey
}

// New a validator instance
func NewFabSimValidator(logger logrus.FieldLogger) *FabSimValidator {
	return &FabSimValidator{
		logger: logger,
		pkMap:  make(map[string]*ecdsa.PublicKey),
	}
}

// Verify will check whether the transaction info is valid
func (vlt *FabSimValidator) Verify(address, from string, proof, payload []byte, validators string) (bool, error) {
	artifact, err := validatorlib.PreCheck(proof, payload, "broker")
	if err != nil {
		return false, err
	}

	signatureSet := validatorlib.GetSignatureSet(artifact)

	var pk *ecdsa.PublicKey
	pk, ok := vlt.pkMap["111"]
	if !ok {
		pemCert, _ := pem.Decode([]byte(validators))
		cert, err := x509.ParseCertificate(pemCert.Bytes)
		if err != nil {
			return false, err
		}
		pk = cert.PublicKey.(*ecdsa.PublicKey)
		vlt.pkMap["111"] = pk
	}
	r, s, err := unmarshalECDSASignature(signatureSet[0].Signature)
	if err != nil {
		return false, err
	}

	h := sha256.New()
	h.Write(signatureSet[0].Data)
	ret := h.Sum(nil)
	isValid := ecdsa.Verify(pk, ret, r, s)

	return isValid, nil
}

type ECDSASignature struct {
	R, S *big.Int
}

func unmarshalECDSASignature(raw []byte) (*big.Int, *big.Int, error) {
	sig := new(ECDSASignature)
	_, err := asn1.Unmarshal(raw, sig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed unmashalling signature [%s]", err)
	}

	// Validate sig
	if sig.R == nil {
		return nil, nil, fmt.Errorf("invalid signature, r must be different from nil")
	}
	if sig.S == nil {
		return nil, nil, fmt.Errorf("invalid signature, s must be different from nil")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, fmt.Errorf("invalid signature, r must be larger than zero")
	}
	if sig.S.Sign() != 1 {
		return nil, nil, fmt.Errorf("invalid signature, s must be larger than zero")
	}

	return sig.R, sig.S, nil
}
