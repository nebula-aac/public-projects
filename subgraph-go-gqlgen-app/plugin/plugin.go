package main

import (
	"github.com/hashicorp/go-plugin"
	"github.com/nebula-aac/public-projects/shared"
)

func main() {
	var handshakeConfig = plugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "GRAPHQL_PLUGIN",
		MagicCookieValue: "hello",
	}

	pluginMap := map[string]plugin.Plugin{
		"graphql": shared.NewGraphQLServerPlugin().Plugin,
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
	})
}
