package validator

import (
	"sync"

	"github.com/sirupsen/logrus"
)

const (
	FabricRuleAddr    = "0x00000000000000000000000000000000000000a0"
	SimFabricRuleAddr = "0x00000000000000000000000000000000000000a1"
)

// Validator is the instance that can use wasm to verify transaction validity
type ValidationEngine struct {
	instances       *sync.Map
	fabValidator    Validator
	simFabValidator Validator

	ledger Ledger
	logger logrus.FieldLogger
}

// New a validator instance
func NewValidationEngine(ledger Ledger, logger logrus.FieldLogger) *ValidationEngine {
	return &ValidationEngine{
		ledger:          ledger,
		logger:          logger,
		fabValidator:    NewFabV14Validator(logger),
		simFabValidator: NewFabSimValidator(logger),
	}
}

// Verify will check whether the transaction info is valid
func (ve *ValidationEngine) Validate(address, from string, proof, payload []byte, validators string) (bool, error) {
	vlt := ve.getValidator(address)

	return vlt.Verify(address, from, proof, payload, validators)
}

func (ve *ValidationEngine) getValidator(address string) Validator {
	if address == FabricRuleAddr {
		return ve.fabValidator
	}

	if address == SimFabricRuleAddr {
		return ve.simFabValidator
	}

	if ve.instances == nil {
		ve.instances = &sync.Map{}
	}

	return NewWasmValidator(ve.ledger, ve.logger, ve.instances)
}
