//go:generate tinygo build -o plugin-morning/morning.wasm -scheduler=none -target=wasi --no-debug plugin-morning/morning.go

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/nebula-aac/public-projects/wasm-go-plugin/greeting"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	ctx := context.Background()
	p, err := greeting.NewGreeterPlugin(ctx)
	if err != nil {
		return err
	}

	morningPlugin, err := p.Load(ctx, "plugin-morning/morning.wasm")
	if err != nil {
		return err
	}
	defer morningPlugin.Close(ctx)

	reply, err := morningPlugin.Greet(ctx, &greeting.GreetRequest{
		Name: "go-plugin",
	})
	if err != nil {
		return err
	}

	fmt.Println(reply.GetMessage())

	return nil
}
