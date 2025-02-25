// Code generated by ent, DO NOT EDIT.

package orderitem

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the orderitem type in the database.
	Label = "order_item"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProductName holds the string denoting the productname field in the database.
	FieldProductName = "product_name"
	// FieldQuantity holds the string denoting the quantity field in the database.
	FieldQuantity = "quantity"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"
	// Table holds the table name of the orderitem in the database.
	Table = "order_items"
	// OrderTable is the table that holds the order relation/edge.
	OrderTable = "order_items"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
	// OrderColumn is the table column denoting the order relation/edge.
	OrderColumn = "order_order_item"
)

// Columns holds all SQL columns for orderitem fields.
var Columns = []string{
	FieldID,
	FieldProductName,
	FieldQuantity,
	FieldPrice,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "order_items"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"order_order_item",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the OrderItem queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProductName orders the results by the productName field.
func ByProductName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProductName, opts...).ToFunc()
}

// ByQuantity orders the results by the quantity field.
func ByQuantity(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldQuantity, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByOrderField orders the results by order field.
func ByOrderField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOrderStep(), sql.OrderByField(field, opts...))
	}
}
func newOrderStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OrderInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, OrderTable, OrderColumn),
	)
}
