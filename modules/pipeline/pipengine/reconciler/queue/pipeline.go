package queue

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/pipengine/queue/enhancedqueue"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

// addQueueToManager add one queue to mgr.
func (mgr *Manager) addQueueToManager(pipelineQueue *apistructs.PipelineQueue) {
	mgr.qLock.Lock()
	defer mgr.qLock.Unlock()

	queue := &Queue{
		pq: pipelineQueue,
		eq: enhancedqueue.NewEnhancedQueue(pipelineQueue.Concurrency),
	}
	mgr.Queues = append(mgr.Queues, queue)
	mgr.QueueByID[pipelineQueue.ID] = queue
}

// AddPipelineToQueue add pipeline to queue.
// return: beginCh, needRetryIfErr, err
func (mgr *Manager) AddPipelineToQueue(pipelineID uint64) <-chan struct{} {
	// channel: send done signal when pipeline pop from the queue.
	doneCh := make(chan struct{})

	// get pipeline related queue
	mgr.cacheLock.Lock()
	defer mgr.cacheLock.Unlock()
	var (
		needRetry     = true
		pipelineQueue *apistructs.PipelineQueue
		err           error
	)
	for needRetry {
		pipelineQueue, needRetry, err = mgr.getPipelineQueueDetail(pipelineID)
		if err != nil {
			logrus.Error(err)
		}
		time.Sleep(time.Second * 10)
	}
	// pipeline doesn't bind queue, continue reconcile
	if pipelineQueue == nil {
		defer func() {
			doneCh <- struct{}{}
		}()
		return doneCh
	}

	// add queue to manager
	mgr.addQueueToManager(pipelineQueue)

	// add pipeline to queue
	mgr.QueueByID[pipelineQueue.ID].AddPipelineIntoQueue(mgr.PipelineCaches[pipelineID])

	// return channel when pipeline pop from queue
	return doneCh
}

func (mgr *Manager) getPipeline(pipelineID uint64) (*spec.Pipeline, bool, error) {
	mgr.cacheLock.Lock()
	defer mgr.cacheLock.Unlock()

	p, ok := mgr.PipelineCaches[pipelineID]
	if ok {
		return p, true, nil
	}

	_p, exist, err := mgr.dbClient.GetPipelineWithExistInfo(pipelineID)
	if err != nil {
		return nil, false, fmt.Errorf("failed to get pipeline: %d, err: %v", pipelineID, err)
	}
	if !exist {
		return nil, false, nil
	}

	return &_p, true, nil
}

// getPipelineQueueDetail
// return: queue, needRetryIfErr, err
func (mgr *Manager) getPipelineQueueDetail(pipelineID uint64) (*apistructs.PipelineQueue, bool, error) {
	// get pipeline detail
	p, pExist, err := mgr.dbClient.GetPipelineWithExistInfo(pipelineID)
	if err != nil {
		return nil, true, fmt.Errorf("failed to get pipeline detail, pipelineID: %d, err: %v", pipelineID, err)
	}
	if !pExist {
		return nil, false, fmt.Errorf("pipeline not found, pipelineID: %d", pipelineID)
	}
	// add pipeline to cache
	mgr.PipelineCaches[pipelineID] = &p

	// get queue detail
	queueID, pExist := p.GetPipelineQueueID()
	queue, qExist, err := mgr.dbClient.GetPipelineQueue(queueID)
	if err != nil {
		return nil, true, fmt.Errorf("failed to get pipeline queue, pipelineID: %d, queueID: %d, err: %v", pipelineID, queueID, err)
	}
	if !qExist {
		return nil, false, nil
	}
	return queue, false, nil
}
