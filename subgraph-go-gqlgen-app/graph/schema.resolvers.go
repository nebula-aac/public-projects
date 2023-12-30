package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"

	"github.com/nebula-aac/public-projects/graph/generated"
	"github.com/nebula-aac/public-projects/graph/model"
)

// Foo is the resolver for the foo field.
func (r *queryResolver) Foo(ctx context.Context, id string) (*model.Foo, error) {
	return FindFoo(id)
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
