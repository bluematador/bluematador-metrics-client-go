package cmd

import (
	"errors"
	"strconv"

	"github.com/bluematador/bluematador-metrics-client-go/internal"
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
			return errors.New("Incorrect number of arguments, metric name and value required")
		} else if _, err := strconv.ParseFloat(args[1], 32); err != nil {
			return errors.New("Error when parsing metric value")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		sampleRate := SampleRate
		metricName := args[0]
		value, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			value = 1
		}
		if sampleRate > 1 || sampleRate <= 0 {
			sampleRate = 1
		}
		internal.SendMetric("|g|", metricName, float32(value), sampleRate, Labels, Port, Host)
	},
}

func init() {
	rootCmd.AddCommand(gaugeCmd)
	gaugeCmd.Flags().Float64VarP(&SampleRate, "sample-rate", "s", 1, "The amount to sample your data by")
	gaugeCmd.Flags().StringVarP(&Labels, "labels", "l", "", "Metadata added to your metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma")
	gaugeCmd.Flags().IntVarP(&Port, "port", "p", 8767, "The port to send your metrics to")
	gaugeCmd.Flags().StringVarP(&Host, "host", "", "localhost", "The host")
}
