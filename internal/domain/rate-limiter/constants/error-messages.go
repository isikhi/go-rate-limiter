package constants

type RateLimiterErrorType string

const (
	ErrorRateLimitOptionsNotFound RateLimiterErrorType = "error.rate_limit_options.not_found"
	ErrorRateLimitNotFound        RateLimiterErrorType = "error.rate_limit.not_found"
	ErrorRateLimitExceeded        RateLimiterErrorType = "error.rate_limit.exceeded"
)
