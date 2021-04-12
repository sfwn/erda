// Copyright (c) 2021 Terminus, Inc.
//
// This program is free software: you can use, redistribute, and/or modify
// it under the terms of the GNU Affero General Public License, version 3
// or later ("AGPL"), as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package manager

import (
	"sync"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/dbclient"
	"github.com/erda-project/erda/modules/pipeline/pipengine/reconciler/queuemanage/types"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

// defaultManager is the default manager.
type defaultManager struct {
	// queues holds all queues order by priority.
	queues    []types.Queue
	queueByID map[uint64]types.Queue // key: pq id
	qLock     sync.Mutex

	pipelineCaches map[uint64]*spec.Pipeline
	pCacheLock     sync.Mutex

	dbClient *dbclient.Client
}

func (mgr *defaultManager) ValidatePipeline(queue types.Queue, p *spec.Pipeline) apistructs.PipelineQueueValidateResult {
	// capacity
	result := queue.ValidateCapacity(mgr.pipelineCaches, p)
	if result.IsFailed() {
		return result
	}
	// free resources
	result = queue.ValidateFreeResources(mgr.pipelineCaches, p)
	if result.IsFailed() {
		return result
	}

	// default result
	return types.SuccessValidateResult
}

// New return a new queue manager.
func New(ops ...Option) types.QueueManager {
	var mgr defaultManager

	mgr.queues = []types.Queue{}
	mgr.queueByID = make(map[uint64]types.Queue)

	mgr.pipelineCaches = make(map[uint64]*spec.Pipeline)

	// apply options
	for _, op := range ops {
		op(&mgr)
	}

	return &mgr
}

type Option func(manager *defaultManager)

func WithDBClient(dbClient *dbclient.Client) Option {
	return func(mgr *defaultManager) {
		mgr.dbClient = dbClient
	}
}
