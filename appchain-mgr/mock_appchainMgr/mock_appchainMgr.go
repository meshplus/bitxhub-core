// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mock_appchainMgr is a generated GoMock package.
package mock_appchainMgr

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	logrus "github.com/sirupsen/logrus"
)

// MockAppchainMgr is a mock of AppchainMgr interface.
type MockAppchainMgr struct {
	ctrl     *gomock.Controller
	recorder *MockAppchainMgrMockRecorder
}

// MockAppchainMgrMockRecorder is the mock recorder for MockAppchainMgr.
type MockAppchainMgrMockRecorder struct {
	mock *MockAppchainMgr
}

// NewMockAppchainMgr creates a new mock instance.
func NewMockAppchainMgr(ctrl *gomock.Controller) *MockAppchainMgr {
	mock := &MockAppchainMgr{ctrl: ctrl}
	mock.recorder = &MockAppchainMgrMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAppchainMgr) EXPECT() *MockAppchainMgrMockRecorder {
	return m.recorder
}

// Appchain mocks base method.
func (m *MockAppchainMgr) Appchain() (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Appchain")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Appchain indicates an expected call of Appchain.
func (mr *MockAppchainMgrMockRecorder) Appchain() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Appchain", reflect.TypeOf((*MockAppchainMgr)(nil).Appchain))
}

// Appchains mocks base method.
func (m *MockAppchainMgr) Appchains() (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Appchains")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Appchains indicates an expected call of Appchains.
func (mr *MockAppchainMgrMockRecorder) Appchains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Appchains", reflect.TypeOf((*MockAppchainMgr)(nil).Appchains))
}

// Audit mocks base method.
func (m *MockAppchainMgr) Audit(proposer string, isApproved int32, desc string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Audit", proposer, isApproved, desc)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Audit indicates an expected call of Audit.
func (mr *MockAppchainMgrMockRecorder) Audit(proposer, isApproved, desc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Audit", reflect.TypeOf((*MockAppchainMgr)(nil).Audit), proposer, isApproved, desc)
}

// ChangeStatus mocks base method.
func (m *MockAppchainMgr) ChangeStatus(id, trigger string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChangeStatus", id, trigger)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// ChangeStatus indicates an expected call of ChangeStatus.
func (mr *MockAppchainMgrMockRecorder) ChangeStatus(id, trigger interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChangeStatus", reflect.TypeOf((*MockAppchainMgr)(nil).ChangeStatus), id, trigger)
}

// CountAppchains mocks base method.
func (m *MockAppchainMgr) CountAppchains() (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAppchains")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// CountAppchains indicates an expected call of CountAppchains.
func (mr *MockAppchainMgrMockRecorder) CountAppchains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAppchains", reflect.TypeOf((*MockAppchainMgr)(nil).CountAppchains))
}

// CountAvailableAppchains mocks base method.
func (m *MockAppchainMgr) CountAvailableAppchains() (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountAvailableAppchains")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// CountAvailableAppchains indicates an expected call of CountAvailableAppchains.
func (mr *MockAppchainMgrMockRecorder) CountAvailableAppchains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountAvailableAppchains", reflect.TypeOf((*MockAppchainMgr)(nil).CountAvailableAppchains))
}

// DeleteAppchain mocks base method.
func (m *MockAppchainMgr) DeleteAppchain(id string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAppchain", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// DeleteAppchain indicates an expected call of DeleteAppchain.
func (mr *MockAppchainMgrMockRecorder) DeleteAppchain(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAppchain", reflect.TypeOf((*MockAppchainMgr)(nil).DeleteAppchain), id)
}

// FetchAuditRecords mocks base method.
func (m *MockAppchainMgr) FetchAuditRecords(id string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchAuditRecords", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// FetchAuditRecords indicates an expected call of FetchAuditRecords.
func (mr *MockAppchainMgrMockRecorder) FetchAuditRecords(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchAuditRecords", reflect.TypeOf((*MockAppchainMgr)(nil).FetchAuditRecords), id)
}

// GetAppchain mocks base method.
func (m *MockAppchainMgr) GetAppchain(id string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAppchain", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// GetAppchain indicates an expected call of GetAppchain.
func (mr *MockAppchainMgrMockRecorder) GetAppchain(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAppchain", reflect.TypeOf((*MockAppchainMgr)(nil).GetAppchain), id)
}

// GetPubKeyByChainID mocks base method.
func (m *MockAppchainMgr) GetPubKeyByChainID(id string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKeyByChainID", id)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// GetPubKeyByChainID indicates an expected call of GetPubKeyByChainID.
func (mr *MockAppchainMgrMockRecorder) GetPubKeyByChainID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKeyByChainID", reflect.TypeOf((*MockAppchainMgr)(nil).GetPubKeyByChainID), id)
}

// Register mocks base method.
func (m *MockAppchainMgr) Register(id, validators, consensusType, chainType, name, desc, version, pubkey string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", id, validators, consensusType, chainType, name, desc, version, pubkey)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAppchainMgrMockRecorder) Register(id, validators, consensusType, chainType, name, desc, version, pubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAppchainMgr)(nil).Register), id, validators, consensusType, chainType, name, desc, version, pubkey)
}

// UpdateAppchain mocks base method.
func (m *MockAppchainMgr) UpdateAppchain(id, validators, consensusType, chainType, name, desc, version, pubkey string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAppchain", id, validators, consensusType, chainType, name, desc, version, pubkey)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// UpdateAppchain indicates an expected call of UpdateAppchain.
func (mr *MockAppchainMgrMockRecorder) UpdateAppchain(id, validators, consensusType, chainType, name, desc, version, pubkey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAppchain", reflect.TypeOf((*MockAppchainMgr)(nil).UpdateAppchain), id, validators, consensusType, chainType, name, desc, version, pubkey)
}

// MockPersister is a mock of Persister interface.
type MockPersister struct {
	ctrl     *gomock.Controller
	recorder *MockPersisterMockRecorder
}

// MockPersisterMockRecorder is the mock recorder for MockPersister.
type MockPersisterMockRecorder struct {
	mock *MockPersister
}

// NewMockPersister creates a new mock instance.
func NewMockPersister(ctrl *gomock.Controller) *MockPersister {
	mock := &MockPersister{ctrl: ctrl}
	mock.recorder = &MockPersisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPersister) EXPECT() *MockPersisterMockRecorder {
	return m.recorder
}

// Caller mocks base method.
func (m *MockPersister) Caller() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Caller")
	ret0, _ := ret[0].(string)
	return ret0
}

// Caller indicates an expected call of Caller.
func (mr *MockPersisterMockRecorder) Caller() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Caller", reflect.TypeOf((*MockPersister)(nil).Caller))
}

// Delete mocks base method.
func (m *MockPersister) Delete(key string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", key)
}

// Delete indicates an expected call of Delete.
func (mr *MockPersisterMockRecorder) Delete(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockPersister)(nil).Delete), key)
}

// Get mocks base method.
func (m *MockPersister) Get(key string) (bool, []byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([]byte)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockPersisterMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockPersister)(nil).Get), key)
}

