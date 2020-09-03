package validator

import (
	"github.com/sirupsen/logrus"
	"github.com/wasmerio/go-ext-wasm/wasmer"
)

const (
	FabricRuleAddr    = "0x00000000000000000000000000000000000000a0"
	SimFabricRuleAddr = "0x00000000000000000000000000000000000000a1"
)

// Validator is the instance that can use wasm to verify transaction validity
type ValidationEngine struct {
	instances       map[string]wasmer.Instance
	fabValidator    Validator
	simFabValidator Validator

	ledger Ledger
	logger logrus.FieldLogger
}

// New a validator instance
func NewValidationEngine(ledger Ledger, instances *sync.Map, logger logrus.FieldLogger) *ValidationEngine {
	return &ValidationEngine{
		ledger:          ledger,
		logger:          logger,
		fabValidator:    NewFabV14Validator(logger),
		simFabValidator: NewFabSimValidator(logger),
		instances:       instances,
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
		ve.instances = make(map[string]wasmer.Instance)
	}

	return NewWasmValidator(ve.ledger, ve.logger, ve.instances)
}
