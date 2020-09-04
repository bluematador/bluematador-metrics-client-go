# Blue Matador Bash Metrics Client

**Send StatsD-style custom metrics to your Blue Matador dashboard** 

## Installation

To install the bash client download the [bluematador-client](https://github.com/bluematador/bluematador-metrics-client-go/releases/tag/1.0.0) executable and make it a global command for your environment.

## Setup


### Global Flags
  * `--host` specifies the host to send the custom metrics to. If no host is specified, `localhost` is the default host.
  * `--port or -p` specifies the port to send the custom metrics to. If no port is specified, `8767` is the default port. 


**Note:** The client will automatically detect if you have set `BLUEMATADOR_AGENT_HOST` and `BLUEMATADOR_AGENT_PORT` in the config file for your agent. These variables will be used over manually setting the host or port.   

## Commands

### Gauge
`gauge metricName value [flags]`
  * `metricName: (required)` The metric name e.g. 'myapp.request.size'. Cannot contain ':' or '|'
  * `value: (required)` The latest value to set for the metric

  `flags`: 
  * `--labels or -l` adds metadata to a metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma e.g 'env:dev,account_id:12354'

The following are all valid ways to send a gauge metric:

```
# gauge 100
bluematador-client gauge request.size 100

# gauge 100 with labels
bluematador-client gauge request.size 100 --labels 'env:dev,api'

# gauge 100 and set host and port 
bluematador-client gauge request.size 100 --host localhost --port 8181

```

### Count
`count metricName [flags]`
  * `metricName: (required)` The metric name e.g. 'myapp.request.size'. Cannot contain ':' or '|'

  `flags`: 
  * `--value` The amount to increment your metric by, the default value is 1.
  * `--labels or -l` adds metadata to a metric, should be formatted as a key-value pair string with a colon separator e.g 'env:dev'. To send multiple labels seperate each label with a comma e.g 'env:dev,account_id:12354'

The following are all valid ways to send a count metric:

```
# count 1
bluematador-client count app.homepage.clicks

# count 2
bluematador-client count app.homepage.clicks --value 2

# count 2 with labels
bluematador-client count app.homepage.clicks --value 25 --labels 'env:dev'

# count 1 and set host and port
bluematador-client count app.homepage.clicks --host localhost --port 8181

```

# Additional Information

To install the Blue Matador agent, visit the [Integrations](https://app.bluematador.com/ur/app#/setup/integrations) page in the Blue Matador app.

For more information about Blue Matador, visit the [Blue Matador Website](https://www.bluematador.com)


# License

More details in [LICENSE.](https://github.com/bluematador/bluematador-metrics-client-go/blob/master/LICENSE)

Copyright (c) 2020 [Blue Matador, Inc.](https://www.bluematador.com/)
