package main

import (
	"net/http"
	"strconv"

	"github.com/felixge/httpsnoop"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// Counter for calculate HTTP responses by status code and by method
	// app_http_requests_total{"code": "200", "method": "GET"}
	// https://godoc.org/github.com/prometheus/client_golang/prometheus
	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "puppetdb_proxy_http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and method.",
		},
		[]string{"code", "method"},
	)
	// requestDuration collects sets of histograms for measure HTTP request latencies,
	// partitioned by method, URI and status code.
	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "puppetdb_proxy_http_reuests_duration_seconds",
			Help:    "Time in seconds spent serving HTTP requests.",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "uri", "status_code"},
	)
)

func init() {
	// Register the collectors with Prometheus's default registry.
	prometheus.MustRegister(httpReqs)
	prometheus.MustRegister(requestDuration)
}

func (s *server) metricsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(h, w, r)
		// Increase total counter.
		httpReqs.WithLabelValues(strconv.Itoa(m.Code), r.Method).Inc()
		// Measures histograms.
		requestDuration.WithLabelValues(r.Method, r.URL.Path,
			strconv.Itoa(m.Code)).Observe(m.Duration.Seconds())
	})
}
