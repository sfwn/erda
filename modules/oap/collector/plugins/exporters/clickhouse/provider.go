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

package clickhouse

import (
	"context"
	"fmt"

	"github.com/erda-project/erda-infra/base/logs"
	"github.com/erda-project/erda-infra/base/servicehub"
	"github.com/erda-project/erda-infra/providers/clickhouse"
	"github.com/erda-project/erda/modules/core/monitor/log"
	"github.com/erda-project/erda/modules/core/monitor/metric"
	"github.com/erda-project/erda/modules/msp/apm/trace"
	"github.com/erda-project/erda/modules/oap/collector/core/model/odata"
	"github.com/erda-project/erda/modules/oap/collector/plugins/exporters/clickhouse/span"
)

type config struct {
	CurrencyNum int    `file:"currency_num" default:"20" ENV:"EXPORTER_CH_CURRENCY_NUM"`
	RetryNum    int    `file:"retry_num" default:"5" ENV:"EXPORTER_CH_RETRY_NUM"`
	Database    string `file:"database" default:"monitor"`
}

// +provider
type provider struct {
	Cfg        *config
	Log        logs.Logger
	ch         clickhouse.Interface
	ctx        context.Context
	cancelFunc context.CancelFunc

	spanWriter *span.WriteSpan
}

func (p *provider) ExportRaw(items ...*odata.Raw) error        { return nil }
func (p *provider) ExportMetric(items ...*metric.Metric) error { return nil }
func (p *provider) ExportLog(items ...*log.Log) error          { return nil }

// TODO currency
func (p *provider) ExportSpan(items ...*trace.Span) error {
	p.spanWriter.AddBatch(items)
	return nil
}

func (p *provider) ComponentConfig() interface{} {
	return p.Cfg
}

func (p *provider) Connect() error {
	return nil
}

// Run this is optional
func (p *provider) Init(ctx servicehub.Context) error {
	svc := ctx.Service("clickhouse@span")
	if svc == nil {
		svc = ctx.Service("clickhouse")
	}
	if svc == nil {
		return fmt.Errorf("service clickhouse is required")
	}
	p.ch = svc.(clickhouse.Interface)
	p.ctx, p.cancelFunc = context.WithCancel(context.Background())

	p.spanWriter = span.NewWriteSpan(span.Config{
		Database: p.Cfg.Database,
		Retry:    p.Cfg.RetryNum,
		Logger:   p.Log.Sub("spanWriter"),
		Client:   p.ch.Client(),
	})

	return nil
}

func (p *provider) Start() error {
	span.InitCurrencyLimiter(p.Cfg.CurrencyNum)
	if err := span.InitSeriesIDMap(p.ch.Client(), p.Cfg.Database); err != nil {
		return fmt.Errorf("cannot init seriesIDMap: %w", err)
	}
	p.spanWriter.Start(p.ctx)
	return nil
}

func (p *provider) Close() error {
	p.cancelFunc()
	return nil
}

func init() {
	servicehub.Register("erda.oap.collector.exporter.clickhouse", &servicehub.Spec{
		Services: []string{
			"erda.oap.collector.exporter.clickhouse",
		},
		Description:  "here is description of erda.oap.collector.exporter.clickhouse",
		Dependencies: []string{"clickhouse", "clickhouse.table.initializer@span"},
		ConfigFunc: func() interface{} {
			return &config{}
		},
		Creator: func() servicehub.Provider {
			return &provider{}
		},
	})
}
