package graphql

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		fx.Annotated{
			Name:   "query",
			Target: newGraphiQlQuery,
		},
		fx.Annotated{
			Name:   "mutation",
			Target: newGraphiQLMutation,
		},
		fx.Annotated{
			Name:   "subscription",
			Target: newGraphiQLSubscription,
		},
		newGraphiQlSchema,
	),
)
