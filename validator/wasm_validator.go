package validator

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/meshplus/bitxhub-core/wasm"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type WasmValidator struct {
	wasm      *wasm.Wasm
	input     []byte
	ledger    Ledger
	logger    logrus.FieldLogger
	instances *sync.Map
	gasLimit  uint64
}

// New a validator instance
func NewWasmValidator(ledger Ledger, logger logrus.FieldLogger, instances *sync.Map, gasLimit uint64) *WasmValidator {
	return &WasmValidator{
		ledger:    ledger,
		logger:    logger,
		instances: instances,
		gasLimit:  gasLimit,
	}
}

// Verify will check whether the transaction info is valid
func (vlt *WasmValidator) Verify(address, from string, proof, payload []byte, validators string) (bool, uint64, error) {
	ruleHash, err := vlt.initRule(address, from, proof, payload, validators)
	if err != nil {
		return false, 0, err
	}

	ret, gasUsed, err := vlt.wasm.Execute(vlt.input, vlt.gasLimit)
	if err != nil {
		return false, 0, err
	}
	// put wasm instance into pool
	v, ok := vlt.instances.Load(ruleHash)
	if !ok {
		return false, 0, fmt.Errorf("load wasm instance failed")
	}
	v.(*sync.Pool).Put(vlt.wasm.Instance)

	// check execution status
	result, err := strconv.Atoi(string(ret))
	if err != nil {
		return false, 0, err
	}

	if result == 0 {
		return false, 0, nil
	}

	return true, gasUsed, nil
}

// InitRule can import a specific rule for validator to verify the transaction
func (vlt *WasmValidator) initRule(address, from string, proof, payload []byte, validators string) (string, error) {
	err := vlt.setTransaction(address, from, proof, validators, payload)
	if err != nil {
		return "", err
	}

	imports := validatorlib.New()
	contractByte := vlt.ledger.GetCode(types.NewAddressByStr(address))

	if contractByte == nil {
		return "", fmt.Errorf("this rule address does not exist")
	}

	wasmInstance, err := wasm.New(contractByte, imports, vlt.instances)
	if err != nil {
		return "", err
	}
	vlt.wasm = wasmInstance

	contract := &wasm.Contract{}
	if err := json.Unmarshal(contractByte, contract); err != nil {
		return "", fmt.Errorf("contract byte not correct")
	}
	return contract.Hash.String(), nil
}

func (vlt *WasmValidator) setTransaction(address, from string, proof []byte, validators string, payload []byte) error {
	invokePayload := &pb.InvokePayload{
		Method: "start_verify",
		Args: []*pb.Arg{
			{Type: pb.Arg_Bytes, Value: proof},
			{Type: pb.Arg_Bytes, Value: []byte(validators)},
			{Type: pb.Arg_Bytes, Value: payload},
		},
	}
	input, _ := proto.Marshal(invokePayload)

	vlt.input = input

	return nil
}
