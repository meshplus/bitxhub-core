package validator

import (
	"strconv"

	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/meshplus/bitxhub-core/wasm"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
	"github.com/wasmerio/wasmer-go/wasmer"
)

// Validator is the instance that can use wasm to verify transaction validity
type WasmValidator struct {
	wasm     *wasm.Wasm
	input    []byte
	ledger   Ledger
	logger   logrus.FieldLogger
	gasLimit uint64
}

// New a validator instance
func NewWasmValidator(module *wasmer.Module, store *wasmer.Store, logger logrus.FieldLogger, gasLimit uint64) *WasmValidator {
	imports := validatorlib.New()
	wasmInstance, err := wasm.New(imports, module, store)
	if err != nil {
		return nil
	}
	return &WasmValidator{
		wasm:     wasmInstance,
		logger:   logger,
		gasLimit: gasLimit,
	}
}

// Verify will check whether the transaction info is valid
func (vlt *WasmValidator) Verify(from string, proof, payload []byte, validators string) (bool, uint64, error) {
	err := vlt.initRule(from, proof, payload, validators)
	if err != nil {
		return false, 0, err
	}

	ret, gasUsed, err := vlt.wasm.Execute(vlt.input, vlt.gasLimit)
	if err != nil {
		return false, gasUsed, err
	}
	// put wasm instance into pool

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
func (vlt *WasmValidator) initRule(from string, proof, payload []byte, validators string) error {
	err := vlt.setTransaction(from, proof, validators, payload)
	if err != nil {
		return err
	}

	return nil
}

func (vlt *WasmValidator) setTransaction(from string, proof []byte, validators string, payload []byte) error {
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
