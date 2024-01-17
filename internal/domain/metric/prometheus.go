package metric

import "github.com/prometheus/client_golang/prometheus"

// Örnek metrik
var (
	HttpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "HTTP requests total count",
		},
		[]string{"method"},
	)
	GrpcRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "GRPC Requests total count",
		},
		[]string{"rpc_call"},
	)
	RateLimitExceededRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rate_limit_exceeded_requests_total",
			Help: "Rate limit exceeded request count",
		},
		[]string{"client_id"},
	)
	RateLimitProcessedRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rate_limit_processed_requests_total",
			Help: "Rate limit processed requests count",
		},
		[]string{"client_id"},
	)
	RateLimitFailedRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rate_limit_failed_requests_total",
			Help: "Rate limit failed requests count",
		},
		[]string{"client_id"},
	)
	RateLimitNotFoundRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rate_limit_not_found_requests_total",
			Help: "Rate limit not found requests count",
		},
		[]string{"client_id"},
	)
)

func init() {
	prometheus.MustRegister(HttpRequestsTotal)
	prometheus.MustRegister(GrpcRequestsTotal)
	prometheus.MustRegister(RateLimitExceededRequestsTotal)
	prometheus.MustRegister(RateLimitProcessedRequestsTotal)
	prometheus.MustRegister(RateLimitNotFoundRequestsTotal)
	prometheus.MustRegister(RateLimitFailedRequestsTotal)
}
