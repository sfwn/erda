// 每个队列持有一定资源，目前不与真实资源联动，是用户期望的资源。
// 队列之间有优先级，它的用处是：所有队列中优先级最高的 pending 流水线，根据优先级先后进行处理。
package queue

import (
	"sync"
	"time"

	"github.com/erda-project/erda/modules/pipeline/dbclient"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

// Manager manage all queues.
type Manager struct {
	// Queues holds all queues order by priority.
	Queues    []*Queue
	QueueByID map[uint64]*Queue
	qLock     sync.Mutex

	PipelineCaches map[uint64]*spec.Pipeline
	cacheLock      sync.Mutex

	dbClient *dbclient.Client
}

// NewManager return a new queue manager.
func NewManager(ops ...Option) *Manager {
	var mgr Manager

	mgr.Queues = []*Queue{}
	mgr.QueueByID = make(map[uint64]*Queue)

	mgr.PipelineCaches = make(map[uint64]*spec.Pipeline)

	for _, op := range ops {
		op(&mgr)
	}

	return &mgr
}

type Option func(*Manager)

func WithDBClient(dbClient *dbclient.Client) Option {
	return func(mgr *Manager) {
		mgr.dbClient = dbClient
	}
}

// Start manger.
func (mgr *Manager) Start() {
	ticker := time.NewTicker(time.Second * 3)
	go func() {
		for {
			select {
			case <-ticker.C:
				// interval: 3s
				mgr.handleAllQueues()
			}
		}
	}()
}

// handleAllQueues
func (mgr *Manager) handleAllQueues() {
	mgr.qLock.Lock()
	defer mgr.qLock.Unlock()

	// iterate all queues
	for _, queue := range mgr.Queues {

		// every queue handle in one goroutine
		go func(queue *Queue) {
			mgr.handleOneQueue(queue)
		}(queue)
	}
}
