package validator

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

const (
	FabricRuleAddr = "0x00000000000000000000000000000000000000a0"
	HpcRuleAddr    = "0x00000000000000000000000000000000000000a1"
)

// Validator is the instance that can use wasm to verify transaction validity
type ValidationEngine struct {
	ledger Ledger
	logger logrus.FieldLogger
}

// New a validator instance
func NewValidationEngine(ledger Ledger, logger logrus.FieldLogger) *ValidationEngine {
	return &ValidationEngine{
		ledger: ledger,
		logger: logger,
	}
}

// Verify will check whether the transaction info is valid
func (ve *ValidationEngine) Validate(address, from string, proof, payload []byte, validators string) (bool, error) {
	vlt := ve.getValidator(address)

	if vlt == nil {
		return false, fmt.Errorf("wrong rule address")
	}

	return vlt.Verify(address, from, proof, payload, validators)
}

func (ve *ValidationEngine) getValidator(address string) Validator {
	if address == FabricRuleAddr {
		return NewFabV14Validator(ve.logger)
	}

	if address == HpcRuleAddr {
		return NewHpcValidator(ve.logger)
	}

	return nil
}
