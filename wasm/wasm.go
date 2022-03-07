package wasm

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/bytecodealliance/wasmtime-go"
	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
)

const (
	CONTEXT_ARGMAP    = "argmap"
	CONTEXT_INTERFACE = "interface"
	ERROR             = "error"
)

var (
	errorLackOfMethod = fmt.Errorf("wasm execute: lack of method name")
	errorNoSuchMethod = fmt.Errorf("wasm execute: no such method")
)

// Wasm represents the wasm vm in BitXHub
type Wasm struct {
	// wasm instance
	Instance *wasmtime.Instance
	Store    *wasmtime.Store

	context map[string]interface{}

	sync.RWMutex
}

// Contract represents the smart contract structure used in the wasm vm
type Contract struct {
	// contract byte
	Code []byte `json:"code"`

	// contract hash
	Hash *types.Hash `json:"hash"`
}

func NewWithStore(code []byte, context map[string]interface{}, libs []*wasmlib.ImportLib, store *wasmtime.Store) (*Wasm, error) {
	wasm := &Wasm{}
	linker := wasmtime.NewLinker(store.Engine)
	for _, lib := range libs {
		err := linker.DefineFunc(store, lib.Module, lib.Name, lib.Func)
		if err != nil {
			return nil, err
		}
	}
	for _, lib := range wasmlib.NewWasmLibs(context, store) {
		err := linker.DefineFunc(store, lib.Module, lib.Name, lib.Func)
		if err != nil {
			return nil, err
		}
	}
	module, err := wasmtime.NewModule(store.Engine, code)
	if err != nil {
		return nil, err
	}
	instance, err := linker.Instantiate(store, module)
	if err != nil {
		return nil, err
	}
	wasm.Instance = instance
	wasm.Store = store
	wasm.context = context

	return wasm, nil
}

func NewStore() *wasmtime.Store {
	cfg := wasmtime.NewConfig()
	cfg.SetWasmReferenceTypes(true)
	cfg.SetConsumeFuel(true)
	return wasmtime.NewStore(wasmtime.NewEngineWithConfig(cfg))
}

// New creates a wasm vm instance
func New(code []byte, context map[string]interface{}, libs []*wasmlib.ImportLib) (*Wasm, error) {
	wasm := &Wasm{}

	cfg := wasmtime.NewConfig()
	cfg.SetWasmReferenceTypes(true)
	cfg.SetConsumeFuel(true)
	store := wasmtime.NewStore(wasmtime.NewEngineWithConfig(cfg))
	linker := wasmtime.NewLinker(store.Engine)
	for _, lib := range libs {
		err := linker.DefineFunc(store, lib.Module, lib.Name, lib.Func)
		if err != nil {
			return nil, err
		}
	}
	err := linker.DefineFunc(
		store,
		"env",
		"set_data",
		func(caller *wasmtime.Caller, key_ptr int64, key_len int64, value_ptr int64, value_len int64) {
			mem := caller.GetExport("memory").Memory()
			buf := mem.UnsafeData(store)
			key := buf[key_ptr : key_ptr+key_len]
			value := buf[value_ptr : value_ptr+value_len]
			context[string(key)] = value
		})
	if err != nil {
		return nil, err
	}
	module, err := wasmtime.NewModule(store.Engine, code)
	if err != nil {
		return nil, err
	}
	instance, err := linker.Instantiate(store, module)
	if err != nil {
		return nil, err
	}
	wasm.Instance = instance
	wasm.Store = store
	wasm.context = context

	return wasm, nil
}

func (w *Wasm) Execute(input []byte, wasmGasLimit uint64) (ret []byte, gasUsed uint64, err error) {
	w.SetContext("result", []byte(""))
	w.SetContext(CONTEXT_ARGMAP, make(map[int32]int32))
	if err := w.Store.AddFuel(wasmGasLimit); err != nil {
		return nil, 0, err
	}

	payload := &pb.InvokePayload{}
	if err := proto.Unmarshal(input, payload); err != nil {
		return nil, 0, err
	}

	if payload.Method == "" {
		return nil, 0, errorLackOfMethod
	}

	methodName := w.Instance.GetFunc(w.Store, payload.Method)
	if methodName == nil {
		return nil, 0, errorNoSuchMethod
	}
	slice := make([]interface{}, len(payload.Args))
	for i := range slice {
		arg := payload.Args[i]
		var temp interface{}
		var err error
		switch arg.Type {
		case pb.Arg_I32:
			temp, err = strconv.Atoi(string(arg.Value))
		case pb.Arg_I64:
			temp, err = strconv.ParseInt(string(arg.Value), 10, 64)
		case pb.Arg_F32:
			temp, err = strconv.ParseFloat(string(arg.Value), 32)
		case pb.Arg_F64:
			temp, err = strconv.ParseFloat(string(arg.Value), 64)
		case pb.Arg_String:
			temp, err = w.SetString(string(arg.Value))
		case pb.Arg_Bytes:
			temp, err = w.SetBytes(arg.Value)
		case pb.Arg_Bool:
			temp, err = strconv.Atoi(string(arg.Value))
		default:
			err = fmt.Errorf("input type not support")
		}
		if err != nil {
			gasUsed, _ := w.Store.FuelConsumed()
			return nil, gasUsed, err
		}
		slice[i] = temp
	}

	// w.Instance.SetContextData(w.context)

	result, err := methodName.Call(w.Store, slice...)
	if err != nil {
		ret = nil
	} else {
		ret = []byte(strconv.Itoa(int(result.(int32))))
	}
	if string(w.GetContext("result").([]byte)) != "" {
		ret = w.GetContext("result").([]byte)
	}

	for i := range slice {
		arg := payload.Args[i]
		switch arg.Type {
		case pb.Arg_String:
			if err1 := w.FreeString(slice[i], string(arg.Value)); err1 != nil {
				err = err1
			}
		case pb.Arg_Bytes:
			if err1 := w.FreeBytes(slice[i], arg.Value); err1 != nil {
				err = err1
			}
		}
	}
	w.context = make(map[string]interface{})

	gasUsed, _ = w.Store.FuelConsumed()
	return ret, gasUsed, err
}

func (w *Wasm) SetContext(key string, value interface{}) {
	w.Lock()
	defer w.Unlock()

	w.context[key] = value
}

func (w *Wasm) GetContext(key string) interface{} {
	w.Lock()
	defer w.Unlock()

	return w.context[key]
}
