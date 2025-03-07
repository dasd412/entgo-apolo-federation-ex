// Code generated by ent, DO NOT EDIT.

package deliveryitem

import (
	"delivery/pkg/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldLTE(FieldID, id))
}

// ProductName applies equality check predicate on the "productName" field. It's identical to ProductNameEQ.
func ProductName(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEQ(FieldProductName, v))
}

// Quantity applies equality check predicate on the "quantity" field. It's identical to QuantityEQ.
func Quantity(v int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEQ(FieldQuantity, v))
}

// Price applies equality check predicate on the "price" field. It's identical to PriceEQ.
func Price(v float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEQ(FieldPrice, v))
}

// ProductNameEQ applies the EQ predicate on the "productName" field.
func ProductNameEQ(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEQ(FieldProductName, v))
}

// ProductNameNEQ applies the NEQ predicate on the "productName" field.
func ProductNameNEQ(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldNEQ(FieldProductName, v))
}

// ProductNameIn applies the In predicate on the "productName" field.
func ProductNameIn(vs ...string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldIn(FieldProductName, vs...))
}

// ProductNameNotIn applies the NotIn predicate on the "productName" field.
func ProductNameNotIn(vs ...string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldNotIn(FieldProductName, vs...))
}

// ProductNameGT applies the GT predicate on the "productName" field.
func ProductNameGT(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldGT(FieldProductName, v))
}

// ProductNameGTE applies the GTE predicate on the "productName" field.
func ProductNameGTE(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldGTE(FieldProductName, v))
}

// ProductNameLT applies the LT predicate on the "productName" field.
func ProductNameLT(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldLT(FieldProductName, v))
}

// ProductNameLTE applies the LTE predicate on the "productName" field.
func ProductNameLTE(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldLTE(FieldProductName, v))
}

// ProductNameContains applies the Contains predicate on the "productName" field.
func ProductNameContains(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldContains(FieldProductName, v))
}

// ProductNameHasPrefix applies the HasPrefix predicate on the "productName" field.
func ProductNameHasPrefix(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldHasPrefix(FieldProductName, v))
}

// ProductNameHasSuffix applies the HasSuffix predicate on the "productName" field.
func ProductNameHasSuffix(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldHasSuffix(FieldProductName, v))
}

// ProductNameEqualFold applies the EqualFold predicate on the "productName" field.
func ProductNameEqualFold(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEqualFold(FieldProductName, v))
}

// ProductNameContainsFold applies the ContainsFold predicate on the "productName" field.
func ProductNameContainsFold(v string) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldContainsFold(FieldProductName, v))
}

// QuantityEQ applies the EQ predicate on the "quantity" field.
func QuantityEQ(v int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEQ(FieldQuantity, v))
}

// QuantityNEQ applies the NEQ predicate on the "quantity" field.
func QuantityNEQ(v int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldNEQ(FieldQuantity, v))
}

// QuantityIn applies the In predicate on the "quantity" field.
func QuantityIn(vs ...int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldIn(FieldQuantity, vs...))
}

// QuantityNotIn applies the NotIn predicate on the "quantity" field.
func QuantityNotIn(vs ...int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldNotIn(FieldQuantity, vs...))
}

// QuantityGT applies the GT predicate on the "quantity" field.
func QuantityGT(v int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldGT(FieldQuantity, v))
}

// QuantityGTE applies the GTE predicate on the "quantity" field.
func QuantityGTE(v int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldGTE(FieldQuantity, v))
}

// QuantityLT applies the LT predicate on the "quantity" field.
func QuantityLT(v int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldLT(FieldQuantity, v))
}

// QuantityLTE applies the LTE predicate on the "quantity" field.
func QuantityLTE(v int) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldLTE(FieldQuantity, v))
}

// PriceEQ applies the EQ predicate on the "price" field.
func PriceEQ(v float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldEQ(FieldPrice, v))
}

// PriceNEQ applies the NEQ predicate on the "price" field.
func PriceNEQ(v float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldNEQ(FieldPrice, v))
}

// PriceIn applies the In predicate on the "price" field.
func PriceIn(vs ...float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldIn(FieldPrice, vs...))
}

// PriceNotIn applies the NotIn predicate on the "price" field.
func PriceNotIn(vs ...float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldNotIn(FieldPrice, vs...))
}

// PriceGT applies the GT predicate on the "price" field.
func PriceGT(v float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldGT(FieldPrice, v))
}

// PriceGTE applies the GTE predicate on the "price" field.
func PriceGTE(v float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldGTE(FieldPrice, v))
}

// PriceLT applies the LT predicate on the "price" field.
func PriceLT(v float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldLT(FieldPrice, v))
}

// PriceLTE applies the LTE predicate on the "price" field.
func PriceLTE(v float64) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.FieldLTE(FieldPrice, v))
}

// HasDelivery applies the HasEdge predicate on the "delivery" edge.
func HasDelivery() predicate.DeliveryItem {
	return predicate.DeliveryItem(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, DeliveryTable, DeliveryColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeliveryWith applies the HasEdge predicate on the "delivery" edge with a given conditions (other predicates).
func HasDeliveryWith(preds ...predicate.Delivery) predicate.DeliveryItem {
	return predicate.DeliveryItem(func(s *sql.Selector) {
		step := newDeliveryStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.DeliveryItem) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.DeliveryItem) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.DeliveryItem) predicate.DeliveryItem {
	return predicate.DeliveryItem(sql.NotPredicates(p))
}
