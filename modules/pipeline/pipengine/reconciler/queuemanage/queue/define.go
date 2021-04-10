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

package queue

import (
	"sync"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/pipengine/queue/enhancedqueue"
)

// defaultQueue is used to implement Queue.
type defaultQueue struct {
	// pq is original pipeline queue.
	pq *apistructs.PipelineQueue

	// eq is enhanced priority queue, transfer from pq.
	eq *enhancedqueue.EnhancedQueue

	// doneChannels
	doneChanByPipelineID map[uint64]chan struct{}

	lock sync.Mutex
}

func New(pq *apistructs.PipelineQueue) *defaultQueue {
	newQueue := defaultQueue{
		pq:                   pq,
		eq:                   enhancedqueue.NewEnhancedQueue(pq.Concurrency),
		doneChanByPipelineID: make(map[uint64]chan struct{}),
	}
	return &newQueue
}
