package repository

import (
	"context"
	"fmt"
	"github.com/go-redis/redismock/v9"
	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
	"time"
)

const (
	SelectOneFromRateLimitOptionsMock = "SELECT \\* FROM rate_limit_options WHERE client_id = \\$1 ORDER BY created_at DESC"
	InsertIntoRateLimitOptionsMock    = "INSERT INTO rate_limit_options (.+) RETURNING id"
	UpdateRateLimitOptionsMock        = "UPDATE rate_limit_options SET token_count = \\$1, duration_in_seconds = \\$2, throttle_percentage = \\$3 WHERE client_id = \\$4 RETURNING id"
	SelectFromRateLimitOptionsMock    = "SELECT \\* FROM rate_limit_options ORDER BY created_at DESC"
	SelectRateLimitOptionsByIDMock    = "SELECT \\* FROM rate_limit_options where id = \\$1"
)

func TestRateLimitRepository_CreateRateLimitOptions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	redisMock := redis.NewClient(&redis.Options{})

	repo := New(sqlxDB, redisMock)

	request := &rate_limiter.CreateRateLimitOptionsRequest{
		ClientID:           "test-client",
		TokenCount:         10,
		DurationInSeconds:  60,
		ThrottlePercentage: 20,
	}

	mock.ExpectQuery(InsertIntoRateLimitOptionsMock).
		WithArgs(request.ClientID, request.TokenCount, request.DurationInSeconds, request.ThrottlePercentage).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1)) // Call the function being tested
	resultID, err := repo.CreateRateLimitOptions(context.Background(), request)

	fmt.Println(resultID)
	assert.NoError(t, err)
	assert.Equal(t, 1, resultID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRateLimitRepository_PatchRateLimitOptions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	redisMock := redis.NewClient(&redis.Options{})

	repo := New(sqlxDB, redisMock)

	request := &rate_limiter.PatchRateLimitOptionsRequest{
		ClientID:           "test-client",
		TokenCount:         15,
		DurationInSeconds:  120,
		ThrottlePercentage: 30,
	}

	mock.ExpectQuery(SelectOneFromRateLimitOptionsMock).
		WithArgs(request.ClientID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "client_id", "token_count", "duration_in_seconds", "throttle_percentage", "created_at"}).
			AddRow(1, request.ClientID, 10, 60, 20, time.Now()))

	mock.ExpectExec(UpdateRateLimitOptionsMock).
		WithArgs(request.TokenCount, request.DurationInSeconds, request.ThrottlePercentage, request.ClientID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	result, err := repo.PatchRateLimitOptions(context.Background(), request)
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, request.TokenCount, result.TokenCount)
	assert.Equal(t, request.DurationInSeconds, result.DurationInSeconds)
	assert.Equal(t, request.ThrottlePercentage, result.ThrottlePercentage)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRateLimitRepository_ListRateLimitOptions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	redisMock := redis.NewClient(&redis.Options{})
	repo := New(sqlxDB, redisMock)

	expectedOptions := []*rate_limiter.RateLimitOptionsSchema{
		{ID: 1, ClientID: "client-1", TokenCount: 10, DurationInSeconds: 60, ThrottlePercentage: 20},
		{ID: 2, ClientID: "client-2", TokenCount: 15, DurationInSeconds: 120, ThrottlePercentage: 30},
	}

	rows := sqlmock.NewRows([]string{"id", "client_id", "token_count", "duration_in_seconds", "throttle_percentage"}).
		AddRow(expectedOptions[0].ID, expectedOptions[0].ClientID, expectedOptions[0].TokenCount, expectedOptions[0].DurationInSeconds, expectedOptions[0].ThrottlePercentage).
		AddRow(expectedOptions[1].ID, expectedOptions[1].ClientID, expectedOptions[1].TokenCount, expectedOptions[1].DurationInSeconds, expectedOptions[1].ThrottlePercentage)

	mock.ExpectQuery(SelectFromRateLimitOptionsMock).WillReturnRows(rows)

	resultOptions, err := repo.ListRateLimitOptions(context.Background())

	assert.NoError(t, err)
	assert.NotNil(t, resultOptions)
	assert.Equal(t, len(expectedOptions), len(resultOptions))
	assert.Equal(t, expectedOptions, resultOptions)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRateLimitRepository_ReadRateLimitOptions(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	redisMock := redis.NewClient(&redis.Options{})

	repo := New(sqlxDB, redisMock)

	rateLimitOptionsID := 1
	expectedOption := &rate_limiter.RateLimitOptionsSchema{
		ID:                 uint64(rateLimitOptionsID),
		ClientID:           "client-1",
		TokenCount:         10,
		DurationInSeconds:  60,
		ThrottlePercentage: 20,
	}

	rows := sqlmock.NewRows([]string{"id", "client_id", "token_count", "duration_in_seconds", "throttle_percentage"}).
		AddRow(expectedOption.ID, expectedOption.ClientID, expectedOption.TokenCount, expectedOption.DurationInSeconds, expectedOption.ThrottlePercentage)

	mock.ExpectQuery(SelectRateLimitOptionsByIDMock).WithArgs(rateLimitOptionsID).WillReturnRows(rows)

	resultOption, err := repo.ReadRateLimitOptions(context.Background(), rateLimitOptionsID)

	assert.NoError(t, err)
	assert.NotNil(t, resultOption)
	assert.Equal(t, expectedOption, resultOption)

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestRateLimitRepository_GetRateLimitTokens(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error setting up mock database: %v", err)
	}
	defer db.Close()

	sqlxDB := sqlx.NewDb(db, "sqlmock")

	redisClient, redisMock := redismock.NewClientMock()
	repo := New(sqlxDB, redisClient)
	clientID := "test-client"
	expectedTokens := &rate_limiter.RateLimitSchema{
		RateLimitOptionsId: "1",
		ClientID:           clientID,
		RemainingTokens:    10,
		MaxToken:           10,
		ExpireAt:           1722123425,
		LastRequestTime:    1722123365,
	}

	redisMock.ExpectGet(fmt.Sprintf("%s_%s", "rate_limit", clientID)).
		SetVal("{\"rate_limit_options_id\":\"1\",\"client_id\":\"test-client\",\"remaining_tokens\":10,\"max_token\":10,\"expire_at\":1722123425,\"last_request_time\":1722123365}")
	resultTokens, err := repo.GetRateLimitTokens(context.Background(), clientID)
	fmt.Println(resultTokens)
	fmt.Println(clientID)
	assert.NoError(t, err)
	assert.NotNil(t, resultTokens)
	assert.Equal(t, expectedTokens, resultTokens)

	assert.NoError(t, mock.ExpectationsWereMet())
}
