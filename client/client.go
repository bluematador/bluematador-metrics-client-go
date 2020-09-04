package client

import (
	"os"
	"strconv"
)

type BlueMatadorClient struct {
	host   string
	port   int
	prefix string
}

func NewBlueMatadorClient() *BlueMatadorClient {
	return &BlueMatadorClient{
		host:   getHostEnv(),
		port:   getPortEnv(),
		prefix: "",
	}
}

func NewBlueMatadorClientWithAddress(host string, port int) *BlueMatadorClient {
	return &BlueMatadorClient{
		host:   host,
		port:   port,
		prefix: "",
	}
}

func NewBlueMatadorClientWithPrefix(prefix string) *BlueMatadorClient {
	metricPrefix := Sanitize(prefix, ":")

	return &BlueMatadorClient{
		host:   getHostEnv(),
		port:   getPortEnv(),
		prefix: metricPrefix,
	}
}

func NewBlueMatadorClientWithAddressAndPrefix(host string, port int, prefix string) *BlueMatadorClient {
	metricPrefix := Sanitize(prefix, ":")

	return &BlueMatadorClient{
		host:   host,
		port:   port,
		prefix: metricPrefix,
	}
}

func getPortEnv() int {
	port, err := strconv.Atoi(os.Getenv("BLUEMATADOR_AGENT_PORT"))
	if err != nil {
		port = 8767
	}
	return port
}

func getHostEnv() string {
	host := os.Getenv("BLUEMATADOR_AGENT_HOST")
	if host == "" {
		host = "localhost"
	}
	return host
}
