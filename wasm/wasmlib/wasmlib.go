package wasmlib

import "github.com/wasmerio/wasmer-go/wasmer"

type WasmEnv struct {
	Instance *wasmer.Instance
	Store    *wasmer.Store
	Module   *wasmer.Module

	Ctx map[string]interface{}
}

type WasmImport interface {
	ImportLib(wasmEnv *WasmEnv)
	GetImportObject() *wasmer.ImportObject
}
