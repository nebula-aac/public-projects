package pkg

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

func NewMux(lc fx.Lifecycle, logger *slog.Logger) *echo.Echo {
	logger.Info("Executing NewMux")

	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{"GET", "POST", "HEAD", "OPTIONS"},
	}))

	server := &http.Server{Addr: fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")), Handler: e}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting HTTP server.")
			go func() {
				if err := server.ListenAndServe(); err != nil {
					logger.Error(err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("Stopping HTTP server.")
			return server.Shutdown(ctx)
		},
	})

	return e
}
