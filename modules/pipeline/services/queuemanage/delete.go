package queuemanage

import (
	"github.com/erda-project/erda/modules/pipeline/services/apierrors"
)

// DeletePipelineQueue delete a pipeline queue by id.
func (qm *QueueManage) DeletePipelineQueue(queueID uint64) error {
	err := qm.dbClient.DeletePipelineQueue(queueID)
	if err != nil {
		return apierrors.ErrDeletePipelineQueue.InternalError(err)
	}
	return nil
}
