// Package prometheus exposes a Prometheus metrics middleware with custom metrics
package prometheus

import (
	echoprom "github.com/labstack/echo-contrib/prometheus"
)

// Prometheus creates an echo middleware Prometheus metrics endpoint,
// with cutom metrics
func Prometheus() *echoprom.Prometheus {
	return echoprom.NewPrometheus("echo", nil)
}
