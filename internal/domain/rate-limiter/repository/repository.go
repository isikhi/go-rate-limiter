package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter"
	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/constants"
	"github.com/isikhi/go-rate-limiter/internal/utility/message"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"log"
	"strconv"
	"time"
)

type RateLimitOptions interface {
	CreateRateLimitOptions(ctx context.Context, rateLimitOptions *rate_limiter.CreateRateLimitOptionsRequest) (int, error)
	ListRateLimitOptions(ctx context.Context) ([]*rate_limiter.RateLimitOptionsSchema, error)
	ReadRateLimitOptions(ctx context.Context, rateLimitOptionsId int) (*rate_limiter.RateLimitOptionsSchema, error)
	GetRateLimitTokens(ctx context.Context, clientId string) (*rate_limiter.RateLimitSchema, error)
	SetRateLimitTokens(ctx context.Context, clientId string) (*rate_limiter.RateLimitSchema, error)
	DecreaseRateLimitToken(ctx context.Context, clientId string) (*rate_limiter.RateLimitSchema, error)
}

type RateLimitRepository struct {
	db          *sqlx.DB
	redisClient *redis.Client
}

const (
	InsertIntoRateLimitOptions         = "INSERT INTO rate_limit_options (client_id, token_count, duration) VALUES ($1, $2, $3) RETURNING id"
	SelectFromRateLimitOptions         = "SELECT * FROM rate_limit_options ORDER BY created_at DESC"
	SelectFromRateLimitOptionsPaginate = "SELECT * FROM rate_limit_options ORDER BY created_at DESC LIMIT $1 OFFSET $2"
	SelectRateLimitOptionsByID         = "SELECT * FROM rate_limit_options where id = $1"
)
const (
	RATE_LIMITER_PREFIX                 = "rate_limit"
	RATE_LIMIT_OPTIONS_BY_CLIENT_PREFIX = "rate_limit_options"
)

func getRateLimiterKey(clientID string) string {
	return fmt.Sprintf("%s_%s", RATE_LIMITER_PREFIX, clientID)
}
func getRateLimiterOptionsByClientKey(clientID string) string {
	return fmt.Sprintf("%s_%s", RATE_LIMIT_OPTIONS_BY_CLIENT_PREFIX, clientID)
}
func getDurationFromMinutes(minutes int64) time.Duration {
	myInt64Value := int64(minutes)
	durationResult := time.Duration(myInt64Value) * time.Minute
	return durationResult
}

func New(db *sqlx.DB, redisClient *redis.Client) *RateLimitRepository {
	return &RateLimitRepository{db: db, redisClient: redisClient}
}

func (r *RateLimitRepository) CreateRateLimitOptions(ctx context.Context, req *rate_limiter.CreateRateLimitOptionsRequest) (rateLimitOptsId int, err error) {
	if err = r.db.QueryRowContext(ctx, InsertIntoRateLimitOptions, req.ClientID, req.TokenCount, req.Duration).Scan(&rateLimitOptsId); err != nil {
		log.Println(err)
		return 0, errors.New("repository.RateLimitOptions.CreateRateLimitOptions")
	}
	rateLimitOptions := &rate_limiter.RateLimitOptionsSchema{
		ID:         uint64(rateLimitOptsId),
		ClientID:   req.ClientID,
		TokenCount: req.TokenCount,
		Duration:   req.Duration,
		CreatedAt:  time.Time{},
	}
	_, err = r.SetRateLimitOptions(ctx, req.ClientID, rateLimitOptions)
	if err != nil {
		log.Printf("rate limit can not set to redis %v\n", err.Error())
		return rateLimitOptsId, nil
	}

	return rateLimitOptsId, nil
}

func (r *RateLimitRepository) ListRateLimitOptions(ctx context.Context) ([]*rate_limiter.RateLimitOptionsSchema, error) {
	var rlo []*rate_limiter.RateLimitOptionsSchema
	err := r.db.SelectContext(ctx, &rlo, SelectFromRateLimitOptions)
	if err != nil {
		return nil, message.ErrFetchingRateLimits
	}

	return rlo, nil
}

func (r *RateLimitRepository) ReadRateLimitOptions(ctx context.Context, rateLimitOptionsId int) (*rate_limiter.RateLimitOptionsSchema, error) {
	var b rate_limiter.RateLimitOptionsSchema
	err := r.db.GetContext(ctx, &b, SelectRateLimitOptionsByID, rateLimitOptionsId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, message.ErrBadRequest
		}
		return nil, err
	}

	return &b, err
}

