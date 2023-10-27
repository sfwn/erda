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

package batchOperationTipModal

import (
	"fmt"
	"testing"

	"github.com/erda-project/erda-infra/providers/component-protocol/cptype"
	"github.com/erda-project/erda-infra/providers/i18n"
)

type NopTranslator struct{}

func (t NopTranslator) Get(lang i18n.LanguageCodes, key, def string) string { return key }

func (t NopTranslator) Text(lang i18n.LanguageCodes, key string) string { return key }

func (t NopTranslator) Sprintf(lang i18n.LanguageCodes, key string, args ...interface{}) string {
	return fmt.Sprintf(key, args...)
}

func TestTransfer(t *testing.T) {
	type TestData struct {
		Name string
		Age  int
	}

	// 准备测试数据
	a := TestData{Name: "John", Age: 30}
	var b TestData

	// 调用被测试的方法
	err := Transfer(a, &b)

	// 检查返回的错误
	if err != nil {
		t.Errorf("Transfer() returned an error: %v", err)
	}

	// 检查b是否与a相等
	if b != a {
		t.Errorf("Transfer() did not transfer data correctly. Expected: %v, Got: %v", a, b)
	}

	// 测试传递nil值的情况
	err = Transfer(nil, &b)
	if err == nil {
		t.Error("Transfer() should return an error when passing nil value")
	}

	// 测试传递非指针类型的情况
	err = Transfer(a, b)
	if err == nil {
		t.Error("Transfer() should return an error when passing non-pointer value")
	}
}

func TestBatchOperationTipModal_getContent(t *testing.T) {
	type args struct {
		tip     string
		nodeIDs []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				tip:     "1",
				nodeIDs: []string{"1"},
			},
			want: "1\n1\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := &BatchOperationTipModal{}
			if got := bot.getContent(tt.args.tip, tt.args.nodeIDs); got != tt.want {
				t.Errorf("getContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBatchOperationTipModal_getOperations(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := &BatchOperationTipModal{}
			bot.SDK = &cptype.SDK{}
			bot.SDK.Tran = &NopTranslator{}
			bot.getOperations()
		})
	}
}

func TestBatchOperationTipModal_getProps(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bot := &BatchOperationTipModal{}
			bot.SDK = &cptype.SDK{}
			bot.SDK.Tran = &NopTranslator{}
			bot.getProps()
		})
	}
}
