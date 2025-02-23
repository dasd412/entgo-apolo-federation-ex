//go:build ignore

package main

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"federation"
	"log"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		entgql.WithRelaySpec(true),
		entgql.WithConfigPath("../graph/gqlgen.yml"),
		entgql.WithSchemaPath("../graph/ent.graphql"),
		entgql.WithSchemaHook(federation.RemoveNodeGoModel, federation.RemoveNodeQueries),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(
			ex,
		),
		//entc.FeatureNames("privacy", "schema/snapshot"),
		entc.TemplateDir("./template"),
	}
	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
