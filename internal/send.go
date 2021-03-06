package internal

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func SendMetric(metricType string, metric string, value float32, sampleRate float32, labels string, port int, host string) error {
	metricString := createMetricString(metricType, metric, value, sampleRate, labels)
	portString := strconv.Itoa(port)
	address := net.JoinHostPort(host, portString)
	conn, err := net.Dial("udp", address)
	if err != nil {
		return err
	}
	defer conn.Close()

	metricByte := []byte(metricString)
	conn.Write(metricByte)

	return nil
}

func createMetricString(metricType string, metric string, value float32, sampleRate float32, labels string) string {
	if labels != "" {
		formattedLabels := strings.ReplaceAll(labels, ",", "#")
		formattedLabels = strings.ReplaceAll(formattedLabels, " ", "")
		return fmt.Sprintf("%s:%f|%s|@%f|#%s", metric, value, metricType, sampleRate, formattedLabels)
	} else {
		return fmt.Sprintf("%s:%f|%s|@%f", metric, value, metricType, sampleRate)
	}
}
