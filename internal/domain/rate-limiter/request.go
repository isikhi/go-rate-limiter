package rate_limiter

type CreateRateLimitOptionsRequest struct {
	ClientID   string `json:"client_id" validate:"required"`
	TokenCount int    `json:"token_count" validate:"required"`
	Duration   int64  `json:"duration" validate:"required"`
	CreatedAt  int64  `json:"created_at"`
}
