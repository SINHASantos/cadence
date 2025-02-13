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
// Source: interface.go
//
// Generated by this command:
//
//	mockgen -package messaging -source interface.go -destination client_mock.go
//

// Package messaging is a generated GoMock package.
package messaging

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// NewConsumer mocks base method.
func (m *MockClient) NewConsumer(appName, consumerName string) (Consumer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewConsumer", appName, consumerName)
	ret0, _ := ret[0].(Consumer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewConsumer indicates an expected call of NewConsumer.
func (mr *MockClientMockRecorder) NewConsumer(appName, consumerName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewConsumer", reflect.TypeOf((*MockClient)(nil).NewConsumer), appName, consumerName)
}

// NewProducer mocks base method.
func (m *MockClient) NewProducer(appName string) (Producer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewProducer", appName)
	ret0, _ := ret[0].(Producer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewProducer indicates an expected call of NewProducer.
func (mr *MockClientMockRecorder) NewProducer(appName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewProducer", reflect.TypeOf((*MockClient)(nil).NewProducer), appName)
}

// MockConsumer is a mock of Consumer interface.
type MockConsumer struct {
	ctrl     *gomock.Controller
	recorder *MockConsumerMockRecorder
	isgomock struct{}
}

// MockConsumerMockRecorder is the mock recorder for MockConsumer.
type MockConsumerMockRecorder struct {
	mock *MockConsumer
}

// NewMockConsumer creates a new mock instance.
func NewMockConsumer(ctrl *gomock.Controller) *MockConsumer {
	mock := &MockConsumer{ctrl: ctrl}
	mock.recorder = &MockConsumerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsumer) EXPECT() *MockConsumerMockRecorder {
	return m.recorder
}

// Messages mocks base method.
func (m *MockConsumer) Messages() <-chan Message {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Messages")
	ret0, _ := ret[0].(<-chan Message)
	return ret0
}

// Messages indicates an expected call of Messages.
func (mr *MockConsumerMockRecorder) Messages() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Messages", reflect.TypeOf((*MockConsumer)(nil).Messages))
}

// Start mocks base method.
func (m *MockConsumer) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockConsumerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockConsumer)(nil).Start))
}

// Stop mocks base method.
func (m *MockConsumer) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockConsumerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockConsumer)(nil).Stop))
}

// MockMessage is a mock of Message interface.
type MockMessage struct {
	ctrl     *gomock.Controller
	recorder *MockMessageMockRecorder
	isgomock struct{}
}

// MockMessageMockRecorder is the mock recorder for MockMessage.
type MockMessageMockRecorder struct {
	mock *MockMessage
}

// NewMockMessage creates a new mock instance.
func NewMockMessage(ctrl *gomock.Controller) *MockMessage {
	mock := &MockMessage{ctrl: ctrl}
	mock.recorder = &MockMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessage) EXPECT() *MockMessageMockRecorder {
	return m.recorder
}

// Ack mocks base method.
func (m *MockMessage) Ack() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ack")
	ret0, _ := ret[0].(error)
	return ret0
}

// Ack indicates an expected call of Ack.
func (mr *MockMessageMockRecorder) Ack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ack", reflect.TypeOf((*MockMessage)(nil).Ack))
}

// Nack mocks base method.
func (m *MockMessage) Nack() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nack")
	ret0, _ := ret[0].(error)
	return ret0
}

// Nack indicates an expected call of Nack.
func (mr *MockMessageMockRecorder) Nack() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nack", reflect.TypeOf((*MockMessage)(nil).Nack))
}

// Offset mocks base method.
func (m *MockMessage) Offset() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Offset")
	ret0, _ := ret[0].(int64)
	return ret0
}

// Offset indicates an expected call of Offset.
func (mr *MockMessageMockRecorder) Offset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Offset", reflect.TypeOf((*MockMessage)(nil).Offset))
}

// Partition mocks base method.
func (m *MockMessage) Partition() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Partition")
	ret0, _ := ret[0].(int32)
	return ret0
}

// Partition indicates an expected call of Partition.
func (mr *MockMessageMockRecorder) Partition() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Partition", reflect.TypeOf((*MockMessage)(nil).Partition))
}

// Value mocks base method.
func (m *MockMessage) Value() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockMessageMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockMessage)(nil).Value))
}

// MockProducer is a mock of Producer interface.
type MockProducer struct {
	ctrl     *gomock.Controller
	recorder *MockProducerMockRecorder
	isgomock struct{}
}

// MockProducerMockRecorder is the mock recorder for MockProducer.
type MockProducerMockRecorder struct {
	mock *MockProducer
}

