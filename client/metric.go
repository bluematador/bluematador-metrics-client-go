package client

import (
	"strings"
)

type Metric struct {
	name       string
	value      float32
	sampleRate float32
	labels     string
}

func CreateMetric(name string, value float32, sampleRate float32, labels string) *Metric {
	metricName := Sanitize(name, ":")
	metricLabels := Sanitize(labels, "#")
	metricSampleRate := sampleRate
	if sampleRate < 0 || sampleRate > 1 {
		metricSampleRate = float32(1)
	}
	return &Metric{
		name:       metricName,
		value:      value,
		sampleRate: metricSampleRate,
		labels:     metricLabels,
	}
}

func Sanitize(name string, character string) string {
	sanitizedString := strings.ReplaceAll(name, character, "_")
	return strings.ReplaceAll(sanitizedString, "|", "_")
}
