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

package settings

import (
	"context"
	"reflect"
	"testing"

	"bou.ke/monkey"

	"github.com/erda-project/erda-infra/base/servicehub"
	commonpb "github.com/erda-project/erda-proto-go/common/pb"
	"github.com/erda-project/erda-proto-go/core/monitor/settings/pb"
	"github.com/erda-project/erda/apistructs"
	"github.com/erda-project/erda/bundle"
)

func Test_settingsService_GetSettings(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.GetSettingsRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.GetSettingsResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		// 		{
		// 			"case 1",
		// 			"erda.core.monitor.settings.SettingsService",
		// 			`
		// erda.core.monitor.settings:
		// `,
		// 			args{
		// 				context.TODO(),
		// 				&pb.GetSettingsRequest{
		// 					// TODO: setup fields
		// 				},
		// 			},
		// 			&pb.GetSettingsResponse{
		// 				// TODO: setup fields.
		// 			},
		// 			false,
		// 		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.SettingsServiceServer)
			got, err := srv.GetSettings(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("settingsService.GetSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("settingsService.GetSettings() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_settingsService_PutSettings(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.PutSettingsRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.PutSettingsResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		// 		{
		// 			"case 1",
		// 			"erda.core.monitor.settings.SettingsService",
		// 			`
		// erda.core.monitor.settings:
		// `,
		// 			args{
		// 				context.TODO(),
		// 				&pb.PutSettingsRequest{
		// 					// TODO: setup fields
		// 				},
		// 			},
		// 			&pb.PutSettingsResponse{
		// 				// TODO: setup fields.
		// 			},
		// 			false,
		// 		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.SettingsServiceServer)
			got, err := srv.PutSettings(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("settingsService.PutSettings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("settingsService.PutSettings() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_settingsService_RegisterMonitorConfig(t *testing.T) {
	type args struct {
		ctx context.Context
		req *pb.RegisterMonitorConfigRequest
	}
	tests := []struct {
		name     string
		service  string
		config   string
		args     args
		wantResp *pb.RegisterMonitorConfigResponse
		wantErr  bool
	}{
		// TODO: Add test cases.
		// 		{
		// 			"case 1",
		// 			"erda.core.monitor.settings.SettingsService",
		// 			`
		// erda.core.monitor.settings:
		// `,
		// 			args{
		// 				context.TODO(),
		// 				&pb.RegisterMonitorConfigRequest{
		// 					// TODO: setup fields
		// 				},
		// 			},
		// 			&pb.RegisterMonitorConfigResponse{
		// 				// TODO: setup fields.
		// 			},
		// 			false,
		// 		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hub := servicehub.New()
			events := hub.Events()
			go func() {
				hub.RunWithOptions(&servicehub.RunOptions{Content: tt.config})
			}()
			err := <-events.Started()
			if err != nil {
				t.Error(err)
				return
			}
			srv := hub.Service(tt.service).(pb.SettingsServiceServer)
			got, err := srv.RegisterMonitorConfig(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("settingsService.RegisterMonitorConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantResp) {
				t.Errorf("settingsService.RegisterMonitorConfig() = %v, want %v", got, tt.wantResp)
			}
		})
	}
}

func Test_checkOrgPermission(t *testing.T) {
	type args struct {
		identityInfo *commonpb.IdentityInfo
		orgID        uint64
		action       string
	}
	bdl := bundle.New()
	pm1 := monkey.PatchInstanceMethod(reflect.TypeOf(bdl), "CheckPermission", func(_ *bundle.Bundle, req *apistructs.PermissionCheckRequest) (*apistructs.PermissionCheckResponseData, error) {
		if req.UserID == "1" {
			return &apistructs.PermissionCheckResponseData{Access: false}, nil
		}
		return &apistructs.PermissionCheckResponseData{Access: true}, nil
	})
	defer pm1.Unpatch()
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid orgID",
			args: args{
				identityInfo: &commonpb.IdentityInfo{},
				orgID:        0,
				action:       apistructs.GetAction,
			},
			wantErr: true,
		},
		{
			name: "invalid identityInfo",
			args: args{
				orgID:        1,
				action:       apistructs.GetAction,
				identityInfo: nil,
			},
			wantErr: true,
		},
		{
			name: "permission less user",
			args: args{
				orgID:  1,
				action: apistructs.GetAction,
				identityInfo: &commonpb.IdentityInfo{
					UserID: "1",
				},
			},
			wantErr: true,
		},
		{
			name: "sys admin user",
			args: args{
				orgID:  1,
				action: apistructs.GetAction,
				identityInfo: &commonpb.IdentityInfo{
					UserID: "2",
				},
			},
			wantErr: false,
		},
	}
	e := &settingsService{
		bundle: bdl,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := e.checkOrgPermission(tt.args.identityInfo, tt.args.orgID, tt.args.action); (err != nil) != tt.wantErr {
				t.Errorf("checkOrgPermission() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
