package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"federation"
	"time"
)

// Delivery 엔티티 정의
type Delivery struct {
	ent.Schema
}

// Fields of the Delivery.
func (Delivery) Fields() []ent.Field {
	return []ent.Field{
		field.Int("order_id").Unique(),
		field.Int("user_id"),
		field.Enum("status").Values("pending", "in_transit", "delivered"),
		field.String("tracking_number").Optional(),
		field.Time("created_at").Default(time.Now),
	}
}
func (Delivery) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("delivery_item", DeliveryItem.Type).
			Annotations(entgql.RelayConnection()), // Order -> OrderItem (1:N 관계)
	}
}

func (Delivery) Annotations() []schema.Annotation {
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
