package validatorlib

import (
	"github.com/meshplus/bitxhub-core/usegas"
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Imports struct {
	imports *usegas.Imports
}

func New() wasmlib.WasmImport {
	imports := &Imports{
		imports: usegas.New(),
	}
	return imports
}

func (imports *Imports) ImportLib(wasmEnv *wasmlib.WasmEnv) {
	imports.imports.ImportLib(wasmEnv)
	imports.importECDSA(wasmEnv.Store, wasmEnv)
	imports.importFabricV14(wasmEnv.Store, wasmEnv)
	imports.importFabricV13(wasmEnv.Store, wasmEnv)
}

func (imports *Imports) GetImportObject() *wasmer.ImportObject {
	return imports.imports.GetImportObject()
}
