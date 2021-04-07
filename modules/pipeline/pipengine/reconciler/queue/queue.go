package queue

import (
	"strconv"
	"time"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/pipengine/queue/enhancedqueue"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

// Queue
type Queue struct {
	// pq is original pipeline queue.
	pq *apistructs.PipelineQueue

	// eq is enhanced priority queue, transfer from pq.
	eq *enhancedqueue.EnhancedQueue

	// doneChannels
	doneChanByPipelineID map[uint64]chan struct{}
}

// IsStrictMode
func (q *Queue) IsStrictMode() bool {
	return q.pq.Mode == apistructs.PipelineQueueModeStrict
}

// AddPipelineIntoQueue add pipeline into queue.
func (q *Queue) AddPipelineIntoQueue(p *spec.Pipeline) {
	// use pipeline id as key
	key := strconv.FormatUint(p.ID, 10)
	// use custom priority
	priority := q.pq.Priority
	customPriority, err := strconv.ParseInt(p.GetLabel(apistructs.LabelBindPipelineQueueInsidePriority), 10, 64)
	if err == nil {
		priority = customPriority
	}
	// createdTime
	createdTime := p.TimeCreated
	if createdTime == nil {
		now := time.Now()
		createdTime = &now
	}

	q.eq.Add(key, priority, *createdTime)
}
