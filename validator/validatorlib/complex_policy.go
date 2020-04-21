package validatorlib

// #include <stdlib.h>
//
// extern int32_t fabric_validate_v14(void *context, long long proof_ptr, long long validator_ptr, long long payload_ptr);
import "C"
import (
	"unsafe"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

//export fabric_validate_v14
func fabric_validate_v14(context unsafe.Pointer, proof_ptr int64, validator_ptr int64, payload_ptr int64) int32 {
	ctx := wasmer.IntoInstanceContext(context)
	data := ctx.Data().(map[int]int)
	memory := ctx.Memory()
	proof := memory.Data()[proof_ptr : proof_ptr+int64(data[int(proof_ptr)])]
	validator := memory.Data()[validator_ptr : validator_ptr+int64(data[int(validator_ptr)])]
	payload := memory.Data()[payload_ptr : payload_ptr+int64(data[int(payload_ptr)])]
	vInfo, err := UnmarshalValidatorInfo(validator)
	if err != nil {
		return 0
	}
	err = ValidateV14(proof, payload, []byte(vInfo.Policy), vInfo.ConfByte, vInfo.Cid)
	if err != nil {
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
