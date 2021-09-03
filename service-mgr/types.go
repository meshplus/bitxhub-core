package service_mgr

import "github.com/meshplus/bitxhub-core/governance"

//go:generate mockgen -destination mock_serviceMgr/mock_serviceMgr.go -package mock_serviceMgr -source types.go
type ServiceMgr interface {
	governance.Governance

	// Register registers object info, return object id and error
	Register(serviceInfo *Service) (bool, []byte)

	// Update updates available or frozen object
	Update(serviceInfo *Service) (bool, []byte)
}
