package timeout

// Code generated by gowrap. DO NOT EDIT.
// template: ../../templates/timeout.tmpl
// gowrap: http://github.com/hexdigest/gowrap

import (
	"context"
	"time"

	"go.uber.org/yarpc"

	"github.com/uber/cadence/client/sharddistributorexecutor"
	"github.com/uber/cadence/common/types"
)

var _ sharddistributorexecutor.Client = (*sharddistributorexecutorClient)(nil)

// sharddistributorexecutorClient implements the sharddistributorexecutor.Client interface instrumented with timeouts
type sharddistributorexecutorClient struct {
	client  sharddistributorexecutor.Client
	timeout time.Duration
}

// NewShardDistributorExecutorClient creates a new sharddistributorexecutorClient instance
func NewShardDistributorExecutorClient(
	client sharddistributorexecutor.Client,
	timeout time.Duration,
) sharddistributorexecutor.Client {
	return &sharddistributorexecutorClient{
		client:  client,
		timeout: timeout,
	}
}

func (c *sharddistributorexecutorClient) Heartbeat(ctx context.Context, ep1 *types.ExecutorHeartbeatRequest, p1 ...yarpc.CallOption) (ep2 *types.ExecutorHeartbeatResponse, err error) {
	ctx, cancel := createContext(ctx, c.timeout)
	defer cancel()
	return c.client.Heartbeat(ctx, ep1, p1...)
}
