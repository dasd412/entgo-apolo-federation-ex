package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/vektah/gqlparser/v2/ast"
	"time"
)

// Order 엔티티 정의
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id"), // User ID (Foreign Key 역할)
		field.Enum("status").Values("pending", "paid", "shipped", "canceled"),
		field.Float("total_price"),
		field.Time("created_at").Default(time.Now),
	}
}
func (Order) Annotations() []schema.Annotation {
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
