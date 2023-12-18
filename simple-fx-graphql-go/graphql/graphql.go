package graphql

import (
	"log/slog"

	"github.com/graphql-go/graphql"
	"go.uber.org/fx"
)

func newGraphiQlQuery() *graphql.Object {
	return graphql.NewObject((graphql.ObjectConfig{
		Name:        "Query",
		Description: "Root query",
		Fields:      graphql.Fields{},
	}))
}

func newGraphiQLMutation() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Mutation",
		Description: "Root mutation",
		Fields:      graphql.Fields{},
	})
}

func newGraphiQLSubscription() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name:        "Subscription",
		Description: "Root subscription",
		Fields:      graphql.Fields{},
	})
}

type graphqlDependencies struct {
	fx.In
	Query        *graphql.Object `name:"query"`
	Mutation     *graphql.Object `name:"mutation"`
	Subscription *graphql.Object `name:"subscription"`
}

func newGraphiQlSchema(dependencies graphqlDependencies, logger *slog.Logger) *graphql.Schema {
	query := dependencies.Query
	mutation := dependencies.Mutation
	subscription := dependencies.Subscription

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: func() *graphql.Object {
			if len(query.Fields()) == 0 {
				return nil
			}
			return query
		}(),
		Mutation: func() *graphql.Object {
			if len(mutation.Fields()) == 0 {
				return nil
			}
			return mutation
		}(),
		Subscription: func() *graphql.Object {
			if len(subscription.Fields()) == 0 {
				return nil
			}
			return subscription
		}(),
	})
	if err != nil {
		logger.Error("Failed to create schema", err)
	}
	return &schema
}
