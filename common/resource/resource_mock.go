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
// Source: types.go
//
// Generated by this command:
//
//	mockgen -package resource -source types.go -destination resource_mock.go -self_package github.com/uber/cadence/common/resource
//

// Package resource is a generated GoMock package.
package resource

import (
	reflect "reflect"

	workflowserviceclient "go.uber.org/cadence/.gen/go/cadence/workflowserviceclient"
	gomock "go.uber.org/mock/gomock"
	yarpc "go.uber.org/yarpc"

	client "github.com/uber/cadence/client"
	admin "github.com/uber/cadence/client/admin"
	frontend "github.com/uber/cadence/client/frontend"
	history "github.com/uber/cadence/client/history"
	matching "github.com/uber/cadence/client/matching"
	archiver "github.com/uber/cadence/common/archiver"
	provider "github.com/uber/cadence/common/archiver/provider"
	queue "github.com/uber/cadence/common/asyncworkflow/queue"
	blobstore "github.com/uber/cadence/common/blobstore"
	cache "github.com/uber/cadence/common/cache"
	clock "github.com/uber/cadence/common/clock"
	cluster "github.com/uber/cadence/common/cluster"
	domain "github.com/uber/cadence/common/domain"
	configstore "github.com/uber/cadence/common/dynamicconfig/configstore"
	isolationgroup "github.com/uber/cadence/common/isolationgroup"
	log "github.com/uber/cadence/common/log"
	membership "github.com/uber/cadence/common/membership"
	messaging "github.com/uber/cadence/common/messaging"
	metrics "github.com/uber/cadence/common/metrics"
	partition "github.com/uber/cadence/common/partition"
	persistence "github.com/uber/cadence/common/persistence"
	client0 "github.com/uber/cadence/common/persistence/client"
	rpc "github.com/uber/cadence/common/quotas/global/rpc"
	service "github.com/uber/cadence/common/service"
)

// MockResourceFactory is a mock of ResourceFactory interface.
type MockResourceFactory struct {
	ctrl     *gomock.Controller
	recorder *MockResourceFactoryMockRecorder
	isgomock struct{}
}

// MockResourceFactoryMockRecorder is the mock recorder for MockResourceFactory.
type MockResourceFactoryMockRecorder struct {
	mock *MockResourceFactory
}

// NewMockResourceFactory creates a new mock instance.
func NewMockResourceFactory(ctrl *gomock.Controller) *MockResourceFactory {
	mock := &MockResourceFactory{ctrl: ctrl}
	mock.recorder = &MockResourceFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResourceFactory) EXPECT() *MockResourceFactoryMockRecorder {
	return m.recorder
}

// NewResource mocks base method.
func (m *MockResourceFactory) NewResource(params *Params, serviceName string, serviceConfig *service.Config) (Resource, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewResource", params, serviceName, serviceConfig)
	ret0, _ := ret[0].(Resource)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewResource indicates an expected call of NewResource.
func (mr *MockResourceFactoryMockRecorder) NewResource(params, serviceName, serviceConfig any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewResource", reflect.TypeOf((*MockResourceFactory)(nil).NewResource), params, serviceName, serviceConfig)
}

// MockResource is a mock of Resource interface.
type MockResource struct {
	ctrl     *gomock.Controller
	recorder *MockResourceMockRecorder
	isgomock struct{}
}

// MockResourceMockRecorder is the mock recorder for MockResource.
type MockResourceMockRecorder struct {
	mock *MockResource
}

// NewMockResource creates a new mock instance.
func NewMockResource(ctrl *gomock.Controller) *MockResource {
	mock := &MockResource{ctrl: ctrl}
	mock.recorder = &MockResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResource) EXPECT() *MockResourceMockRecorder {
	return m.recorder
}

// GetArchivalMetadata mocks base method.
func (m *MockResource) GetArchivalMetadata() archiver.ArchivalMetadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchivalMetadata")
	ret0, _ := ret[0].(archiver.ArchivalMetadata)
	return ret0
}

// GetArchivalMetadata indicates an expected call of GetArchivalMetadata.
func (mr *MockResourceMockRecorder) GetArchivalMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchivalMetadata", reflect.TypeOf((*MockResource)(nil).GetArchivalMetadata))
}

// GetArchiverProvider mocks base method.
func (m *MockResource) GetArchiverProvider() provider.ArchiverProvider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetArchiverProvider")
	ret0, _ := ret[0].(provider.ArchiverProvider)
	return ret0
}

