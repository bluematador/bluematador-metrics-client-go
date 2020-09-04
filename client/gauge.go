package client

import (
	"github.com/bluematador/bluematador-metrics-client-go/internal"
)

func (this *BlueMatadorClient) Gauge(name string, value float32, labels string) {
	metric := this.CreateMetric(name, value, float32(1), labels)

	internal.SendMetric("g", metric.name, metric.value, metric.sampleRate, metric.labels, this.port, this.host)
}

func (this *BlueMatadorClient) GaugeWithSampleRate(name string, value float32, labels string, sampleRate float32) {
	metric := this.CreateMetric(name, value, sampleRate, labels)

	internal.SendMetric("g", metric.name, metric.value, metric.sampleRate, metric.labels, this.port, this.host)
}
