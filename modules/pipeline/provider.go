package pipeline

import (
	"github.com/erda-project/erda-infra/base/servicehub"
)

const (
	servicePipeline = "pipeline"
)

type provider struct{}

func (p *provider) Service() []string      { return []string{servicePipeline} }
func (p *provider) Dependencies() []string { return []string{"http-endpoints"} }

func (p *provider) Init(ctx servicehub.Context) error {
	panic("implement me")
}

func (p *provider) Start() error {
	panic("implement me")
}

func (p *provider) Creator() servicehub.Creator {
	return func() servicehub.Provider { return &provider{} }
}

func init() {
	servicehub.RegisterProvider(servicePipeline, &provider{})
}
