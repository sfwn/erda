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

package addWaitButton

import (
	"context"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cpregister/base"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/apps/dop/component-protocol/components/auto-test-scenes/common/gshelper"
	"github.com/erda-project/erda/modules/apps/dop/component-protocol/types"
)

type ComponentAction struct {
	sdk *cptype.SDK
	bdl *bundle.Bundle
}

func init() {
	base.InitProviderWithCreator("auto-test-scenes", "addWaitButton",
		func() servicehub.Provider { return &ComponentAction{} })
}

func (ca *ComponentAction) Render(ctx context.Context, c *cptype.Component, scenario cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	gh := gshelper.NewGSHelper(gs)
	ca.sdk = cputil.SDK(ctx)
	ca.bdl = ctx.Value(types.GlobalCtxKeyBundle).(*bundle.Bundle)

	switch event.Operation {
	case cptype.OperationKey(apistructs.ExecuteAddApiOperationKey):
		var req apistructs.AutotestSceneRequest
		req.Target = -1
		req.GroupID = -1
		req.Type = apistructs.StepTypeWait

		var autotestSceneRequest apistructs.AutotestSceneRequest
		autotestSceneRequest.UserID = ca.sdk.Identity.UserID
		autotestSceneRequest.ID = gh.GetFileTreeSceneID()
		autotestSceneRequest.SceneID = gh.GetFileTreeSceneID()
		result, err := ca.bdl.GetAutoTestScene(autotestSceneRequest)
		if err != nil {
			return err
		}

		req.SceneID = result.ID
		req.SpaceID = result.SpaceID
		req.CreatorID = ca.sdk.Identity.UserID
		req.UpdaterID = ca.sdk.Identity.UserID
		req.UserID = ca.sdk.Identity.UserID
		req.Value = "{\"waitTimeSec\":1}"
		stepID, err := ca.bdl.CreateAutoTestSceneStep(req)
		if err != nil {
			return err
		}
		c.State["createStepID"] = stepID
		c.State["showWaitEditorDrawer"] = true
	case cptype.InitializeOperation, cptype.RenderingOperation:
		c.Type = "Button"
		c.Props = map[string]interface{}{
			"text": "+ " + ca.sdk.I18n("wait"),
		}
		c.Operations = make(map[string]interface{})
		c.Operations = map[string]interface{}{
			"click": map[string]interface{}{
				"key":    "addApi",
				"reload": true,
			},
		}
	}
	return nil
}
