// Package metrics exposes prometheus metrics for access by the whole app.
package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
)

// SecondaryStoreCounter keeps a count of the occurrences of a read-through
// to the secondary store
var SecondaryStoreCounter = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "secondary_store_read_through_total",
	Help: "The total requests that read through to the secondary store.",
})
