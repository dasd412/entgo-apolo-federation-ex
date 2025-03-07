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

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order_item", OrderItem.Type).
			Annotations(entgql.RelayConnection()), // Order -> OrderItem (1:N 관계)
	}
}

func (Order) Annotations() []schema.Annotation {
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
