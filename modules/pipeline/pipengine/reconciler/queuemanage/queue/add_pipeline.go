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
	"strconv"
	"time"

	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/modules/pipeline/spec"
)

func (q *defaultQueue) AddPipelineIntoQueue(p *spec.Pipeline, doneCh chan struct{}) {
	q.lock.Lock()
	defer q.lock.Unlock()

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
	q.doneChanByPipelineID[p.ID] = doneCh
}

