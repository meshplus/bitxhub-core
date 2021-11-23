package service_mgr

import "github.com/meshplus/bitxhub-core/governance"

//go:generate mockgen -destination mock_serviceMgr/mock_serviceMgr.go -package mock_serviceMgr -source types.go
type ServiceMgr interface {
	governance.Governance

	PackageServiceInfo(chainID, serviceID, name, typ, intro string, ordered bool, permits, details string, createTime int64, status governance.GovernanceStatus) (*Service, error)

	RegisterPre(info *Service)

	// Register registers object info, return object id and error
	Register(serviceInfo *Service)

	// Update updates available or frozen object
	Update(serviceInfo *Service) (bool, []byte)

	GetIDListByChainID(chainID string) ([]string, error)

	GetIDListByType(typ string) ([]string, error)

	GetServiceIDByName(name string) (string, error)
}
