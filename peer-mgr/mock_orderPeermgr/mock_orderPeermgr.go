// Code generated by MockGen. DO NOT EDIT.
// Source: peermgr.go

// Package mock_orderPeermgr is a generated GoMock package.
package mock_orderPeermgr

import (
	reflect "reflect"

	event "github.com/ethereum/go-ethereum/event"
	gomock "github.com/golang/mock/gomock"
	peer "github.com/libp2p/go-libp2p-core/peer"
	peer_mgr "github.com/meshplus/bitxhub-core/peer-mgr"
	pb "github.com/meshplus/bitxhub-model/pb"
)

// MockBasicPeerManager is a mock of BasicPeerManager interface.
type MockBasicPeerManager struct {
	ctrl     *gomock.Controller
	recorder *MockBasicPeerManagerMockRecorder
}

// MockBasicPeerManagerMockRecorder is the mock recorder for MockBasicPeerManager.
type MockBasicPeerManagerMockRecorder struct {
	mock *MockBasicPeerManager
}

// NewMockBasicPeerManager creates a new mock instance.
func NewMockBasicPeerManager(ctrl *gomock.Controller) *MockBasicPeerManager {
	mock := &MockBasicPeerManager{ctrl: ctrl}
	mock.recorder = &MockBasicPeerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBasicPeerManager) EXPECT() *MockBasicPeerManagerMockRecorder {
	return m.recorder
}

// AsyncSend mocks base method.
func (m *MockBasicPeerManager) AsyncSend(arg0 peer_mgr.KeyType, arg1 *pb.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsyncSend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AsyncSend indicates an expected call of AsyncSend.
func (mr *MockBasicPeerManagerMockRecorder) AsyncSend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncSend", reflect.TypeOf((*MockBasicPeerManager)(nil).AsyncSend), arg0, arg1)
}

// CountConnectedPeers mocks base method.
func (m *MockBasicPeerManager) CountConnectedPeers() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountConnectedPeers")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// CountConnectedPeers indicates an expected call of CountConnectedPeers.
func (mr *MockBasicPeerManagerMockRecorder) CountConnectedPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountConnectedPeers", reflect.TypeOf((*MockBasicPeerManager)(nil).CountConnectedPeers))
}

// Peers mocks base method.
func (m *MockBasicPeerManager) Peers() map[string]*peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers")
	ret0, _ := ret[0].(map[string]*peer.AddrInfo)
	return ret0
}

// Peers indicates an expected call of Peers.
func (mr *MockBasicPeerManagerMockRecorder) Peers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockBasicPeerManager)(nil).Peers))
}

// Send mocks base method.
func (m *MockBasicPeerManager) Send(arg0 peer_mgr.KeyType, arg1 *pb.Message) (*pb.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*pb.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockBasicPeerManagerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockBasicPeerManager)(nil).Send), arg0, arg1)
}

// Start mocks base method.
func (m *MockBasicPeerManager) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockBasicPeerManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockBasicPeerManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockBasicPeerManager) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockBasicPeerManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBasicPeerManager)(nil).Stop))
}

// MockOrderPeerManager is a mock of OrderPeerManager interface.
type MockOrderPeerManager struct {
	ctrl     *gomock.Controller
	recorder *MockOrderPeerManagerMockRecorder
}

// MockOrderPeerManagerMockRecorder is the mock recorder for MockOrderPeerManager.
type MockOrderPeerManagerMockRecorder struct {
	mock *MockOrderPeerManager
}

// NewMockOrderPeerManager creates a new mock instance.
func NewMockOrderPeerManager(ctrl *gomock.Controller) *MockOrderPeerManager {
	mock := &MockOrderPeerManager{ctrl: ctrl}
	mock.recorder = &MockOrderPeerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderPeerManager) EXPECT() *MockOrderPeerManagerMockRecorder {
	return m.recorder
}

// AddNode mocks base method.
func (m *MockOrderPeerManager) AddNode(newNodeID uint64, vpInfo *pb.VpInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AddNode", newNodeID, vpInfo)
}

// AddNode indicates an expected call of AddNode.
func (mr *MockOrderPeerManagerMockRecorder) AddNode(newNodeID, vpInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNode", reflect.TypeOf((*MockOrderPeerManager)(nil).AddNode), newNodeID, vpInfo)
}

// AsyncSend mocks base method.
func (m *MockOrderPeerManager) AsyncSend(arg0 peer_mgr.KeyType, arg1 *pb.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AsyncSend", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AsyncSend indicates an expected call of AsyncSend.
func (mr *MockOrderPeerManagerMockRecorder) AsyncSend(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AsyncSend", reflect.TypeOf((*MockOrderPeerManager)(nil).AsyncSend), arg0, arg1)
}

// Broadcast mocks base method.
func (m *MockOrderPeerManager) Broadcast(arg0 *pb.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockOrderPeerManagerMockRecorder) Broadcast(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockOrderPeerManager)(nil).Broadcast), arg0)
}

