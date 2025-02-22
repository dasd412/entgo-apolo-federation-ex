package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/vektah/gqlparser/v2/ast"
	"time"
)

// Delivery 엔티티 정의
type Delivery struct {
	ent.Schema
}

// Fields of the Delivery.
func (Delivery) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_id"), // Order ID (Foreign Key 역할)
		field.Enum("status").Values("pending", "in_transit", "delivered"),
		field.String("tracking_number").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}

func (Delivery) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.MultiOrder(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),

		GraphKeyDirective("id"),
	}
}

func GraphKeyDirective(fields string) entgql.Annotation {
	return entgql.Directives(keyDirective(fields))
}

func keyDirective(fields string) entgql.Directive {
	var args []*ast.Argument
	if fields != "" {
		args = append(args, &ast.Argument{
			Name: "fields",
			Value: &ast.Value{
				Raw:  fields,
				Kind: ast.StringValue,
			},
		})
	}

	return entgql.NewDirective("key", args...)
}
