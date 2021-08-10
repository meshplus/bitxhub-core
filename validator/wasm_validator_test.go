package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
	"testing"

	"github.com/meshplus/bitxhub-core/validator/validatorlib"
	"github.com/meshplus/bitxhub-core/wasm"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/stretchr/testify/require"
)

func TestWasmValidator(t *testing.T) {
	wasmBytes, err := ioutil.ReadFile("./testdata/hpc_demo.wasm")
	require.Nil(t, err)

	proof, err := ioutil.ReadFile("./testdata/proof_1.0.0_rc")
	require.Nil(t, err)

	validators, err := ioutil.ReadFile("./testdata/validator_1.0.0_rc")
	require.Nil(t, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "mychannel&transfer",
		Func:          "interchainCharge",
		Args:          [][]byte{[]byte("Alice"), []byte("Alice"), []byte("1")},
		Callback:      "interchainConfirm",
	}

	bytes, err := content.Marshal()
	require.Nil(t, err)

	payload := &pb.Payload{
		Encrypted: false,
		Content:   bytes,
	}

	body, err := payload.Marshal()
	require.Nil(t, err)

	validator := &WasmValidator{
		instances: &sync.Map{},
	}

	wasmCode := &wasm.Contract{
		Code: wasmBytes,
	}
	contractBytes, err := json.Marshal(wasmCode)
	require.Nil(t, err)
	imports := validatorlib.New()
	require.Nil(t, err)
	wasm, err := wasm.New(contractBytes, imports, validator.instances, wasmGasLimit)
	require.Nil(t, err)
	validator.wasm = wasm
	err = validator.setTransaction("", "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", proof, string(validators), body)
	require.Nil(t, err)
	ret, err := validator.wasm.Execute(validator.input)
	require.Nil(t, err)
	result, err := strconv.Atoi(string(ret))
	require.Nil(t, err)
	fmt.Println(result)
}

func BenchmarkHpcWasm_Verify(b *testing.B) {
	wasmBytes, err := ioutil.ReadFile("./testdata/hpc_demo.wasm")
	require.Nil(b, err)

	// proof, err := ioutil.ReadFile("./testdata/proof_1.0.0_rc")
	// require.Nil(b, err)

	// validators, err := ioutil.ReadFile("./testdata/validator_1.0.0_rc")
	// require.Nil(b, err)

	content := &pb.Content{
		SrcContractId: "mychannel&transfer",
		DstContractId: "mychannel&transfer",
		Func:          "interchainCharge",
		Args:          [][]byte{[]byte("Alice"), []byte("Alice"), []byte("1")},
		Callback:      "interchainConfirm",
	}

	bytes, err := content.Marshal()
	require.Nil(b, err)

	payload := &pb.Payload{
		Encrypted: false,
		Content:   bytes,
	}

	body, err := payload.Marshal()
	require.Nil(b, err)

	validator := &WasmValidator{
		instances: &sync.Map{},
	}

	wasmCode := &wasm.Contract{
		Code: wasmBytes,
	}
	contractBytes, err := json.Marshal(wasmCode)
	require.Nil(b, err)
	imports := validatorlib.New()
	require.Nil(b, err)
	wasm, err := wasm.New(contractBytes, imports, validator.instances, wasmGasLimit)
	require.Nil(b, err)
	validator.wasm = wasm
	err = validator.setTransaction("", "0xe02d8fdacd59020d7f292ab3278d13674f5c404d", []byte("111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"), "111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111", body)
	require.Nil(b, err)
	// for i := 0; i < 400000; i++ {
	// 	_, err := validator.wasm.Execute(validator.input)
	// 	require.Nil(b, err)
	// }
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := validator.wasm.Execute(validator.input)
		require.Nil(b, err)
	}
}
