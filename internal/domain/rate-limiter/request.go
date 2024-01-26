package rate_limiter

type CreateRateLimitOptionsRequest struct {
	ClientID           string `json:"client_id" validate:"required"`
	TokenCount         int    `json:"token_count" validate:"required"`
	DurationInSeconds  int64  `json:"duration_in_seconds" validate:"required"`
	ThrottlePercentage int    `json:"throttle_percentage"`
	CreatedAt          int64  `json:"created_at"`
}
type PatchRateLimitOptionsRequest struct {
	ClientID           string `json:"client_id" validate:"required"`
	TokenCount         int    `json:"token_count"`
	DurationInSeconds  int64  `json:"duration_in_seconds"`
	ThrottlePercentage int    `json:"throttle_percentage"`
	CreatedAt          int64  `json:"created_at"`
}
