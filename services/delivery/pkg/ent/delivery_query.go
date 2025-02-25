// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"delivery/pkg/ent/delivery"
	"delivery/pkg/ent/deliveryitem"
	"delivery/pkg/ent/predicate"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeliveryQuery is the builder for querying Delivery entities.
type DeliveryQuery struct {
	config
	ctx                   *QueryContext
	order                 []delivery.OrderOption
	inters                []Interceptor
	predicates            []predicate.Delivery
	withDeliveryItem      *DeliveryItemQuery
	modifiers             []func(*sql.Selector)
	loadTotal             []func(context.Context, []*Delivery) error
	withNamedDeliveryItem map[string]*DeliveryItemQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DeliveryQuery builder.
func (dq *DeliveryQuery) Where(ps ...predicate.Delivery) *DeliveryQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit the number of records to be returned by this query.
func (dq *DeliveryQuery) Limit(limit int) *DeliveryQuery {
	dq.ctx.Limit = &limit
	return dq
}

// Offset to start from.
func (dq *DeliveryQuery) Offset(offset int) *DeliveryQuery {
	dq.ctx.Offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DeliveryQuery) Unique(unique bool) *DeliveryQuery {
	dq.ctx.Unique = &unique
	return dq
}

// Order specifies how the records should be ordered.
func (dq *DeliveryQuery) Order(o ...delivery.OrderOption) *DeliveryQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryDeliveryItem chains the current query on the "delivery_item" edge.
func (dq *DeliveryQuery) QueryDeliveryItem() *DeliveryItemQuery {
	query := (&DeliveryItemClient{config: dq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(delivery.Table, delivery.FieldID, selector),
			sqlgraph.To(deliveryitem.Table, deliveryitem.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, delivery.DeliveryItemTable, delivery.DeliveryItemColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Delivery entity from the query.
// Returns a *NotFoundError when no Delivery was found.
func (dq *DeliveryQuery) First(ctx context.Context) (*Delivery, error) {
	nodes, err := dq.Limit(1).All(setContextOp(ctx, dq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{delivery.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DeliveryQuery) FirstX(ctx context.Context) *Delivery {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Delivery ID from the query.
// Returns a *NotFoundError when no Delivery ID was found.
func (dq *DeliveryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(1).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{delivery.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DeliveryQuery) FirstIDX(ctx context.Context) int {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Delivery entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Delivery entity is found.
// Returns a *NotFoundError when no Delivery entities are found.
func (dq *DeliveryQuery) Only(ctx context.Context) (*Delivery, error) {
	nodes, err := dq.Limit(2).All(setContextOp(ctx, dq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{delivery.Label}
	default:
		return nil, &NotSingularError{delivery.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DeliveryQuery) OnlyX(ctx context.Context) *Delivery {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Delivery ID in the query.
// Returns a *NotSingularError when more than one Delivery ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DeliveryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = dq.Limit(2).IDs(setContextOp(ctx, dq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{delivery.Label}
	default:
		err = &NotSingularError{delivery.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DeliveryQuery) OnlyIDX(ctx context.Context) int {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Deliveries.
func (dq *DeliveryQuery) All(ctx context.Context) ([]*Delivery, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryAll)
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Delivery, *DeliveryQuery]()
	return withInterceptors[[]*Delivery](ctx, dq, qr, dq.inters)
}

// AllX is like All, but panics if an error occurs.
func (dq *DeliveryQuery) AllX(ctx context.Context) []*Delivery {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Delivery IDs.
func (dq *DeliveryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if dq.ctx.Unique == nil && dq.path != nil {
		dq.Unique(true)
	}
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryIDs)
	if err = dq.Select(delivery.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DeliveryQuery) IDsX(ctx context.Context) []int {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DeliveryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryCount)
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, dq, querierCount[*DeliveryQuery](), dq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DeliveryQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DeliveryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, dq.ctx, ent.OpQueryExist)
	switch _, err := dq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DeliveryQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DeliveryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DeliveryQuery) Clone() *DeliveryQuery {
	if dq == nil {
		return nil
	}
	return &DeliveryQuery{
		config:           dq.config,
		ctx:              dq.ctx.Clone(),
		order:            append([]delivery.OrderOption{}, dq.order...),
		inters:           append([]Interceptor{}, dq.inters...),
		predicates:       append([]predicate.Delivery{}, dq.predicates...),
		withDeliveryItem: dq.withDeliveryItem.Clone(),
		// clone intermediate query.
		sql:  dq.sql.Clone(),
		path: dq.path,
	}
}

// WithDeliveryItem tells the query-builder to eager-load the nodes that are connected to
// the "delivery_item" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DeliveryQuery) WithDeliveryItem(opts ...func(*DeliveryItemQuery)) *DeliveryQuery {
	query := (&DeliveryItemClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	dq.withDeliveryItem = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderID int `json:"order_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Delivery.Query().
//		GroupBy(delivery.FieldOrderID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DeliveryQuery) GroupBy(field string, fields ...string) *DeliveryGroupBy {
	dq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &DeliveryGroupBy{build: dq}
	grbuild.flds = &dq.ctx.Fields
	grbuild.label = delivery.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderID int `json:"order_id,omitempty"`
//	}
//
//	client.Delivery.Query().
//		Select(delivery.FieldOrderID).
//		Scan(ctx, &v)
func (dq *DeliveryQuery) Select(fields ...string) *DeliverySelect {
	dq.ctx.Fields = append(dq.ctx.Fields, fields...)
	sbuild := &DeliverySelect{DeliveryQuery: dq}
	sbuild.label = delivery.Label
	sbuild.flds, sbuild.scan = &dq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a DeliverySelect configured with the given aggregations.
func (dq *DeliveryQuery) Aggregate(fns ...AggregateFunc) *DeliverySelect {
	return dq.Select().Aggregate(fns...)
}

func (dq *DeliveryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range dq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, dq); err != nil {
				return err
			}
		}
	}
	for _, f := range dq.ctx.Fields {
		if !delivery.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DeliveryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Delivery, error) {
	var (
		nodes       = []*Delivery{}
		_spec       = dq.querySpec()
		loadedTypes = [1]bool{
			dq.withDeliveryItem != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Delivery).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Delivery{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withDeliveryItem; query != nil {
		if err := dq.loadDeliveryItem(ctx, query, nodes,
			func(n *Delivery) { n.Edges.DeliveryItem = []*DeliveryItem{} },
			func(n *Delivery, e *DeliveryItem) { n.Edges.DeliveryItem = append(n.Edges.DeliveryItem, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range dq.withNamedDeliveryItem {
		if err := dq.loadDeliveryItem(ctx, query, nodes,
			func(n *Delivery) { n.appendNamedDeliveryItem(name) },
			func(n *Delivery, e *DeliveryItem) { n.appendNamedDeliveryItem(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range dq.loadTotal {
		if err := dq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DeliveryQuery) loadDeliveryItem(ctx context.Context, query *DeliveryItemQuery, nodes []*Delivery, init func(*Delivery), assign func(*Delivery, *DeliveryItem)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Delivery)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.DeliveryItem(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(delivery.DeliveryItemColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.delivery_delivery_item
		if fk == nil {
			return fmt.Errorf(`foreign-key "delivery_delivery_item" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "delivery_delivery_item" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (dq *DeliveryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	if len(dq.modifiers) > 0 {
		_spec.Modifiers = dq.modifiers
	}
	_spec.Node.Columns = dq.ctx.Fields
	if len(dq.ctx.Fields) > 0 {
		_spec.Unique = dq.ctx.Unique != nil && *dq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DeliveryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(delivery.Table, delivery.Columns, sqlgraph.NewFieldSpec(delivery.FieldID, field.TypeInt))
	_spec.From = dq.sql
	if unique := dq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if dq.path != nil {
		_spec.Unique = true
	}
	if fields := dq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, delivery.FieldID)
		for i := range fields {
			if fields[i] != delivery.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DeliveryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(delivery.Table)
	columns := dq.ctx.Fields
	if len(columns) == 0 {
		columns = delivery.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.ctx.Unique != nil && *dq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedDeliveryItem tells the query-builder to eager-load the nodes that are connected to the "delivery_item"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (dq *DeliveryQuery) WithNamedDeliveryItem(name string, opts ...func(*DeliveryItemQuery)) *DeliveryQuery {
	query := (&DeliveryItemClient{config: dq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if dq.withNamedDeliveryItem == nil {
		dq.withNamedDeliveryItem = make(map[string]*DeliveryItemQuery)
	}
	dq.withNamedDeliveryItem[name] = query
	return dq
}

// DeliveryGroupBy is the group-by builder for Delivery entities.
type DeliveryGroupBy struct {
	selector
	build *DeliveryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DeliveryGroupBy) Aggregate(fns ...AggregateFunc) *DeliveryGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the selector query and scans the result into the given value.
func (dgb *DeliveryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, dgb.build.ctx, ent.OpQueryGroupBy)
	if err := dgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeliveryQuery, *DeliveryGroupBy](ctx, dgb.build, dgb, dgb.build.inters, v)
}

func (dgb *DeliveryGroupBy) sqlScan(ctx context.Context, root *DeliveryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*dgb.flds)+len(dgb.fns))
		for _, f := range *dgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*dgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// DeliverySelect is the builder for selecting fields of Delivery entities.
type DeliverySelect struct {
	*DeliveryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ds *DeliverySelect) Aggregate(fns ...AggregateFunc) *DeliverySelect {
	ds.fns = append(ds.fns, fns...)
	return ds
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DeliverySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ds.ctx, ent.OpQuerySelect)
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*DeliveryQuery, *DeliverySelect](ctx, ds.DeliveryQuery, ds, ds.inters, v)
}

func (ds *DeliverySelect) sqlScan(ctx context.Context, root *DeliveryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ds.fns))
	for _, fn := range ds.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ds.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
