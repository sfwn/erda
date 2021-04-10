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
	"fmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/pipengine/reconciler/queuemanage/queue"
	"github.com/erda-project/erda/modules/pipeline/pipengine/reconciler/queuemanage/types"
	"github.com/erda-project/erda/modules/pipeline/spec"
	"github.com/erda-project/erda/pkg/loop"
)

// PutPipelineIntoQueue put pipeline into queue.
// return: popCh, needRetryIfErr, err
func (mgr *defaultManager) PutPipelineIntoQueue(pipelineID uint64) (<-chan struct{}, bool, error) {
	// channel: send done signal when pipeline pop from the queue.
	popCh := make(chan struct{})

	// query pipeline detail
	p := mgr.EnsureQueryPipelineDetail(pipelineID)
	if p == nil {
		return nil, false, fmt.Errorf("pipeline not found, pipelineID: %d", pipelineID)
	}

	// query pipeline queue detail
	pq := mgr.EnsureQueryPipelineQueueDetail(pipelineID)
	if pq == nil {
		// pipeline doesn't bind queue, can reconcile directly
		go func() {
			popCh <- struct{}{}
			close(popCh)
		}()
		return popCh, false, nil
	}

	// add queue to manager
	q := mgr.addQueueToManager(pq)

	// add pipeline to queue
	q.AddPipelineIntoQueue(p, popCh)

	// return channel when pipeline pop from queue
	return popCh, false, nil
}

// addQueueToManager add one queue to mgr.
func (mgr *defaultManager) addQueueToManager(pq *apistructs.PipelineQueue) types.Queue {
	mgr.qLock.Lock()
	defer mgr.qLock.Unlock()

	existQueue, ok := mgr.queueByID[pq.ID]
	if ok {
		// update queue
		existQueue.Update(pq)
		mgr.queueByID[pq.ID] = existQueue
		return existQueue
	}

	// not exist, construct and add new queue
	newQueue := queue.New(pq)
	mgr.queues = append(mgr.queues, newQueue)
	mgr.queueByID[pq.ID] = newQueue

	return newQueue
}

// EnsureQueryPipelineDetail handle err properly.
// return: pipeline or nil
func (mgr *defaultManager) EnsureQueryPipelineDetail(pipelineID uint64) *spec.Pipeline {
	mgr.pCacheLock.Lock()
	defer mgr.pCacheLock.Unlock()

	// try to get from pCache
	cachedP, ok := mgr.pipelineCaches[pipelineID]
	if ok {
		return cachedP
	}

	// query from db
	var p *spec.Pipeline
	_ = loop.New(loop.WithDeclineLimit(time.Second*10), loop.WithDeclineRatio(2)).Do(func() (abort bool, err error) {
		_p, exist, err := mgr.dbClient.GetPipelineWithExistInfo(pipelineID)
		if err != nil {
			err = fmt.Errorf("failed to query pipeline: %d, err: %v", pipelineID, err)
			logrus.Error(err)
			return false, err
		}
		if !exist {
			return true, nil
		}
		p = &_p
		return true, nil
	})
	if p == nil {
		return nil
	}

	// add into cache
	mgr.pipelineCaches[pipelineID] = p

	return p
}

// EnsureQueryPipelineQueueDetail
// return: queue or nil
func (mgr *defaultManager) EnsureQueryPipelineQueueDetail(pipelineID uint64) *apistructs.PipelineQueue {
	// get pipeline detail
	p := mgr.EnsureQueryPipelineDetail(pipelineID)
	if p == nil {
		return nil
	}

	// get queue id
	queueID, exist := p.GetPipelineQueueID()
	if !exist {
		return nil
	}

	// query from db
	var pq *apistructs.PipelineQueue
	_ = loop.New(loop.WithDeclineLimit(time.Second*10), loop.WithDeclineRatio(2)).Do(func() (abort bool, err error) {
		_pq, exist, err := mgr.dbClient.GetPipelineQueue(queueID)
		if err != nil {
			err = fmt.Errorf("failed to query pipeline queue, queueID: %d, err: %v", queueID, err)
			logrus.Error(err)
			return false, err
		}
		if !exist {
			return true, nil
		}
		pq = _pq
		return true, nil
	})
	if pq == nil {
		return nil
	}

	return pq
}
