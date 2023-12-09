package main

import "github.com/labstack/echo/v4"

type App struct {
	Echo *echo.Echo
}

func NewApp(e *echo.Echo) App {
	return App{
		Echo: e,
	}
}

func NewEcho() *echo.Echo {
	e := echo.New()

	return e
}

func (a App) Start() {
	a.Echo.Logger.Fatal(a.Echo.Start(":8080"))
}

func main() {
	e := InitializeEcho()

	e.Start()
}
