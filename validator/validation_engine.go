package validator

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/meshplus/bitxhub-core/wasm"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/sirupsen/logrus"
)

const (
	FabricRuleAddr    = "0x00000000000000000000000000000000000000a0"
	SimFabricRuleAddr = "0x00000000000000000000000000000000000000a1"
	HappyRuleAddr     = "0x00000000000000000000000000000000000000a2"
	MultiSignAddr     = "0x00000000000000000000000000000000000000a3"
)

// Validator is the instance that can use wasm to verify transaction validity
type ValidationEngine struct {
	pools           *ValidatorPools
	fabValidator    Validator
	simFabValidator Validator
	happyValidator  Validator
	wasmGasLimit    uint64
	ledger          Ledger
	logger          logrus.FieldLogger

	sync.RWMutex
}

// New a validator instance
func NewValidationEngine(ledger Ledger, instances *sync.Map, logger logrus.FieldLogger, gasLimit uint64) *ValidationEngine {
	return &ValidationEngine{
		ledger:          ledger,
		logger:          logger,
		fabValidator:    NewFabV14Validator(logger),
		simFabValidator: NewFabSimValidator(logger),
		happyValidator:  &HappyValidator{},
		pools:           NewValidationPools(),
		wasmGasLimit:    gasLimit,
	}
}

// Verify will check whether the transaction info is valid
func (ve *ValidationEngine) Validate(address, from string, proof, payload []byte, validators string) (bool, uint64, error) {
	pool, ok := ve.pools.GetPool(address)
	if !ok {
		ve.Lock()
		defer ve.Unlock()
		ve.pools.SetPool(address, NewValidationPool(10))
		vlt, err := ve.getValidator(address)
		if err != nil {
			return false, 0, err
		}
		ve.pools.pools[address].Add(vlt)
		pool, _ = ve.pools.GetPool(address)
	} else {
		// ve.Lock()
		// defer ve.Unlock()
		if pool.length < pool.size {
			vlt, err := ve.getValidator(address)
			if err != nil {
				return false, 0, err
			}
			pool.Add(vlt)
		}
	}

	vlt := pool.GetValidator()

	ok, gasUsed, err := vlt.Verify(from, proof, payload, validators)
	pool.SetValidator(vlt)
	return ok, gasUsed, err
}

func (ve *ValidationEngine) getValidator(address string) (Validator, error) {
	if address == FabricRuleAddr {
		return NewFabV14Validator(ve.logger), nil
	}

	if address == SimFabricRuleAddr {
		return NewFabSimValidator(ve.logger), nil
	}

	if address == MultiSignAddr {
		return NewMultiSignValidator(ve.logger), nil
	}

	if address == HappyRuleAddr {
		return &HappyValidator{}, nil
	}

	contractByte := ve.ledger.GetCode(types.NewAddressByStr(address))
	if contractByte == nil {
		return nil, fmt.Errorf("this rule address %s does not exist", address)
	}
	contract := &wasm.Contract{}
	if err := json.Unmarshal(contractByte, contract); err != nil {
		return nil, fmt.Errorf("contract byte not correct")
	}

	return NewWasmValidator(contract.Code, ve.logger, ve.wasmGasLimit), nil
}
