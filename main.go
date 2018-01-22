// Copyright © 2018 Kenneth Herner <kherner@navistone.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"net/http"

	"github.com/chosenken/prometheus-kairosdb-adapter/cmd"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Setup the prometheus handler
	http.Handle("/metrics", promhttp.Handler())
	cmd.Execute()
}