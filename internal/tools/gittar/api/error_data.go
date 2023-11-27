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

package api

import "errors"

type Map map[string]interface{}

var (
	ERROR_CODE_INTERNAL     = "500"
	ERROR_CODE_NOT_FOUND    = "404"
	ERROR_CODE_INVALID_ARGS = "400"
)

var (
	ERROR_NOT_FILE       = errors.New("path not file")
	ERROR_PATH_NOT_FOUND = errors.New("path not exist")
	ERROR_DB             = errors.New("db error")
	ERROR_ARG_ID         = errors.New("id parse failed")
	ERROR_HOOK_NOT_FOUND = errors.New("hook not found")
	ERROR_LOCKED_DENIED  = errors.New("locked denied")
	ERROR_REPO_LOCKED    = errors.New("repo locked")
	ERROR_AI_NOT_ENABLED = errors.New("ai not enabled")
)
