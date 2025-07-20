// metrics/metrics.go

package metrics

import "net/http"

// Counter is an interface for a metric that only goes up.
type Counter interface {
	Inc(labels ...string)
}

// Gauge is an interface for a metric that can go up and down.
type Gauge interface {
	Set(value float64, labels ...string)
	Inc(labels ...string)
	Dec(labels ...string)
}

// Histogram is an interface for a metric that samples observations.
type Histogram interface {
	Observe(value float64, labels ...string)
}

// Collector is the main interface for creating and managing metrics.
type Collector interface {
	// NewCounter creates and registers a new Counter metric.
	NewCounter(name, help string, labels ...string) (Counter, error)

	// NewGauge creates and registers a new Gauge metric.
	NewGauge(name, help string, labels ...string) (Gauge, error)

	// NewHistogram creates and registers a new Histogram metric.
	NewHistogram(name, help string, buckets []float64, labels ...string) (Histogram, error)

	// Handler returns an http.Handler that exposes the metrics for scraping.
	Handler() http.Handler
}

type AppMetrics struct {
	RequestsTotal    Counter
	RequestDuration  Histogram
	RequestsInFlight Gauge
}

// NewAppMetrics creates and registers the default application metrics.
func NewAppMetrics(collector Collector) (*AppMetrics, error) {
	var err error
	appMetrics := &AppMetrics{}

	appMetrics.RequestsTotal, err = collector.NewCounter(
		"http_requests_total",
		"Total number of HTTP requests.",
		"method", "path", "status_code",
	)
	if err != nil {
		return nil, err
	}

	appMetrics.RequestDuration, err = collector.NewHistogram(
		"http_request_duration_seconds",
		"Histogram of HTTP request latencies.",
		[]float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 2.5},
		"method", "path",
	)
	if err != nil {
		return nil, err
	}

	appMetrics.RequestsInFlight, err = collector.NewGauge(
		"http_requests_in_flight",
		"Number of current in-flight HTTP requests.",
	)
	if err != nil {
		return nil, err
	}

	return appMetrics, nil
}