// GetArchiverProvider indicates an expected call of GetArchiverProvider.
func (mr *MockResourceMockRecorder) GetArchiverProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetArchiverProvider", reflect.TypeOf((*MockResource)(nil).GetArchiverProvider))
}

// GetAsyncWorkflowQueueProvider mocks base method.
func (m *MockResource) GetAsyncWorkflowQueueProvider() queue.Provider {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAsyncWorkflowQueueProvider")
	ret0, _ := ret[0].(queue.Provider)
	return ret0
}

// GetAsyncWorkflowQueueProvider indicates an expected call of GetAsyncWorkflowQueueProvider.
func (mr *MockResourceMockRecorder) GetAsyncWorkflowQueueProvider() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAsyncWorkflowQueueProvider", reflect.TypeOf((*MockResource)(nil).GetAsyncWorkflowQueueProvider))
}

// GetBlobstoreClient mocks base method.
func (m *MockResource) GetBlobstoreClient() blobstore.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBlobstoreClient")
	ret0, _ := ret[0].(blobstore.Client)
	return ret0
}

// GetBlobstoreClient indicates an expected call of GetBlobstoreClient.
func (mr *MockResourceMockRecorder) GetBlobstoreClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBlobstoreClient", reflect.TypeOf((*MockResource)(nil).GetBlobstoreClient))
}

// GetClientBean mocks base method.
func (m *MockResource) GetClientBean() client.Bean {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientBean")
	ret0, _ := ret[0].(client.Bean)
	return ret0
}

// GetClientBean indicates an expected call of GetClientBean.
func (mr *MockResourceMockRecorder) GetClientBean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientBean", reflect.TypeOf((*MockResource)(nil).GetClientBean))
}

// GetClusterMetadata mocks base method.
func (m *MockResource) GetClusterMetadata() cluster.Metadata {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterMetadata")
	ret0, _ := ret[0].(cluster.Metadata)
	return ret0
}

// GetClusterMetadata indicates an expected call of GetClusterMetadata.
func (mr *MockResourceMockRecorder) GetClusterMetadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterMetadata", reflect.TypeOf((*MockResource)(nil).GetClusterMetadata))
}

// GetDispatcher mocks base method.
func (m *MockResource) GetDispatcher() *yarpc.Dispatcher {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDispatcher")
	ret0, _ := ret[0].(*yarpc.Dispatcher)
	return ret0
}

// GetDispatcher indicates an expected call of GetDispatcher.
func (mr *MockResourceMockRecorder) GetDispatcher() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDispatcher", reflect.TypeOf((*MockResource)(nil).GetDispatcher))
}

// GetDomainCache mocks base method.
func (m *MockResource) GetDomainCache() cache.DomainCache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainCache")
	ret0, _ := ret[0].(cache.DomainCache)
	return ret0
}

// GetDomainCache indicates an expected call of GetDomainCache.
func (mr *MockResourceMockRecorder) GetDomainCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainCache", reflect.TypeOf((*MockResource)(nil).GetDomainCache))
}

// GetDomainManager mocks base method.
func (m *MockResource) GetDomainManager() persistence.DomainManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainManager")
	ret0, _ := ret[0].(persistence.DomainManager)
	return ret0
}

// GetDomainManager indicates an expected call of GetDomainManager.
func (mr *MockResourceMockRecorder) GetDomainManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainManager", reflect.TypeOf((*MockResource)(nil).GetDomainManager))
}

// GetDomainMetricsScopeCache mocks base method.
func (m *MockResource) GetDomainMetricsScopeCache() cache.DomainMetricsScopeCache {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainMetricsScopeCache")
	ret0, _ := ret[0].(cache.DomainMetricsScopeCache)
	return ret0
}

// GetDomainMetricsScopeCache indicates an expected call of GetDomainMetricsScopeCache.
func (mr *MockResourceMockRecorder) GetDomainMetricsScopeCache() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainMetricsScopeCache", reflect.TypeOf((*MockResource)(nil).GetDomainMetricsScopeCache))
}

// GetDomainReplicationQueue mocks base method.
func (m *MockResource) GetDomainReplicationQueue() domain.ReplicationQueue {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDomainReplicationQueue")
	ret0, _ := ret[0].(domain.ReplicationQueue)
	return ret0
}

