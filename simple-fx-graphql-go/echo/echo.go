package echo

import (
	"context"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

type EchoParams struct {
	fx.In
	Logger *slog.Logger
}

func newEcho(params EchoParams) *echo.Echo {
	e := echo.New()

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				logger.LogAttrs(context.Background(), slog.LevelInfo, "REQUEST",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
				)
			} else {
				logger.LogAttrs(context.Background(), slog.LevelError, "REQUEST_ERROR",
					slog.String("uri", v.URI),
					slog.Int("status", v.Status),
					slog.String("err", v.Error.Error()),
				)
			}
			return nil
		},
	}))

	return e
}

func useEcho(lifecycle fx.Lifecycle, e *echo.Echo, logger *slog.Logger) {
	// addr := env.StringWithDefault("FIBER_ADDR", ":3000")

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := e.Start(":8080"); err != nil {
					logger.Error("", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := e.Shutdown(ctx); err != nil {
				logger.Error("", err)
			}
			return nil
		},
	})
}
