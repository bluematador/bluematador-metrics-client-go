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
	"strconv"

	"github.com/spf13/cobra"
)

var gaugeCmd = &cobra.Command{
	Use:   "gauge metricName value",
	Short: "Send a gauge metric",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Not enought arguments, metric name and value required")
		} else if _, err := strconv.ParseFloat(args[1], 64); err != nil {
			return errors.New("Error when parsing metric value")
		} else if len(args) > 2 {
			return errors.New("Too many arguments, metric name and value only supported arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, s := range args {
			fmt.Println(s)
		}
		fmt.Println(SampleRate, Labels)
		fmt.Println("sending gauge metric")
	},
}

func init() {
	rootCmd.AddCommand(gaugeCmd)
	gaugeCmd.Flags().Float64VarP(&SampleRate, "sampleRate", "s", 1, "The amount to sample your data by")
	gaugeCmd.Flags().StringVarP(&Labels, "labels", "l", "", "Metadata added to your metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma")
	gaugeCmd.Flags().IntVarP(&Port, "port", "p", 8767, "The port to send your metrics to")
}
