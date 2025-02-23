package service

import (
	"context"
	"entgo.io/contrib/entgql"
	"order/internal/repository"
	"order/pkg/ent"
	"order/pkg/ent/order"
)

type (
	orderService struct {
		orderRepository repository.OrderRepository
	}

	OrderService interface {
		FindOrder(
			ctx context.Context,
			client *ent.Client,
			id int,
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
		CreateOrder(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateOrderInput,
		) (*ent.Order, error)
		UpdateOrder(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateOrderInput,
		) (*ent.Order, error)
		DeleteOrder(ctx context.Context, client *ent.Client, id int) (
			bool,
			error,
		)
	}
)

func NewOrderService(orderRepository repository.OrderRepository) OrderService {
	return &orderService{
		orderRepository: orderRepository,
	}
}

func (s *orderService) FindOrder(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.Order, error) {
	return s.orderRepository.FindOne(
		ctx, client, func(query *ent.OrderQuery) {
			query.Where(order.ID(id))
		},
	)
}

func (s *orderService) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.OrderWhereInput) (*ent.OrderConnection, error) {
	return s.orderRepository.Paginate(ctx, client, after, first, before, last, where)
}

func (s *orderService) CreateOrder(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateOrderInput,
) (*ent.Order, error) {
	return s.orderRepository.CreateOne(
		ctx, client, input,
	)
}

func (s *orderService) UpdateOrder(
	ctx context.Context,
	client *ent.Client,
	id int,
	input ent.UpdateOrderInput,
) (*ent.Order, error) {
	return s.orderRepository.UpdateOne(
		ctx, client, id, input,
	)
}

func (s *orderService) DeleteOrder(
	ctx context.Context,
	client *ent.Client,
	id int,
) (bool, error) {
	var success = false

	err := s.orderRepository.DeleteOne(ctx, client, id)

	if err == nil {
		success = true
	}

	return success, err
}
