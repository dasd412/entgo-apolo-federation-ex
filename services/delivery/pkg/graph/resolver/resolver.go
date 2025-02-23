package resolver

import (
	"delivery/internal/repository"
	"delivery/internal/service"
	"delivery/pkg/ent"
	"delivery/pkg/graph/gen"
	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	entClient       *ent.Client
	deliveryService service.DeliveryService
}

func NewSchema(entClient *ent.Client) graphql.ExecutableSchema {
	return gen.NewExecutableSchema(
		gen.Config{
			Resolvers: &Resolver{
				entClient:       entClient,
				deliveryService: service.NewDeliveryService(repository.NewDeliveryRepository()),
			},
		},
	)
}
