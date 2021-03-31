package queuemanage

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/services/apierrors"
)

// UpdatePipelineQueue update queue all fields by id.
func (qm *QueueManage) UpdatePipelineQueue(req apistructs.PipelineQueueUpdateRequest) (*apistructs.PipelineQueue, error) {
	// validate
	if err := req.Validate(); err != nil {
		return nil, apierrors.ErrUpdatePipelineQueue.InvalidParameter(err)
	}

	// do update
	queue, err := qm.dbClient.UpdatePipelineQueue(req)
	if err != nil {
		return nil, apierrors.ErrUpdatePipelineQueue.InternalError(err)
	}

	return queue, nil
}
