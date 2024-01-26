package usecase

import (
	"context"
	"errors"
	rate_limiter "github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter"
	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/constants"
	"log"

	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/repository"
)

type RateLimiter interface {
	CreateRateLimitOptions(ctx context.Context, rateLimitOpts *rate_limiter.CreateRateLimitOptionsRequest) (*rate_limiter.RateLimitOptionsSchema, error)
	PatchRateLimitOptions(ctx context.Context, rateLimitOpts *rate_limiter.PatchRateLimitOptionsRequest) (*rate_limiter.RateLimitOptionsSchema, error)
	ListRateLimitOptions(ctx context.Context) ([]*rate_limiter.RateLimitOptionsSchema, error)
}

type RateLimiterUseCase struct {
	rateLimitRepo repository.RateLimitOptions
}

func New(rateLimitRepo repository.RateLimitOptions) *RateLimiterUseCase {
	return &RateLimiterUseCase{
		rateLimitRepo: rateLimitRepo,
	}
}

func (u *RateLimiterUseCase) CreateRateLimitOptions(ctx context.Context, rateLimitOptions *rate_limiter.CreateRateLimitOptionsRequest) (*rate_limiter.RateLimitOptionsSchema, error) {
	rateLimitOptsID, err := u.rateLimitRepo.CreateRateLimitOptions(ctx, rateLimitOptions)
	if err != nil {
		return nil, err
	}
	rateLimitOptsFound, err := u.rateLimitRepo.ReadRateLimitOptions(ctx, rateLimitOptsID)
	if err != nil {
		return nil, err
	}
	return rateLimitOptsFound, err
}

func (u *RateLimiterUseCase) PatchRateLimitOptions(ctx context.Context, rateLimitOptions *rate_limiter.PatchRateLimitOptionsRequest) (*rate_limiter.RateLimitOptionsSchema, error) {
	rateLimitOpts, err := u.rateLimitRepo.PatchRateLimitOptions(ctx, rateLimitOptions)
	if err != nil {
		return nil, err
	}
	return rateLimitOpts, err
}

func (u *RateLimiterUseCase) ListRateLimitOptions(ctx context.Context) ([]*rate_limiter.RateLimitOptionsSchema, error) {
	return u.rateLimitRepo.ListRateLimitOptions(ctx)
}

func (u *RateLimiterUseCase) CheckRateLimit(ctx context.Context, clientId string) (*rate_limiter.RateLimitSchema, error) {
	var rateLimit *rate_limiter.RateLimitSchema
	rateLimit, err := u.rateLimitRepo.GetRateLimitTokens(ctx, clientId)
	if err != nil {
		log.Println(err)
		return &rate_limiter.RateLimitSchema{}, err
	}

	if rateLimit == nil {
		rateLimit, err = u.rateLimitRepo.SetRateLimitTokens(ctx, clientId)
		if err != nil {
			log.Println(err)
			return &rate_limiter.RateLimitSchema{}, err
		}
	}

	if rateLimit.RemainingTokens <= 0 {
		return rateLimit, errors.New(string(constants.ErrorRateLimitExceeded))
	}
	rateLimit, err = u.rateLimitRepo.DecreaseRateLimitToken(ctx, clientId)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rateLimit, nil

}