func (r *RateLimitRepository) GetRateLimitTokens(ctx context.Context, clientId string) (*rate_limiter.RateLimitSchema, error) {
	rateLimiterKey := getRateLimiterKey(clientId)
	data, _ := r.redisClient.Get(ctx, rateLimiterKey).Result()
	if data == "" {
		return nil, nil
	}
	byteData := []byte(data)
	var rateLimitInfo rate_limiter.RateLimitSchema
	err := json.Unmarshal(byteData, &rateLimitInfo)
	if err != nil {
		return nil, err
	}

	return &rateLimitInfo, nil
}

func (r *RateLimitRepository) SetRateLimitTokens(ctx context.Context, clientId string) (*rate_limiter.RateLimitSchema, error) {
	rateLimiterKey := getRateLimiterKey(clientId)
	rateLimitOptions, err := r.getRateLimitOptionsForClient(ctx, clientId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if rateLimitOptions == nil {
		return nil, errors.New(string(constants.ErrorRateLimitOptionsNotFound))
	}
	rateLimitInfo := rate_limiter.RateLimitSchema{
		RateLimitOptionsId: strconv.FormatUint(rateLimitOptions.ID, 10),
		ClientID:           clientId,
		RemainingToken:     rateLimitOptions.TokenCount,
		MaxToken:           rateLimitOptions.TokenCount,
		ExpireAt:           time.Now().Add(getDurationFromMinutes(rateLimitOptions.Duration)).Unix(),
		LastRequestTime:    time.Now().Unix(),
	}
	rateLimitInfoJson, _ := json.Marshal(rateLimitInfo)
	err = r.redisClient.Set(ctx, rateLimiterKey, rateLimitInfoJson, getDurationFromMinutes(rateLimitOptions.Duration)).Err()
	if err != nil {
		return nil, err
	}
	return &rateLimitInfo, nil
}
func (r *RateLimitRepository) SetRateLimitOptions(ctx context.Context, clientId string, rateLimitOptions *rate_limiter.RateLimitOptionsSchema) (bool, error) {
	rateLimiterOptionsByClientKey := getRateLimiterOptionsByClientKey(clientId)
	rateLimitOptsJson, _ := json.Marshal(rateLimitOptions)
	err := r.redisClient.Set(ctx, rateLimiterOptionsByClientKey, rateLimitOptsJson, -1).Err()
	if err != nil {
		log.Println(err)
		return false, err
	}
	_, err = r.SetRateLimitTokens(ctx, clientId) // If i combine without calling again it will be faster. But for now not important.
	if err != nil {
		return false, err
	}
	return true, nil
}
func (r *RateLimitRepository) DecreaseRateLimitToken(ctx context.Context, clientId string) (*rate_limiter.RateLimitSchema, error) {
	rateLimiterKey := getRateLimiterKey(clientId)
	var rateLimitInfo *rate_limiter.RateLimitSchema
	val, err := r.redisClient.Get(ctx, rateLimiterKey).Result()
	if err != nil {
		log.Println(err)
	}
	err = json.Unmarshal([]byte(val), &rateLimitInfo)
	if err != nil {
		return nil, err
	}
	if rateLimitInfo.RemainingToken <= 0 {
		return nil, nil
	}
	rateLimitInfo.RemainingToken--
	rateLimitInfo.LastRequestTime = time.Now().Unix()
	newVal, err := json.Marshal(rateLimitInfo)
	if err != nil {
		return nil, err
	}
	err = r.redisClient.Set(ctx, rateLimiterKey, newVal, time.Until(time.Unix(rateLimitInfo.ExpireAt, 0))).Err()
	if err != nil {
		return nil, err
	}
	return rateLimitInfo, nil
}
func (r *RateLimitRepository) getRateLimitOptionsForClient(ctx context.Context, clientId string) (*rate_limiter.RateLimitOptionsSchema, error) {
	rateLimiterOptionsByClientKey := getRateLimiterOptionsByClientKey(clientId)
	data, _ := r.redisClient.Get(ctx, rateLimiterOptionsByClientKey).Result()
	if data == "" {
		return nil, nil
	}
	var rateLimitOptions = rate_limiter.RateLimitOptionsSchema{}
	err := json.Unmarshal([]byte(data), &rateLimitOptions)
	if err != nil {
		return nil, err
	}
	return &rateLimitOptions, nil
}
