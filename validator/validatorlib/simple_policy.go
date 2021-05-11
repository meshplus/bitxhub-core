package validatorlib

import (
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func fabric_validate_v13(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	return []wasmer.Value{wasmer.NewI32(1)}, nil
}

func (im *Imports) importFabricV13(store *wasmer.Store, wasmEnv *wasmlib.WasmEnv) {
	function := wasmer.NewFunctionWithEnvironment(
		store,
		wasmer.NewFunctionType(
			wasmer.NewValueTypes(wasmer.I64, wasmer.I64),
			wasmer.NewValueTypes(wasmer.I32),
		),
		wasmEnv,
		fabric_validate_v13,
	)
	im.imports.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"fabric_validate_v13": function,
		},
	)
}