// GetDomainReplicationQueue indicates an expected call of GetDomainReplicationQueue.
func (mr *MockResourceMockRecorder) GetDomainReplicationQueue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDomainReplicationQueue", reflect.TypeOf((*MockResource)(nil).GetDomainReplicationQueue))
}

// GetExecutionManager mocks base method.
func (m *MockResource) GetExecutionManager(arg0 int) (persistence.ExecutionManager, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExecutionManager", arg0)
	ret0, _ := ret[0].(persistence.ExecutionManager)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExecutionManager indicates an expected call of GetExecutionManager.
func (mr *MockResourceMockRecorder) GetExecutionManager(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExecutionManager", reflect.TypeOf((*MockResource)(nil).GetExecutionManager), arg0)
}

// GetFrontendClient mocks base method.
func (m *MockResource) GetFrontendClient() frontend.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrontendClient")
	ret0, _ := ret[0].(frontend.Client)
	return ret0
}

// GetFrontendClient indicates an expected call of GetFrontendClient.
func (mr *MockResourceMockRecorder) GetFrontendClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrontendClient", reflect.TypeOf((*MockResource)(nil).GetFrontendClient))
}

// GetFrontendRawClient mocks base method.
func (m *MockResource) GetFrontendRawClient() frontend.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFrontendRawClient")
	ret0, _ := ret[0].(frontend.Client)
	return ret0
}

// GetFrontendRawClient indicates an expected call of GetFrontendRawClient.
func (mr *MockResourceMockRecorder) GetFrontendRawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFrontendRawClient", reflect.TypeOf((*MockResource)(nil).GetFrontendRawClient))
}

// GetHistoryClient mocks base method.
func (m *MockResource) GetHistoryClient() history.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryClient")
	ret0, _ := ret[0].(history.Client)
	return ret0
}

// GetHistoryClient indicates an expected call of GetHistoryClient.
func (mr *MockResourceMockRecorder) GetHistoryClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryClient", reflect.TypeOf((*MockResource)(nil).GetHistoryClient))
}

// GetHistoryManager mocks base method.
func (m *MockResource) GetHistoryManager() persistence.HistoryManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryManager")
	ret0, _ := ret[0].(persistence.HistoryManager)
	return ret0
}

// GetHistoryManager indicates an expected call of GetHistoryManager.
func (mr *MockResourceMockRecorder) GetHistoryManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryManager", reflect.TypeOf((*MockResource)(nil).GetHistoryManager))
}

// GetHistoryRawClient mocks base method.
func (m *MockResource) GetHistoryRawClient() history.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistoryRawClient")
	ret0, _ := ret[0].(history.Client)
	return ret0
}

// GetHistoryRawClient indicates an expected call of GetHistoryRawClient.
func (mr *MockResourceMockRecorder) GetHistoryRawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistoryRawClient", reflect.TypeOf((*MockResource)(nil).GetHistoryRawClient))
}

// GetHostInfo mocks base method.
func (m *MockResource) GetHostInfo() membership.HostInfo {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostInfo")
	ret0, _ := ret[0].(membership.HostInfo)
	return ret0
}

// GetHostInfo indicates an expected call of GetHostInfo.
func (mr *MockResourceMockRecorder) GetHostInfo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostInfo", reflect.TypeOf((*MockResource)(nil).GetHostInfo))
}

// GetHostName mocks base method.
func (m *MockResource) GetHostName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHostName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetHostName indicates an expected call of GetHostName.
func (mr *MockResourceMockRecorder) GetHostName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHostName", reflect.TypeOf((*MockResource)(nil).GetHostName))
}

// GetIsolationGroupState mocks base method.
func (m *MockResource) GetIsolationGroupState() isolationgroup.State {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIsolationGroupState")
	ret0, _ := ret[0].(isolationgroup.State)
	return ret0
}

// GetIsolationGroupState indicates an expected call of GetIsolationGroupState.
func (mr *MockResourceMockRecorder) GetIsolationGroupState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsolationGroupState", reflect.TypeOf((*MockResource)(nil).GetIsolationGroupState))
}

// GetIsolationGroupStore mocks base method.
func (m *MockResource) GetIsolationGroupStore() configstore.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIsolationGroupStore")
	ret0, _ := ret[0].(configstore.Client)
	return ret0
}

