package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"delivery/pkg/ent"
	"delivery/pkg/graph/gen"
	"delivery/pkg/graph/gen/graphqlmodel"
	"fmt"
)

// FindDeliveryByID is the resolver for the findDeliveryByID field.
func (r *entityResolver) FindDeliveryByID(ctx context.Context, id int) (*ent.Delivery, error) {
	return r.deliveryService.FindDelivery(ctx, r.entClient, id)
}

// FindDeliveryItemByID is the resolver for the findDeliveryItemByID field.
func (r *entityResolver) FindDeliveryItemByID(ctx context.Context, id int) (*ent.DeliveryItem, error) {
	panic(fmt.Errorf("not implemented: FindDeliveryItemByID - findDeliveryItemByID"))
}

// FindOrderByID is the resolver for the findOrderByID field.
func (r *entityResolver) FindOrderByID(ctx context.Context, id int) (*graphqlmodel.Order, error) {
	delivery, err := r.deliveryService.FindDeliveryByOrderId(ctx, r.entClient, id)

	if ent.IsNotFound(err) {
		return &graphqlmodel.Order{
			ID:       id,
			Delivery: nil,
		}, nil
	}

	if err != nil {
		return nil, err
	}

	return &graphqlmodel.Order{
		ID:       id,
		Delivery: delivery,
	}, nil
}

// FindUserByID is the resolver for the findUserByID field.
func (r *entityResolver) FindUserByID(ctx context.Context, id int) (*graphqlmodel.User, error) {
	deliveries, err := r.deliveryService.FindDeliveriesByUserId(ctx, r.entClient, id)

	if err != nil {
		return nil, err
	}

	return &graphqlmodel.User{
		ID:         id,
		Deliveries: deliveries,
	}, nil
}

// Entity returns gen.EntityResolver implementation.
func (r *Resolver) Entity() gen.EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
