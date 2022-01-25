package node_mgr

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/iancoleman/orderedmap"
	"github.com/looplab/fsm"
	"github.com/meshplus/bitxhub-core/boltvm"
	"github.com/meshplus/bitxhub-core/governance"
	"github.com/sirupsen/logrus"
)

type NodeType string

const (
	NODEPREFIX              = "node"
	NODETYPE_PREFIX         = "type"
	VP_NODE_ID_PREFIX       = "vp-id"
	VP_NODE_PID_PREFIX      = "vp-pid"
	NVP_NODE_NAME_PREFIX    = "nvp-name"
	NODE_OCCUPY_PID_PREFIX  = "occupy-node-pid"
	NODE_OCCUPY_NAME_PREFIX = "occupy-node-name"

	VPNode  NodeType = "vpNode"
	NVPNode NodeType = "nvpNode"
)

type NodeManager struct {
	governance.Persister
}

type Node struct {
	Account  string   `toml:"account" json:"account"`
	NodeType NodeType `toml:"node_type" json:"node_type"`

	// VP Node Info
	Pid      string `toml:"pid" json:"pid"`
	VPNodeId uint64 `toml:"id" json:"id"`
	Primary  bool   `toml:"primary" json:"primary"`

	// NVP Node
	Name           string              `toml:"name" json:"name"`
	Permissions    map[string]struct{} `toml:"permissions" json:"permissions"`
	AuditAdminAddr string              `toml:"audit_admin_addr" json:"audit_admin_addr"`

	Status governance.GovernanceStatus `toml:"status" json:"status"`
	FSM    *fsm.FSM                    `json:"fsm"`
}

var nodeAvailableMap = map[governance.GovernanceStatus]struct{}{
	governance.GovernanceAvailable: {},
	governance.GovernanceLogouting: {},
	governance.GovernanceBinding:   {},
	governance.GovernanceBinded:    {},
	governance.GovernanceUpdating:  {},
}

var nodeStateMap = map[governance.EventType][]governance.GovernanceStatus{
	governance.EventRegister: {governance.GovernanceUnavailable},
	governance.EventUpdate:   {governance.GovernanceAvailable, governance.GovernanceBinded},
	governance.EventBind:     {governance.GovernanceAvailable},
	// unbind binding node:
	// If the audit admin is logouted during the binding between the audit node and the audit admin, the audit node can be restored to the available state through the unbind event
	governance.EventUnbind: {governance.GovernanceBinded, governance.GovernanceBinding},
	governance.EventLogout: {governance.GovernanceAvailable, governance.GovernanceBinding, governance.GovernanceBinded, governance.GovernanceUpdating},
}

func New(persister governance.Persister) NodeMgr {
	return &NodeManager{persister}
}

func (n *Node) IsAvailable() bool {
	if _, ok := nodeAvailableMap[n.Status]; ok {
		return true
	} else {
		return false
	}
}

func (node *Node) setFSM(lastStatus governance.GovernanceStatus) {
	node.FSM = fsm.NewFSM(
		string(node.Status),
		fsm.Events{
			// register 3
			{Name: string(governance.EventRegister), Src: []string{string(governance.GovernanceUnavailable)}, Dst: string(governance.GovernanceRegisting)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceRegisting)}, Dst: string(governance.GovernanceAvailable)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceRegisting)}, Dst: string(lastStatus)},

			// update 2
			{Name: string(governance.EventUpdate), Src: []string{string(governance.GovernanceAvailable), string(governance.GovernanceBinded), string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceUpdating)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceUpdating)}, Dst: string(lastStatus)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceUpdating)}, Dst: string(lastStatus)},

			// bind 1
			{Name: string(governance.EventBind), Src: []string{string(governance.GovernanceAvailable), string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceBinding)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceBinding)}, Dst: string(governance.GovernanceBinded)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceBinding)}, Dst: string(governance.GovernanceAvailable)},

			// unbind 1
			{Name: string(governance.EventUnbind), Src: []string{string(governance.GovernanceBinded), string(governance.GovernanceBinding)}, Dst: string(governance.GovernanceAvailable)},

			// logout 3
			{Name: string(governance.EventLogout), Src: []string{string(governance.GovernanceAvailable), string(governance.GovernanceBinding), string(governance.GovernanceBinded), string(governance.GovernanceUpdating)}, Dst: string(governance.GovernanceLogouting)},
			{Name: string(governance.EventApprove), Src: []string{string(governance.GovernanceLogouting)}, Dst: string(governance.GovernanceForbidden)},
			{Name: string(governance.EventReject), Src: []string{string(governance.GovernanceLogouting)}, Dst: string(lastStatus)},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) {
				node.Status = governance.GovernanceStatus(node.FSM.Current())
				if node.Status == governance.GovernanceAvailable && node.NodeType == NVPNode {
					node.AuditAdminAddr = ""
				}
			},
		},
	)
}

