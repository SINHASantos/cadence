// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package task

import (
	"fmt"

	"github.com/uber/cadence/common/backoff"
	"github.com/uber/cadence/common/dynamicconfig"
)

// WeightedRoundRobinTaskSchedulerOptions configs WRR task scheduler
type WeightedRoundRobinTaskSchedulerOptions[K comparable] struct {
	QueueSize            int
	WorkerCount          dynamicconfig.IntPropertyFn
	DispatcherCount      int
	RetryPolicy          backoff.RetryPolicy
	TaskToChannelKeyFn   func(PriorityTask) K
	ChannelKeyToWeightFn func(K) int
}

func (o *WeightedRoundRobinTaskSchedulerOptions[K]) String() string {
	return fmt.Sprintf("{QueueSize: %v, WorkerCount: %v, DispatcherCount: %v}", o.QueueSize, o.WorkerCount(), o.DispatcherCount)
}
