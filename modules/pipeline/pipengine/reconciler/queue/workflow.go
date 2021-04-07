package queue

import (
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/modules/pipeline/aop"
	"github.com/erda-project/erda/modules/pipeline/aop/aoptypes"
	"github.com/erda-project/erda/modules/pipeline/pipengine/queue/priorityqueue"
)

// handleOneQueue handle all pipelines inside this queue.
func (mgr *Manager) handleOneQueue(queue *Queue) {
	queue.eq.PendingQueue().Range(func(item priorityqueue.Item) (stopRange bool) {
		// item key is pipeline id
		pipelineID, err := strconv.ParseUint(item.Key(), 10, 64)
		if err != nil {
			stopRange = false
			logrus.Errorf("failed to parse pipelineID from item key: %s, err: %v", item.Key(), err)
			// remove this invalid item
			queue.eq.PendingQueue().Remove(item.Key())
			return
		}

		// get pipeline
		p, exist, err := mgr.getPipeline(pipelineID)
		if err != nil {
			stopRange = false
			logrus.Error(err)
			// retry next time
			return
		}
		if !exist {
			// not exist, remove this invalid item
			queue.eq.PendingQueue().Remove(item.Key())
			return
		}

		// precheck before run
		ctx := aop.NewContextForPipeline(*p, aoptypes.TuneTriggerPipelineBeforeExecPrecheck)
		aopErr := aop.Handle(ctx)
		if aopErr != nil {
			// retry next time
			stopRange = false
			logrus.Error(aopErr)
			return
		}
		// TODO get check detail from ctx
		//ctx.TryGet()
		checkResult, ok := ctx.TryGet("check_result")
		if !ok {
			// no result, retry next time
			stopRange = false
			return
		}
		switch checkResult {
		case "success":
			// pop now
			queue.eq.PopPending()
			queue.doneChanByPipelineID[]
		default:
			// not pop
			// according to queue mode, check next pipeline or not
			if queue.IsStrictMode() {
				// strict mode, no need try next
				stopRange = true
				return
			}
			// loose mode
			stopRange = false
			return
		}

		stopRange = false
		return
	})
}
