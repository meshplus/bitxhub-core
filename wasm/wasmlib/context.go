package wasmlib

import (
	"github.com/bytecodealliance/wasmtime-go"
)

func set_data(context map[string]interface{}, store *wasmtime.Store) *ImportLib {
	return &ImportLib{
		Module: "env",
		Name:   "set_data",
		Func: func(caller *wasmtime.Caller, value_ptr int32, value_len int32) {
			context[CONTEXT_ARGMAP].(map[int32]int32)[value_ptr] = value_len
		},
	}
}

func set_result(context map[string]interface{}, store *wasmtime.Store) *ImportLib {
	return &ImportLib{
		Module: "env",
		Name:   "set_result",
		Func: func(caller *wasmtime.Caller, result_ptr int32, result_len int32) {
			mem := caller.GetExport("memory").Memory()
			buf := mem.UnsafeData(store)
			result := buf[result_ptr : result_ptr+result_len]
			context[RESULT] = result
		},
	}
}

func ImportWasmLib(context map[string]interface{}, store *wasmtime.Store) []*ImportLib {
	var libs []*ImportLib
	libs = append(libs, set_data(context, store))
	libs = append(libs, set_result(context, store))

	return libs
}
