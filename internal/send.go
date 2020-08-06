package internal

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func SendMetric(metricType string, metric string, value float32, sampleRate float64, labels string, port int, host string) error {
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

func createMetricString(metricType string, metric string, value float32, sampleRate float64, labels string) string {
	stringValue := fmt.Sprintf("%f", value)
	stringSample := fmt.Sprintf("%f", sampleRate)
	formattedLabels := strings.ReplaceAll(labels, ",", "#")
	formattedLabels = strings.ReplaceAll(formattedLabels, " ", "")

	if sampleRate != 1 {
		return metric + ":" + stringValue + metricType + "@" + stringSample + "|" + "#" + formattedLabels
	} else {
		return metric + ":" + stringValue + metricType + "#" + formattedLabels
	}
}
