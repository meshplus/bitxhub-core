package node_mgr

import "github.com/meshplus/bitxhub-core/governance"

//go:generate mockgen -destination mock_nodeMgr/mock_nodeMgr.go -package mock_nodeMgr -source types.go
type NodeMgr interface {
	governance.Governance

	// Register registers node info, return node id and error
	Register(info []byte) (bool, []byte)

	// GetIdByPid query node id by node pid
	GetIdByPid(pid string) (string, error)
}
