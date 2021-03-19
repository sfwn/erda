package actionagentsvc

import (
	"github.com/erda-project/erda/bundle"
	"github.com/erda-project/erda/modules/pipeline/dbclient"
	"terminus.io/dice/dice/pkg/jsonstore"
	"terminus.io/dice/dice/pkg/jsonstore/etcd"
)

type ActionAgentSvc struct {
	dbClient        *dbclient.Client
	bdl             *bundle.Bundle
	accessibleCache jsonstore.JsonStore
	etcdctl         *etcd.Store
}

func New(dbClient *dbclient.Client, bdl *bundle.Bundle, js jsonstore.JsonStore, etcdctl *etcd.Store) *ActionAgentSvc {
	s := ActionAgentSvc{}
	s.dbClient = dbClient
	s.bdl = bdl
	s.accessibleCache = js
	s.etcdctl = etcdctl
	return &s
}
