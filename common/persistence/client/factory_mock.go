// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: factory.go
//
// Generated by this command:
//
//	mockgen -package client -source factory.go -destination factory_mock.go
//

// Package client is a generated GoMock package.
package client

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"

	persistence "github.com/uber/cadence/common/persistence"
	service "github.com/uber/cadence/common/service"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
	isgomock struct{}
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockFactory) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockFactoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFactory)(nil).Close))
}

// NewConfigStoreManager mocks base method.
func (m *MockFactory) NewConfigStoreManager() (persistence.ConfigStoreManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConfigStoreManager")
	ret0, _ := ret[0].(persistence.ConfigStoreManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewConfigStoreManager indicates an expected call of NewConfigStoreManager.
func (mr *MockFactoryMockRecorder) NewConfigStoreManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConfigStoreManager", reflect.TypeOf((*MockFactory)(nil).NewConfigStoreManager))
}

// NewDomainManager mocks base method.
func (m *MockFactory) NewDomainManager() (persistence.DomainManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDomainManager")
	ret0, _ := ret[0].(persistence.DomainManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDomainManager indicates an expected call of NewDomainManager.
func (mr *MockFactoryMockRecorder) NewDomainManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDomainManager", reflect.TypeOf((*MockFactory)(nil).NewDomainManager))
}

// NewDomainReplicationQueueManager mocks base method.
func (m *MockFactory) NewDomainReplicationQueueManager() (persistence.QueueManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDomainReplicationQueueManager")
	ret0, _ := ret[0].(persistence.QueueManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDomainReplicationQueueManager indicates an expected call of NewDomainReplicationQueueManager.
func (mr *MockFactoryMockRecorder) NewDomainReplicationQueueManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDomainReplicationQueueManager", reflect.TypeOf((*MockFactory)(nil).NewDomainReplicationQueueManager))
}

// NewExecutionManager mocks base method.
func (m *MockFactory) NewExecutionManager(shardID int) (persistence.ExecutionManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewExecutionManager", shardID)
	ret0, _ := ret[0].(persistence.ExecutionManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewExecutionManager indicates an expected call of NewExecutionManager.
func (mr *MockFactoryMockRecorder) NewExecutionManager(shardID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewExecutionManager", reflect.TypeOf((*MockFactory)(nil).NewExecutionManager), shardID)
}

// NewHistoryManager mocks base method.
func (m *MockFactory) NewHistoryManager() (persistence.HistoryManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHistoryManager")
	ret0, _ := ret[0].(persistence.HistoryManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewHistoryManager indicates an expected call of NewHistoryManager.
func (mr *MockFactoryMockRecorder) NewHistoryManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistoryManager", reflect.TypeOf((*MockFactory)(nil).NewHistoryManager))
}

// NewShardManager mocks base method.
func (m *MockFactory) NewShardManager() (persistence.ShardManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewShardManager")
	ret0, _ := ret[0].(persistence.ShardManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewShardManager indicates an expected call of NewShardManager.
func (mr *MockFactoryMockRecorder) NewShardManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewShardManager", reflect.TypeOf((*MockFactory)(nil).NewShardManager))
}

// NewTaskManager mocks base method.
func (m *MockFactory) NewTaskManager() (persistence.TaskManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTaskManager")
	ret0, _ := ret[0].(persistence.TaskManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTaskManager indicates an expected call of NewTaskManager.
func (mr *MockFactoryMockRecorder) NewTaskManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTaskManager", reflect.TypeOf((*MockFactory)(nil).NewTaskManager))
}

// NewVisibilityManager mocks base method.
func (m *MockFactory) NewVisibilityManager(params *Params, serviceConfig *service.Config) (persistence.VisibilityManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewVisibilityManager", params, serviceConfig)
	ret0, _ := ret[0].(persistence.VisibilityManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewVisibilityManager indicates an expected call of NewVisibilityManager.
func (mr *MockFactoryMockRecorder) NewVisibilityManager(params, serviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewVisibilityManager", reflect.TypeOf((*MockFactory)(nil).NewVisibilityManager), params, serviceConfig)
}

// MockDataStoreFactory is a mock of DataStoreFactory interface.
type MockDataStoreFactory struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreFactoryMockRecorder
	isgomock struct{}
}

// MockDataStoreFactoryMockRecorder is the mock recorder for MockDataStoreFactory.
type MockDataStoreFactoryMockRecorder struct {
	mock *MockDataStoreFactory
}

// NewMockDataStoreFactory creates a new mock instance.
func NewMockDataStoreFactory(ctrl *gomock.Controller) *MockDataStoreFactory {
	mock := &MockDataStoreFactory{ctrl: ctrl}
	mock.recorder = &MockDataStoreFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStoreFactory) EXPECT() *MockDataStoreFactoryMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDataStoreFactory) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockDataStoreFactoryMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDataStoreFactory)(nil).Close))
}

// NewConfigStore mocks base method.
func (m *MockDataStoreFactory) NewConfigStore() (persistence.ConfigStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConfigStore")
	ret0, _ := ret[0].(persistence.ConfigStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewConfigStore indicates an expected call of NewConfigStore.
func (mr *MockDataStoreFactoryMockRecorder) NewConfigStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConfigStore", reflect.TypeOf((*MockDataStoreFactory)(nil).NewConfigStore))
}

// NewDomainStore mocks base method.
func (m *MockDataStoreFactory) NewDomainStore() (persistence.DomainStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDomainStore")
	ret0, _ := ret[0].(persistence.DomainStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewDomainStore indicates an expected call of NewDomainStore.
func (mr *MockDataStoreFactoryMockRecorder) NewDomainStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDomainStore", reflect.TypeOf((*MockDataStoreFactory)(nil).NewDomainStore))
}

// NewExecutionStore mocks base method.
func (m *MockDataStoreFactory) NewExecutionStore(shardID int) (persistence.ExecutionStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewExecutionStore", shardID)
	ret0, _ := ret[0].(persistence.ExecutionStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewExecutionStore indicates an expected call of NewExecutionStore.
func (mr *MockDataStoreFactoryMockRecorder) NewExecutionStore(shardID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewExecutionStore", reflect.TypeOf((*MockDataStoreFactory)(nil).NewExecutionStore), shardID)
}

// NewHistoryStore mocks base method.
func (m *MockDataStoreFactory) NewHistoryStore() (persistence.HistoryStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewHistoryStore")
	ret0, _ := ret[0].(persistence.HistoryStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewHistoryStore indicates an expected call of NewHistoryStore.
func (mr *MockDataStoreFactoryMockRecorder) NewHistoryStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewHistoryStore", reflect.TypeOf((*MockDataStoreFactory)(nil).NewHistoryStore))
}

// NewQueue mocks base method.
func (m *MockDataStoreFactory) NewQueue(queueType persistence.QueueType) (persistence.Queue, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewQueue", queueType)
	ret0, _ := ret[0].(persistence.Queue)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewQueue indicates an expected call of NewQueue.
func (mr *MockDataStoreFactoryMockRecorder) NewQueue(queueType any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewQueue", reflect.TypeOf((*MockDataStoreFactory)(nil).NewQueue), queueType)
}

// NewShardStore mocks base method.
func (m *MockDataStoreFactory) NewShardStore() (persistence.ShardStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewShardStore")
	ret0, _ := ret[0].(persistence.ShardStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewShardStore indicates an expected call of NewShardStore.
func (mr *MockDataStoreFactoryMockRecorder) NewShardStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewShardStore", reflect.TypeOf((*MockDataStoreFactory)(nil).NewShardStore))
}

// NewTaskStore mocks base method.
func (m *MockDataStoreFactory) NewTaskStore() (persistence.TaskStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewTaskStore")
	ret0, _ := ret[0].(persistence.TaskStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewTaskStore indicates an expected call of NewTaskStore.
func (mr *MockDataStoreFactoryMockRecorder) NewTaskStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewTaskStore", reflect.TypeOf((*MockDataStoreFactory)(nil).NewTaskStore))
}

// NewVisibilityStore mocks base method.
func (m *MockDataStoreFactory) NewVisibilityStore(sortByCloseTime bool) (persistence.VisibilityStore, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewVisibilityStore", sortByCloseTime)
	ret0, _ := ret[0].(persistence.VisibilityStore)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewVisibilityStore indicates an expected call of NewVisibilityStore.
func (mr *MockDataStoreFactoryMockRecorder) NewVisibilityStore(sortByCloseTime any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewVisibilityStore", reflect.TypeOf((*MockDataStoreFactory)(nil).NewVisibilityStore), sortByCloseTime)
}
