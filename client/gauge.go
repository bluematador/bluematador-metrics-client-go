package client

import (
	"fmt"
)

func Gauge(args ...interface{}) error {

	if len(args) < 2 {
		return fmt.Errorf("Not enough parameters to create metric")
	}

	metric := &Metric{
		metricType: Metrics.Gauge,
		sampleRate: 1,
		port:       8767,
		host:       "localhost",
	}

	for i, arg := range args {
		switch i {
		case 0:
			name, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Metric name must be a string")
			}
			metric.name = name
		case 1:
			value, ok := arg.(float64)
			if !ok {
				intValue, ok := arg.(int)
				if !ok {
					return fmt.Errorf("Metric value must be a number")
				}
				value = float64(intValue)
				ok = true
			}
			metric.value = float32(value)
		case 2:
			sampleRate, ok := arg.(float64)
			if !ok {
				intRate, ok := arg.(int)
				if !ok {
					return fmt.Errorf("Metric sample rate must be a number between 0 and 1")
				}
				sampleRate = float64(intRate)
				ok = true
			}

			if sampleRate < 0 || sampleRate > 1 {
				return fmt.Errorf("Metric sample rate must be a number between 0 and 1")
			}
			metric.sampleRate = float32(sampleRate)
		case 3:
			labels, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Metric labels must be a string")
			}
			metric.labels = labels
		case 4:
			port, ok := arg.(int)
			if !ok {
				return fmt.Errorf("Port must be a number")
			}
			metric.port = port
		case 5:
			host, ok := arg.(string)
			if !ok {
				return fmt.Errorf("Host must be a string")
			}
			metric.host = host
		default:
			return fmt.Errorf("Too many arguments to create metric")
		}
	}

	fmt.Println("metric name", metric.name)
	fmt.Println("metric value", metric.value)
	fmt.Println("metric sampleRate", metric.sampleRate)
	fmt.Println("metric labels", metric.labels)
	fmt.Println("metric port", metric.port)
	fmt.Println("metric host", metric.host)
	return nil
}
