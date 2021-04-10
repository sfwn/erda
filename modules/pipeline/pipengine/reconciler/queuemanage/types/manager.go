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

package types

import (
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

// QueueManager manage all queues and related pipelines.
type QueueManager interface {
	PutPipelineIntoQueue(pipelineID uint64) (popCh <-chan struct{}, needRetryIfErr bool, err error)
	HandleAllQueues()
	Start()
	EnsureQueryPipelineDetail(pipelineID uint64) *spec.Pipeline
	ValidatePipeline(queue Queue, p *spec.Pipeline) apistructs.PipelineQueueValidateResult
}
