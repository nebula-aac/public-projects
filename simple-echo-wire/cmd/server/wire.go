//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
)

func InitializeEcho() App {
	wire.Build(NewEcho, NewApp, NewLogger, NewEchoHandler)
	return App{}
}
