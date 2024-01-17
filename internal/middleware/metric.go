package middleware

import (
	"github.com/isikhi/go-rate-limiter/internal/domain/metric"
	"net/http"
)

func Metric(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metric.HttpRequestsTotal.WithLabelValues(r.Method).Inc()
		next.ServeHTTP(w, r)
	})
}
