// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package deftype

import (
	"time"

	"github.com/erda-project/erda/apistructs"
)

type ProjectPipelineListExecHistory struct {
	Name      string
	Executors []string
	AppNames  []string
	Statuses  []string

	PageNo         uint64
	PageSize       uint64
	StartTimeBegin time.Time
	StartTimeEnd   time.Time

	DescCols []string
	AscCols  []string

	ProjectID uint64
	apistructs.IdentityInfo
}

type ProjectPipelineListExecHistoryResult struct {
	Data *apistructs.PipelinePageListData
}
