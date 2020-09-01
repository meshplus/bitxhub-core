package validator

import (
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	FabricRuleAddr = "0x00000000000000000000000000000000000000a0"
)

// Validator is the instance that can use wasm to verify transaction validity
type ValidationEngine struct {
	ledger    Ledger
	logger    logrus.FieldLogger
	instances *sync.Map
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

	return vlt.Verify(address, from, proof, payload, validators)
}

func (ve *ValidationEngine) getValidator(address string) Validator {
	if address == FabricRuleAddr {
		return NewFabV14Validator(ve.logger)
	}

	if ve.instances == nil {
		ve.instances = &sync.Map{}
	}

	return NewWasmValidator(ve.ledger, ve.logger, ve.instances)
}
