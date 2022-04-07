package validatorlib

import (
	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym/ecdsa"
)

func EcdsaVerify(context map[string]interface{}) *wasmlib.ImportLib {
	return &wasmlib.ImportLib{
		Module: "env",
		Name:   "ecdsa",
		Func: func(signatureKey, digestKey, pubKey, optKey string) int64 {
			signature := context[signatureKey].([]byte)
			digest := context[digestKey].([]byte)
			pubkey, err := ecdsa.UnmarshalPublicKey(context[pubKey].([]byte), crypto.CryptoNameType[optKey])
			if err != nil {
				context["error"] = err
				return 0
			}
			isValid, err := pubkey.Verify(digest, signature)
			if err != nil {
				context["error"] = err
				return 0
			}

			if isValid {
				return 1
			}
			return 0
		},
	}
}
