package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"github.com/hashicorp/go-plugin"
	"github.com/nebula-aac/public-projects/hashicorp-go-plugin/shared"
)

func run() error {
	// We're a host. Start by launching the plugin process.
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  shared.Handshake,
		Plugins:          shared.PluginMap,
		Cmd:              exec.Command("sh", "-c", os.Getenv("AUTH_PLUGIN")), // Update to use AUTH_PLUGIN
		AllowedProtocols: []plugin.Protocol{plugin.ProtocolNetRPC},
	})
	defer client.Kill()

	// Connect via RPC
	rpcClient, err := client.Client()
	if err != nil {
		return err
	}

	// Request the plugin
	raw, err := rpcClient.Dispense("auth")
	if err != nil {
		return err
	}

	// We should have an Auth instance now!
	auth := raw.(shared.Auth)

	os.Args = os.Args[1:]
	switch os.Args[0] {
	case "authenticate":
		if len(os.Args) < 3 {
			return errors.New("usage: authenticate <username> <password>")
		}

		success, token, err := auth.Authenticate(os.Args[1], os.Args[2])
		if err != nil {
			return err
		}

		if success {
			fmt.Printf("Authentication successful! Token: %s\n", token)
		} else {
			fmt.Println("Authentication failed.")
		}

	default:
		return fmt.Errorf("Please only use 'authenticate', given: %q", os.Args[0])
	}

	return nil
}

func main() {
	// We don't want to see the plugin logs.
	log.SetOutput(ioutil.Discard)

	if err := run(); err != nil {
		fmt.Printf("error: %+v\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
