package wasmlib

import (
	"github.com/wasmerio/wasmer-go/wasmer"
)

type Imports struct {
	imports *wasmer.ImportObject
}

func New() WasmImport {
	imports := &Imports{
		imports: wasmer.NewImportObject(),
	}
	return imports
}

func (imports *Imports) ImportLib(wasmEnv *WasmEnv) {
	imports.importWasmLib(wasmEnv.Store, wasmEnv)
}

func (imports *Imports) GetImportObject() *wasmer.ImportObject {
	return imports.imports
}
