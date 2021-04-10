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
	"github.com/erda-project/erda/modules/pipeline/pipengine/reconciler/queuemanage/types"
)

// HandleAllQueues handle queues in order.
func (mgr *defaultManager) HandleAllQueues() {
	mgr.qLock.Lock()
	defer mgr.qLock.Unlock()

	// iterate all queues
	for _, queue := range mgr.queues {

		// every queue handle in one goroutine
		go func(queue types.Queue) {
			mgr.handleOneQueue(queue)
		}(queue)
	}
}

// handleOneQueue handle all pipelines inside this queue.
func (mgr *defaultManager) handleOneQueue(queue types.Queue) {
	queue.RangePendingQueue(mgr)
}
