// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"delivery/pkg/ent/delivery"
	"delivery/pkg/ent/deliveryitem"
	"fmt"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (d *DeliveryQuery) CollectFields(ctx context.Context, satisfies ...string) (*DeliveryQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return d, nil
	}
	if err := d.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return d, nil
}

func (d *DeliveryQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(delivery.Columns))
		selectedFields = []string{delivery.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "deliveryItem":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DeliveryItemClient{config: d.config}).Query()
			)
			args := newDeliveryItemPaginateArgs(fieldArgs(ctx, new(DeliveryItemWhereInput), path...))
			if err := validateFirstLast(args.first, args.last); err != nil {
				return fmt.Errorf("validate first and last in path %q: %w", path, err)
			}
			pager, err := newDeliveryItemPager(args.opts, args.last != nil)
			if err != nil {
				return fmt.Errorf("create new pager in path %q: %w", path, err)
			}
			if query, err = pager.applyFilter(query); err != nil {
				return err
			}
			ignoredEdges := !hasCollectedField(ctx, append(path, edgesField)...)
			if hasCollectedField(ctx, append(path, totalCountField)...) || hasCollectedField(ctx, append(path, pageInfoField)...) {
				hasPagination := args.after != nil || args.first != nil || args.before != nil || args.last != nil
				if hasPagination || ignoredEdges {
					query := query.Clone()
					d.loadTotal = append(d.loadTotal, func(ctx context.Context, nodes []*Delivery) error {
						ids := make([]driver.Value, len(nodes))
						for i := range nodes {
							ids[i] = nodes[i].ID
						}
						var v []struct {
							NodeID int `sql:"delivery_delivery_item"`
							Count  int `sql:"count"`
						}
						query.Where(func(s *sql.Selector) {
							s.Where(sql.InValues(s.C(delivery.DeliveryItemColumn), ids...))
						})
						if err := query.GroupBy(delivery.DeliveryItemColumn).Aggregate(Count()).Scan(ctx, &v); err != nil {
							return err
						}
						m := make(map[int]int, len(v))
						for i := range v {
							m[v[i].NodeID] = v[i].Count
						}
						for i := range nodes {
							n := m[nodes[i].ID]
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				} else {
					d.loadTotal = append(d.loadTotal, func(_ context.Context, nodes []*Delivery) error {
						for i := range nodes {
							n := len(nodes[i].Edges.DeliveryItem)
							if nodes[i].Edges.totalCount[0] == nil {
								nodes[i].Edges.totalCount[0] = make(map[string]int)
							}
							nodes[i].Edges.totalCount[0][alias] = n
						}
						return nil
					})
				}
			}
			if ignoredEdges || (args.first != nil && *args.first == 0) || (args.last != nil && *args.last == 0) {
				continue
			}
			if query, err = pager.applyCursors(query, args.after, args.before); err != nil {
				return err
			}
			path = append(path, edgesField, nodeField)
			if field := collectedField(ctx, path...); field != nil {
				if err := query.collectField(ctx, false, opCtx, *field, path, mayAddCondition(satisfies, deliveryitemImplementors)...); err != nil {
					return err
				}
			}
			if limit := paginateLimit(args.first, args.last); limit > 0 {
				if oneNode {
					pager.applyOrder(query.Limit(limit))
				} else {
					modify := entgql.LimitPerRow(delivery.DeliveryItemColumn, limit, pager.orderExpr(query))
					query.modifiers = append(query.modifiers, modify)
				}
			} else {
				query = pager.applyOrder(query)
			}
			d.WithNamedDeliveryItem(alias, func(wq *DeliveryItemQuery) {
				*wq = *query
			})
		case "orderID":
			if _, ok := fieldSeen[delivery.FieldOrderID]; !ok {
				selectedFields = append(selectedFields, delivery.FieldOrderID)
				fieldSeen[delivery.FieldOrderID] = struct{}{}
			}
		case "userID":
			if _, ok := fieldSeen[delivery.FieldUserID]; !ok {
				selectedFields = append(selectedFields, delivery.FieldUserID)
				fieldSeen[delivery.FieldUserID] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[delivery.FieldStatus]; !ok {
				selectedFields = append(selectedFields, delivery.FieldStatus)
				fieldSeen[delivery.FieldStatus] = struct{}{}
			}
		case "trackingNumber":
			if _, ok := fieldSeen[delivery.FieldTrackingNumber]; !ok {
				selectedFields = append(selectedFields, delivery.FieldTrackingNumber)
				fieldSeen[delivery.FieldTrackingNumber] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[delivery.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, delivery.FieldCreatedAt)
				fieldSeen[delivery.FieldCreatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		d.Select(selectedFields...)
	}
	return nil
}

type deliveryPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []DeliveryPaginateOption
}

func newDeliveryPaginateArgs(rv map[string]any) *deliveryPaginateArgs {
	args := &deliveryPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*DeliveryWhereInput); ok {
		args.opts = append(args.opts, WithDeliveryFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (di *DeliveryItemQuery) CollectFields(ctx context.Context, satisfies ...string) (*DeliveryItemQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return di, nil
	}
	if err := di.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return di, nil
}

func (di *DeliveryItemQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(deliveryitem.Columns))
		selectedFields = []string{deliveryitem.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "delivery":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&DeliveryClient{config: di.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, deliveryImplementors)...); err != nil {
				return err
			}
			di.withDelivery = query
		case "productname":
			if _, ok := fieldSeen[deliveryitem.FieldProductName]; !ok {
				selectedFields = append(selectedFields, deliveryitem.FieldProductName)
				fieldSeen[deliveryitem.FieldProductName] = struct{}{}
			}
		case "quantity":
			if _, ok := fieldSeen[deliveryitem.FieldQuantity]; !ok {
				selectedFields = append(selectedFields, deliveryitem.FieldQuantity)
				fieldSeen[deliveryitem.FieldQuantity] = struct{}{}
			}
		case "price":
			if _, ok := fieldSeen[deliveryitem.FieldPrice]; !ok {
				selectedFields = append(selectedFields, deliveryitem.FieldPrice)
				fieldSeen[deliveryitem.FieldPrice] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		di.Select(selectedFields...)
	}
	return nil
}

type deliveryitemPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []DeliveryItemPaginateOption
}

func newDeliveryItemPaginateArgs(rv map[string]any) *deliveryitemPaginateArgs {
	args := &deliveryitemPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*DeliveryItemWhereInput); ok {
		args.opts = append(args.opts, WithDeliveryItemFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok || v == nil {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
