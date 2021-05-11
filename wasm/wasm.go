package wasm

import (
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/gogo/protobuf/proto"
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/meshplus/bitxhub-model/pb"
	"github.com/wasmerio/wasmer-go/wasmer"
)

const (
	CONTEXT_ARGMAP    = "argmap"
	CONTEXT_INTERFACE = "interface"

	ACCOUNT   = "account"
	ALLOC_MEM = "allocate"
)

var (
	errorLackOfMethod = fmt.Errorf("wasm execute: lack of method name")
)

// Wasm represents the wasm vm in BitXHub
type Wasm struct {
	// wasm instance
	Instance *wasmer.Instance

	env     *wasmlib.WasmEnv
	context map[string]interface{}
	argMap  map[int]int

	sync.RWMutex
}

// Contract represents the smart contract structure used in the wasm vm
type Contract struct {
	// contract byte
	Code []byte

	// contract hash
	Hash *types.Hash
}

func getInstance(contract *Contract, imports wasmlib.WasmImport, env *wasmlib.WasmEnv, instances *sync.Map) (*wasmer.Instance, error) {
	var (
		instance *wasmer.Instance
		pool     *sync.Pool
	)
	v, ok := instances.Load(contract.Hash.String())
	if !ok {
		v = &sync.Pool{
			New: func() interface{} {
				engine := wasmer.NewEngine()
				store := wasmer.NewStore(engine)
				module, err := wasmer.NewModule(store, contract.Code)
				if err != nil {
					return nil
				}
				env.Store = store
				imports.ImportLib(env)
				instance, err = wasmer.NewInstance(module, imports.GetImportObject())
				env.Instance = instance
				if err != nil {
					return nil
				}
				return instance
			},
		}
		instances.Store(contract.Hash.String(), v)
	}

	pool = v.(*sync.Pool)
	rawInstance := pool.Get()
	if rawInstance == nil {
		engine := wasmer.NewEngine()
		store := wasmer.NewStore(engine)
		module, err := wasmer.NewModule(store, contract.Code)
		if err != nil {
			return &wasmer.Instance{}, err
		}
		env.Store = store
		imports.ImportLib(env)
		instance, err = wasmer.NewInstance(module, imports.GetImportObject())
		if err != nil {
			return &wasmer.Instance{}, err
		}
		env.Instance = instance
	} else {
		instance = rawInstance.(*wasmer.Instance)
	}

	return instance, nil
}

// New creates a wasm vm instance
func New(contractByte []byte, imports wasmlib.WasmImport, instances *sync.Map) (*Wasm, error) {
	wasm := &Wasm{}

	contract := &Contract{}
	if err := json.Unmarshal(contractByte, contract); err != nil {
		return wasm, fmt.Errorf("contract byte not correct")
	}

	if len(contract.Code) == 0 {
		return wasm, fmt.Errorf("contract byte is empty")
	}

	env := &wasmlib.WasmEnv{}
	instance, err := getInstance(contract, imports, env, instances)
	if err != nil {
		return nil, err
	}

	wasm.Instance = instance
	wasm.argMap = make(map[int]int)
	wasm.context = make(map[string]interface{})
	env.Ctx = make(map[string]interface{})
	wasm.env = env

	return wasm, nil
}

func EmptyImports() (*wasmer.ImportObject, error) {
	return wasmer.NewImportObject(), nil
}

func (w *Wasm) Execute(input []byte) ([]byte, error) {
	payload := &pb.InvokePayload{}
	if err := proto.Unmarshal(input, payload); err != nil {
		return nil, err
	}

	if payload.Method == "" {
		return nil, errorLackOfMethod
	}

	// alloc, err := w.Instance.Exports.GetFunction(ALLOC_MEM)
	// if err != nil {
	// 	return nil, err
	// }
	// if alloc == nil {
	// 	return nil, fmt.Errorf("not found allocate method")
	// }
	// w.context[ALLOC_MEM] = alloc
	methodName, err := w.Instance.Exports.GetFunction(payload.Method)
	if err != nil {
		return nil, err
	}
	slice := make([]interface{}, len(payload.Args))
	for i := range slice {
		arg := payload.Args[i]
		switch arg.Type {
		case pb.Arg_I32:
			temp, err := strconv.Atoi(string(arg.Value))
			if err != nil {
				return nil, err
			}
			slice[i] = temp
		case pb.Arg_I64:
			temp, err := strconv.ParseInt(string(arg.Value), 10, 64)
			if err != nil {
				return nil, err
			}
			slice[i] = temp
		case pb.Arg_F32:
			temp, err := strconv.ParseFloat(string(arg.Value), 32)
			if err != nil {
				return nil, err
			}
			slice[i] = temp
		case pb.Arg_F64:
			temp, err := strconv.ParseFloat(string(arg.Value), 64)
			if err != nil {
				return nil, err
			}
			slice[i] = temp
		case pb.Arg_String:
			inputPointer, err := w.SetString(string(arg.Value))
			if err != nil {
				return nil, err
			}
			slice[i] = inputPointer
		case pb.Arg_Bytes:
			inputPointer, err := w.SetBytes(arg.Value)
			if err != nil {
				return nil, err
			}
			slice[i] = inputPointer
		case pb.Arg_Bool:
			inputPointer, err := strconv.Atoi(string(arg.Value))
			if err != nil {
				return nil, err
			}
			slice[i] = inputPointer
		default:
			return nil, fmt.Errorf("input type not support")
		}
	}

	w.env.Ctx[CONTEXT_ARGMAP] = w.argMap
	// w.Instance.SetContextData(w.context)

	result, err := methodName(slice...)
	if err != nil {
		return nil, err
	}
	for i := range slice {
		arg := payload.Args[i]
		switch arg.Type {
		case pb.Arg_String:
			if err := w.FreeString(slice[i], string(arg.Value)); err != nil {
				return nil, err
			}
		case pb.Arg_Bytes:
			if err := w.FreeBytes(slice[i], arg.Value); err != nil {
				return nil, err
			}
		}
	}

	return []byte(strconv.Itoa(int(result.(int32)))), err
}

func (w *Wasm) SetContext(key string, value interface{}) {
	w.Lock()
	defer w.Unlock()

	w.env.Ctx[key] = value
}

func (w *Wasm) GetContext(key string) interface{} {
	w.Lock()
	defer w.Unlock()

	return w.env.Ctx[key]
}