// NewMockProducer creates a new mock instance.
func NewMockProducer(ctrl *gomock.Controller) *MockProducer {
	mock := &MockProducer{ctrl: ctrl}
	mock.recorder = &MockProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducer) EXPECT() *MockProducerMockRecorder {
	return m.recorder
}

// Publish mocks base method.
func (m *MockProducer) Publish(ctx context.Context, message any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockProducerMockRecorder) Publish(ctx, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockProducer)(nil).Publish), ctx, message)
}

// MockCloseableProducer is a mock of CloseableProducer interface.
type MockCloseableProducer struct {
	ctrl     *gomock.Controller
	recorder *MockCloseableProducerMockRecorder
	isgomock struct{}
}

// MockCloseableProducerMockRecorder is the mock recorder for MockCloseableProducer.
type MockCloseableProducerMockRecorder struct {
	mock *MockCloseableProducer
}

// NewMockCloseableProducer creates a new mock instance.
func NewMockCloseableProducer(ctrl *gomock.Controller) *MockCloseableProducer {
	mock := &MockCloseableProducer{ctrl: ctrl}
	mock.recorder = &MockCloseableProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCloseableProducer) EXPECT() *MockCloseableProducerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockCloseableProducer) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockCloseableProducerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockCloseableProducer)(nil).Close))
}

// Publish mocks base method.
func (m *MockCloseableProducer) Publish(ctx context.Context, message any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", ctx, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// Publish indicates an expected call of Publish.
func (mr *MockCloseableProducerMockRecorder) Publish(ctx, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockCloseableProducer)(nil).Publish), ctx, message)
}

// MockAckManager is a mock of AckManager interface.
type MockAckManager struct {
	ctrl     *gomock.Controller
	recorder *MockAckManagerMockRecorder
	isgomock struct{}
}

// MockAckManagerMockRecorder is the mock recorder for MockAckManager.
type MockAckManagerMockRecorder struct {
	mock *MockAckManager
}

// NewMockAckManager creates a new mock instance.
func NewMockAckManager(ctrl *gomock.Controller) *MockAckManager {
	mock := &MockAckManager{ctrl: ctrl}
	mock.recorder = &MockAckManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAckManager) EXPECT() *MockAckManagerMockRecorder {
	return m.recorder
}

// AckItem mocks base method.
func (m *MockAckManager) AckItem(id int64) int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AckItem", id)
	ret0, _ := ret[0].(int64)
	return ret0
}

// AckItem indicates an expected call of AckItem.
func (mr *MockAckManagerMockRecorder) AckItem(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AckItem", reflect.TypeOf((*MockAckManager)(nil).AckItem), id)
}

// GetAckLevel mocks base method.
func (m *MockAckManager) GetAckLevel() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAckLevel")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetAckLevel indicates an expected call of GetAckLevel.
func (mr *MockAckManagerMockRecorder) GetAckLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAckLevel", reflect.TypeOf((*MockAckManager)(nil).GetAckLevel))
}

// GetBacklogCount mocks base method.
func (m *MockAckManager) GetBacklogCount() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBacklogCount")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetBacklogCount indicates an expected call of GetBacklogCount.
func (mr *MockAckManagerMockRecorder) GetBacklogCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBacklogCount", reflect.TypeOf((*MockAckManager)(nil).GetBacklogCount))
}

// GetReadLevel mocks base method.
func (m *MockAckManager) GetReadLevel() int64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReadLevel")
	ret0, _ := ret[0].(int64)
	return ret0
}

// GetReadLevel indicates an expected call of GetReadLevel.
func (mr *MockAckManagerMockRecorder) GetReadLevel() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReadLevel", reflect.TypeOf((*MockAckManager)(nil).GetReadLevel))
}

// ReadItem mocks base method.
func (m *MockAckManager) ReadItem(id int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadItem", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReadItem indicates an expected call of ReadItem.
func (mr *MockAckManagerMockRecorder) ReadItem(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadItem", reflect.TypeOf((*MockAckManager)(nil).ReadItem), id)
}

// SetAckLevel mocks base method.
func (m *MockAckManager) SetAckLevel(ackLevel int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAckLevel", ackLevel)
}

// SetAckLevel indicates an expected call of SetAckLevel.
func (mr *MockAckManagerMockRecorder) SetAckLevel(ackLevel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAckLevel", reflect.TypeOf((*MockAckManager)(nil).SetAckLevel), ackLevel)
}

// SetReadLevel mocks base method.
func (m *MockAckManager) SetReadLevel(readLevel int64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReadLevel", readLevel)
}

// SetReadLevel indicates an expected call of SetReadLevel.
func (mr *MockAckManagerMockRecorder) SetReadLevel(readLevel any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReadLevel", reflect.TypeOf((*MockAckManager)(nil).SetReadLevel), readLevel)
}
