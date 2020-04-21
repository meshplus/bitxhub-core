package validatorlib

// #include <stdlib.h>
//
// extern int32_t fabric_validate_v13(void *context, long long proof_ptr, long long validator_ptr);
import "C"
import (
	"unsafe"
)

//export fabric_validate_v13
func fabric_validate_v13(context unsafe.Pointer, proof_ptr int64, validator_ptr int64) int32 {
	return 1
}

func (im *Imports) importFabricV13() {
	var err error
	im.imports, err = im.imports.Append("fabric_validate_v13", fabric_validate_v13, C.fabric_validate_v13)
	if err != nil {
		return
	}
}
