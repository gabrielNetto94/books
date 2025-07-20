package prometheus

// metrics/prometheus.go

import (
	"books/pkg/observability/metrics"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// promCollector is the Prometheus implementation of the Collector interface.
type promCollector struct {
	reg        *prometheus.Registry
	counters   map[string]*prometheus.CounterVec
	gauges     map[string]*prometheus.GaugeVec
	histograms map[string]*prometheus.HistogramVec
	mu         sync.RWMutex
}

// NewPrometheusCollector creates a new Prometheus-backed metrics collector.
// It returns the Collector interface, hiding the concrete implementation.
func NewPrometheusCollector() metrics.Collector {
	return &promCollector{
		reg:        prometheus.NewRegistry(),
		counters:   make(map[string]*prometheus.CounterVec),
		gauges:     make(map[string]*prometheus.GaugeVec),
		histograms: make(map[string]*prometheus.HistogramVec),
	}
}

// NewCounter creates and registers a new Counter metric.
func (c *promCollector) NewCounter(name, help string, labels ...string) (metrics.Counter, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.counters[name]; ok {
		return nil, fmt.Errorf("counter metric '%s' already exists", name)
	}

	counterVec := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	if err := c.reg.Register(counterVec); err != nil {
		return nil, fmt.Errorf("could not register counter '%s': %w", name, err)
	}
	c.counters[name] = counterVec
	return &promCounter{vec: counterVec}, nil
}

// NewGauge creates and registers a new Gauge metric.
func (c *promCollector) NewGauge(name, help string, labels ...string) (metrics.Gauge, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.gauges[name]; ok {
		return nil, fmt.Errorf("gauge metric '%s' already exists", name)
	}

	gaugeVec := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: name,
			Help: help,
		},
		labels,
	)

	if err := c.reg.Register(gaugeVec); err != nil {
		return nil, fmt.Errorf("could not register gauge '%s': %w", name, err)
	}
	c.gauges[name] = gaugeVec
	return &promGauge{vec: gaugeVec}, nil
}

// NewHistogram creates and registers a new Histogram metric.
func (c *promCollector) NewHistogram(name, help string, buckets []float64, labels ...string) (metrics.Histogram, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.histograms[name]; ok {
		return nil, fmt.Errorf("histogram metric '%s' already exists", name)
	}

	histogramVec := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    name,
			Help:    help,
			Buckets: buckets,
		},
		labels,
	)

	if err := c.reg.Register(histogramVec); err != nil {
		return nil, fmt.Errorf("could not register histogram '%s': %w", name, err)
	}
	c.histograms[name] = histogramVec
	return &promHistogram{vec: histogramVec}, nil
}

// Handler returns an http.Handler that exposes the metrics for scraping.
func (c *promCollector) Handler() http.Handler {
	return promhttp.HandlerFor(c.reg, promhttp.HandlerOpts{})
}

// --- Metric Type Wrappers ---
// These structs wrap the underlying Prometheus types to satisfy the interfaces.

// promCounter wraps prometheus.CounterVec to implement the Counter interface.
type promCounter struct {
	vec *prometheus.CounterVec
}

// Inc increments the counter for the given label values.
func (c *promCounter) Inc(labels ...string) {
	c.vec.WithLabelValues(labels...).Inc()
}

// promGauge wraps prometheus.GaugeVec to implement the Gauge interface.
type promGauge struct {
	vec *prometheus.GaugeVec
}

// Set sets the gauge to a specific value for the given label values.
func (g *promGauge) Set(value float64, labels ...string) {
	g.vec.WithLabelValues(labels...).Set(value)
}

// Inc increments the gauge for the given label values.
func (g *promGauge) Inc(labels ...string) {
	g.vec.WithLabelValues(labels...).Inc()
}

// Dec decrements the gauge for the given label values.
func (g *promGauge) Dec(labels ...string) {
	g.vec.WithLabelValues(labels...).Dec()
}

// promHistogram wraps prometheus.HistogramVec to implement the Histogram interface.
type promHistogram struct {
	vec *prometheus.HistogramVec
}

// Observe records a new observation for the given label values.
func (h *promHistogram) Observe(value float64, labels ...string) {
	h.vec.WithLabelValues(labels...).Observe(value)
}

func MetricsMiddleware(appMetrics *metrics.AppMetrics) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		appMetrics.RequestsInFlight.Inc()
		defer appMetrics.RequestsInFlight.Dec()

		c.Next()

		// After request, record metrics
		duration := time.Since(start).Seconds()
		statusCode := strconv.Itoa(c.Writer.Status())
		path := c.Request.URL.Path // Or use c.FullPath() for the matched route path

		appMetrics.RequestDuration.Observe(duration, c.Request.Method, path)
		appMetrics.RequestsTotal.Inc(c.Request.Method, path, statusCode)
	}
}
