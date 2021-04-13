package queuemanage

import (
	"github.com/erda-project/erda/modules/pipeline/dbclient"
)

type QueueManage struct {
	dbClient *dbclient.Client
}

func New(ops ...Option) *QueueManage {
	qm := QueueManage{}
	for _, op := range ops {
		op(&qm)
	}
	return &qm
}

type Option func(*QueueManage)

func WithDBClient(dbClient *dbclient.Client) Option {
	return func(qm *QueueManage) {
		qm.dbClient = dbClient
	}
}
