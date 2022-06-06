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

package addPipelineBtn

import (
	"context"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cpregister/base"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda/internal/apps/dop/component-protocol/components/project-pipeline/common"
)

func init() {
	base.InitProviderWithCreator(common.ScenarioKey, "addPipelineBtn", func() servicehub.Provider {
		return &AddPipelineBtn{}
	})
}

func (a *AddPipelineBtn) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	if err := a.InitFromProtocol(ctx, c); err != nil {
		return err
	}
	a.SetType()
	a.SetName()
	a.SetProps(ctx)
	return a.SetToProtocolComponent(c)
}
