package usegas

import (
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Imports struct {
	imports *wasmer.ImportObject
}

func New() *Imports {
	imports := &Imports{
		imports: wasmer.NewImportObject(),
	}
	return imports
}

func (imports *Imports) ImportLib(wasmEnv *wasmlib.WasmEnv) {
	imports.importGasLib(wasmEnv.Store, wasmEnv)
}

func (imports *Imports) GetImportObject() *wasmer.ImportObject {
	return imports.imports
}
