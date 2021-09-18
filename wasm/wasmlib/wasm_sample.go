package wasmlib

import (
	"github.com/wasmerio/wasmer-go/wasmer"
)

func test_verify(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	_, err := env.(*WasmEnv).Instance.Exports.GetMemory("memory")
	if err != nil {
		return nil, err
	}
	_ = env.(*WasmEnv).Ctx["argmap"].(map[int]int)
	// proof := memory.Data()[proof_ptr : proof_ptr+int64(data[int(proof_ptr)])]
	// validator := memory.Data()[validator_ptr : validator_ptr+int64(data[int(validator_ptr)])]
	// payload := memory.Data()[payload_ptr : payload_ptr+int64(data[int(payload_ptr)])]

	// fmt.Println(proof)
	// fmt.Println(validator)
	// fmt.Println(payload)
	// fmt.Println(ctxMap["hello"].(string))

	return []wasmer.Value{wasmer.NewI32(1)}, nil
}

func (im *Imports) importWasmLib(store *wasmer.Store, wasmEnv *WasmEnv) {
	function := wasmer.NewFunctionWithEnvironment(
		store,
		wasmer.NewFunctionType(
			wasmer.NewValueTypes(wasmer.I64, wasmer.I64, wasmer.I64),
			wasmer.NewValueTypes(wasmer.I32),
		),
		wasmEnv,
		test_verify,
	)
	im.imports.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"test_verify": function,
		},
	)
}