// GovernancePre checks if the appchain can do the event. (only check, not modify infomation)
// return *node, extra info, error
func (nm *NodeManager) GovernancePre(nodeAccount string, event governance.EventType, _ []byte) (interface{}, *boltvm.BxhError) {
	node := &Node{}
	if ok := nm.GetObject(NodeKey(nodeAccount), node); !ok {
		if event == governance.EventRegister {
			return nil, nil
		} else {
			return nil, boltvm.BError(boltvm.NodeNonexistentNodeCode, fmt.Sprintf(string(boltvm.NodeNonexistentNodeMsg), nodeAccount))
		}
	}

	for _, s := range nodeStateMap[event] {
		if node.Status == s {
			return node, nil
		}
	}

	return nil, boltvm.BError(boltvm.NodeStatusErrorCode, fmt.Sprintf(string(boltvm.NodeStatusErrorMsg), nodeAccount, string(node.Status), string(event)))
}

func (nm *NodeManager) ChangeStatus(nodeAccount string, trigger, lastStatus string, _ []byte) (bool, []byte) {
	node := &Node{}
	if ok := nm.GetObject(NodeKey(nodeAccount), node); !ok {
		return false, []byte("this node does not exist")
	}

	node.setFSM(governance.GovernanceStatus(lastStatus))
	err := node.FSM.Event(trigger)
	if err != nil {
		return false, []byte(fmt.Sprintf("change status error: %v", err))
	}

	nm.SetObject(NodeKey(nodeAccount), *node)
	return true, nil
}

func (nm *NodeManager) RegisterPre(node *Node) {
	nm.SetObject(NodeKey(node.Account), *node)
}

// Register record node info
func (nm *NodeManager) Register(node *Node) {
	// 1. store node info
	nm.SetObject(NodeKey(node.Account), *node)

	// 2. store node type
	nodeAccountMap := orderedmap.New()
	_ = nm.GetObject(NodeTypeKey(string(node.NodeType)), nodeAccountMap)
	nodeAccountMap.Set(node.Account, struct{}{})
	nm.SetObject(NodeTypeKey(string(node.NodeType)), *nodeAccountMap)

	// 3. store vpid, pid, name
	switch node.NodeType {
	case VPNode:
		nm.SetObject(VpNodeIdKey(strconv.Itoa(int(node.VPNodeId))), node.Account)
		nm.SetObject(VpNodePidKey(node.Pid), node.Account)
	case NVPNode:
		nm.SetObject(NvpNodeNameKey(node.Name), node.Account)
	}

	nm.Logger().WithFields(logrus.Fields{
		"account":  node.Account,
		"nodeType": node.NodeType,
	}).Info("Node is registering")
}

func (nm *NodeManager) Update(nodeInfo *Node) (bool, []byte) {
	node := &Node{}
	ok := nm.GetObject(NodeKey(nodeInfo.Account), node)
	if !ok {
		return false, []byte(fmt.Sprintf("the node is not exist: %s", nodeInfo.Account))
	}

	oldName := node.Name
	node.Name = nodeInfo.Name
	node.Permissions = nodeInfo.Permissions
	nm.SetObject(NodeKey(nodeInfo.Account), *node)
	switch node.NodeType {
	case VPNode:
	case NVPNode:
		if oldName != node.Name {
			nm.Delete(NvpNodeNameKey(oldName))
			nm.SetObject(NvpNodeNameKey(node.Name), node.Account)
		}
	}

	nm.Logger().WithFields(logrus.Fields{
		"account": node.Account,
	}).Info("node is updating")

	return true, nil
}

