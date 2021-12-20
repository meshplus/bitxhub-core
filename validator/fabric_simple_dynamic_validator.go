package validator

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type FabSimDynamicValidator struct {
	logger          logrus.FieldLogger
	simFabValidator Validator
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
		logger:          logger,
		simFabValidator: NewFabSimValidator(logger),
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
	return vlt.simFabValidator.Verify(from, dynamicProof.Proof, payload, validators)
}
