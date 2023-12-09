package main

import (
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	slogEcho "github.com/samber/slog-echo"
)

type App struct {
	Logger    *slog.Logger
	Echo      *echo.Echo
	EchoRoute *EchoHandler
}

// NewApp creates a new App instance
func NewApp(logger *slog.Logger, e *echo.Echo, eh *EchoHandler) App {
	return App{
		Logger:    logger,
		Echo:      e,
		EchoRoute: eh,
	}
}

// NewLogger creates a new *slog.Logger with a default handler
func NewLogger() *slog.Logger {
	logger := slog.NewJSONHandler(os.Stderr, nil)

	slog.SetDefault(slog.New(logger))

	return slog.Default()
}

// NewEcho creates a new Echo instance with middleware
func NewEcho(logger *slog.Logger) *echo.Echo {
	e := echo.New()
	e.Use(slogEcho.New(logger))
	e.Use(middleware.Recover())
	return e
}

// InitializeEcho initializes the echo server
func (a App) Start() {
	a.Echo.GET("/hello", a.EchoRoute.helloHandler)

	a.Echo.Logger.Fatal(a.Echo.Start(":8080"))
}

func main() {
	e := InitializeEcho()

	e.Start()
}

type EchoHandler struct {
	log *slog.Logger
}

func NewEchoHandler(log *slog.Logger) *EchoHandler {
	return &EchoHandler{log: log}
}

func (eh *EchoHandler) helloHandler(c echo.Context) error {
	eh.log.Info("Handling Hello request")
	return c.String(200, "Hello, Echo!")
}
