// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "paid", "shipped", "canceled"}},
		{Name: "total_price", Type: field.TypeFloat64},
		{Name: "created_at", Type: field.TypeTime},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
	}
	// OrderItemsColumns holds the columns for the "order_items" table.
	OrderItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "product_name", Type: field.TypeString},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "price", Type: field.TypeFloat64},
		{Name: "order_order_item", Type: field.TypeInt, Nullable: true},
	}
	// OrderItemsTable holds the schema information for the "order_items" table.
	OrderItemsTable = &schema.Table{
		Name:       "order_items",
		Columns:    OrderItemsColumns,
		PrimaryKey: []*schema.Column{OrderItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "order_items_orders_order_item",
				Columns:    []*schema.Column{OrderItemsColumns[4]},
				RefColumns: []*schema.Column{OrdersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrdersTable,
		OrderItemsTable,
	}
)

func init() {
	OrderItemsTable.ForeignKeys[0].RefTable = OrdersTable
}
