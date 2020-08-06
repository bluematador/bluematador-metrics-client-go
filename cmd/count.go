package cmd

import (
	"errors"
	"os"
	"strconv"

	"github.com/bluematador/bluematador-metrics-client-go/internal"
	"github.com/spf13/cobra"
)

var countCmd = &cobra.Command{
	Use:   "count metricName",
	Short: "Send a counter metric",
	Long: `Send a counter metric to your BlueMatador dashboard. 

The count command takes the metric name as a single argument and defaults to incrementing the metric by 1. 
count app.homepage.clicks
To manualy set the amount to increment by use the flag --value or -v
count app.homepage.clicks -v 2
By default all metrics sent to the Blue Matador agent will be stored. 
To send a sample of data for a give metric use the flag --sample-rate or -s with a value between 0 and 1
count app.homepage.clicks -s .5 // 50% of data from this metric will be sent to the agent
To add metadata to your metrics use the flag --labels or -l. Format your labels as single string of key value colon seperated labels with each label being separated by a comma.
count app.homepage.clicks -l 'env:dev, account_id:12354'`,
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
		port, err := strconv.Atoi(os.Getenv("BLUEMATADOR_AGENT_PORT"))
		if err != nil {
			port = Port
		}
		host := os.Getenv("BLUEMATADOR_AGENT_HOST")
		if host == "" {
			host = Host
		}
		internal.SendMetric("|c|", metricName, Value, sampleRate, Labels, port, host)
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
