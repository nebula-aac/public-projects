// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

// Injectors from wire.go:

func InitializeEcho() App {
	echo := NewEcho()
	app := NewApp(echo)
	return app
}
