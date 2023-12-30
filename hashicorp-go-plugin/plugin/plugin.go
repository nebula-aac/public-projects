package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/nebula-aac/public-projects/hashicorp-go-plugin/shared"
)

// Auth is a real implementation of Auth that validates a user with a dummy check.
type Auth struct{}

func (Auth) Authenticate(username, password string) (bool, string, error) {
	// Dummy authentication logic (replace with your actual logic)
	return username == "user" && password == "pass", "dummy-token", nil
}

func main() {
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: shared.Handshake,
		Plugins: map[string]plugin.Plugin{
			"authenticate": &shared.AuthGRPCPlugin{Impl: &Auth{}},
		},

		// A non-nil value here enables gRPC serving for this plugin...
		GRPCServer: plugin.DefaultGRPCServer,
	})
}
