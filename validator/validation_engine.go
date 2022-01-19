package validator

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/meshplus/bitxhub-core/wasm"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/sirupsen/logrus"
	"github.com/wasmerio/wasmer-go/wasmer"
)

const (
	FabricRuleAddr    = "0x00000000000000000000000000000000000000a0"
	SimFabricRuleAddr = "0x00000000000000000000000000000000000000a1"
	HappyRuleAddr     = "0x00000000000000000000000000000000000000a2"
)

// Validator is the instance that can use wasm to verify transaction validity
type ValidationEngine struct {
	pools           *ValidatorPools
	wasmStore       *wasmer.Store
	modules         *WasmModuleMap
	fabValidator    Validator
	simFabValidator Validator
	happyValidator  Validator
	wasmGasLimit    uint64
	ledger          Ledger
	logger          logrus.FieldLogger

	sync.RWMutex
}

type WasmModuleMap struct {
	modules map[string]*wasmer.Module
	sync.RWMutex
}

// New a validator instance
func NewValidationEngine(ledger Ledger, instances *sync.Map, logger logrus.FieldLogger, gasLimit uint64) *ValidationEngine {
	wasmModuleMap := &WasmModuleMap{
		modules: make(map[string]*wasmer.Module),
	}
	return &ValidationEngine{
		ledger:          ledger,
		logger:          logger,
		wasmStore:       wasm.NewStore(),
		modules:         wasmModuleMap,
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
		ve.pools.SetPool(address, NewValidationPool(100))
		vlt, err := ve.getValidator(address)
		if err != nil {
			return false, 0, err
		}
		err = ve.pools.pools[address].Add(vlt)
		if err != nil {
			return false, 0, err
		}
		pool, _ = ve.pools.GetPool(address)
	} else {
		ve.Lock()
		defer ve.Unlock()
		if pool.length < pool.size {
			vlt, err := ve.getValidator(address)
			if err != nil {
				return false, 0, err
			}
			err = pool.Add(vlt)
			if err != nil {
				return false, 0, err
			}
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

	if address == HappyRuleAddr {
		return &HappyValidator{}, nil
	}

	var module *wasmer.Module
	var ok bool
	module, ok = ve.modules.GetModule(address)
	if !ok {
		ve.modules.Lock()
		defer ve.modules.Unlock()

		contractByte := ve.ledger.GetCode(types.NewAddressByStr(address))
		if contractByte == nil {
			return nil, fmt.Errorf("this rule address %s does not exist", address)
		}
		contract := &wasm.Contract{}
		if err := json.Unmarshal(contractByte, contract); err != nil {
			return nil, fmt.Errorf("contract byte not correct")
		}
		var err error
		module, err = ve.modules.SetModule(contract.Code, address, ve.wasmStore)
		if err != nil {
			return nil, err
		}
	}

	fmt.Println("new wasm validator")
	return NewWasmValidator(module, ve.wasmStore, ve.logger, ve.wasmGasLimit), nil
}

func (wm *WasmModuleMap) SetModule(code []byte, address string, store *wasmer.Store) (*wasmer.Module, error) {
	module, err := wasm.NewModule(code, store)
	wm.modules[address] = module
	return module, err
}

func (wm *WasmModuleMap) GetModule(address string) (*wasmer.Module, bool) {
	wm.Lock()
	defer wm.Unlock()

	module, ok := wm.modules[address]
	return module, ok
}
