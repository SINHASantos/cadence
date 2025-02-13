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
// Source: pprof.go
//
// Generated by this command:
//
//	mockgen -package common -source pprof.go -destination pprof_mock.go -package common github.com/uber/cadence/common PProfInitializer
//

// Package common is a generated GoMock package.
package common

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockPProfInitializer is a mock of PProfInitializer interface.
type MockPProfInitializer struct {
	ctrl     *gomock.Controller
	recorder *MockPProfInitializerMockRecorder
	isgomock struct{}
}

// MockPProfInitializerMockRecorder is the mock recorder for MockPProfInitializer.
type MockPProfInitializerMockRecorder struct {
	mock *MockPProfInitializer
}

// NewMockPProfInitializer creates a new mock instance.
func NewMockPProfInitializer(ctrl *gomock.Controller) *MockPProfInitializer {
	mock := &MockPProfInitializer{ctrl: ctrl}
	mock.recorder = &MockPProfInitializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPProfInitializer) EXPECT() *MockPProfInitializerMockRecorder {
	return m.recorder
}

// Start mocks base method.
func (m *MockPProfInitializer) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockPProfInitializerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockPProfInitializer)(nil).Start))
}
