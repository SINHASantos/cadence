// Copyright (c) 2019 Uber Technologies, Inc.
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

package main

import (
	"os"

	"github.com/uber/cadence/cmd/server/cadence"
	"github.com/uber/cadence/common/metrics"
	"github.com/uber/cadence/tools/common/commoncli"

	_ "github.com/uber/cadence/common/archiver/gcloud"                                      // needed to load the optional gcloud archiver plugin
	_ "github.com/uber/cadence/common/asyncworkflow/queue/kafka"                            // needed to load kafka asyncworkflow queue
	_ "github.com/uber/cadence/common/persistence/nosql/nosqlplugin/cassandra"              // needed to load cassandra plugin
	_ "github.com/uber/cadence/common/persistence/nosql/nosqlplugin/cassandra/gocql/public" // needed to load the default gocql client
	_ "github.com/uber/cadence/common/persistence/sql/sqlplugin/mysql"                      // needed to load mysql plugin
	_ "github.com/uber/cadence/common/persistence/sql/sqlplugin/postgres"                   // needed to load postgres plugin
	_ "github.com/uber/cadence/common/persistence/sql/sqlplugin/sqlite"                     // needed to load sqlite plugin
	_ "github.com/uber/cadence/service/sharddistributor/store/etcd"                         // needed for shard distributor shard/heartbeat and leader election
)

// main entry point for the cadence server
func main() {
	app := cadence.BuildCLI(metrics.ReleaseVersion, metrics.Revision)
	commoncli.ExitHandler(app.Run(os.Args))
}
