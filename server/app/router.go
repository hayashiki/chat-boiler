package app

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
)

func Routing(r *chi.Mux, d *Dependency) {
	r.Use(middleware.Recoverer)
	acceptOrigins := []string{
		"http://localhost:3000",
		"https://hoge.com",
	}

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   acceptOrigins,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "HEAD", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		Debug:            false,
	}).Handler)

	r.Get("/health", HealthCheck)
	r.Group(func(r chi.Router) {
		//r.Route("query", func(r chi.Router) {
		r.Use(middleware.RequestID)
		r.Handle("/", playground.Handler("GraphQL playground", "/query"))
		r.Handle("/query", d.graphQLHandler)
		//})
	})
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("."))
}
