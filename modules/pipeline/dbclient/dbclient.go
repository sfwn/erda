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

package dbclient

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/erda-project/erda-infra/providers/mysqlxorm"
)

type Client struct {
	mysqlxorm.Interface
}

var (
	ErrRecordNotFound = errors.New("not found")
)

func New(p mysqlxorm.Interface) (*Client, error) {
	if p == nil {
		return nil, fmt.Errorf("mysqlxorm is nil")
	}
	client := Client{
		Interface: p,
	}
	return &client, nil
}

func MustNew(p mysqlxorm.Interface) *Client {
	client, err := New(p)
	if err != nil {
		panic(err)
	}
	return client
}
