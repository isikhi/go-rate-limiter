package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter"
	"github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/usecase"
	"github.com/isikhi/go-rate-limiter/internal/utility/message"
	"github.com/isikhi/go-rate-limiter/internal/utility/respond"
	"github.com/isikhi/go-rate-limiter/internal/utility/validate"
)

type Handler struct {
	useCase  usecase.RateLimiterUseCase
	validate *validator.Validate
}

func NewHandler(useCase usecase.RateLimiterUseCase, validate *validator.Validate) *Handler {
	return &Handler{
		useCase:  useCase,
		validate: validate,
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	var rateLimitOpts rate_limiter.CreateRateLimitOptionsRequest
	err := json.NewDecoder(r.Body).Decode(&rateLimitOpts)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	errs := validate.Validate(h.validate, rateLimitOpts)
	if errs != nil {
		respond.Errors(w, http.StatusBadRequest, errs)
		return
	}
	rlo, err := h.useCase.CreateRateLimitOptions(r.Context(), &rateLimitOpts)
	if err != nil {
		if err == sql.ErrNoRows {
			respond.Error(w, http.StatusBadRequest, message.ErrBadRequest)
			return
		}
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.Json(w, http.StatusCreated, rlo)
}

func (h *Handler) List(w http.ResponseWriter, r *http.Request) {

	var rateLimitOpts []*rate_limiter.RateLimitOptionsSchema
	ctx := r.Context()

	resp, err := h.useCase.ListRateLimitOptions(ctx)
	if err != nil {
		if errors.Is(err, message.ErrFetchingRateLimits) {
			respond.Error(w, http.StatusInternalServerError, err)
			return
		}
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	rateLimitOpts = resp

	if err != nil {
		respond.Error(w, http.StatusInternalServerError, message.ErrFormingResponse)
		return
	}

	respond.Json(w, http.StatusOK, rateLimitOpts)
}

func (h *Handler) Patch(w http.ResponseWriter, r *http.Request) {
	var rateLimitOpts rate_limiter.PatchRateLimitOptionsRequest
	err := json.NewDecoder(r.Body).Decode(&rateLimitOpts)
	if err != nil {
		respond.Error(w, http.StatusBadRequest, err)
		return
	}

	errs := validate.Validate(h.validate, rateLimitOpts)
	if errs != nil {
		respond.Errors(w, http.StatusBadRequest, errs)
		return
	}
	rlo, err := h.useCase.PatchRateLimitOptions(r.Context(), &rateLimitOpts)
	if err != nil {
		fmt.Println(err)
		if err == sql.ErrNoRows {
			respond.Error(w, http.StatusBadRequest, message.ErrBadRequest)
			return
		}
		respond.Error(w, http.StatusInternalServerError, err)
		return
	}
	respond.Json(w, http.StatusOK, rlo)
}
