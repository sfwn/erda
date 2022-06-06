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

package workloadStatus

import (
	"context"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/internal/apps/cmp"
)

type ComponentWorkloadStatus struct {
	sdk    *cptype.SDK
	ctx    context.Context
	server cmp.SteveServer

	Type  string `json:"type,omitempty"`
	Props Props  `json:"props,omitempty"`
	State State  `json:"state,omitempty"`
}

type Props struct {
	Text      string `json:"text,omitempty"`
	Status    string `json:"status,omitempty"`
	Breathing bool   `json:"breathing"`
}

type Labels struct {
	Label string `json:"label,omitempty"`
	Color string `json:"color,omitempty"`
}

type State struct {
	ClusterName string `json:"clusterName,omitempty"`
	WorkloadID  string `json:"workloadId,omitempty"`
}