func (nm *NodeManager) Bind(nodeAccount, auditAdminAddr string) (bool, []byte) {
	node := &Node{}
	ok := nm.GetObject(NodeKey(nodeAccount), node)
	if !ok {
		return false, []byte(fmt.Sprintf("the node is not exist: %s", nodeAccount))
	}

	node.AuditAdminAddr = auditAdminAddr
	nm.SetObject(NodeKey(nodeAccount), *node)
	nm.Logger().WithFields(logrus.Fields{
		"account":    node.Account,
		"auditAdmin": auditAdminAddr,
	}).Info("node is binding")

	return true, nil
}

// CountAvailable counts all available nodes (available、logouting)
func (nm *NodeManager) CountAvailable(nodeType []byte) (bool, []byte) {
	count := 0
	accountMap := nm.GetAccountMapByType(string(nodeType))

	for _, account := range accountMap.Keys() {
		node, err := nm.QueryById(account, nil)
		if err != nil {
			return false, []byte(fmt.Sprintf("the node %s is not exist", account))
		}
		if node.(*Node).IsAvailable() {
			count++
		}
	}

	return true, []byte(strconv.Itoa(count))
}

func (nm *NodeManager) CountAll(nodeType []byte) (bool, []byte) {
	accountMap := nm.GetAccountMapByType(string(nodeType))
	return true, []byte(strconv.Itoa(len(accountMap.Keys())))
}

// All returns all nodes
func (nm *NodeManager) All(_ []byte) (interface{}, error) {
	ret := make([]*Node, 0)
	ok, value := nm.Query(NODEPREFIX)
	if ok {
		for _, data := range value {
			node := &Node{}
			if err := json.Unmarshal(data, node); err != nil {
				return nil, err
			}
			ret = append(ret, node)
		}
	}

	return ret, nil
}

func (nm *NodeManager) QueryById(nodeAccount string, _ []byte) (interface{}, error) {
	var node Node
	ok := nm.GetObject(NodeKey(nodeAccount), &node)
	if !ok {
		return nil, fmt.Errorf("this node does not exist")
	}

	return &node, nil
}

func (nm *NodeManager) GetAccountMapByType(typ string) *orderedmap.OrderedMap {
	nodeAccountMap := orderedmap.New()
	_ = nm.GetObject(NodeTypeKey(typ), nodeAccountMap)
	return nodeAccountMap
}

func (nm *NodeManager) GetAccountByVpId(vpNodeId string) (string, error) {
	ok, data := nm.Get(VpNodeIdKey(vpNodeId))
	if !ok {
		return "", fmt.Errorf("this node does not exist")
	}

	return string(data), nil
}

func (nm *NodeManager) GetAccountByPid(pid string) (string, error) {
	ok, data := nm.Get(VpNodePidKey(pid))
	if !ok {
		return "", fmt.Errorf("this node does not exist")
	}

	return string(data), nil
}

func (nm *NodeManager) GetAccountByName(name string) (string, error) {
	ok, data := nm.Get(NvpNodeNameKey(name))
	if !ok {
		return "", fmt.Errorf("this node does not exist")
	}

	return string(data), nil
}

func NodeKey(account string) string {
	return fmt.Sprintf("%s-%s", NODEPREFIX, account)
}

func NodeTypeKey(typ string) string {
	return fmt.Sprintf("%s-%s", NODETYPE_PREFIX, typ)
}

func VpNodeIdKey(id string) string {
	return fmt.Sprintf("%s-%s", VP_NODE_ID_PREFIX, id)
}

func VpNodePidKey(pid string) string {
	return fmt.Sprintf("%s-%s", VP_NODE_PID_PREFIX, pid)
}

func NvpNodeNameKey(name string) string {
	return fmt.Sprintf("%s-%s", NVP_NODE_NAME_PREFIX, name)
}

func NodeOccupyPidKey(pid string) string {
	return fmt.Sprintf("%s-%s", NODE_OCCUPY_PID_PREFIX, pid)
}

func NodeOccupyNameKey(name string) string {
	return fmt.Sprintf("%s-%s", NODE_OCCUPY_NAME_PREFIX, name)
}
