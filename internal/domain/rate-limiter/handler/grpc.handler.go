package handler

import (
	"context"
	"github.com/isikhi/go-rate-limiter/internal/domain/metric"
	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/constants"
	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/usecase"
	pb "github.com/isikhi/go-rate-limiter/proto/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GrpcHandler struct {
	pb.UnimplementedRateLimitServer
	useCase usecase.RateLimiterUseCase
}

func NewGrpcHandler(uc usecase.RateLimiterUseCase) *GrpcHandler {
	return &GrpcHandler{useCase: uc}
}
func (s *GrpcHandler) CheckRateLimit(ctx context.Context, in *pb.CheckRequest) (*pb.CheckResponse, error) {
	metric.GrpcRequestsTotal.WithLabelValues("CheckRateLimit").Inc()
	checkRateLimit, err := s.useCase.CheckRateLimit(ctx, in.ClientId)
	if err != nil {

		switch err.Error() {
		case string(constants.ErrorRateLimitOptionsNotFound):
			metric.RateLimitNotFoundRequestsTotal.WithLabelValues(in.ClientId).Inc()
			return nil, status.Error(codes.NotFound, "Rate Limit not found.")
		case string(constants.ErrorRateLimitExceeded):
			metric.RateLimitExceededRequestsTotal.WithLabelValues(in.ClientId).Inc()
			return nil, status.Error(codes.ResourceExhausted, "Rate Limit exceeded.")
		default:
			metric.RateLimitFailedRequestsTotal.WithLabelValues(in.ClientId).Inc()
			return nil, status.Errorf(codes.Unknown, "Unknown error.")
		}
	}

	if checkRateLimit != nil && checkRateLimit.RemainingToken > 0 {
		metric.RateLimitProcessedRequestsTotal.WithLabelValues(in.ClientId).Inc()
		return &pb.CheckResponse{
			ClientId:       checkRateLimit.ClientID,
			RemainingToken: int64(checkRateLimit.RemainingToken),
			MaxToken:       int64(checkRateLimit.MaxToken),
			ExpireAt:       checkRateLimit.ExpireAt,
		}, nil
	}
	metric.RateLimitExceededRequestsTotal.WithLabelValues(in.ClientId).Inc()
	return nil, status.Errorf(codes.ResourceExhausted, "Rate Limit exceeded.")
}
