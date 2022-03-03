package validatorlib

import (
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
)

func NewValidatorLibs(context map[string]interface{}) []*wasmlib.ImportLib {
	var libs []*wasmlib.ImportLib
	libs = append(libs, FabricV14Verify(context))
	libs = append(libs, EcdsaVerify(context))
	return libs
}
