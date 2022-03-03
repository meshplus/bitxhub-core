package wasm

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/stretchr/testify/assert"
)

const wasmGasLimit = 5000000000000000

func TestExecute(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/wasm_test.wasm")
	assert.Nil(t, err)

	libs := []*wasmlib.ImportLib{}
	context := make(map[string]interface{})
	wasm, err := New(data, context, libs)
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
	ret, _, err := wasm.Execute(inputBytes, wasmGasLimit)
	assert.Nil(t, err)
	fmt.Println(string(ret))
	hash := types.NewHashByStr("")
	fmt.Println(hash)
}

func TestImportExecute(t *testing.T) {
	data, err := ioutil.ReadFile("./testdata/test_demo.wasm")
	assert.Nil(t, err)
	hello := "hello world"

	libs := []*wasmlib.ImportLib{}
	context := make(map[string]interface{})
	wasm, err := New(data, context, libs)
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
	ret, _, err := wasm.Execute(inputBytes, wasmGasLimit)
	assert.Nil(t, err)
	_, _, err = wasm.Execute(inputBytes, wasmGasLimit)
	assert.Nil(t, err)
	_, _, err = wasm.Execute(inputBytes, wasmGasLimit)
	assert.Nil(t, err)
	_, _, err = wasm.Execute(inputBytes, wasmGasLimit)
	assert.Nil(t, err)
	_, _, err = wasm.Execute(inputBytes, wasmGasLimit)
	assert.Nil(t, err)
	_, _, err = wasm.Execute(inputBytes, wasmGasLimit)
	assert.Nil(t, err)
	fmt.Println(string(ret))
	hash := types.NewHashByStr("")
	fmt.Println(hash)
}

func BenchmarkImportExecute(b *testing.B) {
	data, err := ioutil.ReadFile("./testdata/test_demo.wasm")
	assert.Nil(b, err)
	hello := "hello world"

	libs := []*wasmlib.ImportLib{}
	context := make(map[string]interface{})
	wasm, err := New(data, context, libs)
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
		_, _, err := wasm.Execute(inputBytes, wasmGasLimit)
		assert.Nil(b, err)
		// fmt.Println(string(ret))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, err := wasm.Execute(inputBytes, wasmGasLimit)
		assert.Nil(b, err)
		// fmt.Println(string(ret))
	}
}
