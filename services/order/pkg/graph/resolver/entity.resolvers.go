package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	"order/pkg/ent"
	"order/pkg/graph/gen"
)

// FindOrderByID is the resolver for the findOrderByID field.
func (r *entityResolver) FindOrderByID(ctx context.Context, id int) (*ent.Order, error) {
	panic(fmt.Errorf("not implemented: FindOrderByID - findOrderByID"))
}

// Entity returns gen.EntityResolver implementation.
func (r *Resolver) Entity() gen.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
