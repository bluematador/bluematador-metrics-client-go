# Blue Matador Go Metrics Client

**Send StatsD-style custom metrics to your Blue Matador dashboard** 

## Installation

To install the go client package run the following command in your environment: `go get github.com/bluematador/bluematador-metrics-client-go`

## Setup

To start using the Blue Matador Metrics Client import the client package. 

```
import (
    "github.com/bluematador/bluematador-metrics-client-go/client"
)
```

Once you have the client packaged imported you can instantiate the Blue Matador Client. 

The Blue Matador Client has four builder methods to choose from. 

```
func NewBlueMatadorClient() *BlueMatadorClient

```
`NewBlueMatadorClient` instantiates a BlueMatadorClient with the default host ("localhost") and port (8767) address. If you have set BLUEMATADOR_AGENT_HOST and BLUEMATADOR_AGENT_PORT as env variables, they will be used for the address. 
 

```
blueMatador := client.NewBlueMatadorClient()

```

```
func NewBlueMatadorClientWithAddress(host string, port int) *BlueMatadorClient

```

`NewBlueMatadorClientWithAddress` allows you to manually set the connection address. Using this method will override the environment variables in the agent config file. 

```
blueMatador := client.NewBlueMatadorClientWithAddress("localhost", 8767)
```

```
func NewBlueMatadorClientWithPrefix(prefix string) *BlueMatadorClient

```

`NewBlueMatadorClientWithPrefix` allows you to set a prefix that will be applied to each metric you send. If you set the prefix to "app" and then submit a metric named "requests.size" the resulting metric name is "app.requests.size"

**Note:** The prefix cannot contain ":" or "|". If either of these characters are contained in the prefix they will be replaced with "_" before being sent to the Agent. 

```
blueMatador := client.NewBlueMatadorClientWithPrefix("app")
```

```
func NewBlueMatadorClientWithAddressAndPrefix(host string, port int, prefix string) *BlueMatadorClient

```

`NewBlueMatadorClientWithAddressAndPrefix` allows you to manually set the connection address and set a metric prefix. 

```
blueMatador := NewBlueMatadorClientWithAddressAndPrefix("localhost", 8767, "app")
```

Once you have a Blue Matador Client instantiated you can start sending metrics to your Agent. 

### Gauge

```
func (*BlueMatadorClient) Gauge(name string, value float32, labels string)
```

`Gauge` requires a metric name, value and labels. Metric labels adds metadata to the metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma e.g 'env:dev,account_id:12354'. If you do not wish to send a metric with labels use an empty string as the argument. 

**Note:** The labels cannot contain "#" or "|". If either of these characters are contained in the labels they will be replaced with "_" before being sent to the Agent. 

```
client.Gauge("app.load.size", 100, "env:dev")
```

```
func (*BlueMatadorClient) GaugeWithSampleRate(name string, value float32, labels string, sampleRate float32)
```

`GaugeWithSampleRate` works the same as `Gauge` but allows you to send only a sample of your data e.g. `0.5` indicates 50% of data being sent to the agent. Only useful to cut down on network usage or agent resource usage on extremely high-volume metrics. Default value is 1

**Note:** The sampleRate must be a number between 0 and 1. If the give sampleRate is greater than 1 or less than 0 it will default back to 1.

```
client.GaugeWithSampleRate("app.load.size", 100, "env:dev", 0.5)
```

### Count

```
func (*BlueMatadorClient) Count(name string, labels string)
```

`Count` requires a metric name and labels. The metric value for `Count` is 1. Metric labels adds metadata to the metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma e.g 'env:dev,account_id:12354'. If you do not wish to send a metric with labels use an empty string as the argument. 

**Note:** The labels cannot contain "#" or "|". If either of these characters are contained in the labels they will be replaced with "_" before being sent to the Agent. 

```
client.Count("app.requests", "")
```

```
func (*BlueMatadorClient) CountWithValue(name string, value float32, labels string)
```

`CountWithValue` is the same as `Count` but allows you to set the value to increment your metric by. 

```
client.CountWithValue("app.requests", 10, "env:dev")
```

```
func (*BlueMatadorClient) CountWithSampleRate(name string, labels string, sampleRate float32)
```

`CountWithSampleRate` is the same as `Count` but allows you to send only a sample of your data e.g. `0.5` indicates 50% of data being sent to the agent. Only useful to cut down on network usage or agent resource usage on extremely high-volume metrics. Default value is 1

**Note:** The sampleRate must be a number between 0 and 1. If the give sampleRate is greater than 1 or less than 0 it will default back to 1.

```
client.CountWithSampleRate("app.requests", "", 0.5)
```

```
func (*BlueMatadorClient) CountWithValueAndSample(name string, value float32, labels string, sampleRate float32)
```

`CountWithValueAndSample` allows you to combine a custom value to increment your metric with and a percentage of the data you wish to sample. 

```
client.CountWithValueAndSample("app.requests", 10, "env:dev", 0.5)
```

# Additional Information

To install the Blue Matador agent, visit the [Integrations](https://app.bluematador.com/ur/app#/setup/integrations) page in the Blue Matador app.

For more information about Blue Matador, visit the [Blue Matador Website](https://www.bluematador.com)


# License

More details in [LICENSE.](https://github.com/bluematador/bluematador-metrics-client-go/blob/master/LICENSE)

Copyright (c) 2020 [Blue Matador, Inc.](https://www.bluematador.com/)
