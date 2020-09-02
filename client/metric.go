package client

type MetricType string

type metricTypeList struct {
	Gauge   MetricType
	Counter MetricType
}

var Metrics = &metricTypeList{
	Gauge:   "g",
	Counter: "c",
}

type Metric struct {
	metricType MetricType
	name       string
	value      float32
	sampleRate float32
	labels     string
	port       int
	host       string
}
