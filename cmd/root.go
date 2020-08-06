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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Value float32
var Labels string
var Port int
var Host string

var rootCmd = &cobra.Command{
	Use:   "bluematador-client",
	Short: "Send metrics to Blue Matador",
	Long: `Send metrics to Blue Matador.

Send counter or gauge metrics
	bluematador-client count app.homepage.clicks
	//
	bluematador-client gauge app.request.size 100 

To set the port use the flag --port or -p. The default is 8767
	bluematador-client --port 8767
To set the host use the flag --host. The default is 'localhost'
	bluematador-client --host localhost

**Note: If you have set BLUEMATADOR_AGENT_HOST and BLUEMATADOR_AGENT_PORT in the config file for your agent there is no need to set either using flags.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().IntVarP(&Port, "port", "p", 8767, "The port to send the custom metrics to")
	rootCmd.PersistentFlags().StringVarP(&Host, "host", "", "localhost", "The host to send the custom metrics to")
}
