package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"user/internal/repository"
	"user/internal/service"
	"user/pkg/ent"
	"user/pkg/graph/gen"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	entClient   *ent.Client
	userService service.UserService
}

func NewSchema(entClient *ent.Client) graphql.ExecutableSchema {

	return gen.NewExecutableSchema(
		gen.Config{
			Resolvers: &Resolver{
				entClient:   entClient,
				userService: service.NewUserService(repository.NewUserRepository()),
			},
		},
	)
}
