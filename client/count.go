package client

// func Count(args ...interface{}) error {
// 	if len(args) < 2 {
// 		return errors.New("Not enough parameters to create metric")
// 	}

// 	metric := &Metric{
// 		metricType	Metrics.Count
// 		sampleRate  1
// 		port 		8767
// 		host		"localhost"
// 	}

// 	for i, arg := range args {
// 		switch i {
// 		case 0:
// 			name, ok := arg.(string)
// 			if !ok {
// 				return errors.New("Metric name must be a string")
// 			}
// 			metric.name = name
// 		case 1:
// 			value, ok := arg.(float32)
// 			if !ok {
// 				return errors.New("Metric value must be a number")
// 			}
// 			metric.value = value
// 		case 2:
// 			sampleRate, ok := arg.(float32)
// 			if !ok {
// 				return errors.New("Metric sample rate must be a number between 0 and 1")
// 			}

// 			if sampleRate < 0 || sampleRate > 1 {
// 				return errors.New("Metric sample rate must be a number between 0 and 1")
// 			}
// 			metric.sampleRate = sampleRate
// 		case 3:
// 			labels, ok := arg.(string)
// 			if !ok {
// 				return errors.New("Metric labels must be a string")
// 			}
// 			metric.labels = labels
// 		case 4:
// 			port, ok := arg.(int)
// 			if !ok {
// 				return errors.New("Port must be a number")
// 			}
// 			metric.port = port
// 		case 5:
// 			host, ok := arg.(string)
// 			if !ok {
// 				return errors.New("Host must be a string")
// 			}
// 			metric.host = host
// 		default:
// 			return errors.New("Too many arguments to create metric")
// 		}
// 	}

// 	fmt.Printf("metric name", metric.name)
// 	fmt.Printf("metric value", metric.value)
// 	fmt.Printf("metric sampleRate", metric.sampleRate)
// 	fmt.Printf("metric labels", metric.labels)
// 	fmt.Printf("metric port", metric.port)
// 	fmt.Printf("metric host", metric.host)
// }
