package validator

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-kit/wasm"
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
}

// New a validator instance
func NewWasmValidator(ledger Ledger, logger logrus.FieldLogger, instances *sync.Map) *WasmValidator {
	return &WasmValidator{
		ledger:    ledger,
		logger:    logger,
		instances: instances,
	}
}

// Verify will check whether the transaction info is valid
func (vlt *WasmValidator) Verify(address, from string, proof, payload []byte, validators string) (bool, error) {
	err := vlt.initRule(address, from, proof, payload, validators)
	if err != nil {
		return false, err
	}

	ret, err := vlt.wasm.Execute(vlt.input)
	if err != nil {
		return false, err
	}

	// put wasm instance into pool
	contractByte := vlt.ledger.GetCode(types.String2Address(address))
	if contractByte == nil {
		return false, fmt.Errorf("this rule address does not exist")
	}
	hash := sha256.Sum256(contractByte)
	v, ok := vlt.instances.Load(string(hash[:]))
	if !ok {
		return false, fmt.Errorf("free wasm instance failed")
	}
	v.(*sync.Pool).Put(vlt.wasm.Instance)

	// check execution status
	result, err := strconv.Atoi(string(ret))
	if err != nil {
		return false, err
	}

	if result == 0 {
		return false, nil
	}

	return true, nil
}

// InitRule can import a specific rule for validator to verify the transaction
func (vlt *WasmValidator) initRule(address, from string, proof, payload []byte, validators string) error {
	err := vlt.setTransaction(address, from, proof, validators, payload)
	if err != nil {
		return err
	}

	imports, err := validatorlib.New()
	if err != nil {
		return err
	}
	contractByte := vlt.ledger.GetCode(types.String2Address(address))

	if contractByte == nil {
		return fmt.Errorf("this rule address does not exist")
	}

	wasm, err := wasm.New(contractByte, imports, vlt.instances)
	if err != nil {
		return err
	}
	vlt.wasm = wasm

	return nil
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
