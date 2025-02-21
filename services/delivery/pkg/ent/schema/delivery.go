package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
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
