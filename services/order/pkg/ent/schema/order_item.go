package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"federation"
)

// OrderItem 엔터티 정의
type OrderItem struct {
	ent.Schema
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("productName"),
		field.Int("quantity"),
		field.Float("price"),
	}
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("order_item").
			Unique(), // OrderItem -> Order (N:1 관계)
	}
}

func (OrderItem) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
		entgql.Mutations(
			entgql.MutationCreate(),
			entgql.MutationUpdate(),
		),
		federation.GraphKeyDirective("id"),
	}
}
