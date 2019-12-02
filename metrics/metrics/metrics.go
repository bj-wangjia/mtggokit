package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
)
// Counter describes a metric that accumulates values monotonically.
// An example of a counter is the number of received HTTP requests.
type Counter interface {
    With(lables prometheus.Labels) Counter
    Add(delta float64)
}

// Gauge describes a metric that takes specific values over time.
// An example of a gauge is the current depth of a job queue.
type Gauge interface {
    With(lables prometheus.Labels) Gauge
    Set(value float64)
    Add(delta float64)
}

// Histogram describes a metric that takes repeated observations of the same
// kind of thing, and produces a statistical summary of those observations,
// typically expressed as quantiles or buckets. An example of a histogram is
// HTTP request latencies.
type Histogram interface {
    With(lables prometheus.Labels) Histogram
    Observe(value float64)
}
