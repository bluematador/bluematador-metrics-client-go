package cmd

import (
	"errors"
	"os"
	"strconv"

	"github.com/bluematador/bluematador-metrics-client-go/client"
	"github.com/bluematador/bluematador-metrics-client-go/internal"
	"github.com/spf13/cobra"
)

var gaugeCmd = &cobra.Command{
	Use:   "gauge metricName value",
	Short: "Send a gauge metric",
	Long: `Send a gauge metric to your BlueMatador dashboard. 

The gauge command takes the metric name and value as its sole arguments. 
	gauge app.request.size 100
To add metadata to your metrics use the flag --labels or -l. Format your labels as single string of key value colon seperated labels with each label being separated by a comma.
	gauge app.request.size 100 -l 'env:dev,account_id:12354'`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Incorrect number of arguments, metric name and value required")
		} else if _, err := strconv.ParseFloat(args[1], 32); err != nil {
			return errors.New("Error when parsing metric value")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		metricName := args[0]
		value, err := strconv.ParseFloat(args[1], 32)
		if err != nil {
			value = 1
		}
		port, err := strconv.Atoi(os.Getenv("BLUEMATADOR_AGENT_PORT"))
		if err != nil {
			port = Port
		}
		host := os.Getenv("BLUEMATADOR_AGENT_HOST")
		if host == "" {
			host = Host
		}
		internal.SendMetric("g", client.Sanitize(metricName, ":"), float32(value), float32(1), client.Sanitize(Labels, "#"), port, host)
	},
}

func init() {
	rootCmd.AddCommand(gaugeCmd)
	gaugeCmd.Flags().StringVarP(&Labels, "labels", "l", "", "Metadata added to your metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma")
}
