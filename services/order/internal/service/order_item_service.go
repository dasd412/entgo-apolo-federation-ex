package service

import (
	"context"
	"entgo.io/contrib/entgql"
	"order/internal/repository"
	"order/pkg/ent"
	"order/pkg/ent/orderitem"
)

type (
	orderItemService struct {
		orderItemRepository repository.OrderItemRepository
	}

	OrderItemService interface {
		FindOrderItem(
			ctx context.Context,
			client *ent.Client,
			id int,
		) (*ent.OrderItem, error)
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.OrderItemWhereInput,
		) (*ent.OrderItemConnection, error)
		CreateOrderItem(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateOrderItemInput,
		) (*ent.OrderItem, error)
		UpdateOrderItem(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateOrderItemInput,
		) (*ent.OrderItem, error)
		DeleteOrderItem(ctx context.Context, client *ent.Client, id int) (
			bool,
			error,
		)
	}
)

func NewOrderItemService(orderItemRepository repository.OrderItemRepository) OrderItemService {
	return &orderItemService{
		orderItemRepository: orderItemRepository,
	}
}

func (s *orderItemService) FindOrderItem(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.OrderItem, error) {
	return s.orderItemRepository.FindOne(
		ctx, client, func(query *ent.OrderItemQuery) {
			query.Where(orderitem.ID(id))
		},
	)
}

func (s *orderItemService) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.OrderItemWhereInput) (*ent.OrderItemConnection, error) {
	return s.orderItemRepository.Paginate(ctx, client, after, first, before, last, where)
}

func (s *orderItemService) CreateOrderItem(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateOrderItemInput,
) (*ent.OrderItem, error) {
	return s.orderItemRepository.CreateOne(
		ctx, client, input,
	)
}

func (s *orderItemService) UpdateOrderItem(
	ctx context.Context,
	client *ent.Client,
	id int,
	input ent.UpdateOrderItemInput,
) (*ent.OrderItem, error) {
	return s.orderItemRepository.UpdateOne(
		ctx, client, id, input,
	)
}

func (s *orderItemService) DeleteOrderItem(
	ctx context.Context,
	client *ent.Client,
	id int,
) (bool, error) {
	var success = false

	err := s.orderItemRepository.DeleteOne(ctx, client, id)

	if err == nil {
		success = true
	}

	return success, err
}
