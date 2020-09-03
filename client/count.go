package client

import (
	"github.com/bluematador/bluematador-metrics-client-go/internal"
)

func (this *BlueMatadorClient) Count(name string, labels string) {
	metric := CreateMetric(name, float32(1), float32(1), labels)

	internal.SendMetric("c", metric.name, metric.value, metric.sampleRate, metric.labels, this.port, this.host)
}

func (this *BlueMatadorClient) CountWithSampleRate(name string, labels string, sampleRate float32) {
	metric := CreateMetric(name, float32(1), sampleRate, labels)

	internal.SendMetric("c", metric.name, metric.value, metric.sampleRate, metric.labels, this.port, this.host)
}

func (this *BlueMatadorClient) CountWithValue(name string, value float32, labels string) {
	metric := CreateMetric(name, value, float32(1), labels)

	internal.SendMetric("c", metric.name, metric.value, metric.sampleRate, metric.labels, this.port, this.host)
}

func (this *BlueMatadorClient) CountWithValueAndSample(name string, value float32, labels string, sampleRate float32) {
	metric := CreateMetric(name, value, sampleRate, labels)

	internal.SendMetric("c", metric.name, metric.value, metric.sampleRate, metric.labels, this.port, this.host)
}
