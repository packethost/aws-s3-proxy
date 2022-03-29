// Package prometheus exposes a Prometheus metrics middleware with custom metrics
package prometheus

import (
	echoprom "github.com/labstack/echo-contrib/prometheus"
)

var secondaryStoreCounter = echoprom.Metric{
	Type:        "counter",
	Name:        "secondary_store_read_through_total",
	ID:          "secStoreReadThruReq",
	Description: "The total requests that read through to the secondary store.",
}

var customMetricList = []*echoprom.Metric{&secondaryStoreCounter}

// Prometheus creates an echo middleware Prometheus metrics endpoint,
// with cutom metrics
func Prometheus() *echoprom.Prometheus {
	return echoprom.NewPrometheus("echo", nil, customMetricList)
}
