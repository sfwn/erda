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

package at_num_api

import (
	"context"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/modules/dop/component-protocol/components/test-dashboard/common"
	"github.com/erda-project/erda/modules/dop/component-protocol/components/test-dashboard/common/gshelper"
	"github.com/erda-project/erda/modules/dop/component-protocol/components/test-dashboard/overview_group/blocks/at/pkg"
	"github.com/erda-project/erda/modules/dop/component-protocol/types"
	autotestv2 "github.com/erda-project/erda/modules/dop/services/autotest_v2"
	"github.com/erda-project/erda/modules/openapi/component-protocol/components/base"
	"github.com/erda-project/erda/pkg/strutil"
)

func init() {
	base.InitProviderWithCreator(common.ScenarioKeyTestDashboard, "at_num_api", func() servicehub.Provider {
		return &Text{}
	})
}

type Text struct {
	base.DefaultProvider
}

func (t *Text) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	h := gshelper.NewGSHelper(gs)
	atSvc := ctx.Value(types.AutoTestPlanService).(*autotestv2.Service)

	scenes, _, err := atSvc.ListSceneBySceneSetID(func() []uint64 {
		setIDs := make([]uint64, 0, len(h.GetBlockAtStep()))
		for _, v := range h.GetBlockAtStep() {
			setIDs = append(setIDs, v.SceneSetID)
		}
		return setIDs
	}()...)
	if err != nil {
		return err
	}
	sceneSteps, err := atSvc.ListAutoTestSceneSteps(func() []uint64 {
		sceneIDs := make([]uint64, 0, len(scenes))
		for _, v := range scenes {
			sceneIDs = append(sceneIDs, v.ID)
		}
		return sceneIDs
	}())
	if err != nil {
		return err
	}
	tv := pkg.TextValue{
		Value: strutil.String(func() int {
			var count int
			for _, v := range sceneSteps {
				if v.Type.IsEffectiveStepType() && v.Name != "" {
					count++
				}
			}
			return count
		}()),
		Kind: cputil.I18n(ctx, "auto-test-api-num"),
	}
	c.Props = tv.ConvertToProps()
	return nil
}
