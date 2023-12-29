//go:build tinygo.wasm

package main

import (
	"context"

	"github.com/nebula-aac/public-projects/wasm-go-plugin/greeting"
)

func main() {
	greeting.RegisterGreeter(GoodMorning{})
}

type GoodMorning struct{}

var _ greeting.Greeter = (*GoodMorning)(nil)

func (m GoodMorning) Greet(_ context.Context, request *greeting.GreetRequest) (*greeting.GreetReply, error) {
	return &greeting.GreetReply{
		Message: "Good morning, " + request.GetName(),
	}, nil
}
