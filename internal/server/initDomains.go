package server

import (
	"github.com/isikhi/go-rate-limiter/internal/domain/metric"
	"io/fs"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/isikhi/go-rate-limiter/internal/domain/authentication"
	"github.com/isikhi/go-rate-limiter/internal/domain/health"
	rateLimiterHandler "github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/handler"
	rateLimiterRepo "github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/repository"
	rateLimiterUseCase "github.com/isikhi/go-rate-limiter/internal/domain/rate-limiter/usecase"
	"github.com/isikhi/go-rate-limiter/internal/middleware"
	"github.com/isikhi/go-rate-limiter/internal/utility/respond"
)

func (s *Server) InitDomains() {
	s.initVersion()
	s.initSwagger()
	s.initAuthentication()
	s.initMetric()
	s.initHealth()
	s.initRateLimiter()
}

func (s *Server) initVersion() {
	s.router.Route("/version", func(router chi.Router) {
		router.Use(middleware.Json)
		router.Get("/", func(w http.ResponseWriter, r *http.Request) {
			respond.Json(w, http.StatusOK, map[string]string{"version": s.Version})
		})
	})
}

func (s *Server) initHealth() {
	newHealthRepo := health.NewRepo(s.sqlx)
	newHealthUseCase := health.New(newHealthRepo)
	health.RegisterHTTPEndPoints(s.router, newHealthUseCase)
}
func (s *Server) initMetric() {
	metricUseCase := metric.New()
	metric.RegisterHTTPEndPoints(s.router, metricUseCase)
}

func (s *Server) initSwagger() {
	if s.Config().Api.RunSwagger {
		docsPath, err := fs.Sub(swaggerDocsAssetPath, "docs")
		if err != nil {
			panic(err)
		}

		fileServer := http.FileServer(http.FS(docsPath))

		s.router.HandleFunc("/swagger", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
		})
		s.router.Handle("/swagger/", http.StripPrefix("/swagger", middleware.ContentType(fileServer)))
		s.router.Handle("/swagger/*", http.StripPrefix("/swagger", middleware.ContentType(fileServer)))
	}
}

func (s *Server) initRateLimiter() {
	newRateLimiterRepo := rateLimiterRepo.New(s.sqlx, s.cache)
	rateLimiterUseCase.New(newRateLimiterRepo)
	newRateLimiterUseCase := rateLimiterUseCase.New(newRateLimiterRepo)
	rateLimiterHandler.RegisterHTTPEndPoints(s.router, s.validator, *newRateLimiterUseCase)
	rateLimiterHandler.RegisterGRPCEndpoints(s.rpcServer, *newRateLimiterUseCase)
}

func (s *Server) initAuthentication() {
	repo := authentication.NewRepo(s.ent, s.db, s.session)
	authentication.RegisterHTTPEndPoints(s.router, s.session, repo, s.cfg)
}
