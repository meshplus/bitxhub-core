package validator

import (
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/sirupsen/logrus"
)

// Validator is the instance that can use wasm to verify transaction validity
type WasmValidator struct {
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

	return true, nil
}

// InitRule can import a specific rule for validator to verify the transaction
func (vlt *WasmValidator) initRule(address, from string, proof, payload []byte, validators string) (string, error) {
	err := vlt.setTransaction(address, from, proof, validators, payload)
	if err != nil {
		return "", err
	}

	// imports, err := validatorlib.New()
	// if err != nil {
	// 	return "", err
	// }
	// contractByte := vlt.ledger.GetCode(types.NewAddressByStr(address))
	return address, nil
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
