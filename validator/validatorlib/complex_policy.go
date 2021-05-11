package validatorlib

import (
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/wasmerio/wasmer-go/wasmer"
)

func fabric_validate_v14(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	ctx := env.(*wasmlib.WasmEnv).Ctx
	memory, err := env.(*wasmlib.WasmEnv).Instance.Exports.GetMemory("memory")
	if err != nil {
		return nil, err
	}
	proof_ptr := args[0].I64()
	validator_ptr := args[1].I64()
	payload_ptr := args[2].I64()
	data := ctx["argmap"].(map[int]int)
	proof := memory.Data()[proof_ptr : proof_ptr+int64(data[int(proof_ptr)])]
	validator := memory.Data()[validator_ptr : validator_ptr+int64(data[int(validator_ptr)])]
	payload := memory.Data()[payload_ptr : payload_ptr+int64(data[int(payload_ptr)])]
	vInfo, err := UnmarshalValidatorInfo(validator)
	if err != nil {
		return []wasmer.Value{wasmer.NewI32(0)}, nil
	}
	artifact, err := extractValidationArtifacts(proof)
	if err != nil {
		return []wasmer.Value{wasmer.NewI32(0)}, nil
	}

	if err := ValidateChainCodeID(artifact.prp, vInfo.Cid); err != nil {
		return []wasmer.Value{wasmer.NewI32(0)}, nil
	}

	if err := ValidatePayload(artifact.payload, payload); err != nil {
		return []wasmer.Value{wasmer.NewI32(0)}, nil
	}

	signatureSet := GetSignatureSet(artifact)

	pe, ok := ctx[FABRIC_EVALUATOR].(*PolicyEvaluator)
	if !ok {
		pe, err = NewPolicyEvaluator(vInfo.ConfByte)
		if err != nil {
			return []wasmer.Value{wasmer.NewI32(0)}, nil
		}
		env.(*wasmlib.WasmEnv).Ctx[FABRIC_EVALUATOR] = pe
	}
	if err = pe.Evaluate([]byte(vInfo.Policy), signatureSet); err != nil {
		return []wasmer.Value{wasmer.NewI32(0)}, nil
	}

	return []wasmer.Value{wasmer.NewI32(1)}, nil
}

func (im *Imports) importFabricV14(store *wasmer.Store, wasmEnv *wasmlib.WasmEnv) {
	function := wasmer.NewFunctionWithEnvironment(
		store,
		wasmer.NewFunctionType(
			wasmer.NewValueTypes(wasmer.I64, wasmer.I64, wasmer.I64),
			wasmer.NewValueTypes(wasmer.I32),
		),
		wasmEnv,
		fabric_validate_v14,
	)
	im.imports.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"fabric_validate_v14": function,
		},
	)
}
