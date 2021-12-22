package validator

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/meshplus/bitxhub-kit/types"
	"github.com/sirupsen/logrus"
)

const (
	FabricRuleAddr    = "0x00000000000000000000000000000000000000a0"
	SimFabricRuleAddr = "0x00000000000000000000000000000000000000a1"
	HappyRuleAddr     = "0x00000000000000000000000000000000000000a2"
	Fab10RuleAddr     = "0x0000000000000000000000000000000000000a10"
	Fab20RuleAddr     = "0x0000000000000000000000000000000000000a20"
	Fab50RuleAddr     = "0x0000000000000000000000000000000000000a50"
	Fab100RuleAddr    = "0x000000000000000000000000000000000000a100"
	Fab200RuleAddr    = "0x000000000000000000000000000000000000a200"
	Fab400RuleAddr    = "0x000000000000000000000000000000000000a400"
	Fab600RuleAddr    = "0x000000000000000000000000000000000000a600"
	Fab800RuleAddr    = "0x000000000000000000000000000000000000a800"
	Fab1000RuleAddr   = "0x00000000000000000000000000000000000a1000"
	Fab1500RuleAddr   = "0x00000000000000000000000000000000000a1500"
	Fab2000RuleAddr   = "0x00000000000000000000000000000000000a2000"
)

// Validator is the instance that can use wasm to verify transaction validity
type ValidationEngine struct {
	instances        *sync.Map
	fabValidator     Validator
	simFabValidator  Validator
	happyValidator   Validator
	fab10Validator   Validator
	fab20Validator   Validator
	fab50Validator   Validator
	fab100Validator  Validator
	fab200Validator  Validator
	fab400Validator  Validator
	fab600Validator  Validator
	fab800Validator  Validator
	fab1000Validator Validator
	fab1500Validator Validator
	fab2000Validator Validator
	wasmGasLimit     uint64
	ledger           Ledger
	logger           logrus.FieldLogger
}

// New a validator instance
func NewValidationEngine(ledger Ledger, instances *sync.Map, logger logrus.FieldLogger, gasLimit uint64) *ValidationEngine {
	return &ValidationEngine{
		ledger:           ledger,
		logger:           logger,
		fabValidator:     NewFabV14Validator(logger),
		simFabValidator:  NewFabSimValidator(logger),
		happyValidator:   &HappyValidator{},
		fab10Validator:   &Fab10Validator{},
		fab20Validator:   &Fab20Validator{},
		fab50Validator:   &Fab50Validator{},
		fab100Validator:  &Fab100Validator{},
		fab200Validator:  &Fab200Validator{},
		fab400Validator:  &Fab400Validator{},
		fab600Validator:  &Fab600Validator{},
		fab800Validator:  &Fab800Validator{},
		fab1000Validator: &Fab1000Validator{},
		fab1500Validator: &Fab1500Validator{},
		fab2000Validator: &Fab2000Validator{},

		instances:    instances,
		wasmGasLimit: gasLimit,
	}
}

// Verify will check whether the transaction info is valid
func (ve *ValidationEngine) Validate(address, from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error) {
	time1 := time.Now().UnixNano()

	vlt, err := ve.getValidator(address)
	if err != nil {
		return false, 0, err
	}
	time2 := time.Now().UnixNano()
	ve.logger.WithFields(logrus.Fields{
		"height": height,
		"index":  index,
		"time":   time2 - time1,
	}).Debug("------------------ get validator end")
	ve.logger.WithFields(logrus.Fields{}).Debugf("----- get %d", time2-time1)

	ok, gasUsed, err := vlt.Verify(from, proof, payload, validators, index, height)
	time3 := time.Now().UnixNano()
	ve.logger.WithFields(logrus.Fields{
		"height": height,
		"index":  index,
		"time":   time3 - time2,
	}).Debug("------------------ verify end")
	ve.logger.WithFields(logrus.Fields{}).Debugf("----- ver %d", time3-time2)
	ve.instances = nil
	runtime.GC()
	return ok, gasUsed, err
}

func (ve *ValidationEngine) getValidator(address string) (Validator, error) {
	switch address {
	case FabricRuleAddr:
		return ve.fabValidator, nil
	case SimFabricRuleAddr:
		return ve.simFabValidator, nil
	case HappyRuleAddr:
		return ve.happyValidator, nil
	case Fab10RuleAddr:
		return ve.fab10Validator, nil
	case Fab20RuleAddr:
		return ve.fab20Validator, nil
	case Fab50RuleAddr:
		return ve.fab50Validator, nil
	case Fab100RuleAddr:
		return ve.fab100Validator, nil
	case Fab200RuleAddr:
		return ve.fab200Validator, nil
	case Fab400RuleAddr:
		return ve.fab400Validator, nil
	case Fab600RuleAddr:
		return ve.fab600Validator, nil
	case Fab800RuleAddr:
		return ve.fab800Validator, nil
	case Fab1000RuleAddr:
		return ve.fab1000Validator, nil
	case Fab1500RuleAddr:
		return ve.fab1500Validator, nil
	case Fab2000RuleAddr:
		return ve.fab2000Validator, nil
	default:
		if ve.instances == nil {
			ve.instances = &sync.Map{}
		}

		contractByte := ve.ledger.GetCode(types.NewAddressByStr(address))
		if contractByte == nil {
			return nil, fmt.Errorf("this rule address %s does not exist", address)
		}

		return NewWasmValidator(contractByte, ve.logger, ve.instances, ve.wasmGasLimit), nil
	}
}
