package metric

import (
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func RegisterHTTPEndPoints(router *chi.Mux, uc UseCase) *Handler {
	h := NewHandler(uc)

	router.Route("/metric", func(router chi.Router) {
		router.Handle("/", promhttp.Handler())
	})

	return h
}