// GetIsolationGroupStore indicates an expected call of GetIsolationGroupStore.
func (mr *MockResourceMockRecorder) GetIsolationGroupStore() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIsolationGroupStore", reflect.TypeOf((*MockResource)(nil).GetIsolationGroupStore))
}

// GetLogger mocks base method.
func (m *MockResource) GetLogger() log.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLogger")
	ret0, _ := ret[0].(log.Logger)
	return ret0
}

// GetLogger indicates an expected call of GetLogger.
func (mr *MockResourceMockRecorder) GetLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLogger", reflect.TypeOf((*MockResource)(nil).GetLogger))
}

// GetMatchingClient mocks base method.
func (m *MockResource) GetMatchingClient() matching.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchingClient")
	ret0, _ := ret[0].(matching.Client)
	return ret0
}

// GetMatchingClient indicates an expected call of GetMatchingClient.
func (mr *MockResourceMockRecorder) GetMatchingClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchingClient", reflect.TypeOf((*MockResource)(nil).GetMatchingClient))
}

// GetMatchingRawClient mocks base method.
func (m *MockResource) GetMatchingRawClient() matching.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMatchingRawClient")
	ret0, _ := ret[0].(matching.Client)
	return ret0
}

// GetMatchingRawClient indicates an expected call of GetMatchingRawClient.
func (mr *MockResourceMockRecorder) GetMatchingRawClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMatchingRawClient", reflect.TypeOf((*MockResource)(nil).GetMatchingRawClient))
}

// GetMembershipResolver mocks base method.
func (m *MockResource) GetMembershipResolver() membership.Resolver {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMembershipResolver")
	ret0, _ := ret[0].(membership.Resolver)
	return ret0
}

// GetMembershipResolver indicates an expected call of GetMembershipResolver.
func (mr *MockResourceMockRecorder) GetMembershipResolver() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMembershipResolver", reflect.TypeOf((*MockResource)(nil).GetMembershipResolver))
}

// GetMessagingClient mocks base method.
func (m *MockResource) GetMessagingClient() messaging.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessagingClient")
	ret0, _ := ret[0].(messaging.Client)
	return ret0
}

// GetMessagingClient indicates an expected call of GetMessagingClient.
func (mr *MockResourceMockRecorder) GetMessagingClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessagingClient", reflect.TypeOf((*MockResource)(nil).GetMessagingClient))
}

// GetMetricsClient mocks base method.
func (m *MockResource) GetMetricsClient() metrics.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetricsClient")
	ret0, _ := ret[0].(metrics.Client)
	return ret0
}

// GetMetricsClient indicates an expected call of GetMetricsClient.
func (mr *MockResourceMockRecorder) GetMetricsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetricsClient", reflect.TypeOf((*MockResource)(nil).GetMetricsClient))
}

// GetPartitioner mocks base method.
func (m *MockResource) GetPartitioner() partition.Partitioner {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartitioner")
	ret0, _ := ret[0].(partition.Partitioner)
	return ret0
}

// GetPartitioner indicates an expected call of GetPartitioner.
func (mr *MockResourceMockRecorder) GetPartitioner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartitioner", reflect.TypeOf((*MockResource)(nil).GetPartitioner))
}

// GetPayloadSerializer mocks base method.
func (m *MockResource) GetPayloadSerializer() persistence.PayloadSerializer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPayloadSerializer")
	ret0, _ := ret[0].(persistence.PayloadSerializer)
	return ret0
}

// GetPayloadSerializer indicates an expected call of GetPayloadSerializer.
func (mr *MockResourceMockRecorder) GetPayloadSerializer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPayloadSerializer", reflect.TypeOf((*MockResource)(nil).GetPayloadSerializer))
}

// GetPersistenceBean mocks base method.
func (m *MockResource) GetPersistenceBean() client0.Bean {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPersistenceBean")
	ret0, _ := ret[0].(client0.Bean)
	return ret0
}

// GetPersistenceBean indicates an expected call of GetPersistenceBean.
func (mr *MockResourceMockRecorder) GetPersistenceBean() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPersistenceBean", reflect.TypeOf((*MockResource)(nil).GetPersistenceBean))
}

// GetRatelimiterAggregatorsClient mocks base method.
func (m *MockResource) GetRatelimiterAggregatorsClient() rpc.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRatelimiterAggregatorsClient")
	ret0, _ := ret[0].(rpc.Client)
	return ret0
}

