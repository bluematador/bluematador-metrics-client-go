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
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count metricName",
	Short: "Send a counter metric",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Metric Name required to send metric")
		} else if len(args) > 1 {
			return errors.New("Too many arguments, send a single argument for the metric name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range args {
			fmt.Println(s)
		}
		fmt.Println(Value, SampleRate, Labels)
		fmt.Println("sending count metric")
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().Float64VarP(&Value, "value", "v", 1, "The value to increase your metric by")
	countCmd.Flags().Float64VarP(&SampleRate, "sampleRate", "s", 1, "The amount to sample your data by")
	countCmd.Flags().StringVarP(&Labels, "labels", "l", "", "Metadata added to your metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma")
}
