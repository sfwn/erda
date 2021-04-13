package queuemanage

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/services/apierrors"
)

// GetPipelineQueue get a pipeline queue by id.
func (qm *QueueManage) GetPipelineQueue(queueID uint64) (*apistructs.PipelineQueue, error) {
	queue, exist, err := qm.dbClient.GetPipelineQueue(queueID)
	if err != nil {
		return nil, apierrors.ErrGetPipelineQueue.InternalError(err)
	}
	if !exist {
		return nil, apierrors.ErrGetPipelineQueue.NotFound()
	}
	return queue, nil
}

