package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "app_http_requests_total",
			Help: "Total number of HTTP requests handled by the app",
		},
		[]string{"path", "method", "code"},
	)

	reqDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "app_request_duration_seconds",
			Help:    "Histogram of request durations for the app",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

func init() {
	// register metrics (safe to call multiple times across builds)
	prometheus.MustRegister(httpReqs)
	prometheus.MustRegister(reqDuration)
}

// Instrument wraps an http.Handler and records request count and duration.
// It does not alter the handler's response body or status codes.
func Instrument(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// capture status code
		lrw := &loggingResponseWriter{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(lrw, r)

		duration := time.Since(start).Seconds()
		reqDuration.WithLabelValues(r.URL.Path, r.Method).Observe(duration)
		httpReqs.WithLabelValues(r.URL.Path, r.Method, http.StatusText(lrw.status)).Inc()
	})
}

// ExposePrometheusHandler returns the handler that should be mounted at /metrics.
func ExposePrometheusHandler() http.Handler {
	return promhttp.Handler()
}
