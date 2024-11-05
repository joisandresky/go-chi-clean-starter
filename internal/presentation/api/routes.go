package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"

	mw "github.com/joisandresky/go-chi-clean-starter/internal/presentation/middleware"
)

func SetupRoutes(
	testMw mw.TestMiddleware,
	postHttpApi PostHttpApi,
) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is Running..."))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	// all routes registration
	postHttpApi.RegisterRoutes(r, testMw)

	return r
}
