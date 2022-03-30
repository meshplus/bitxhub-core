package conversion

import (
	"crypto/elliptic"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
)

func BytesToHashString(msg []byte) (string, error) {
	h := sha256.New()
	_, err := h.Write(msg)
	if err != nil {
		return "", fmt.Errorf("fail to caculate sha256 hash: %w", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func MsgToHashInt(msg []byte) (*big.Int, error) {
	return hashToInt(msg, btcec.S256()), nil
}

func MsgToHashString(msg []byte) (string, error) {
	if len(msg) == 0 {
		return "", errors.New("empty message")
	}
	h := sha256.New()
	_, err := h.Write(msg)
	if err != nil {
		return "", fmt.Errorf("fail to caculate sha256 hash: %w", err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func hashToInt(hash []byte, c elliptic.Curve) *big.Int {
	orderBits := c.Params().N.BitLen()
	orderBytes := (orderBits + 7) / 8
	if len(hash) > orderBytes {
		hash = hash[:orderBytes]
	}

	ret := new(big.Int).SetBytes(hash)
	excess := len(hash)*8 - orderBits
	if excess > 0 {
		ret.Rsh(ret, uint(excess))
	}
	return ret
}

type Status byte

const (
	NA Status = iota
	Success
	Fail
)
