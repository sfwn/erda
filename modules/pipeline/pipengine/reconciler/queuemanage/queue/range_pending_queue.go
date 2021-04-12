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
	"strconv"

	"github.com/erda-project/erda/apistructs"

	"github.com/erda-project/erda/modules/pipeline/aop"
	"github.com/erda-project/erda/modules/pipeline/aop/aoptypes"
	"github.com/erda-project/erda/modules/pipeline/pipengine/queue/priorityqueue"
	"github.com/erda-project/erda/modules/pipeline/pipengine/reconciler/queuemanage/types"
	"github.com/erda-project/erda/modules/pipeline/pipengine/reconciler/rlog"
)

func (q *defaultQueue) RangePendingQueue(mgr types.QueueManager) {
	q.eq.PendingQueue().Range(func(item priorityqueue.Item) (stopRange bool) {

		pipelineID := parsePipelineIDFromQueueItem(item)

		// get pipeline
		p := mgr.EnsureQueryPipelineDetail(pipelineID)
		if p == nil {
			// pipeline not exist, remove this invalid item, continue handle next pipeline inside the queue
			stopRange = false
			rlog.PWarnf(pipelineID, "queueManager: failed to handle pipeline inside queue, pipeline not exist, pop from pending queue")
			q.eq.PendingQueue().Remove(item.Key())
			return
		}

		// queue validate
		validateResult := mgr.ValidatePipeline(q, p)
		if !validateResult.Success {
			rlog.PDebugf(pipelineID, "queueManager: failed to validate pipeline before pop from pending queue, reason: %s", validateResult.Reason)
			// stopRange if queue is strict mode
			return q.IsStrictMode()
		}

		// precheck before run
		customKVsOfAOP := map[interface{}]interface{}{}
		ctx := aop.NewContextForPipeline(*p, aoptypes.TuneTriggerPipelineInQueuePrecheckBeforePop, customKVsOfAOP)
		_ = aop.Handle(ctx)
		checkResultI, ok := ctx.TryGet("precheck_result")
		if !ok {
			// no result, pop now
			stopRange = true
			return
		}
		checkResult, ok := checkResultI.(apistructs.PipelineQueueValidateResult)
		if !ok {
			// invalid result, log and pop now
			rlog.PWarnf(pipelineID, "queue precheck result type is not expected, detail: %#v", checkResult)
			stopRange = true
			return
		}
		// check result
		if !checkResult.Success {
			// not success
			// according to queue mode, check next pipeline or skip
			return q.IsStrictMode()
		}
		// do pop
		stopRange = q.doPop(item)
		return
	})
}

func (q *defaultQueue) doPop(item priorityqueue.Item) (stopRange bool) {
	// pop now
	poppedKey := q.eq.PopPendingKey(item.Key())
	// queue cannot pop item anymore
	if poppedKey == "" {
		stopRange = true
		return
	}
	// send popped signal to channel
	pipelineID, _ := strconv.ParseUint(item.Key(), 10, 64)
	ch, ok := q.doneChanByPipelineID[pipelineID]
	if ok {
		ch <- struct{}{}
		close(ch)
		delete(q.doneChanByPipelineID, pipelineID)
	}
	// according to queue mode, check next pipeline or not
	return q.IsStrictMode()
}
