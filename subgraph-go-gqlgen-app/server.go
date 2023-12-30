package main

import (
	"fmt"
	"log"
	"os/exec"

	hashiplugin "github.com/hashicorp/go-plugin"
	"github.com/nebula-aac/public-projects/shared"
)

func main() {
	var handshakeConfig = hashiplugin.HandshakeConfig{
		ProtocolVersion:  1,
		MagicCookieKey:   "GRAPHQL_PLUGIN",
		MagicCookieValue: "hello",
	}

	pluginMap := map[string]hashiplugin.Plugin{
		"graphql": shared.NewGraphQLServerPlugin().Plugin,
	}

	client := hashiplugin.NewClient(&hashiplugin.ClientConfig{
		HandshakeConfig: handshakeConfig,
		Plugins:         pluginMap,
		Cmd:             exec.Command("./plugin/graphql"),
	})
	defer client.Kill()

	rpcClient, err := client.Client()
	if err != nil {
		log.Fatal(err)
	}

	raw, err := rpcClient.Dispense("graphql")
	if err != nil {
		log.Fatal(err)
	}

	graphqlServer := raw.(shared.GraphQLServer)
	fmt.Println(graphqlServer.Run())
}
