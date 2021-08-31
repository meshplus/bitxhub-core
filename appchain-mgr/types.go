package appchain_mgr

import "github.com/meshplus/bitxhub-core/governance"

//go:generate mockgen -destination mock_appchainMgr/mock_appchainMgr.go -package mock_appchainMgr -source types.go
type AppchainMgr interface {
	governance.Governance

	// Register registers object info, return object id and error
	Register(chainInfo *Appchain) (bool, []byte)

	// Update updates available or frozen object
	Update(updateInfo *Appchain) (bool, []byte)

	// Audit bitxhub manager audit appchain register info
	// caller is the bitxhub manager address
	// proposer is the appchain manager address
	Audit(proposer string, isApproved int32, desc string) (bool, []byte)

	//FetchAuditRecords fetches audit records by appchain id
	FetchAuditRecords(id string) (bool, []byte)

	// DeleteAppchain deletes appchain
	DeleteAppchain(id string) (bool, []byte)
}
