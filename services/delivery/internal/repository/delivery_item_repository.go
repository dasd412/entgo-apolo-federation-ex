package repository

import (
	"context"
	"delivery/pkg/ent"
	"entgo.io/contrib/entgql"
)

type (
	deliveryItemRepository struct {
	}

	DeliveryItemRepository interface {
		FindOne(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.DeliveryItemQuery),
		) (*ent.DeliveryItem, error)
		FindAll(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.DeliveryItemQuery),
		) ([]*ent.DeliveryItem, error)
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.DeliveryItemWhereInput,
		) (*ent.DeliveryItemConnection, error)
		CreateOne(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateDeliveryItemInput,
		) (*ent.DeliveryItem, error)
		UpdateOne(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateDeliveryItemInput,
		) (*ent.DeliveryItem, error)
		DeleteOne(ctx context.Context, client *ent.Client, id int) error
	}
)

func NewDeliveryItemRepository() DeliveryItemRepository {
	return &deliveryItemRepository{}
}

func (u *deliveryItemRepository) FindOne(
	ctx context.Context,
	client *ent.Client,
	opts ...func(query *ent.DeliveryItemQuery),
) (*ent.DeliveryItem, error) {
	query := client.DeliveryItem.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.Only(ctx)
}

func (u *deliveryItemRepository) FindAll(ctx context.Context, client *ent.Client, opts ...func(query *ent.DeliveryItemQuery)) ([]*ent.DeliveryItem, error) {
	query := client.DeliveryItem.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.All(ctx)
}

func (u *deliveryItemRepository) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.DeliveryItemWhereInput) (*ent.DeliveryItemConnection, error) {
	return client.DeliveryItem.Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			ent.WithDeliveryItemFilter(where.Filter),
		)
}

func (u *deliveryItemRepository) CreateOne(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateDeliveryItemInput,
) (*ent.DeliveryItem, error) {
	return client.DeliveryItem.
		Create().
		SetInput(input).
		Save(ctx)
}

func (u *deliveryItemRepository) UpdateOne(ctx context.Context, client *ent.Client, id int, input ent.UpdateDeliveryItemInput) (*ent.DeliveryItem, error) {
	return client.DeliveryItem.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (u *deliveryItemRepository) DeleteOne(ctx context.Context, client *ent.Client, id int) error {
	return client.DeliveryItem.
		DeleteOneID(id).
		Exec(ctx)
}
