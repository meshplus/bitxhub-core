package validator

import (
	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type FabV14Validator struct {
	logger logrus.FieldLogger
}

// New a validator instance
func NewFabV14Validator(logger logrus.FieldLogger) *FabV14Validator {
	return &FabV14Validator{
		logger: logger,
	}
}

// Verify will check whether the transaction info is valid
func (vlt *FabV14Validator) Verify(address, from string, proof, payload []byte, validators string) (bool, error) {
	vInfo, err := validatorlib.UnmarshalValidatorInfo([]byte(validators))
	if err != nil {
		return false, err
	}
	err = validatorlib.ValidateV14(proof, payload, []byte(vInfo.Policy), vInfo.ConfByte, vInfo.Cid)
	if err != nil {
		return false, err
	}

	return true, nil
}
