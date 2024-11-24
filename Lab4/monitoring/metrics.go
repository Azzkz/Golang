package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"method", "path"},
	)
	ErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_request_errors_total",
			Help: "Total number of HTTP request errors.",
		},
	)

	requestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)
)

func InitMetrics() {
	prometheus.MustRegister(requestCount, requestDuration)
	http.Handle("/metrics", promhttp.Handler())
}

func TrackMetrics(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timer := prometheus.NewTimer(requestDuration.WithLabelValues(r.Method, r.URL.Path))
		defer timer.ObserveDuration()

		requestCount.WithLabelValues(r.Method, r.URL.Path).Inc()
		handler.ServeHTTP(w, r)
	})
}
