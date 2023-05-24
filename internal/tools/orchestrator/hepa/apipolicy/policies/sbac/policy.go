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

package sbac

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/erda-project/erda/internal/tools/orchestrator/hepa/apipolicy"
	gateway_providers "github.com/erda-project/erda/internal/tools/orchestrator/hepa/gateway-providers"
	providerDto "github.com/erda-project/erda/internal/tools/orchestrator/hepa/gateway-providers/dto"
	mseCommon "github.com/erda-project/erda/internal/tools/orchestrator/hepa/gateway-providers/mse/common"
	"github.com/erda-project/erda/internal/tools/orchestrator/hepa/repository/orm"
	db "github.com/erda-project/erda/internal/tools/orchestrator/hepa/repository/service"
)

func init() {
	if err := apipolicy.RegisterPolicyEngine(apipolicy.Policy_Engine_SBAC, new(Policy)); err != nil {
		panic(err)
	}
}

type Policy struct {
	apipolicy.BasePolicy
}

func (policy Policy) CreateDefaultConfig(ctx map[string]interface{}) apipolicy.PolicyDto {
	return new(PolicyDto)
}

func (policy Policy) UnmarshalConfig(config []byte, gatewayProvider string) (apipolicy.PolicyDto, error, string) {
	var policyDto PolicyDto
	if err := json.Unmarshal(config, &policyDto); err != nil {
		return nil, errors.Wrapf(err, "failed to Unmarshal config: %s", string(config)), "Invalid config"
	}
	if err := policyDto.IsValidDto(gatewayProvider); err != nil {
		return nil, errors.Wrap(err, "Invalid config"), "Invalid config"
	}
	return &policyDto, nil, ""
}

func (policy Policy) buildPluginReq(dto *PolicyDto, gatewayProvider, zoneName string) *providerDto.PluginReqDto {
	return dto.ToPluginReqDto(gatewayProvider, zoneName)
}

