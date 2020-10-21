package wasm

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-kit/wasm/wasmlib"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/stretchr/testify/assert"
	"github.com/wasmerio/go-ext-wasm/wasmer"
)

func TestExecute(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/wasm_test.wasm")
	assert.Nil(t, err)

	contract := &Contract{
		Code: data,
		Hash: types.Hash{},
	}

	bytes, err := json.Marshal(contract)
	assert.Nil(t, err)
	imports := wasmer.NewImports()
	instances := make(map[string]wasmer.Instance)
	wasm, err := New(bytes, imports, instances)
	assert.Nil(t, err)
	input := &pb.InvokePayload{
		Method: "a",
		Args: []*pb.Arg{
			{Type: pb.Arg_I32, Value: []byte(fmt.Sprintf("%d", 1))},
			{Type: pb.Arg_I32, Value: []byte(fmt.Sprintf("%d", 2))},
		},
	}
	inputBytes, err := input.Marshal()
	assert.Nil(t, err)
	ret, err := wasm.Execute(inputBytes)
	assert.Nil(t, err)
	fmt.Println(string(ret))
}

func TestImportExecute(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/test_demo.wasm")
	assert.Nil(t, err)
	hello := "hello world"

	contract := &Contract{
		Code: data,
		Hash: types.Hash{},
	}

	bytes, err := json.Marshal(contract)
	assert.Nil(t, err)
	imports, err := wasmlib.New()
	assert.Nil(t, err)
	instances := make(map[string]wasmer.Instance)
	wasm, err := New(bytes, imports, instances)
	assert.Nil(t, err)
	input := &pb.InvokePayload{
		Method: "start_verify",
		Args: []*pb.Arg{
			{Type: pb.Arg_Bytes, Value: []byte(fmt.Sprintf("%d", 1))},
			{Type: pb.Arg_Bytes, Value: []byte(fmt.Sprintf("%d", 2))},
			{Type: pb.Arg_Bytes, Value: []byte(fmt.Sprintf("%d", 2))},
		},
	}
	inputBytes, err := input.Marshal()
	assert.Nil(t, err)
	wasm.SetContext("hello", hello)
	ret, err := wasm.Execute(inputBytes)
	assert.Nil(t, err)
	fmt.Println(string(ret))
}

func BenchmarkImportExecute(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/test_demo.wasm")
	assert.Nil(b, err)
	hello := "hello world"

	contract := &Contract{
		Code: data,
		Hash: types.Hash{},
	}

	bytes, err := json.Marshal(contract)
	assert.Nil(b, err)
	imports, err := wasmlib.New()
	assert.Nil(b, err)
	instances := make(map[string]wasmer.Instance)
	wasm, err := New(bytes, imports, instances)
	assert.Nil(b, err)
	input := &pb.InvokePayload{
		Method: "start_verify",
		Args: []*pb.Arg{
			{Type: pb.Arg_Bytes, Value: []byte("1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")},
			{Type: pb.Arg_Bytes, Value: []byte("1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111")},
			{Type: pb.Arg_Bytes, Value: []byte(fmt.Sprintf("%d", 2))},
		},
	}
	inputBytes, err := input.Marshal()
	assert.Nil(b, err)
	wasm.SetContext("hello", hello)
	for i := 0; i < 200000; i++ {
		_, err := wasm.Execute(inputBytes)
		assert.Nil(b, err)
		// fmt.Println(string(ret))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := wasm.Execute(inputBytes)
		assert.Nil(b, err)
		// fmt.Println(string(ret))
	}
}
