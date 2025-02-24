package repository

import (
	"context"
	"entgo.io/contrib/entgql"
	"order/pkg/ent"
)

type (
	orderItemRepository struct {
	}

	OrderItemRepository interface {
		FindOne(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.OrderItemQuery),
		) (*ent.OrderItem, error)
		FindAll(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.OrderItemQuery),
		) ([]*ent.OrderItem, error)
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.OrderItemWhereInput,
		) (*ent.OrderItemConnection, error)
		CreateOne(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateOrderItemInput,
		) (*ent.OrderItem, error)
		UpdateOne(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateOrderItemInput,
		) (*ent.OrderItem, error)
		DeleteOne(ctx context.Context, client *ent.Client, id int) error
	}
)

func NewOrderItemRepository() OrderItemRepository {
	return &orderItemRepository{}
}

func (u *orderItemRepository) FindOne(
	ctx context.Context,
	client *ent.Client,
	opts ...func(query *ent.OrderItemQuery),
) (*ent.OrderItem, error) {
	query := client.OrderItem.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.Only(ctx)
}

func (u *orderItemRepository) FindAll(ctx context.Context, client *ent.Client, opts ...func(query *ent.OrderItemQuery)) ([]*ent.OrderItem, error) {
	query := client.OrderItem.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.All(ctx)
}

func (u *orderItemRepository) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.OrderItemWhereInput) (*ent.OrderItemConnection, error) {
	return client.OrderItem.Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			ent.WithOrderItemFilter(where.Filter),
		)
}

func (u *orderItemRepository) CreateOne(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateOrderItemInput,
) (*ent.OrderItem, error) {
	return client.OrderItem.
		Create().
		SetInput(input).
		Save(ctx)
}

func (u *orderItemRepository) UpdateOne(ctx context.Context, client *ent.Client, id int, input ent.UpdateOrderItemInput) (*ent.OrderItem, error) {
	return client.OrderItem.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (u *orderItemRepository) DeleteOne(ctx context.Context, client *ent.Client, id int) error {
	return client.OrderItem.
		DeleteOneID(id).
		Exec(ctx)
}
