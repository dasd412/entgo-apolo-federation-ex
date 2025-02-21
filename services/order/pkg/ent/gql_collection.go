// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"order/pkg/ent/order"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (o *OrderQuery) CollectFields(ctx context.Context, satisfies ...string) (*OrderQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return o, nil
	}
	if err := o.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return o, nil
}

func (o *OrderQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(order.Columns))
		selectedFields = []string{order.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "userID":
			if _, ok := fieldSeen[order.FieldUserID]; !ok {
				selectedFields = append(selectedFields, order.FieldUserID)
				fieldSeen[order.FieldUserID] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[order.FieldStatus]; !ok {
				selectedFields = append(selectedFields, order.FieldStatus)
				fieldSeen[order.FieldStatus] = struct{}{}
			}
		case "totalPrice":
			if _, ok := fieldSeen[order.FieldTotalPrice]; !ok {
				selectedFields = append(selectedFields, order.FieldTotalPrice)
				fieldSeen[order.FieldTotalPrice] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[order.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, order.FieldCreatedAt)
				fieldSeen[order.FieldCreatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		o.Select(selectedFields...)
	}
	return nil
}

type orderPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []OrderPaginateOption
}

func newOrderPaginateArgs(rv map[string]any) *orderPaginateArgs {
	args := &orderPaginateArgs{}
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
	if v, ok := rv[whereField].(*OrderWhereInput); ok {
		args.opts = append(args.opts, WithOrderFilter(v.Filter))
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
		if !ok {
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
