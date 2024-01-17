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
	ListRateLimitOptions(ctx context.Context) ([]*rate_limiter.RateLimitOptionsSchema, error)
}

type RateLimiterUseCase struct {
	rateLimitRepo repository.RateLimitOptions
	//	throttlePercentage int
}

func New(rateLimitRepo repository.RateLimitOptions) *RateLimiterUseCase {
	return &RateLimiterUseCase{
		rateLimitRepo: rateLimitRepo,
	}
}

//func (u *RateLimiterUseCase) SetThrottlePercentage(percentage int) {
//ask about dynamic throttle logic. rate limit service load or client load is it global or client specific etc.
//	u.throttlePercentage = percentage
//	rateLimit.RemainingToken = (u.throttlePercentage * rateLimit.RemainingToken) / 100
//}

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
			return &rate_limiter.RateLimitSchema{}, err
		}
	}
	if rateLimit.RemainingToken <= 0 {
		return rateLimit, errors.New(string(constants.ErrorRateLimitExceeded))
	}
	rateLimit, err = u.rateLimitRepo.DecreaseRateLimitToken(ctx, clientId)

	if err != nil {
		log.Println(err)
		return nil, err
	}
	return rateLimit, nil

}
