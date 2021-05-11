package validatorlib

import (
	"bytes"
	"crypto"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/meshplus/bitxhub-core/wasm/wasmlib"
	"github.com/wasmerio/wasmer-go/wasmer"
)

type AlgorithmOption string

const (
	// Secp256r1 secp256r1 algorithm
	Secp256r1 AlgorithmOption = "Secp256r1"
)

type PrivateKey struct {
	K *ecdsa.PrivateKey
}

// PublicKey ECDSA public key.
// never new(PublicKey), use NewPublicKey()
type PublicKey struct {
	k *ecdsa.PublicKey
}

type ECDSASignature struct {
	R, S *big.Int
}

func ecdsa_verify(env interface{}, args []wasmer.Value) ([]wasmer.Value, error) {
	memory, err := env.(*wasmlib.WasmEnv).Instance.Exports.GetMemory("memory")
	if err != nil {
		return nil, err
	}
	sig_ptr := args[0].I64()
	digest_ptr := args[1].I64()
	pubkey_ptr := args[2].I64()
	data := env.(*wasmlib.WasmEnv).Ctx["data"].(map[int]int)
	signature := memory.Data()[sig_ptr : sig_ptr+70]
	digest := memory.Data()[digest_ptr : digest_ptr+32]
	pubkey := memory.Data()[pubkey_ptr : pubkey_ptr+int64(data[int(pubkey_ptr)])]
	pemCert, _ := pem.Decode(pubkey)
	var cert *x509.Certificate
	cert, err = x509.ParseCertificate(pemCert.Bytes)
	if err != nil {
		return []wasmer.Value{wasmer.NewI32(0)}, nil
	}
	pk := cert.PublicKey
	r, s, err := unmarshalECDSASignature(signature)
	if err != nil {
		return []wasmer.Value{wasmer.NewI32(0)}, nil
	}
	isValid := ecdsa.Verify(pk.(*ecdsa.PublicKey), digest, r, s)

	if isValid {
		return []wasmer.Value{wasmer.NewI32(1)}, nil
	} else {
		return []wasmer.Value{wasmer.NewI32(0)}, nil
	}
}

func unmarshalECDSASignature(raw []byte) (*big.Int, *big.Int, error) {
	sig := new(ECDSASignature)
	_, err := asn1.Unmarshal(raw, sig)
	if err != nil {
		return nil, nil, fmt.Errorf("failed unmashalling signature [%s]", err)
	}

	// Validate sig
	if sig.R == nil {
		return nil, nil, fmt.Errorf("invalid signature, r must be different from nil")
	}
	if sig.S == nil {
		return nil, nil, fmt.Errorf("invalid signature, s must be different from nil")
	}

	if sig.R.Sign() != 1 {
		return nil, nil, fmt.Errorf("invalid signature, r must be larger than zero")
	}
	if sig.S.Sign() != 1 {
		return nil, nil, fmt.Errorf("invalid signature, s must be larger than zero")
	}

	return sig.R, sig.S, nil
}

func (im *Imports) importECDSA(store *wasmer.Store, wasmEnv *wasmlib.WasmEnv) {
	function := wasmer.NewFunctionWithEnvironment(
		store,
		wasmer.NewFunctionType(
			wasmer.NewValueTypes(wasmer.I64, wasmer.I64, wasmer.I64),
			wasmer.NewValueTypes(wasmer.I32),
		),
		wasmEnv,
		ecdsa_verify,
	)
	im.imports.Register(
		"env",
		map[string]wasmer.IntoExtern{
			"ecdsa_verify": function,
		},
	)
}

// Bytes returns a serialized, storable representation of this key
func (priv *PrivateKey) Bytes() ([]byte, error) {
	if priv.K == nil {
		return nil, fmt.Errorf("ECDSAPrivateKey.K is nil, please invoke FromBytes()")
	}
	r := make([]byte, 32)
	a := priv.K.D.Bytes()
	copy(r[32-len(a):], a)
	return r, nil
}

func (pub *PublicKey) Bytes() ([]byte, error) {
	x := pub.k.X.Bytes()
	y := pub.k.Y.Bytes()
	return bytes.Join(
		[][]byte{{0x04},
			make([]byte, 32-len(x)), x, // padding to 32 bytes
			make([]byte, 32-len(y)), y,
		}, nil), nil
}

func UnmarshalPrivateKey(data []byte, opt AlgorithmOption) (crypto.PrivateKey, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty private key data")
	}
	key := &PrivateKey{K: new(ecdsa.PrivateKey)}
	key.K.D = big.NewInt(0)
	key.K.D.SetBytes(data)
	switch opt {
	case Secp256r1:
		key.K.Curve = elliptic.P256()
	default:
		return nil, fmt.Errorf("unsupported algorithm option")
	}

	key.K.PublicKey.X, key.K.PublicKey.Y = key.K.Curve.ScalarBaseMult(data)

	return key, nil
}

func UnmarshalPublicKey(data []byte, opt AlgorithmOption) (crypto.PublicKey, error) {
	if len(data) == 0 {
		return nil, fmt.Errorf("empty public key data")
	}
	key := &PublicKey{k: new(ecdsa.PublicKey)}
	key.k.X = big.NewInt(0)
	key.k.Y = big.NewInt(0)
	if len(data) != 65 {
		return nil, fmt.Errorf("public key data length is not 65")
	}
	key.k.X.SetBytes(data[1:33])
	key.k.Y.SetBytes(data[33:])
	switch opt {
	case Secp256r1:
		key.k.Curve = elliptic.P256()
	}
	return key, nil
}
