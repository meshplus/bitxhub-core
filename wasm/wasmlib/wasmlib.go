package wasmlib

import (
	"github.com/bytecodealliance/wasmtime-go"
)

const (
	CONTEXT_ARGMAP = "argmap"
	RESULT         = "result"
)

type ImportLib struct {
	Module string
	Name   string
	Func   interface{}
}

func NewWasmLibs(context map[string]interface{}, store *wasmtime.Store) []*ImportLib {
	return ImportWasmLib(context, store)
}
