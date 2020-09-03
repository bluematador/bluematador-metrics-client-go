// Copyright © 2020 NAME HERE <EMAIL ADDRESS>
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
	"fmt"

	"github.com/bluematador/bluematador-metrics-client-go/client"
)

func main() {
	// cmd.Execute()
	fmt.Println("executing main function")
	blueMatador := client.NewBlueMatadorClientWithPrefix("prefix:")
	blueMatador.Gauge("app:test.gauge", 10025, "env:#dev#prod, accountId#|:1234")
	blueMatador.GaugeWithSampleRate("app:test.gauge", 10025, "env:#dev#prod, accountId#|:1234", .5)
	blueMatador.Count("app.test.counter|", "env:#dev")
	blueMatador.CountWithSampleRate("app.count.sample", "env:test", .75)
	blueMatador.CountWithValue("app.count.value", 25, "env:yoyo")
	blueMatador.CountWithValueAndSample("app.count.sampleValue", 10, "env:yes", .5)

}
