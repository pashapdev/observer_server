package router

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(middleware.Timeout(time.Minute))
	return r
}
