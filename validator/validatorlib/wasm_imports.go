package validatorlib

import (
	"github.com/wasmerio/go-ext-wasm/wasmer"
)

type Imports struct {
	imports *wasmer.Imports
}

func New() (*wasmer.Imports, error) {
	imports := &Imports{
		imports: wasmer.NewImports(),
	}
	imports.importECDSA()
	imports.importFabricV14()
	imports.importFabricV13()
	imports.importFabricSm2()

	return imports.imports, nil
}
