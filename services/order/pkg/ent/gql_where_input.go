// Code generated by ent, DO NOT EDIT.

package ent

import (
	"errors"
	"fmt"
	"order/pkg/ent/order"
	"order/pkg/ent/orderitem"
	"order/pkg/ent/predicate"
	"time"
)

// OrderWhereInput represents a where input for filtering Order queries.
type OrderWhereInput struct {
	Predicates []predicate.Order  `json:"-"`
	Not        *OrderWhereInput   `json:"not,omitempty"`
	Or         []*OrderWhereInput `json:"or,omitempty"`
	And        []*OrderWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "user_id" field predicates.
	UserID      *int  `json:"userID,omitempty"`
	UserIDNEQ   *int  `json:"userIDNEQ,omitempty"`
	UserIDIn    []int `json:"userIDIn,omitempty"`
	UserIDNotIn []int `json:"userIDNotIn,omitempty"`
	UserIDGT    *int  `json:"userIDGT,omitempty"`
	UserIDGTE   *int  `json:"userIDGTE,omitempty"`
	UserIDLT    *int  `json:"userIDLT,omitempty"`
	UserIDLTE   *int  `json:"userIDLTE,omitempty"`

	// "status" field predicates.
	Status      *order.Status  `json:"status,omitempty"`
	StatusNEQ   *order.Status  `json:"statusNEQ,omitempty"`
	StatusIn    []order.Status `json:"statusIn,omitempty"`
	StatusNotIn []order.Status `json:"statusNotIn,omitempty"`

	// "total_price" field predicates.
	TotalPrice      *float64  `json:"totalPrice,omitempty"`
	TotalPriceNEQ   *float64  `json:"totalPriceNEQ,omitempty"`
	TotalPriceIn    []float64 `json:"totalPriceIn,omitempty"`
	TotalPriceNotIn []float64 `json:"totalPriceNotIn,omitempty"`
	TotalPriceGT    *float64  `json:"totalPriceGT,omitempty"`
	TotalPriceGTE   *float64  `json:"totalPriceGTE,omitempty"`
	TotalPriceLT    *float64  `json:"totalPriceLT,omitempty"`
	TotalPriceLTE   *float64  `json:"totalPriceLTE,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"createdAt,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"createdAtNEQ,omitempty"`
	CreatedAtIn    []time.Time `json:"createdAtIn,omitempty"`
	CreatedAtNotIn []time.Time `json:"createdAtNotIn,omitempty"`
	CreatedAtGT    *time.Time  `json:"createdAtGT,omitempty"`
	CreatedAtGTE   *time.Time  `json:"createdAtGTE,omitempty"`
	CreatedAtLT    *time.Time  `json:"createdAtLT,omitempty"`
	CreatedAtLTE   *time.Time  `json:"createdAtLTE,omitempty"`

	// "order_item" edge predicates.
	HasOrderItem     *bool                  `json:"hasOrderItem,omitempty"`
	HasOrderItemWith []*OrderItemWhereInput `json:"hasOrderItemWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *OrderWhereInput) AddPredicates(predicates ...predicate.Order) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the OrderWhereInput filter on the OrderQuery builder.
func (i *OrderWhereInput) Filter(q *OrderQuery) (*OrderQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyOrderWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyOrderWhereInput is returned in case the OrderWhereInput is empty.
var ErrEmptyOrderWhereInput = errors.New("ent: empty predicate OrderWhereInput")

// P returns a predicate for filtering orders.
// An error is returned if the input is empty or invalid.
func (i *OrderWhereInput) P() (predicate.Order, error) {
	var predicates []predicate.Order
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, order.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.Order, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, order.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.Order, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, order.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, order.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, order.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, order.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, order.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, order.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, order.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, order.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, order.IDLTE(*i.IDLTE))
	}
	if i.UserID != nil {
		predicates = append(predicates, order.UserIDEQ(*i.UserID))
	}
	if i.UserIDNEQ != nil {
		predicates = append(predicates, order.UserIDNEQ(*i.UserIDNEQ))
	}
	if len(i.UserIDIn) > 0 {
		predicates = append(predicates, order.UserIDIn(i.UserIDIn...))
	}
	if len(i.UserIDNotIn) > 0 {
		predicates = append(predicates, order.UserIDNotIn(i.UserIDNotIn...))
	}
	if i.UserIDGT != nil {
		predicates = append(predicates, order.UserIDGT(*i.UserIDGT))
	}
	if i.UserIDGTE != nil {
		predicates = append(predicates, order.UserIDGTE(*i.UserIDGTE))
	}
	if i.UserIDLT != nil {
		predicates = append(predicates, order.UserIDLT(*i.UserIDLT))
	}
	if i.UserIDLTE != nil {
		predicates = append(predicates, order.UserIDLTE(*i.UserIDLTE))
	}
	if i.Status != nil {
		predicates = append(predicates, order.StatusEQ(*i.Status))
	}
	if i.StatusNEQ != nil {
		predicates = append(predicates, order.StatusNEQ(*i.StatusNEQ))
	}
	if len(i.StatusIn) > 0 {
		predicates = append(predicates, order.StatusIn(i.StatusIn...))
	}
	if len(i.StatusNotIn) > 0 {
		predicates = append(predicates, order.StatusNotIn(i.StatusNotIn...))
	}
	if i.TotalPrice != nil {
		predicates = append(predicates, order.TotalPriceEQ(*i.TotalPrice))
	}
	if i.TotalPriceNEQ != nil {
		predicates = append(predicates, order.TotalPriceNEQ(*i.TotalPriceNEQ))
	}
	if len(i.TotalPriceIn) > 0 {
		predicates = append(predicates, order.TotalPriceIn(i.TotalPriceIn...))
	}
	if len(i.TotalPriceNotIn) > 0 {
		predicates = append(predicates, order.TotalPriceNotIn(i.TotalPriceNotIn...))
	}
	if i.TotalPriceGT != nil {
		predicates = append(predicates, order.TotalPriceGT(*i.TotalPriceGT))
	}
	if i.TotalPriceGTE != nil {
		predicates = append(predicates, order.TotalPriceGTE(*i.TotalPriceGTE))
	}
	if i.TotalPriceLT != nil {
		predicates = append(predicates, order.TotalPriceLT(*i.TotalPriceLT))
	}
	if i.TotalPriceLTE != nil {
		predicates = append(predicates, order.TotalPriceLTE(*i.TotalPriceLTE))
	}
	if i.CreatedAt != nil {
		predicates = append(predicates, order.CreatedAtEQ(*i.CreatedAt))
	}
	if i.CreatedAtNEQ != nil {
		predicates = append(predicates, order.CreatedAtNEQ(*i.CreatedAtNEQ))
	}
	if len(i.CreatedAtIn) > 0 {
		predicates = append(predicates, order.CreatedAtIn(i.CreatedAtIn...))
	}
	if len(i.CreatedAtNotIn) > 0 {
		predicates = append(predicates, order.CreatedAtNotIn(i.CreatedAtNotIn...))
	}
	if i.CreatedAtGT != nil {
		predicates = append(predicates, order.CreatedAtGT(*i.CreatedAtGT))
	}
	if i.CreatedAtGTE != nil {
		predicates = append(predicates, order.CreatedAtGTE(*i.CreatedAtGTE))
	}
	if i.CreatedAtLT != nil {
		predicates = append(predicates, order.CreatedAtLT(*i.CreatedAtLT))
	}
	if i.CreatedAtLTE != nil {
		predicates = append(predicates, order.CreatedAtLTE(*i.CreatedAtLTE))
	}

	if i.HasOrderItem != nil {
		p := order.HasOrderItem()
		if !*i.HasOrderItem {
			p = order.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOrderItemWith) > 0 {
		with := make([]predicate.OrderItem, 0, len(i.HasOrderItemWith))
		for _, w := range i.HasOrderItemWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOrderItemWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, order.HasOrderItemWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyOrderWhereInput
	case 1:
		return predicates[0], nil
	default:
		return order.And(predicates...), nil
	}
}

// OrderItemWhereInput represents a where input for filtering OrderItem queries.
type OrderItemWhereInput struct {
	Predicates []predicate.OrderItem  `json:"-"`
	Not        *OrderItemWhereInput   `json:"not,omitempty"`
	Or         []*OrderItemWhereInput `json:"or,omitempty"`
	And        []*OrderItemWhereInput `json:"and,omitempty"`

	// "id" field predicates.
	ID      *int  `json:"id,omitempty"`
	IDNEQ   *int  `json:"idNEQ,omitempty"`
	IDIn    []int `json:"idIn,omitempty"`
	IDNotIn []int `json:"idNotIn,omitempty"`
	IDGT    *int  `json:"idGT,omitempty"`
	IDGTE   *int  `json:"idGTE,omitempty"`
	IDLT    *int  `json:"idLT,omitempty"`
	IDLTE   *int  `json:"idLTE,omitempty"`

	// "productName" field predicates.
	ProductName             *string  `json:"productname,omitempty"`
	ProductNameNEQ          *string  `json:"productnameNEQ,omitempty"`
	ProductNameIn           []string `json:"productnameIn,omitempty"`
	ProductNameNotIn        []string `json:"productnameNotIn,omitempty"`
	ProductNameGT           *string  `json:"productnameGT,omitempty"`
	ProductNameGTE          *string  `json:"productnameGTE,omitempty"`
	ProductNameLT           *string  `json:"productnameLT,omitempty"`
	ProductNameLTE          *string  `json:"productnameLTE,omitempty"`
	ProductNameContains     *string  `json:"productnameContains,omitempty"`
	ProductNameHasPrefix    *string  `json:"productnameHasPrefix,omitempty"`
	ProductNameHasSuffix    *string  `json:"productnameHasSuffix,omitempty"`
	ProductNameEqualFold    *string  `json:"productnameEqualFold,omitempty"`
	ProductNameContainsFold *string  `json:"productnameContainsFold,omitempty"`

	// "quantity" field predicates.
	Quantity      *int  `json:"quantity,omitempty"`
	QuantityNEQ   *int  `json:"quantityNEQ,omitempty"`
	QuantityIn    []int `json:"quantityIn,omitempty"`
	QuantityNotIn []int `json:"quantityNotIn,omitempty"`
	QuantityGT    *int  `json:"quantityGT,omitempty"`
	QuantityGTE   *int  `json:"quantityGTE,omitempty"`
	QuantityLT    *int  `json:"quantityLT,omitempty"`
	QuantityLTE   *int  `json:"quantityLTE,omitempty"`

	// "price" field predicates.
	Price      *float64  `json:"price,omitempty"`
	PriceNEQ   *float64  `json:"priceNEQ,omitempty"`
	PriceIn    []float64 `json:"priceIn,omitempty"`
	PriceNotIn []float64 `json:"priceNotIn,omitempty"`
	PriceGT    *float64  `json:"priceGT,omitempty"`
	PriceGTE   *float64  `json:"priceGTE,omitempty"`
	PriceLT    *float64  `json:"priceLT,omitempty"`
	PriceLTE   *float64  `json:"priceLTE,omitempty"`

	// "order" edge predicates.
	HasOrder     *bool              `json:"hasOrder,omitempty"`
	HasOrderWith []*OrderWhereInput `json:"hasOrderWith,omitempty"`
}

// AddPredicates adds custom predicates to the where input to be used during the filtering phase.
func (i *OrderItemWhereInput) AddPredicates(predicates ...predicate.OrderItem) {
	i.Predicates = append(i.Predicates, predicates...)
}

// Filter applies the OrderItemWhereInput filter on the OrderItemQuery builder.
func (i *OrderItemWhereInput) Filter(q *OrderItemQuery) (*OrderItemQuery, error) {
	if i == nil {
		return q, nil
	}
	p, err := i.P()
	if err != nil {
		if err == ErrEmptyOrderItemWhereInput {
			return q, nil
		}
		return nil, err
	}
	return q.Where(p), nil
}

// ErrEmptyOrderItemWhereInput is returned in case the OrderItemWhereInput is empty.
var ErrEmptyOrderItemWhereInput = errors.New("ent: empty predicate OrderItemWhereInput")

// P returns a predicate for filtering orderitems.
// An error is returned if the input is empty or invalid.
func (i *OrderItemWhereInput) P() (predicate.OrderItem, error) {
	var predicates []predicate.OrderItem
	if i.Not != nil {
		p, err := i.Not.P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'not'", err)
		}
		predicates = append(predicates, orderitem.Not(p))
	}
	switch n := len(i.Or); {
	case n == 1:
		p, err := i.Or[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'or'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		or := make([]predicate.OrderItem, 0, n)
		for _, w := range i.Or {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'or'", err)
			}
			or = append(or, p)
		}
		predicates = append(predicates, orderitem.Or(or...))
	}
	switch n := len(i.And); {
	case n == 1:
		p, err := i.And[0].P()
		if err != nil {
			return nil, fmt.Errorf("%w: field 'and'", err)
		}
		predicates = append(predicates, p)
	case n > 1:
		and := make([]predicate.OrderItem, 0, n)
		for _, w := range i.And {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'and'", err)
			}
			and = append(and, p)
		}
		predicates = append(predicates, orderitem.And(and...))
	}
	predicates = append(predicates, i.Predicates...)
	if i.ID != nil {
		predicates = append(predicates, orderitem.IDEQ(*i.ID))
	}
	if i.IDNEQ != nil {
		predicates = append(predicates, orderitem.IDNEQ(*i.IDNEQ))
	}
	if len(i.IDIn) > 0 {
		predicates = append(predicates, orderitem.IDIn(i.IDIn...))
	}
	if len(i.IDNotIn) > 0 {
		predicates = append(predicates, orderitem.IDNotIn(i.IDNotIn...))
	}
	if i.IDGT != nil {
		predicates = append(predicates, orderitem.IDGT(*i.IDGT))
	}
	if i.IDGTE != nil {
		predicates = append(predicates, orderitem.IDGTE(*i.IDGTE))
	}
	if i.IDLT != nil {
		predicates = append(predicates, orderitem.IDLT(*i.IDLT))
	}
	if i.IDLTE != nil {
		predicates = append(predicates, orderitem.IDLTE(*i.IDLTE))
	}
	if i.ProductName != nil {
		predicates = append(predicates, orderitem.ProductNameEQ(*i.ProductName))
	}
	if i.ProductNameNEQ != nil {
		predicates = append(predicates, orderitem.ProductNameNEQ(*i.ProductNameNEQ))
	}
	if len(i.ProductNameIn) > 0 {
		predicates = append(predicates, orderitem.ProductNameIn(i.ProductNameIn...))
	}
	if len(i.ProductNameNotIn) > 0 {
		predicates = append(predicates, orderitem.ProductNameNotIn(i.ProductNameNotIn...))
	}
	if i.ProductNameGT != nil {
		predicates = append(predicates, orderitem.ProductNameGT(*i.ProductNameGT))
	}
	if i.ProductNameGTE != nil {
		predicates = append(predicates, orderitem.ProductNameGTE(*i.ProductNameGTE))
	}
	if i.ProductNameLT != nil {
		predicates = append(predicates, orderitem.ProductNameLT(*i.ProductNameLT))
	}
	if i.ProductNameLTE != nil {
		predicates = append(predicates, orderitem.ProductNameLTE(*i.ProductNameLTE))
	}
	if i.ProductNameContains != nil {
		predicates = append(predicates, orderitem.ProductNameContains(*i.ProductNameContains))
	}
	if i.ProductNameHasPrefix != nil {
		predicates = append(predicates, orderitem.ProductNameHasPrefix(*i.ProductNameHasPrefix))
	}
	if i.ProductNameHasSuffix != nil {
		predicates = append(predicates, orderitem.ProductNameHasSuffix(*i.ProductNameHasSuffix))
	}
	if i.ProductNameEqualFold != nil {
		predicates = append(predicates, orderitem.ProductNameEqualFold(*i.ProductNameEqualFold))
	}
	if i.ProductNameContainsFold != nil {
		predicates = append(predicates, orderitem.ProductNameContainsFold(*i.ProductNameContainsFold))
	}
	if i.Quantity != nil {
		predicates = append(predicates, orderitem.QuantityEQ(*i.Quantity))
	}
	if i.QuantityNEQ != nil {
		predicates = append(predicates, orderitem.QuantityNEQ(*i.QuantityNEQ))
	}
	if len(i.QuantityIn) > 0 {
		predicates = append(predicates, orderitem.QuantityIn(i.QuantityIn...))
	}
	if len(i.QuantityNotIn) > 0 {
		predicates = append(predicates, orderitem.QuantityNotIn(i.QuantityNotIn...))
	}
	if i.QuantityGT != nil {
		predicates = append(predicates, orderitem.QuantityGT(*i.QuantityGT))
	}
	if i.QuantityGTE != nil {
		predicates = append(predicates, orderitem.QuantityGTE(*i.QuantityGTE))
	}
	if i.QuantityLT != nil {
		predicates = append(predicates, orderitem.QuantityLT(*i.QuantityLT))
	}
	if i.QuantityLTE != nil {
		predicates = append(predicates, orderitem.QuantityLTE(*i.QuantityLTE))
	}
	if i.Price != nil {
		predicates = append(predicates, orderitem.PriceEQ(*i.Price))
	}
	if i.PriceNEQ != nil {
		predicates = append(predicates, orderitem.PriceNEQ(*i.PriceNEQ))
	}
	if len(i.PriceIn) > 0 {
		predicates = append(predicates, orderitem.PriceIn(i.PriceIn...))
	}
	if len(i.PriceNotIn) > 0 {
		predicates = append(predicates, orderitem.PriceNotIn(i.PriceNotIn...))
	}
	if i.PriceGT != nil {
		predicates = append(predicates, orderitem.PriceGT(*i.PriceGT))
	}
	if i.PriceGTE != nil {
		predicates = append(predicates, orderitem.PriceGTE(*i.PriceGTE))
	}
	if i.PriceLT != nil {
		predicates = append(predicates, orderitem.PriceLT(*i.PriceLT))
	}
	if i.PriceLTE != nil {
		predicates = append(predicates, orderitem.PriceLTE(*i.PriceLTE))
	}

	if i.HasOrder != nil {
		p := orderitem.HasOrder()
		if !*i.HasOrder {
			p = orderitem.Not(p)
		}
		predicates = append(predicates, p)
	}
	if len(i.HasOrderWith) > 0 {
		with := make([]predicate.Order, 0, len(i.HasOrderWith))
		for _, w := range i.HasOrderWith {
			p, err := w.P()
			if err != nil {
				return nil, fmt.Errorf("%w: field 'HasOrderWith'", err)
			}
			with = append(with, p)
		}
		predicates = append(predicates, orderitem.HasOrderWith(with...))
	}
	switch len(predicates) {
	case 0:
		return nil, ErrEmptyOrderItemWhereInput
	case 1:
		return predicates[0], nil
	default:
		return orderitem.And(predicates...), nil
	}
}
