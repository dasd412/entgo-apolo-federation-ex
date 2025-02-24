package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"federation"
)

type DeliveryItem struct {
	ent.Schema
}

// Fields of the DeliveryItem.
func (DeliveryItem) Fields() []ent.Field {
	return []ent.Field{
		field.String("productName"),
		field.Int("quantity"),
		field.Float("price"),
	}
}

// Edges of the DeliveryItem.
func (DeliveryItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("delivery", Delivery.Type).
			Ref("delivery_item").
			Unique(),
	}
}

func (DeliveryItem) Annotations() []schema.Annotation {
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
