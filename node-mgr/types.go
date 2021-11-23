package node_mgr

import (
	"github.com/iancoleman/orderedmap"
	"github.com/meshplus/bitxhub-core/governance"
)

//go:generate mockgen -destination mock_nodeMgr/mock_nodeMgr.go -package mock_nodeMgr -source types.go
type NodeMgr interface {
	governance.Governance

	// Register registers node info, return node id and error
	RegisterPre(node *Node)

	Register(node *Node)

	Update(nodeInfo *Node) (bool, []byte)

	Bind(nodeAccount, auditAdminAddr string) (bool, []byte)

	GetAccountMapByType(typ string) *orderedmap.OrderedMap

	GetAccountByVpId(vpNodeId string) (string, error)

	GetAccountByPid(pid string) (string, error)

	GetAccountByName(name string) (string, error)
}
