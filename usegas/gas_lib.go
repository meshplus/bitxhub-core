package usegas

import (
	"fmt"
	"sync"

	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type GasLimit struct {
	limit uint64

	sync.RWMutex
}

func (g *GasLimit) GetLimit() uint64 {
	g.RLock()
	defer g.RUnlock()

	return g.limit
}

func (g *GasLimit) SetLimit(limit uint64) {
	g.Lock()
	defer g.Unlock()

	g.limit = limit
}

func usegas(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	gasPrice := uint64(args[0].I64())
	gasLimit := env.(*wasmlib.WasmEnv).Ctx["gaslimit"].(*GasLimit)

	gasL := gasLimit.GetLimit()
	if gasL < gasPrice {
		fmt.Println(fmt.Sprintf("!!!! insufficient remaining gas(%d) !!!", gasL))
		gasLimit.SetLimit(uint64(0))
		return []wasmer.Value{}, fmt.Errorf("run out of gas limit")
	}
	remain := gasL - gasPrice

	gasLimit.SetLimit(remain)
	return []wasmer.Value{}, nil
}

func setResult(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	value_ptr := args[0].I64()
	value_len := args[1].I64()
	ctx := env.(*wasmlib.WasmEnv).Ctx
	mem, err := env.(*wasmlib.WasmEnv).Instance.Exports.GetMemory("memory")
	if err != nil {
		return []wasmer.Value{}, err
	}
	value := mem.Data()[value_ptr : value_ptr+value_len]
	ctx["result"] = value
	return []wasmer.Value{}, nil
}

func (im *Imports) importGasLib(store *wasmer.Store, wasmEnv *wasmlib.WasmEnv) {
	useGasFunc := wasmer.NewFunctionWithEnvironment(
		store,
		wasmer.NewFunctionType(
			wasmer.NewValueTypes(wasmer.I64),
			wasmer.NewValueTypes(),
		),
		wasmEnv,
		usegas,
	)
	setResultFunc := wasmer.NewFunctionWithEnvironment(
		store,
		wasmer.NewFunctionType(
			wasmer.NewValueTypes(wasmer.I64, wasmer.I64),
			wasmer.NewValueTypes(),
		),
		wasmEnv,
		setResult,
	)
	im.imports.Register(
		"metering",
		map[string]wasmer.IntoExtern{
			"usegas": useGasFunc,
		},
	)
	im.imports.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"set_result": setResultFunc,
		},
	)
}