// GetRatelimiterAggregatorsClient indicates an expected call of GetRatelimiterAggregatorsClient.
func (mr *MockResourceMockRecorder) GetRatelimiterAggregatorsClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRatelimiterAggregatorsClient", reflect.TypeOf((*MockResource)(nil).GetRatelimiterAggregatorsClient))
}

// GetRemoteAdminClient mocks base method.
func (m *MockResource) GetRemoteAdminClient(cluster string) admin.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteAdminClient", cluster)
	ret0, _ := ret[0].(admin.Client)
	return ret0
}

// GetRemoteAdminClient indicates an expected call of GetRemoteAdminClient.
func (mr *MockResourceMockRecorder) GetRemoteAdminClient(cluster any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteAdminClient", reflect.TypeOf((*MockResource)(nil).GetRemoteAdminClient), cluster)
}

// GetRemoteFrontendClient mocks base method.
func (m *MockResource) GetRemoteFrontendClient(cluster string) frontend.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRemoteFrontendClient", cluster)
	ret0, _ := ret[0].(frontend.Client)
	return ret0
}

// GetRemoteFrontendClient indicates an expected call of GetRemoteFrontendClient.
func (mr *MockResourceMockRecorder) GetRemoteFrontendClient(cluster any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRemoteFrontendClient", reflect.TypeOf((*MockResource)(nil).GetRemoteFrontendClient), cluster)
}

// GetSDKClient mocks base method.
func (m *MockResource) GetSDKClient() workflowserviceclient.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSDKClient")
	ret0, _ := ret[0].(workflowserviceclient.Interface)
	return ret0
}

// GetSDKClient indicates an expected call of GetSDKClient.
func (mr *MockResourceMockRecorder) GetSDKClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSDKClient", reflect.TypeOf((*MockResource)(nil).GetSDKClient))
}

// GetServiceName mocks base method.
func (m *MockResource) GetServiceName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetServiceName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetServiceName indicates an expected call of GetServiceName.
func (mr *MockResourceMockRecorder) GetServiceName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetServiceName", reflect.TypeOf((*MockResource)(nil).GetServiceName))
}

// GetShardManager mocks base method.
func (m *MockResource) GetShardManager() persistence.ShardManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShardManager")
	ret0, _ := ret[0].(persistence.ShardManager)
	return ret0
}

// GetShardManager indicates an expected call of GetShardManager.
func (mr *MockResourceMockRecorder) GetShardManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShardManager", reflect.TypeOf((*MockResource)(nil).GetShardManager))
}

// GetTaskManager mocks base method.
func (m *MockResource) GetTaskManager() persistence.TaskManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskManager")
	ret0, _ := ret[0].(persistence.TaskManager)
	return ret0
}

// GetTaskManager indicates an expected call of GetTaskManager.
func (mr *MockResourceMockRecorder) GetTaskManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskManager", reflect.TypeOf((*MockResource)(nil).GetTaskManager))
}

// GetThrottledLogger mocks base method.
func (m *MockResource) GetThrottledLogger() log.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetThrottledLogger")
	ret0, _ := ret[0].(log.Logger)
	return ret0
}

// GetThrottledLogger indicates an expected call of GetThrottledLogger.
func (mr *MockResourceMockRecorder) GetThrottledLogger() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetThrottledLogger", reflect.TypeOf((*MockResource)(nil).GetThrottledLogger))
}

// GetTimeSource mocks base method.
func (m *MockResource) GetTimeSource() clock.TimeSource {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimeSource")
	ret0, _ := ret[0].(clock.TimeSource)
	return ret0
}

// GetTimeSource indicates an expected call of GetTimeSource.
func (mr *MockResourceMockRecorder) GetTimeSource() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimeSource", reflect.TypeOf((*MockResource)(nil).GetTimeSource))
}

// GetVisibilityManager mocks base method.
func (m *MockResource) GetVisibilityManager() persistence.VisibilityManager {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVisibilityManager")
	ret0, _ := ret[0].(persistence.VisibilityManager)
	return ret0
}

// GetVisibilityManager indicates an expected call of GetVisibilityManager.
func (mr *MockResourceMockRecorder) GetVisibilityManager() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVisibilityManager", reflect.TypeOf((*MockResource)(nil).GetVisibilityManager))
}

// Start mocks base method.
func (m *MockResource) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start.
func (mr *MockResourceMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockResource)(nil).Start))
}

// Stop mocks base method.
func (m *MockResource) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockResourceMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockResource)(nil).Stop))
}
