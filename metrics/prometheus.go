// Package prometheus provides Prometheus implementations for metrics.
// Individual metrics are mapped to their Prometheus counterparts, and
// (depending on the constructor used) may be automatically registered in the
// global Prometheus metrics registry.
package metrics

import (
    "github.com/prometheus/client_golang/prometheus"
    "github.com/bj-wangjia/mtggokit/metrics/metrics"
)

// Counter implements Counter, via a Prometheus CounterVec.
type Counter struct {
    cv  *prometheus.CounterVec
}

func NewCounterFromConfig(fileName string, labels []string) Counter {
    v := setViper(fileName)
    baseOpts := prometheus.CounterOpts{
        Namespace: v.GetString("Prometheus.Default.Namespace"),
        Subsystem: v.GetString("Prometheus.Default.Subsystem"),
        Name:      v.GetString("Prometheus.Default.Name"),
        Help:      v.GetString("Prometheus.Default.Help"),
    }
    return NewCounterFrom(baseOpts, labels)
}

// NewCounterFrom constructs and registers a Prometheus CounterVec,
// and returns a usable Counter object.
func NewCounterFrom(opts prometheus.CounterOpts, labelNames []string) *Counter {
    cv := prometheus.NewCounterVec(opts, labelNames)
    prometheus.MustRegister(cv)
    return NewCounter(cv)
}

// NewCounter wraps the CounterVec and returns a usable Counter object.
func NewCounter(cv *prometheus.CounterVec) *Counter {
    return &Counter{
        cv: cv,
    }
}

// With implements Counter.
func (c *Counter) With(labels prometheus.Labels) metrics.Counter {
    c.cv.With(labels)
    return &Counter{
        cv:  c.cv,
    }
}

// Add implements Counter.
func (c *Counter) Add(delta float64) {
    c.cv.Add(delta)
}

// Gauge implements Gauge, via a Prometheus GaugeVec.
type Gauge struct {
    gv  *prometheus.GaugeVec
}

func NewGaugeFromConfig(fileName string, labels []string) Gauge {
    v := setViper(fileName)
    baseOpts := prometheus.GaugeOpts{
        Namespace: v.GetString("Prometheus.Default.Namespace"),
        Subsystem: v.GetString("Prometheus.Default.Subsystem"),
        Name:      v.GetString("Prometheus.Default.Name"),
        Help:      v.GetString("Prometheus.Default.Help"),
    }
    return NewGaugeFrom(baseOpts, labels)
}

// NewGaugeFrom construts and registers a Prometheus GaugeVec,
// and returns a usable Gauge object.
func NewGaugeFrom(opts prometheus.GaugeOpts, labelNames []string) *Gauge {
    gv := prometheus.NewGaugeVec(opts, labelNames)
    prometheus.MustRegister(gv)
    return NewGauge(gv)
}

// NewGauge wraps the GaugeVec and returns a usable Gauge object.
func NewGauge(gv *prometheus.GaugeVec) *Gauge {
    return &Gauge{
        gv: gv,
    }
}

// With implements Gauge.
func (g *Gauge) With(labels prometheus.Labels) metrics.Gauge {
    g.gv.With(labels)
    return &Gauge{
        gv:  g.gv,
    }
}

// Set implements Gauge.
func (g *Gauge) Set(value float64) {
    g.gv.Set(value)
}

// Add is supported by Prometheus GaugeVecs.
func (g *Gauge) Add(delta float64) {
    g.gv.Add(delta)
}

// Summary implements Histogram, via a Prometheus SummaryVec. The difference
// between a Summary and a Histogram is that Summaries don't require predefined
// quantile buckets, but cannot be statistically aggregated.
type Summary struct {
    sv  *prometheus.SummaryVec
}

func NewSummaryFromConfig(fileName string, labels []string) Summary {
    v := setViper(fileName)
    baseOpts := prometheus.SummaryOpts{
        Namespace: v.GetString("Prometheus.Default.Namespace"),
        Subsystem: v.GetString("Prometheus.Default.Subsystem"),
        Name:      v.GetString("Prometheus.Default.Name"),
        Help:      v.GetString("Prometheus.Default.Help"),
    }
    return NewSummaryFrom(baseOpts, labels)
}

// NewSummaryFrom constructs and registers a Prometheus SummaryVec,
// and returns a usable Summary object.
func NewSummaryFrom(opts prometheus.SummaryOpts, labelNames []string) *Summary {
    sv := prometheus.NewSummaryVec(opts, labelNames)
    prometheus.MustRegister(sv)
    return NewSummary(sv)
}

// NewSummary wraps the SummaryVec and returns a usable Summary object.
func NewSummary(sv *prometheus.SummaryVec) *Summary {
    return &Summary{
        sv: sv,
    }
}

// With implements Histogram.
func (s *Summary) With(labels prometheus.Labels) metrics.Histogram {
    s.sv.With(labels)
    return &Summary{
        sv:  s.sv,
    }
}

// Observe implements Histogram.
func (s *Summary) Observe(value float64) {
    s.sv.Observe(value)
}

// Histogram implements Histogram via a Prometheus HistogramVec. The difference
// between a Histogram and a Summary is that Histograms require predefined
// quantile buckets, and can be statistically aggregated.
type Histogram struct {
    hv  *prometheus.HistogramVec
}

// NewHistogramFrom constructs and registers a Prometheus HistogramVec,
// and returns a usable Histogram object.
func NewHistogramFrom(opts prometheus.HistogramOpts, labelNames []string) *Histogram {
    hv := prometheus.NewHistogramVec(opts, labelNames)
    prometheus.MustRegister(hv)
    return NewHistogram(hv)
}

// NewHistogram wraps the HistogramVec and returns a usable Histogram object.
func NewHistogram(hv *prometheus.HistogramVec) *Histogram {
    return &Histogram{
        hv: hv,
    }
}

// With implements Histogram.
func (h *Histogram) With(labels prometheus.Labels) metrics.Histogram {
    h.hv.With(labels)
    return &Histogram{
        hv:  h.hv,
    }
}

// Observe implements Histogram.
func (h *Histogram) Observe(value float64) {
    h.hv.Observe(value)
}

func setViper(fileName string) *viper.Viper {
    configPath, configName := filepath.Split(fileName)
    dotIndex := strings.LastIndex(configName, ".")
    if dotIndex == -1 || configName[dotIndex:] != ".yaml" {
        panic("config file format must be yaml")
    }
    v := viper.New()
    v.AddConfigPath(configPath)
    v.SetConfigName(configName[:dotIndex])
    v.SetConfigType("yaml")
    if err := v.ReadInConfig(); err != nil {
        panic(err)
    }
    return v
}

