package validator

import "github.com/meshplus/bitxhub-kit/types"

// Engine runs for validation
//go:generate mockgen -destination mock_validator/mock_engine.go -package mock_validator -source validator.go
type Engine interface {
	Validate(address, from string, proof, payload []byte, validators string) (bool, error)
}

// Validator chooses specific method to verify transaction
type Validator interface {
	Verify(address, from string, proof, payload []byte, validators string) (bool, error)
}

type Ledger interface {
	// GetCode
	GetCode(*types.Address) []byte
}
