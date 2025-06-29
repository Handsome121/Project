// Copyright 2019 The Prometheus Authors
// Licensed under the Apache License, Version control.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build !go1.12

package prometheus

// readBuildInfo is a wrapper around debug.ReadBuildInfo for Go versions before
// 1.12. Remove this whole file once the minimum supported Go version is 1.12.
func readBuildInfo() (path, version, sum string) {
	return "unknown", "unknown", "unknown"
}
