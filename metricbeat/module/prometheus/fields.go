// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package prometheus

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "prometheus", asset.ModuleFieldsPri, AssetPrometheus); err != nil {
		panic(err)
	}
}

// AssetPrometheus returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/prometheus.
func AssetPrometheus() string {
	return "eJyUkU1u6zAMhPc+xUBv95DkAFr0Ci3QZVEEikXbbPQHkkGQ2xeJjdQBsmi1E78RORpucaSLR5OaySY6aQcYWyIP93Yvug6IpL1wM67F46UDgHcLptBeQqOIQWpGwM8rUImtcrFdB+hUxfZ9LQOPHkNISh0glCgoeYzhqiEzLqN6fDjV5DZwk1lznx0wMKWo/jb3H14lkoAVnFsVC8UwkdAGKRwoKc6cEnKwfsLAoraBTQQhNQQhxHo6JLr12qKETOsEdnOPGwXs0sijHr6ot6U0X/YzOdLlXCUu6ElI17PKJJMJ91jNeOJgFv3Wwuo3D2SfQ2tcxkXm/rs/uryT7cOivgMAAP//FvGzhQ=="
}
