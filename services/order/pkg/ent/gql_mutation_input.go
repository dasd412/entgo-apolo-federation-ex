// Code generated by ent, DO NOT EDIT.

package ent

import (
	"order/pkg/ent/order"
	"time"
)

// CreateOrderInput represents a mutation input for creating orders.
type CreateOrderInput struct {
	UserID       int
	Status       order.Status
	TotalPrice   float64
	CreatedAt    *time.Time
	OrderItemIDs []int
}

// Mutate applies the CreateOrderInput on the OrderMutation builder.
func (i *CreateOrderInput) Mutate(m *OrderMutation) {
	m.SetUserID(i.UserID)
	m.SetStatus(i.Status)
	m.SetTotalPrice(i.TotalPrice)
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if v := i.OrderItemIDs; len(v) > 0 {
		m.AddOrderItemIDs(v...)
	}
}

// SetInput applies the change-set in the CreateOrderInput on the OrderCreate builder.
func (c *OrderCreate) SetInput(i CreateOrderInput) *OrderCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrderInput represents a mutation input for updating orders.
type UpdateOrderInput struct {
	UserID             *int
	Status             *order.Status
	TotalPrice         *float64
	CreatedAt          *time.Time
	ClearOrderItem     bool
	AddOrderItemIDs    []int
	RemoveOrderItemIDs []int
}

// Mutate applies the UpdateOrderInput on the OrderMutation builder.
func (i *UpdateOrderInput) Mutate(m *OrderMutation) {
	if v := i.UserID; v != nil {
		m.SetUserID(*v)
	}
	if v := i.Status; v != nil {
		m.SetStatus(*v)
	}
	if v := i.TotalPrice; v != nil {
		m.SetTotalPrice(*v)
	}
	if v := i.CreatedAt; v != nil {
		m.SetCreatedAt(*v)
	}
	if i.ClearOrderItem {
		m.ClearOrderItem()
	}
	if v := i.AddOrderItemIDs; len(v) > 0 {
		m.AddOrderItemIDs(v...)
	}
	if v := i.RemoveOrderItemIDs; len(v) > 0 {
		m.RemoveOrderItemIDs(v...)
	}
}

// SetInput applies the change-set in the UpdateOrderInput on the OrderUpdate builder.
func (c *OrderUpdate) SetInput(i UpdateOrderInput) *OrderUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrderInput on the OrderUpdateOne builder.
func (c *OrderUpdateOne) SetInput(i UpdateOrderInput) *OrderUpdateOne {
	i.Mutate(c.Mutation())
	return c
}

// CreateOrderItemInput represents a mutation input for creating orderitems.
type CreateOrderItemInput struct {
	ProductName string
	Quantity    int
	Price       float64
	OrderID     *int
}

// Mutate applies the CreateOrderItemInput on the OrderItemMutation builder.
func (i *CreateOrderItemInput) Mutate(m *OrderItemMutation) {
	m.SetProductName(i.ProductName)
	m.SetQuantity(i.Quantity)
	m.SetPrice(i.Price)
	if v := i.OrderID; v != nil {
		m.SetOrderID(*v)
	}
}

// SetInput applies the change-set in the CreateOrderItemInput on the OrderItemCreate builder.
func (c *OrderItemCreate) SetInput(i CreateOrderItemInput) *OrderItemCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateOrderItemInput represents a mutation input for updating orderitems.
type UpdateOrderItemInput struct {
	ProductName *string
	Quantity    *int
	Price       *float64
	ClearOrder  bool
	OrderID     *int
}

// Mutate applies the UpdateOrderItemInput on the OrderItemMutation builder.
func (i *UpdateOrderItemInput) Mutate(m *OrderItemMutation) {
	if v := i.ProductName; v != nil {
		m.SetProductName(*v)
	}
	if v := i.Quantity; v != nil {
		m.SetQuantity(*v)
	}
	if v := i.Price; v != nil {
		m.SetPrice(*v)
	}
	if i.ClearOrder {
		m.ClearOrder()
	}
	if v := i.OrderID; v != nil {
		m.SetOrderID(*v)
	}
}

// SetInput applies the change-set in the UpdateOrderItemInput on the OrderItemUpdate builder.
func (c *OrderItemUpdate) SetInput(i UpdateOrderItemInput) *OrderItemUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateOrderItemInput on the OrderItemUpdateOne builder.
func (c *OrderItemUpdateOne) SetInput(i UpdateOrderItemInput) *OrderItemUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