// GetObject mocks base method.
func (m *MockPersister) GetObject(key string, ret interface{}) bool {
	m.ctrl.T.Helper()
	ret_2 := m.ctrl.Call(m, "GetObject", key, ret)
	ret0, _ := ret_2[0].(bool)
	return ret0
}

// GetObject indicates an expected call of GetObject.
func (mr *MockPersisterMockRecorder) GetObject(key, ret interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetObject", reflect.TypeOf((*MockPersister)(nil).GetObject), key, ret)
}

// Has mocks base method.
func (m *MockPersister) Has(key string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", key)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Has indicates an expected call of Has.
func (mr *MockPersisterMockRecorder) Has(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockPersister)(nil).Has), key)
}

// Logger mocks base method.
func (m *MockPersister) Logger() logrus.FieldLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logger")
	ret0, _ := ret[0].(logrus.FieldLogger)
	return ret0
}

// Logger indicates an expected call of Logger.
func (mr *MockPersisterMockRecorder) Logger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logger", reflect.TypeOf((*MockPersister)(nil).Logger))
}

// Query mocks base method.
func (m *MockPersister) Query(prefix string) (bool, [][]byte) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Query", prefix)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].([][]byte)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockPersisterMockRecorder) Query(prefix interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockPersister)(nil).Query), prefix)
}

// Set mocks base method.
func (m *MockPersister) Set(key string, value []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set.
func (mr *MockPersisterMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockPersister)(nil).Set), key, value)
}

// SetObject mocks base method.
func (m *MockPersister) SetObject(key string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetObject", key, value)
}

// SetObject indicates an expected call of SetObject.
func (mr *MockPersisterMockRecorder) SetObject(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetObject", reflect.TypeOf((*MockPersister)(nil).SetObject), key, value)
}