// CountConnectedPeers mocks base method.
func (m *MockOrderPeerManager) CountConnectedPeers() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountConnectedPeers")
	ret0, _ := ret[0].(uint64)
	return ret0
}

// CountConnectedPeers indicates an expected call of CountConnectedPeers.
func (mr *MockOrderPeerManagerMockRecorder) CountConnectedPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountConnectedPeers", reflect.TypeOf((*MockOrderPeerManager)(nil).CountConnectedPeers))
}

// DelNode mocks base method.
func (m *MockOrderPeerManager) DelNode(delID uint64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DelNode", delID)
}

// DelNode indicates an expected call of DelNode.
func (mr *MockOrderPeerManagerMockRecorder) DelNode(delID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelNode", reflect.TypeOf((*MockOrderPeerManager)(nil).DelNode), delID)
}

// Disconnect mocks base method.
func (m *MockOrderPeerManager) Disconnect(vpInfos map[uint64]*pb.VpInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Disconnect", vpInfos)
}

// Disconnect indicates an expected call of Disconnect.
func (mr *MockOrderPeerManagerMockRecorder) Disconnect(vpInfos interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disconnect", reflect.TypeOf((*MockOrderPeerManager)(nil).Disconnect), vpInfos)
}

// OrderPeers mocks base method.
func (m *MockOrderPeerManager) OrderPeers() map[uint64]*pb.VpInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderPeers")
	ret0, _ := ret[0].(map[uint64]*pb.VpInfo)
	return ret0
}

// OrderPeers indicates an expected call of OrderPeers.
func (mr *MockOrderPeerManagerMockRecorder) OrderPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderPeers", reflect.TypeOf((*MockOrderPeerManager)(nil).OrderPeers))
}

// OtherPeers mocks base method.
func (m *MockOrderPeerManager) OtherPeers() map[uint64]*peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OtherPeers")
	ret0, _ := ret[0].(map[uint64]*peer.AddrInfo)
	return ret0
}

// OtherPeers indicates an expected call of OtherPeers.
func (mr *MockOrderPeerManagerMockRecorder) OtherPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OtherPeers", reflect.TypeOf((*MockOrderPeerManager)(nil).OtherPeers))
}

// Peers mocks base method.
func (m *MockOrderPeerManager) Peers() map[string]*peer.AddrInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Peers")
	ret0, _ := ret[0].(map[string]*peer.AddrInfo)
	return ret0
}

// Peers indicates an expected call of Peers.
func (mr *MockOrderPeerManagerMockRecorder) Peers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Peers", reflect.TypeOf((*MockOrderPeerManager)(nil).Peers))
}

// Send mocks base method.
func (m *MockOrderPeerManager) Send(arg0 peer_mgr.KeyType, arg1 *pb.Message) (*pb.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", arg0, arg1)
	ret0, _ := ret[0].(*pb.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Send indicates an expected call of Send.
func (mr *MockOrderPeerManagerMockRecorder) Send(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockOrderPeerManager)(nil).Send), arg0, arg1)
}

// Start mocks base method.
func (m *MockOrderPeerManager) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockOrderPeerManagerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockOrderPeerManager)(nil).Start))
}

// Stop mocks base method.
func (m *MockOrderPeerManager) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockOrderPeerManagerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockOrderPeerManager)(nil).Stop))
}

// SubscribeOrderMessage mocks base method.
func (m *MockOrderPeerManager) SubscribeOrderMessage(ch chan<- peer_mgr.OrderMessageEvent) event.Subscription {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscribeOrderMessage", ch)
	ret0, _ := ret[0].(event.Subscription)
	return ret0
}

// SubscribeOrderMessage indicates an expected call of SubscribeOrderMessage.
func (mr *MockOrderPeerManagerMockRecorder) SubscribeOrderMessage(ch interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscribeOrderMessage", reflect.TypeOf((*MockOrderPeerManager)(nil).SubscribeOrderMessage), ch)
}

// UpdateRouter mocks base method.
func (m *MockOrderPeerManager) UpdateRouter(vpInfos map[uint64]*pb.VpInfo, isNew bool) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateRouter", vpInfos, isNew)
	ret0, _ := ret[0].(bool)
	return ret0
}

// UpdateRouter indicates an expected call of UpdateRouter.
func (mr *MockOrderPeerManagerMockRecorder) UpdateRouter(vpInfos, isNew interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateRouter", reflect.TypeOf((*MockOrderPeerManager)(nil).UpdateRouter), vpInfos, isNew)
}
