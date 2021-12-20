package validator

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type FabV14DynamicValidator struct {
	logger       logrus.FieldLogger
	fabValidator Validator
}

type V14DynamicValidator struct {
	Height   []uint64 `json:"height"`
	ChainId  string   `json:"chain_id"`
	ConfByte []string `json:"conf_byte"`
	Policy   []string `json:"policy"`
	Cid      string   `json:"cid"`
}

// New a validator instance
func NewFabV14DynamicValidator(logger logrus.FieldLogger) *FabV14DynamicValidator {
	return &FabV14DynamicValidator{
		logger:       logger,
		fabValidator: NewFabV14Validator(logger),
	}
}

// Verify will check whether the transaction info is valid
func (vlt *FabV14DynamicValidator) Verify(from string, proof, payload []byte, validators string) (bool, uint64, error) {
	// ----------------parse proof and validators start----------------------------
	dynamicProof := &DynamicProof{}
	err := json.Unmarshal(proof, dynamicProof)
	if err != nil {
		return false, 0, err
	}
	dynamicValidator := &V14DynamicValidator{}
	err = json.Unmarshal([]byte(validators), dynamicValidator)
	if err != nil {
		return false, 0, err
	}
	// Check that the array is in order and that pem length is equal to height length
	if len(dynamicValidator.Height) == 0 || len(dynamicValidator.Height) != len(dynamicValidator.Policy) {
		return false, 0, fmt.Errorf("len of dynamic validator invalid ")
	}
	for i := 0; i < len(dynamicValidator.Height)-1; i++ {
		if dynamicValidator.Height[i] > dynamicValidator.Height[i+1] {
			return false, 0, fmt.Errorf("dynamic validator height slice not ordered")
		}
	}
	if dynamicProof.height >= dynamicValidator.Height[len(dynamicValidator.Height)-1] {
		validators = dynamicValidator.Policy[len(dynamicValidator.Height)-1]
	} else {
		index := sort.Search(len(dynamicValidator.Height), func(i int) bool { return dynamicValidator.Height[i+1] > dynamicProof.height })
		validators = dynamicValidator.Policy[index]
	}
	validatorInfo, err := validatorlib.MarshalValidatorInfo(&validatorlib.ValidatorInfo{
		ChainId:  dynamicValidator.ChainId,
		ConfByte: dynamicValidator.ConfByte,
		Cid:      dynamicValidator.Cid,
		Policy:   validators,
	})
	if err != nil {
		return false, 0, err
	}
	// ----------------parse proof and validators end----------------------------
	return vlt.fabValidator.Verify(from, dynamicProof.Proof, payload, string(validatorInfo))
}
