package repository

import (
	"context"
	"entgo.io/contrib/entgql"
	"order/pkg/ent"
)

type (
	orderRepository struct {
	}

	OrderRepository interface {
		FindOne(
			ctx context.Context,
			client *ent.Client,
			opts ...func(query *ent.OrderQuery),
		) (*ent.Order, error)
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.OrderWhereInput,
		) (*ent.OrderConnection, error)
		CreateOne(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateOrderInput,
		) (*ent.Order, error)
		UpdateOne(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateOrderInput,
		) (*ent.Order, error)
		DeleteOne(ctx context.Context, client *ent.Client, id int) error
	}
)

func NewOrderRepository() OrderRepository {
	return &orderRepository{}
}

func (u *orderRepository) FindOne(
	ctx context.Context,
	client *ent.Client,
	opts ...func(query *ent.OrderQuery),
) (*ent.Order, error) {
	query := client.Order.
		Query()

	for _, opt := range opts {
		opt(query)
	}

	return query.Only(ctx)
}

func (u *orderRepository) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.OrderWhereInput) (*ent.OrderConnection, error) {
	return client.Order.Query().
		Paginate(
			ctx,
			after,
			first,
			before,
			last,
			ent.WithOrderFilter(where.Filter),
		)
}

func (u *orderRepository) CreateOne(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateOrderInput,
) (*ent.Order, error) {
	return client.Order.
		Create().
		SetInput(input).
		Save(ctx)
}

func (u *orderRepository) UpdateOne(ctx context.Context, client *ent.Client, id int, input ent.UpdateOrderInput) (*ent.Order, error) {
	return client.Order.
		UpdateOneID(id).
		SetInput(input).
		Save(ctx)
}

func (u *orderRepository) DeleteOne(ctx context.Context, client *ent.Client, id int) error {
	return client.Order.
		DeleteOneID(id).
		Exec(ctx)
}
