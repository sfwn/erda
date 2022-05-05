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

package edgereporter

func (p *provider) GetTargetAuthToken() string {
	// if auth token already specified, use it directly
	if len(p.Cfg.Target.AuthToken) > 0 {
		return p.Cfg.Target.AuthToken
	}
	// auth token not specified, get it from edge register every time when you use it, because token might be changed
	return p.EdgeRegister.ClusterAccessKey()
}
