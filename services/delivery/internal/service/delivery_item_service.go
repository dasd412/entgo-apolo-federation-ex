package service

import (
	"context"
	"delivery/internal/repository"
	"delivery/pkg/ent"
	"delivery/pkg/ent/deliveryitem"
	"entgo.io/contrib/entgql"
)

type (
	deliveryItemService struct {
		deliveryItemRepository repository.DeliveryItemRepository
	}

	DeliveryItemService interface {
		FindDeliveryItem(
			ctx context.Context,
			client *ent.Client,
			id int,
		) (*ent.DeliveryItem, error)
		Paginate(
			ctx context.Context,
			client *ent.Client,
			after *entgql.Cursor[int],
			first *int,
			before *entgql.Cursor[int],
			last *int,
			where *ent.DeliveryItemWhereInput,
		) (*ent.DeliveryItemConnection, error)
		CreateDeliveryItem(
			ctx context.Context,
			client *ent.Client,
			input ent.CreateDeliveryItemInput,
		) (*ent.DeliveryItem, error)
		UpdateDeliveryItem(
			ctx context.Context,
			client *ent.Client,
			id int,
			input ent.UpdateDeliveryItemInput,
		) (*ent.DeliveryItem, error)
		DeleteDeliveryItem(ctx context.Context, client *ent.Client, id int) (
			bool,
			error,
		)
	}
)

func NewDeliveryItemService(deliveryItemRepository repository.DeliveryItemRepository) DeliveryItemService {
	return &deliveryItemService{
		deliveryItemRepository: deliveryItemRepository,
	}
}

func (s *deliveryItemService) FindDeliveryItem(
	ctx context.Context,
	client *ent.Client,
	id int,
) (*ent.DeliveryItem, error) {
	return s.deliveryItemRepository.FindOne(
		ctx, client, func(query *ent.DeliveryItemQuery) {
			query.Where(deliveryitem.ID(id))
		},
	)
}

func (s *deliveryItemService) Paginate(ctx context.Context, client *ent.Client, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.DeliveryItemWhereInput) (*ent.DeliveryItemConnection, error) {
	return s.deliveryItemRepository.Paginate(ctx, client, after, first, before, last, where)
}

func (s *deliveryItemService) CreateDeliveryItem(
	ctx context.Context,
	client *ent.Client,
	input ent.CreateDeliveryItemInput,
) (*ent.DeliveryItem, error) {
	return s.deliveryItemRepository.CreateOne(
		ctx, client, input,
	)
}

func (s *deliveryItemService) UpdateDeliveryItem(
	ctx context.Context,
	client *ent.Client,
	id int,
	input ent.UpdateDeliveryItemInput,
) (*ent.DeliveryItem, error) {
	return s.deliveryItemRepository.UpdateOne(
		ctx, client, id, input,
	)
}

func (s *deliveryItemService) DeleteDeliveryItem(
	ctx context.Context,
	client *ent.Client,
	id int,
) (bool, error) {
	var success = false

	err := s.deliveryItemRepository.DeleteOne(ctx, client, id)

	if err == nil {
		success = true
	}

	return success, err
}
