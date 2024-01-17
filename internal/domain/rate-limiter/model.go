package rate_limiter

import (
	"time"
)

type RateLimitOptionsSchema struct {
	ID         uint64    `json:"id"`
	ClientID   string    `json:"client_id" db:"client_id"`
	TokenCount int       `json:"token_count" db:"token_count"`
	Duration   int64     `json:"duration" db:"duration"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}

type RateLimitSchema struct {
	RateLimitOptionsId string `json:"rate_limit_options_id" db:"rate_limit_options_id"`
	ClientID           string `json:"client_id" db:"client_id"`
	RemainingToken     int    `json:"remaining_token" db:"remaining_token"`
	MaxToken           int    `json:"max_token" db:"max_token"`
	ExpireAt           int64  `json:"expire_at" db:"expire_at"`
	LastRequestTime    int64  `json:"last_request_time" db:"last_request_time"`
}
