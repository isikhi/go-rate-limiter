package authentication

import (
	"fmt"
	"github.com/gmhafiz/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/isikhi/go-rate-limiter/config"

	"github.com/isikhi/go-rate-limiter/internal/middleware"
)

func RegisterHTTPEndPoints(router *chi.Mux, session *scs.SessionManager, repo Repo, cfg *config.Config) {
	h := NewHandler(session, repo)

	router.Post("/api/v1/login", h.Login)
	router.Post("/api/v1/register", h.Register)

	router.Route("/api/v1/logout", func(router chi.Router) {
		router.Post("/", h.Logout)
	})

	router.Route("/api/v1/private", func(router chi.Router) {
		switch cfg.Api.AuthType {
		case "jwt":
			fmt.Println("Using JWT authentication")
			router.Use(middleware.AuthenticateWithJWT(cfg.Api))
		case "session":
			fmt.Println("Using Session authentication")
			router.Use(middleware.AuthenticateWithSession(session)) // Burada session'ı nasıl elde edeceğinize bağlıdır
		default:
			fmt.Println("No authentication method specified")
		}
		router.Get("/csrf", h.Csrf)
		router.Get("/", h.Protected)
		router.Get("/me", h.Me)
		router.Post("/logout/{userID}", h.ForceLogout)
	})
}
