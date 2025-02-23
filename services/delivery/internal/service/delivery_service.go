package service

import (
	"context"
	"delivery/internal/repository"
	"delivery/pkg/ent"
	"delivery/pkg/ent/delivery"
	"entgo.io/contrib/entgql"
)

type (
	deliveryService struct {
		deliveryRepository repository.DeliveryRepository
	}

	DeliveryService interface {
		FindDelivery(
			ctx context.Context,
			client *ent.Client,
			id int,
		) (*ent.Delivery, error)
		FindDeliveriesByUserId(
			ctx context.Context,
			client *ent.Client,
			userId int) ([]*ent.Delivery, error)
		FindDeliveryByOrderId(
			ctx context.Context,
			client *ent.Client,
			orderId int) (*ent.Delivery, error)
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.DeliveryWhereInput,
		) (*ent.DeliveryConnection, error)
		CreateDelivery(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateDeliveryInput,
		) (*ent.Delivery, error)
		UpdateDelivery(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateDeliveryInput,
		) (*ent.Delivery, error)
		DeleteDelivery(ctx context.Context, client *ent.Client, id int) (
			bool,
			error,
		)
	}
)

func NewDeliveryService(deliveryRepository repository.DeliveryRepository) DeliveryService {
	return &deliveryService{
		deliveryRepository: deliveryRepository,
	}
}

func (s *deliveryService) FindDelivery(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.Delivery, error) {
	return s.deliveryRepository.FindOne(
		ctx, client, func(query *ent.DeliveryQuery) {
			query.Where(delivery.ID(id))
		},
	)
}

func (s *deliveryService) FindDeliveriesByUserId(ctx context.Context, client *ent.Client, userId int) ([]*ent.Delivery, error) {
	return s.deliveryRepository.FindAll(
		ctx, client, func(query *ent.DeliveryQuery) {
			query.Where(delivery.UserID(userId))
		},
	)
}

func (s *deliveryService) FindDeliveryByOrderId(ctx context.Context, client *ent.Client, orderId int) (*ent.Delivery, error) {
	return s.deliveryRepository.FindOne(
		ctx, client, func(query *ent.DeliveryQuery) {
			query.Where(delivery.OrderID(orderId))
		})
}

func (s *deliveryService) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.DeliveryWhereInput) (*ent.DeliveryConnection, error) {
	return s.deliveryRepository.Paginate(ctx, client, after, first, before, last, where)
}

func (s *deliveryService) CreateDelivery(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateDeliveryInput,
) (*ent.Delivery, error) {
	return s.deliveryRepository.CreateOne(
		ctx, client, input,
	)
}

func (s *deliveryService) UpdateDelivery(
	ctx context.Context,
	client *ent.Client,
	id int,
	input ent.UpdateDeliveryInput,
) (*ent.Delivery, error) {
	return s.deliveryRepository.UpdateOne(
		ctx, client, id, input,
	)
}

func (s *deliveryService) DeleteDelivery(
	ctx context.Context,
	client *ent.Client,
	id int,
) (bool, error) {
	var success = false

	err := s.deliveryRepository.DeleteOne(ctx, client, id)

	if err == nil {
		success = true
	}

	return success, err
}
