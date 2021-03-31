package queuemanage

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/services/apierrors"
)

// PagingPipelineQueues
func (qm *QueueManage) PagingPipelineQueues(req apistructs.PipelineQueuePagingRequest) (*apistructs.PipelineQueuePagingData, error) {
	pagingResult, err := qm.dbClient.PagingPipelineQueues(req)
	if err != nil {
		return nil, apierrors.ErrPagingPipelineQueues.InternalError(err)
	}
	return pagingResult, nil
}
