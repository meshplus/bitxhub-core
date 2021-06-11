package governance

import (
	"github.com/sirupsen/logrus"
)

//go:generate mockgen -destination mock_governance/mock_governance.go -package mock_governance -source interface.go
type Governance interface {
	// ChangeStatus changes state of object
	ChangeStatus(id, trigger, lastStatus string, extra []byte) (bool, []byte)

	// CountAvailable counts all available objects
	CountAvailable(extra []byte) (bool, []byte)

	// CountAll counts all objects including approved, rejected, registered and so on
	CountAll(extra []byte) (bool, []byte)

	// ALL returns all objects
	All(extra []byte) (bool, []byte)

	// QueryById returns object info by id
	QueryById(id string, extra []byte) (bool, []byte)

	// GovernancePre check if the object can do the event
	GovernancePre(id string, event EventType, extra []byte) (bool, []byte)
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
	// GetAccount get ledger account address
	GetAccount(address string) (bool, interface{})
}
