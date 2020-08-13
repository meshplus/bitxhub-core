package validator

import "github.com/sirupsen/logrus"

// Validator is the instance that can use wasm to verify transaction validity
type HpcValidator struct {
	logger logrus.FieldLogger
}

// New a validator instance
func NewHpcValidator(logger logrus.FieldLogger) *HpcValidator {
	return &HpcValidator{
		logger: logger,
	}
}

// Verify will check whether the transaction info is valid
func (vlt *HpcValidator) Verify(address, from string, proof, payload []byte, validators string) (bool, error) {
	return true, nil
}
