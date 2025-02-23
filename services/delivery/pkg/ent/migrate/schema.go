// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// DeliveriesColumns holds the columns for the "deliveries" table.
	DeliveriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "order_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"pending", "in_transit", "delivered"}},
		{Name: "tracking_number", Type: field.TypeString, Nullable: true},
		{Name: "created_at", Type: field.TypeTime},
	}
	// DeliveriesTable holds the schema information for the "deliveries" table.
	DeliveriesTable = &schema.Table{
		Name:       "deliveries",
		Columns:    DeliveriesColumns,
		PrimaryKey: []*schema.Column{DeliveriesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		DeliveriesTable,
	}
)

func init() {
}
