package validatorlib

// #include <stdlib.h>
//
// extern int32_t fabric_validate_v14(void *context, long long proof_ptr, long long validator_ptr, long long payload_ptr);
import "C"
import (
	"unsafe"

	"github.com/meshplus/bitxhub-kit/wasm"
	"github.com/wasmerio/go-ext-wasm/wasmer"
)

//export fabric_validate_v14
func fabric_validate_v14(context unsafe.Pointer, proof_ptr int64, validator_ptr int64, payload_ptr int64) int32 {
	ctx := wasmer.IntoInstanceContext(context)
	ctxMap := ctx.Data().(map[string]interface{})
	data := ctxMap[wasm.CONTEXT_ARGMAP].(map[int]int)
	memory := ctx.Memory()
	proof := memory.Data()[proof_ptr : proof_ptr+int64(data[int(proof_ptr)])]
	validator := memory.Data()[validator_ptr : validator_ptr+int64(data[int(validator_ptr)])]
	payload := memory.Data()[payload_ptr : payload_ptr+int64(data[int(payload_ptr)])]
	vInfo, err := UnmarshalValidatorInfo(validator)
	if err != nil {
		return 0
	}
	artifact, err := extractValidationArtifacts(proof)
	if err != nil {
		return 0
	}

	if err := ValidateChainCodeID(artifact.prp, vInfo.Cid); err != nil {
		return 0
	}

	if err := ValidatePayload(artifact.payload, payload); err != nil {
		return 0
	}

	signatureSet := GetSignatureSet(artifact)

	pe, ok := ctxMap[FABRIC_EVALUATOR].(*PolicyEvaluator)
	if !ok {
		pe, err = NewPolicyEvaluator(vInfo.ConfByte)
		if err != nil {
			return 0
		}
		ctxMap[FABRIC_EVALUATOR] = pe
	}
	if err = pe.Evaluate([]byte(vInfo.Policy), signatureSet); err != nil {
		return 0
	}

	return 1
}

func (im *Imports) importFabricV14() {
	var err error
	im.imports, err = im.imports.Append("fabric_validate_v14", fabric_validate_v14, C.fabric_validate_v14)
	if err != nil {
		return
	}
}
