package validatorlib

// #include <stdlib.h>
//
// extern int32_t fabric_validate_sm2(void *context, long long proof_ptr, long long validator_ptr, long long payload_ptr);
import "C"
import (
	"unsafe"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

//export fabric_validate_sm2
func fabric_validate_sm2(context unsafe.Pointer, proof_ptr int64, validator_ptr int64, payload_ptr int64) int32 {
	ctx := wasmer.IntoInstanceContext(context)
	ctxMap := ctx.Data().(map[string]interface{})
	data := ctxMap["argmap"].(map[int]int)
	memory := ctx.Memory()
	proof := memory.Data()[proof_ptr : proof_ptr+int64(data[int(proof_ptr)])]
	validator := memory.Data()[validator_ptr : validator_ptr+int64(data[int(validator_ptr)])]
	// payload := memory.Data()[payload_ptr : payload_ptr+int64(data[int(payload_ptr)])]
	isPass, err := fabric_sm2_validate(proof, validator)
	if !isPass || err != nil {
		return 0
	}

	return 1
}

func (im *Imports) importFabricSm2() {
	var err error
	im.imports, err = im.imports.Append("fabric_validate_sm2", fabric_validate_sm2, C.fabric_validate_sm2)
	if err != nil {
		return
	}
}
