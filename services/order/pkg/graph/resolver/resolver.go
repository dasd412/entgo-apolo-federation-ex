package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"order/pkg/ent"
	"order/pkg/graph/gen"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	entClient *ent.Client
}

func NewSchema(entClient *ent.Client) graphql.ExecutableSchema {
	return gen.NewExecutableSchema(
		gen.Config{
			Resolvers: &Resolver{
				entClient: entClient,
			},
		},
	)
}
