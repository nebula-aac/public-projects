package router

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/nebula-aac/public-projects/simple-gqlgen-chi/api/handler"
)

// Initialize will handle the routes for the chi server
func initialize() *chi.Mux {
	router := chi.NewRouter()

	//uses chi middleware logic
	router.Use(
		// render.SetContentType(render.ContentTypeJSON),
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)

	//Sets context for all requests
	router.Use(middleware.Timeout(20 * time.Second))

	// "/v1" groups queries for version v1
	// "/" handlers is for graphiQL
	// "/query" is for queries
	router.Route("/v1", func(r chi.Router) {
		r.Mount("/", handler.Routes())
	})
	return router
}

// ServeRouter initialize the router and serves
func ServeRouter() {
	r := initialize()

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Println("Error serving")
	}
}
