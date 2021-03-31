package queuemanage

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/services/apierrors"
)

// CreatePipelineQueue create a pipeline queue.
func (qm *QueueManage) CreatePipelineQueue(req apistructs.PipelineQueueCreateRequest) (*apistructs.PipelineQueue, error) {
	// validate
	if err := req.Validate(); err != nil {
		return nil, apierrors.ErrCreatePipelineQueue.InvalidParameter(err)
	}

	// create
	queue, err := qm.dbClient.CreatePipelineQueue(req)
	if err != nil {
		return nil, apierrors.ErrCreatePipelineQueue.InternalError(err)
	}

	return queue, nil
}
