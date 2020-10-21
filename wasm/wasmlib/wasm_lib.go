package wasmlib

// #include <stdlib.h>
//
// extern int32_t test_verify(void *context, long long proof_ptr, long long validator_ptr, long long payload_ptr);
import "C"
import (
	"unsafe"

	"github.com/wasmerio/go-ext-wasm/wasmer"
)

//export test_verify
func test_verify(context unsafe.Pointer, proof_ptr int64, validator_ptr int64, payload_ptr int64) int32 {
	ctx := wasmer.IntoInstanceContext(context)
	ctxMap := ctx.Data().(map[string]interface{})
	_ = ctxMap["argmap"].(map[int]int)
	_ = ctx.Memory()
	// proof := memory.Data()[proof_ptr : proof_ptr+int64(data[int(proof_ptr)])]
	// validator := memory.Data()[validator_ptr : validator_ptr+int64(data[int(validator_ptr)])]
	// payload := memory.Data()[payload_ptr : payload_ptr+int64(data[int(payload_ptr)])]

	// fmt.Println(proof)
	// fmt.Println(validator)
	// fmt.Println(payload)
	// fmt.Println(ctxMap["hello"].(string))

	return 1
}

func (im *Imports) importWasmLib() {
	var err error
	im.imports, err = im.imports.Append("test_verify", test_verify, C.test_verify)
	if err != nil {
		return
	}
}
