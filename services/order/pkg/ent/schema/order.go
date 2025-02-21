package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
