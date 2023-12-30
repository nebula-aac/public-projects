package shared

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/go-chi/chi/v5"
	"github.com/hashicorp/go-plugin"
	"github.com/nebula-aac/public-projects/graph"
	"github.com/nebula-aac/public-projects/graph/generated"
	"github.com/rs/cors"
)

type GraphQLServer interface {
	Run() error
}

type GraphQLServerPlugin struct {
	plugin.Plugin
	plugin.NetRPCUnsupportedPlugin
}

func NewGraphQLServerPlugin() *GraphQLServerPlugin {
	return &GraphQLServerPlugin{}
}

// Client is required to satisfy the plugin.Plugin interface
func (p *GraphQLServerPlugin) Client(b *plugin.MuxBroker, c *plugin.Client) (interface{}, error) {
	return &GraphQLServerPlugin{}, nil
}

// Server is required to satisfy the plugin.Plugin interface
func (p *GraphQLServerPlugin) Server(b *plugin.MuxBroker) (interface{}, error) {
	return &GraphQLServerPlugin{}, nil
}

const defaultPort = "8080"

func (p *GraphQLServerPlugin) Run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"https://studio.apollographql.com"},
		AllowCredentials: true,
		AllowedHeaders:   []string{"*"},
	}).Handler)

	routerSecret, isSet := os.LookupEnv("ROUTER_SECRET")
	if isSet {
		router.Use(checkRouterAuth(routerSecret))
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))
	router.Handle("/", srv)

	log.Printf("Explore with \"https://studio.apollographql.com/sandbox/explorer?endpoint=http://localhost:" + port + "\"")
	return http.ListenAndServe(":"+port, router)
}

// Produce a middleware to check that the `Router-Authorization` header is set and matches routerSecret
func checkRouterAuth(routerSecret string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Router-Authorization")
			if header != routerSecret {
				http.Error(w, "Authorization required", http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
