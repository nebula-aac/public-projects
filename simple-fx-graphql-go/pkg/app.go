package pkg

import (
	"go.uber.org/fx"
)

func NewApp() *fx.App {
	app := fx.New(
		fx.Provide(
			// NewLogger,
			NewMux,
			// NewDb,
			// graph.NewResolver,
		),
		fx.Invoke(
		// graphql.RegisterRoutes,
		),
	)
	return app
}
