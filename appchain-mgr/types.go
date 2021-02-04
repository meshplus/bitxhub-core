package appchain_mgr

import (
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination mock_appchainMgr/mock_appchainMgr.go -package mock_appchainMgr -source types.go
type AppchainMgr interface {
	// Register appchain managers registers appchain info caller is the appchain
	// manager address return appchain id and error
	Register(id, validators string, consensusType int32, chainType, name, desc, version, pubkey string) (bool, []byte)

	// UpdateAppchain updates approved appchain
	UpdateAppchain(id, validators string, consensusType int32, chainType, name, desc, version, pubkey string) (bool, []byte)

	// Audit bitxhub manager audit appchain register info
	// caller is the bitxhub manager address
	// proposer is the appchain manager address
	Audit(proposer string, isApproved int32, desc string) (bool, []byte)

	//FetchAuditRecords fetches audit records by appchain id
	FetchAuditRecords(id string) (bool, []byte)

	// CountApprovedAppchains counts all approved appchains
	CountApprovedAppchains() (bool, []byte)

	// CountAppchains counts all appchains including approved, rejected or registered
	CountAppchains() (bool, []byte)

	// Appchains returns all appchains
	Appchains() (bool, []byte)

	// DeleteAppchain deletes appchain
	DeleteAppchain(id string) (bool, []byte)

	// Appchain returns appchain info
	Appchain() (bool, []byte)

	// GetAppchain returns appchain info by id
	GetAppchain(id string) (bool, []byte)

	// GetPubKeyByChainID can get aim chain's public key using aim chain ID
	GetPubKeyByChainID(id string) (bool, []byte)
}

type Persister interface {
	// Caller
	Caller() string
	// Logger
	Logger() logrus.FieldLogger
	// Has judges key
	Has(key string) bool
	// Get gets value from datastore by key
	Get(key string) (bool, []byte)
	// GetObject
	GetObject(key string, ret interface{}) bool
	// Set sets k-v
	Set(key string, value []byte)
	// SetObject sets k with object v, v will be marshaled using json
	SetObject(key string, value interface{})
	// Delete deletes k-v
	Delete(key string)
	// QueryByPrefix queries object by prefix
	Query(prefix string) (bool, [][]byte)
}
