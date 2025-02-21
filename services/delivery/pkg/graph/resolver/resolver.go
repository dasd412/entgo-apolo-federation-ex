package resolver

import (
	"delivery/pkg/ent"
	"delivery/pkg/graph/gen"
	"github.com/99designs/gqlgen/graphql"
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
