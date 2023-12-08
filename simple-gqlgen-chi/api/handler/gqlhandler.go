package handler

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/nebula-aac/public-projects/simple-gqlgen-chi/graph"
)

func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Post("/query", graphQLHandler())
	router.Get("/playground", playgroundQLHandler("/v1/query"))

	return router
}

func graphQLHandler() http.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return h.ServeHTTP
}

func playgroundQLHandler(endpoint string) http.HandlerFunc {
	//endpoint argument must be same as graphql handler path
	playgroundHandler := playground.Handler("GraphQL", endpoint)

	return playgroundHandler
}
