// Code generated by ent, DO NOT EDIT.

package delivery

import (
	"delivery/pkg/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Delivery {
	return predicate.Delivery(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Delivery {
	return predicate.Delivery(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Delivery {
	return predicate.Delivery(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Delivery {
	return predicate.Delivery(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Delivery {
	return predicate.Delivery(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Delivery {
	return predicate.Delivery(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Delivery {
	return predicate.Delivery(sql.FieldLTE(FieldID, id))
}

// OrderID applies equality check predicate on the "order_id" field. It's identical to OrderIDEQ.
func OrderID(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldOrderID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldUserID, v))
}

// TrackingNumber applies equality check predicate on the "tracking_number" field. It's identical to TrackingNumberEQ.
func TrackingNumber(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldTrackingNumber, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldCreatedAt, v))
}

// OrderIDEQ applies the EQ predicate on the "order_id" field.
func OrderIDEQ(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldOrderID, v))
}

// OrderIDNEQ applies the NEQ predicate on the "order_id" field.
func OrderIDNEQ(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldNEQ(FieldOrderID, v))
}

// OrderIDIn applies the In predicate on the "order_id" field.
func OrderIDIn(vs ...int) predicate.Delivery {
	return predicate.Delivery(sql.FieldIn(FieldOrderID, vs...))
}

// OrderIDNotIn applies the NotIn predicate on the "order_id" field.
func OrderIDNotIn(vs ...int) predicate.Delivery {
	return predicate.Delivery(sql.FieldNotIn(FieldOrderID, vs...))
}

// OrderIDGT applies the GT predicate on the "order_id" field.
func OrderIDGT(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldGT(FieldOrderID, v))
}

// OrderIDGTE applies the GTE predicate on the "order_id" field.
func OrderIDGTE(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldGTE(FieldOrderID, v))
}

// OrderIDLT applies the LT predicate on the "order_id" field.
func OrderIDLT(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldLT(FieldOrderID, v))
}

// OrderIDLTE applies the LTE predicate on the "order_id" field.
func OrderIDLTE(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldLTE(FieldOrderID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Delivery {
	return predicate.Delivery(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Delivery {
	return predicate.Delivery(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.Delivery {
	return predicate.Delivery(sql.FieldLTE(FieldUserID, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v Status) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v Status) predicate.Delivery {
	return predicate.Delivery(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...Status) predicate.Delivery {
	return predicate.Delivery(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...Status) predicate.Delivery {
	return predicate.Delivery(sql.FieldNotIn(FieldStatus, vs...))
}

// TrackingNumberEQ applies the EQ predicate on the "tracking_number" field.
func TrackingNumberEQ(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldTrackingNumber, v))
}

// TrackingNumberNEQ applies the NEQ predicate on the "tracking_number" field.
func TrackingNumberNEQ(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldNEQ(FieldTrackingNumber, v))
}

// TrackingNumberIn applies the In predicate on the "tracking_number" field.
func TrackingNumberIn(vs ...string) predicate.Delivery {
	return predicate.Delivery(sql.FieldIn(FieldTrackingNumber, vs...))
}

// TrackingNumberNotIn applies the NotIn predicate on the "tracking_number" field.
func TrackingNumberNotIn(vs ...string) predicate.Delivery {
	return predicate.Delivery(sql.FieldNotIn(FieldTrackingNumber, vs...))
}

// TrackingNumberGT applies the GT predicate on the "tracking_number" field.
func TrackingNumberGT(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldGT(FieldTrackingNumber, v))
}

// TrackingNumberGTE applies the GTE predicate on the "tracking_number" field.
func TrackingNumberGTE(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldGTE(FieldTrackingNumber, v))
}

// TrackingNumberLT applies the LT predicate on the "tracking_number" field.
func TrackingNumberLT(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldLT(FieldTrackingNumber, v))
}

// TrackingNumberLTE applies the LTE predicate on the "tracking_number" field.
func TrackingNumberLTE(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldLTE(FieldTrackingNumber, v))
}

// TrackingNumberContains applies the Contains predicate on the "tracking_number" field.
func TrackingNumberContains(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldContains(FieldTrackingNumber, v))
}

// TrackingNumberHasPrefix applies the HasPrefix predicate on the "tracking_number" field.
func TrackingNumberHasPrefix(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldHasPrefix(FieldTrackingNumber, v))
}

// TrackingNumberHasSuffix applies the HasSuffix predicate on the "tracking_number" field.
func TrackingNumberHasSuffix(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldHasSuffix(FieldTrackingNumber, v))
}

// TrackingNumberIsNil applies the IsNil predicate on the "tracking_number" field.
func TrackingNumberIsNil() predicate.Delivery {
	return predicate.Delivery(sql.FieldIsNull(FieldTrackingNumber))
}

// TrackingNumberNotNil applies the NotNil predicate on the "tracking_number" field.
func TrackingNumberNotNil() predicate.Delivery {
	return predicate.Delivery(sql.FieldNotNull(FieldTrackingNumber))
}

// TrackingNumberEqualFold applies the EqualFold predicate on the "tracking_number" field.
func TrackingNumberEqualFold(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldEqualFold(FieldTrackingNumber, v))
}

// TrackingNumberContainsFold applies the ContainsFold predicate on the "tracking_number" field.
func TrackingNumberContainsFold(v string) predicate.Delivery {
	return predicate.Delivery(sql.FieldContainsFold(FieldTrackingNumber, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Delivery {
	return predicate.Delivery(sql.FieldLTE(FieldCreatedAt, v))
}

// HasDeliveryItem applies the HasEdge predicate on the "delivery_item" edge.
func HasDeliveryItem() predicate.Delivery {
	return predicate.Delivery(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DeliveryItemTable, DeliveryItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDeliveryItemWith applies the HasEdge predicate on the "delivery_item" edge with a given conditions (other predicates).
func HasDeliveryItemWith(preds ...predicate.DeliveryItem) predicate.Delivery {
	return predicate.Delivery(func(s *sql.Selector) {
		step := newDeliveryItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Delivery) predicate.Delivery {
	return predicate.Delivery(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Delivery) predicate.Delivery {
	return predicate.Delivery(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Delivery) predicate.Delivery {
	return predicate.Delivery(sql.NotPredicates(p))
}
