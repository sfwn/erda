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

package cmp_dashboard_workloads

import (
	"fmt"
	"strings"

	"github.com/rancher/wrangler/pkg/data"
)

func ParseWorkloadStatus(obj data.Object) (string, string, error) {
	kind := obj.String("kind")
	fields := obj.StringSlice("metadata", "fields")

	switch kind {
	case "Deployment":
		if len(fields) != 8 {
			return "", "", fmt.Errorf("deployment %s has invalid fields length", obj.String("metadata", "name"))
		}
		// up-to-date and available
		if fields[2] == fields[3] {
			return "Active", "green", nil
		} else {
			return "Error", "red", nil
		}
	case "DaemonSet":
		if len(fields) != 11 {
			return "", "", fmt.Errorf("daemonset %s has invalid fields length", obj.String("metadata", "name"))
		}
		// desired and ready
		if fields[1] == fields[3] {
			return "Active", "green", nil
		} else {
			return "Error", "red", nil
		}
	case "StatefulSet":
		if len(fields) != 5 {
			return "", "", fmt.Errorf("statefulSet %s has invalid fields length", obj.String("metadata", "name"))
		}
		//
		readyPods := strings.Split(fields[1], "/")
		if readyPods[0] == readyPods[1] {
			return "Active", "green", nil
		} else {
			return "Error", "red", nil
		}
	case "Job":
		if len(fields) != 7 {
			return "", "", fmt.Errorf("job %s has invalid fields length", obj.String("metadata", "name"))
		}
		active := obj.String("status", "active")
		failed := obj.String("status", "failed")
		if failed != "" {
			return "Failed", "red", nil
		} else if active != "" {
			return "Active", "green", nil
		} else {
			return "Succeeded", "steelBlue", nil
		}
	case "CronJob":
		return "Active", "green", nil
	default:
		return "", "", fmt.Errorf("valid workload kind: %v", kind)
	}
}
