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

package infoTitle

import (
	"context"

	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/component-protocol/cpregister/base"
	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/component-protocol/utils/cputil"
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/internal/apps/cmp/component-protocol/types"
)

func (infoTitle *InfoTitle) Render(ctx context.Context, c *cptype.Component, s cptype.Scenario, event cptype.ComponentEvent, gs *cptype.GlobalStateData) error {
	infoTitle.CtxBdl = ctx.Value(types.GlobalCtxKeyBundle).(*bundle.Bundle)
	infoTitle.Ctx = ctx
	infoTitle.SDK = cputil.SDK(ctx)
	infoTitle.Props = Props{Title: infoTitle.SDK.I18n("nodeInfo"), Size: "small"}
	c.Props = cputil.MustConvertProps(infoTitle.Props)
	return nil
}

func init() {
	base.InitProviderWithCreator("cmp-dashboard-nodeDetail", "infoTitle", func() servicehub.Provider {
		return &InfoTitle{Type: "Title", Props: Props{}}
	})
}
