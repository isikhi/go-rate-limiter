package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/usecase"
	pb "github.com/isikhi/go-rate-limiter/proto/v1"
	"google.golang.org/grpc"
)

func RegisterHTTPEndPoints(router *chi.Mux, validator *validator.Validate, uc usecase.RateLimiterUseCase) *Handler {
	h := NewHandler(uc, validator)
	/**
	Kind of Authentication would be great for below routes.
	*/
	router.Route("/api/v1/rate-limit/options", func(router chi.Router) {
		router.Get("/", h.List)
		router.Post("/", h.Create)
		router.Patch("/", h.Patch)
	})
	return h
}

func RegisterGRPCEndpoints(grpcServer *grpc.Server, uc usecase.RateLimiterUseCase) {
	grpcHandler := NewGrpcHandler(uc)
	pb.RegisterRateLimitServer(grpcServer, grpcHandler)
}
