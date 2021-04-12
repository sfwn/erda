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

package queue

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/pipengine/reconciler/queuemanage/types"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

type CapacityValidator struct {
	mgr types.QueueManager
}

func (q *defaultQueue) ValidateCapacity(pipelineCaches map[uint64]*spec.Pipeline, tryPopP *spec.Pipeline) apistructs.PipelineQueueValidateResult {
	if int64(q.eq.ProcessingQueue().Len()) >= q.pq.Concurrency {
		return apistructs.PipelineQueueValidateResult{
			Success: false,
			Reason:  "Insufficient processing concurrency",
		}
	}
	return types.SuccessValidateResult
}
