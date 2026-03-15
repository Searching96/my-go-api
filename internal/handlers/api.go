package handlers

import (
	"github.com/Searching96/my-go-api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// Global Middleware
	r.Use(chimiddle.StripSlashes)

	r.route("/account", func(router chi.Router) {
		// Middleware for /account route
		router.Use(middleware.Authorization)
		router.Get("/coins", GetCoinBalance)
	})
}
