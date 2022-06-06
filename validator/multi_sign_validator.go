package validator

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/meshplus/bitxhub-kit/crypto"
	"github.com/meshplus/bitxhub-kit/crypto/asym/ecdsa"
	"github.com/meshplus/bitxhub-kit/types"
	"github.com/sirupsen/logrus"
)

type MultiSignValidator struct {
	logger  logrus.FieldLogger
	signMap *sync.Map
}

type MultiValidators struct {
	Validators []string `json:"validators"`
}

type MultiProof struct {
	Signs [][]byte `json:"signs"`
}

// New a validator instance
func NewMultiSignValidator(logger logrus.FieldLogger) *MultiSignValidator {
	return &MultiSignValidator{
		logger:  logger,
		signMap: &sync.Map{},
	}
}

func (m *MultiSignValidator) Verify(from string, proof, payload []byte, validators string) (bool, uint64, error) {
	multiValidators := &MultiValidators{}
	if err := json.Unmarshal([]byte(validators), multiValidators); err != nil {
		return false, 0, err
	}

	multiMap := make(map[string]struct{}, 0)
	for _, val := range multiValidators.Validators {
		multiMap[val] = struct{}{}
	}

	signs := &MultiProof{}
	if err := json.Unmarshal(proof, signs); err != nil {
		return false, 0, err
	}

	hash := sha256.Sum256(payload)
	threshold := (len(multiValidators.Validators) - 1) / 3 // TODO be dynamic
	counter := 0

	for _, sign := range signs.Signs {
		addr, err := recoverSignAddress(sign, hash[:])
		if err != nil {
			m.logger.Warnf("recover sign address failed: %s", err.Error())
			continue
		}

		val := addr.String()
		_, ok := multiMap[val]
		if !ok {
			m.logger.Warnf("wrong validator: %s", val)
			continue
		}

		delete(multiMap, val)
		counter++
		if counter > threshold {
			return true, 0, nil
		}
	}

	return false, 0, fmt.Errorf("proof verify failed: multi signs verify fail, counter: %d", counter)
}

func recoverSignAddress(sig, digest []byte) (*types.Address, error) {
	pubKeyBytes, err := ecdsa.Ecrecover(digest, sig)
	if err != nil {
		return nil, fmt.Errorf("recover public key failed: %w", err)
	}
	pubkey, err := ecdsa.UnmarshalPublicKey(pubKeyBytes, crypto.Secp256k1)
	if err != nil {
		return nil, fmt.Errorf("unmarshal public key error: %w", err)
	}

	return pubkey.Address()
}
