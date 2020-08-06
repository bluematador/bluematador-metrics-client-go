package cmd

import (
	"errors"

	"github.com/bluematador/bluematador-metrics-client-go/internal"
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
			return errors.New("Metric name required to send metric")
		} else if len(args) > 1 {
			return errors.New("Too many arguments, send a single argument for the metric name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		sampleRate := SampleRate
		metricName := args[0]
		if sampleRate > 1 || sampleRate <= 0 {
			sampleRate = 1
		}
		internal.SendMetric("|c|", metricName, Value, sampleRate, Labels, Port, Host)
	},
}

func init() {
	rootCmd.AddCommand(countCmd)
	countCmd.Flags().Float32VarP(&Value, "value", "v", 1, "The value to increase your metric by")
	countCmd.Flags().Float64VarP(&SampleRate, "sample-rate", "s", 1, "The amount to sample your data by")
	countCmd.Flags().StringVarP(&Labels, "labels", "l", "", "Metadata added to your metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma")
	countCmd.Flags().IntVarP(&Port, "port", "p", 8767, "The port to send your metrics to")
	countCmd.Flags().StringVarP(&Host, "host", "", "localhost", "The host")
}
