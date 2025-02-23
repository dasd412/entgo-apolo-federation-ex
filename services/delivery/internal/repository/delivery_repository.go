package repository

import (
	"context"
	"delivery/pkg/ent"
	"entgo.io/contrib/entgql"
)

type (
	deliveryRepository struct {
	}

	DeliveryRepository interface {
		FindOne(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.DeliveryQuery),
		) (*ent.Delivery, error)
		FindAll(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.DeliveryQuery),
		) ([]*ent.Delivery, error)
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.DeliveryWhereInput,
		) (*ent.DeliveryConnection, error)
		CreateOne(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateDeliveryInput,
		) (*ent.Delivery, error)
		UpdateOne(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateDeliveryInput,
		) (*ent.Delivery, error)
		DeleteOne(ctx context.Context, client *ent.Client, id int) error
	}
)

func NewDeliveryRepository() DeliveryRepository {
	return &deliveryRepository{}
}

func (u *deliveryRepository) FindOne(
	ctx context.Context,
	client *ent.Client,
	opts ...func(query *ent.DeliveryQuery),
) (*ent.Delivery, error) {
	query := client.Delivery.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.Only(ctx)
}

func (u *deliveryRepository) FindAll(ctx context.Context, client *ent.Client, opts ...func(query *ent.DeliveryQuery)) ([]*ent.Delivery, error) {
	query := client.Delivery.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.All(ctx)
}

func (u *deliveryRepository) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.DeliveryWhereInput) (*ent.DeliveryConnection, error) {
	return client.Delivery.Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			ent.WithDeliveryFilter(where.Filter),
		)
}

func (u *deliveryRepository) CreateOne(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateDeliveryInput,
) (*ent.Delivery, error) {
	return client.Delivery.
		Create().
		SetInput(input).
		Save(ctx)
}

func (u *deliveryRepository) UpdateOne(ctx context.Context, client *ent.Client, id int, input ent.UpdateDeliveryInput) (*ent.Delivery, error) {
	return client.Delivery.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (u *deliveryRepository) DeleteOne(ctx context.Context, client *ent.Client, id int) error {
	return client.Delivery.
		DeleteOneID(id).
		Exec(ctx)
}
