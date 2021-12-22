package validator

import "github.com/meshplus/bitxhub-kit/types"

// Engine runs for validation
//go:generate mockgen -destination mock_validator/mock_engine.go -package mock_validator -source validator.go
type Engine interface {
	Validate(address, from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error)
}

// Validator chooses specific method to verify transaction
type Validator interface {
	Verify(from string, proof, payload []byte, validators string, index, height uint64) (bool, uint64, error)
}

type Ledger interface {
	// GetCode
	GetCode(*types.Address) []byte
}