// forValidate 用于识别解析的目的，如果解析是用来做 nginx 配置冲突相关的校验，则关于数据表、调用 kong 接口的操作都不会执行
func (policy Policy) ParseConfig(dto apipolicy.PolicyDto, ctx map[string]interface{}, forValidate bool) (apipolicy.PolicyConfig, error) {
	l := logrus.WithField("pluginName", apipolicy.Policy_Engine_SBAC).WithField("func", "ParseConfig")
	l.Infof("dto: %+v", dto)
	res := apipolicy.PolicyConfig{}
	policyDto, ok := dto.(*PolicyDto)
	if !ok {
		return res, errors.Errorf("invalid config:%+v", dto)
	}
	logrus.Infof("sbac policyDto=%+v, ctx=%v", *policyDto, ctx)

	gatewayAdapter, gatewayProvider, err := policy.GetGatewayAdapter(ctx, apipolicy.Policy_Engine_CSRF)
	if err != nil {
		return res, err
	}

	adapter, ok := gatewayAdapter.(gateway_providers.GatewayAdapter)
	if !ok {
		l.Infof("get gateway adpater for set sbac policy failed")
		return res, nil
	}

	gatewayVersion, err := adapter.GetVersion()
	if err != nil {
		return res, errors.Wrap(err, "failed to retrieve Kong version")
	}
	if !strings.HasPrefix(gatewayVersion, "2.") && !strings.HasPrefix(gatewayVersion, "mse-") {
		return res, errors.Errorf("the plugin %s is not supportted on the gateway version %s", apipolicy.Policy_Engine_SBAC, gatewayVersion)
	}

	zone, ok := ctx[apipolicy.CTX_ZONE].(*orm.GatewayZone)
	if !ok {
		return res, errors.Errorf("failed to get identify with %s: %+v", apipolicy.CTX_ZONE, ctx)
	}
	logrus.Infof("set sbac policy for zone=%+v", *zone)
	policyDb, _ := db.NewGatewayPolicyServiceImpl()
	exist, err := policyDb.GetByAny(&orm.GatewayPolicy{
		ZoneId:     zone.Id,
		PluginName: apipolicy.Policy_Engine_SBAC,
	})
	if err != nil {
		return res, err
	}
	if !policyDto.Switch && !forValidate {
		if exist != nil {

			switch gatewayProvider {
			case mseCommon.MseProviderName:
				// 创建网关路由也会走到这里，但这个时候 MSE 那边应该还没有发现到新路由，因此调用此插件会失败，但创建过程又不能一直等待，因此这里做个异步处理
				if zone.Type == db.ZONE_TYPE_PACKAGE_API {
					resp, err := adapter.CreateOrUpdatePluginById(policy.buildPluginReq(policyDto, gatewayProvider, strings.ToLower(zone.Name)))
					if err != nil {
						go policy.nonSwitchUpdateMSEPluginConfig(adapter, policyDto, zone.Name)
					} else {
						logrus.Infof("create or update mse erda-sbac plugin with response: %+v", *resp)
					}
				}
			default:
				// Kong
				err = adapter.RemovePlugin(exist.PluginId)
				if err != nil {
					return res, err
				}
			}

			_ = policyDb.DeleteById(exist.Id)
			res.KongPolicyChange = true
		}
		return res, nil
	}

	req := policy.buildPluginReq(policyDto, gatewayProvider, strings.ToLower(zone.Name))

	packageAPIDB, _ := db.NewGatewayPackageApiServiceImpl()
	packageApi, err := packageAPIDB.GetByAny(&orm.GatewayPackageApi{ZoneId: zone.Id})
	if err != nil {
		l.WithError(err).Warnf("sbac: packageAPIDB.GetByAny(&orm.GatewayPackageApi{ZoneId: %s})", zone.Id)
		return res, err
	}
	if packageApi == nil {
		l.WithError(err).Warnf("sbac: not found packageAPIDB.GetByAny(&orm.GatewayPackageApi{ZoneId: %s})", zone.Id)
		return res, nil
	}

	if gatewayProvider != mseCommon.MseProviderName {
		// Kong 网关，sbac 插件需要关联 routeId，否则就成了全局策略，影响网关其他路由正常访问
		routeDb, _ := db.NewGatewayRouteServiceImpl()
		route, err := routeDb.GetByApiId(packageApi.Id)
		if err != nil {
			l.WithError(err).Warnf("sbac: routeDb.GetByApiId(packageApi.Id:%s)", packageApi.Id)
			return res, nil
		}
		if route == nil {
			l.WithError(err).Warnf("sbac: not found routeDb.GetByApiId(packageApi.Id:%s)", packageApi.Id)
			return res, nil
		}

		if route != nil {
			logrus.Infof("set sbac policy route ID for packageApi id=%s route=%+v\n", packageApi.Id, *route)
			req.RouteId = route.RouteId
			req.Route = &providerDto.KongObj{
				Id: route.RouteId,
			}
		}
	}

	if exist != nil {
		if !forValidate {
			req.Id = exist.PluginId
			resp, err := adapter.CreateOrUpdatePluginById(req)
			if err != nil {
				return res, err
			}
			configByte, err := json.Marshal(resp.Config)
			if err != nil {
				return res, err
			}
			exist.Config = configByte
			logrus.Infof("update exist sbac gateway policy=%+v", *exist)
			err = policyDb.Update(exist)
			if err != nil {
				return res, err
			}
		}
	} else {
		if !forValidate {
			var resp *providerDto.PluginRespDto
			if gatewayProvider == mseCommon.MseProviderName {
				// 如果是 MSE 网关，则仍然是 CreateOrUpdatePluginById
				if zone.Type == db.ZONE_TYPE_PACKAGE_API {
					resp, err = adapter.CreateOrUpdatePluginById(policy.buildPluginReq(policyDto, gatewayProvider, strings.ToLower(zone.Name)))
					if err != nil {
						return res, err
					}
				}
			} else {
				resp, err = adapter.AddPlugin(req)
				if err != nil {
					return res, err
				}
			}

			configByte, err := json.Marshal(resp.Config)
			if err != nil {
				return res, err
			}
			policyDao := &orm.GatewayPolicy{
				ZoneId:     zone.Id,
				PluginName: apipolicy.Policy_Engine_SBAC,
				Category:   apipolicy.Policy_Category_Safety,
				PluginId:   resp.Id,
				Config:     configByte,
				Enabled:    1,
			}
			logrus.Infof("Insert sbac gateway policy=%+v", *policyDao)
			err = policyDb.Insert(policyDao)
			if err != nil {
				return res, err
			}
			res.KongPolicyChange = true
		}
	}
	return res, nil
}

// 初创路由或者关闭路由策略（PolicyDto.Switch == false）的时候,都会进入 ParseConfig 的同一段逻辑中， 但:
// 1. 如果是关闭路由策略，则对应的逻辑里需要清除已经配置的插件策略，一般直接就能处理了，因此进入不了 nonSwitchUpdateMSEPluginConfig() 的逻辑
// 2. 如果是新建路由，实际上是不需要进行处理的（但网关应用默认策略实际上还是会进入 ParseConfig），此时路由还没被 MSE 网关识别到，但可以延时等待拿到对应的新的路由信息，然后进行类似清除路由对应的策略配置的设置即可，但这个过程不能同步等待，因此异步执行，最多重试3次
func (policy Policy) nonSwitchUpdateMSEPluginConfig(mseAdapter gateway_providers.GatewayAdapter, policyDto *PolicyDto, zoneName string) {
	for i := 0; i < 3; i++ {
		time.Sleep(10 * time.Second)
		resp, err := mseAdapter.CreateOrUpdatePluginById(policy.buildPluginReq(policyDto, mseCommon.MseProviderName, strings.ToLower(zoneName)))
		if err != nil {
			if i == 2 {
				logrus.Errorf("can not update mse erda-sbac plugin for 4 times in 30s, err: %v", err)
				return
			}
			continue
		}
		logrus.Infof("create or update mse erda-sbac plugin for zonename=%s with response: %+v", strings.ToLower(zoneName), *resp)
		break
	}
	return
}
