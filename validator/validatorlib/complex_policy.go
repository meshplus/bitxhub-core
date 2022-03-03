package validatorlib

import (
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
)

func FabricV14Verify(context map[string]interface{}) *wasmlib.ImportLib {
	return &wasmlib.ImportLib{
		Module: "env",
		Name:   "fabric_validate_v14",
		Func: func(proofKey, validatorKey, payloadKey string) int64 {
			proof := context[proofKey].([]byte)
			validator := context[validatorKey].([]byte)
			payload := context[payloadKey].([]byte)
			vInfo, err := UnmarshalValidatorInfo(validator)
			if err != nil {
				context["error"] = err
				return 0
			}
			artifact, err := extractValidationArtifacts(proof)
			if err != nil {
				context["error"] = err
				return 0
			}

			if err := ValidateChainCodeID(artifact.prp, vInfo.Cid); err != nil {
				context["error"] = err
				return 0
			}

			if err := ValidatePayload(artifact.payload, payload); err != nil {
				context["error"] = err
				return 0
			}

			signatureSet := GetSignatureSet(artifact)

			pe, ok := context[FABRIC_EVALUATOR].(*PolicyEvaluator)
			if !ok {
				pe, err = NewPolicyEvaluator(vInfo.ConfByte)
				if err != nil {
					context["error"] = err
					return 0
				}
				context[FABRIC_EVALUATOR] = pe
			}
			if err = pe.Evaluate([]byte(vInfo.Policy), signatureSet); err != nil {
				context["error"] = err
				return 0
			}

			return 1
		},
	}
}
