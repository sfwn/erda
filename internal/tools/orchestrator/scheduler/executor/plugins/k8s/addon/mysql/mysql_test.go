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

package mysql

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/erda-project/erda/apistructs"
	mysqlv1 "github.com/erda-project/erda/internal/tools/orchestrator/scheduler/executor/plugins/k8s/addon/mysql/v1"
	"github.com/erda-project/erda/internal/tools/orchestrator/scheduler/executor/plugins/k8s/addon/sourcecov/mock"
	"github.com/erda-project/erda/pkg/http/httpclient"
)

type k8s struct{}

func (k8s) GetK8SAddr() string {
	return ""
}

func TestMysqlOperator(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ns := mock.NewMockNamespaceUtil(ctrl)

	mo := New(new(k8s), ns, nil, nil, httpclient.New())
	sg := new(apistructs.ServiceGroup)
	sg.ID = "abcdefghigklmn"
	mo.Name(sg)
	mo.NamespacedName(sg)
	mo.IsSupported()
	mo.Validate(sg)
	sg.Labels = make(map[string]string)
	sg.Labels["USE_OPERATOR"] = "mysql"
	mo.Validate(sg)
	sg.Services = append(sg.Services, apistructs.Service{
		Name: "mysql",
	})
	sg.Services[0].Env = make(map[string]string)
	mo.Validate(sg)
	sg.Services[0].Env["MYSQL_ROOT_PASSWORD"] = "123"
	mo.Validate(sg)
	mo.Convert(sg)
	mo.Create(sg)
	mo.Inspect(sg)
	mo.Update(sg)
	mo.Remove(sg)
}

func TestConvert(t *testing.T) {
	sg := &apistructs.ServiceGroup{
		Dice: apistructs.Dice{
			ID: "abcdefghigklmnopq",
			Services: []apistructs.Service{
				{
					Name: "mysql",
					Volumes: []apistructs.Volume{
						{
							ID:         "mysql-data",
							VolumePath: "/var/lib/mysql",
							SCVolume: apistructs.SCVolume{
								StorageClassName: "ssd",
								Capacity:         1000,
							},
						},
					},
					Resources: apistructs.Resources{
						Cpu: 1,
						Mem: 1024,
					},
					Scale: 2,
					Env: map[string]string{
						"MYSQL_ROOT_PASSWORD": "123",
						"ADDON_ID":            "addonxxx",
						"DICE_CLUSTER_NAME":   "erda",
						"DICE_ORG_NAME":       "erda",
					},
				},
			},
		},
	}
	op := &MysqlOperator{}
	got := op.Convert(sg)
	mysql, ok := got.(*mysqlv1.Mysql)
	assert.Equal(t, true, ok)
	assert.Equal(t, "erda", mysql.Spec.Labels["DICE_CLUSTER_NAME"])
}
