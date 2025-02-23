package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"delivery/pkg/ent"
	"delivery/pkg/graph/gen"
	"entgo.io/contrib/entgql"
)

// Deliveries is the resolver for the deliveries field.
func (r *queryResolver) Deliveries(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.DeliveryWhereInput) (*ent.DeliveryConnection, error) {
	return r.deliveryService.Paginate(
		ctx,
		r.entClient,
		after,
		first,
		before,
		last,
		where)
}

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
