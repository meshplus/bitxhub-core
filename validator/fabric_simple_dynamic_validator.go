package validator

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"sort"
	"sync"

	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type FabSimDynamicValidator struct {
	logger logrus.FieldLogger
	pkMap  *sync.Map
}

type DynamicProof struct {
	Proof  []byte `json:"proof"`
	height uint64 `json:"height"`
}

type DynamicValidator struct {
	Height []uint64 `json:"height"`
	Pem    []string `json:"pem"`
}

// New a validator instance
func NewFabSimDynamicValidator(logger logrus.FieldLogger) *FabSimDynamicValidator {
	return &FabSimDynamicValidator{
		logger: logger,
		pkMap:  &sync.Map{},
	}
}

// Verify will check whether the transaction info is valid
func (vlt *FabSimDynamicValidator) Verify(from string, proof, payload []byte, validators string) (bool, uint64, error) {
	// ----------------parse proof and validators start----------------------------
	dynamicProof := &DynamicProof{}
	err := json.Unmarshal(proof, dynamicProof)
	if err != nil {
		return false, 0, err
	}
	dynamicValidator := &DynamicValidator{}
	err = json.Unmarshal([]byte(validators), dynamicValidator)
	if err != nil {
		return false, 0, err
	}
	// Check that the array is in order and that pem length is equal to height length
	if len(dynamicValidator.Height) == 0 || len(dynamicValidator.Height) != len(dynamicValidator.Pem) {
		return false, 0, fmt.Errorf("len of dynamic validator invalid ")
	}
	for i := 0; i < len(dynamicValidator.Height)-1; i++ {
		if dynamicValidator.Height[i] > dynamicValidator.Height[i+1] {
			return false, 0, fmt.Errorf("dynamic validator height slice not ordered")
		}
	}
	if dynamicProof.height >= dynamicValidator.Height[len(dynamicValidator.Height)-1] {
		validators = dynamicValidator.Pem[len(dynamicValidator.Height)-1]
	} else {
		index := sort.Search(len(dynamicValidator.Height), func(i int) bool { return dynamicValidator.Height[i+1] > dynamicProof.height })
		validators = dynamicValidator.Pem[index]
	}
	// ----------------parse proof and validators end----------------------------

	artifact, err := validatorlib.PreCheck(dynamicProof.Proof, payload, "broker")
	if err != nil {
		return false, 0, err
	}

	signatureSet := validatorlib.GetSignatureSet(artifact)

	var pk *ecdsa.PublicKey
	raw, ok := vlt.pkMap.Load(from)
	if !ok {
		pemCert, _ := pem.Decode([]byte(validators))
		if pemCert == nil {
			return false, 0, fmt.Errorf("invalid validators information: %s", validators)
		}
		cert, err := x509.ParseCertificate(pemCert.Bytes)
		if err != nil {
			return false, 0, err
		}
		pk = cert.PublicKey.(*ecdsa.PublicKey)
		vlt.pkMap.Store(from, pk)
	} else {
		pk = raw.(*ecdsa.PublicKey)
	}

	r, s, err := unmarshalECDSASignature(signatureSet[0].Signature)
	if err != nil {
		return false, 0, err
	}

	h := sha256.New()
	_, err = h.Write(signatureSet[0].Data)
	if err != nil {
		return false, 0, err
	}
	ret := h.Sum(nil)
	isValid := ecdsa.Verify(pk, ret, r, s)
	if !isValid {
		return false, 0, fmt.Errorf("signature not right")
	}
	return true, 0, nil
}
