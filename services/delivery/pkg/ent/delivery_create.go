// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"delivery/pkg/ent/delivery"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// DeliveryCreate is the builder for creating a Delivery entity.
type DeliveryCreate struct {
	config
	mutation *DeliveryMutation
	hooks    []Hook
}

// SetOrderID sets the "order_id" field.
func (dc *DeliveryCreate) SetOrderID(i int) *DeliveryCreate {
	dc.mutation.SetOrderID(i)
	return dc
}

// SetUserID sets the "user_id" field.
func (dc *DeliveryCreate) SetUserID(i int) *DeliveryCreate {
	dc.mutation.SetUserID(i)
	return dc
}

// SetStatus sets the "status" field.
func (dc *DeliveryCreate) SetStatus(d delivery.Status) *DeliveryCreate {
	dc.mutation.SetStatus(d)
	return dc
}

// SetTrackingNumber sets the "tracking_number" field.
func (dc *DeliveryCreate) SetTrackingNumber(s string) *DeliveryCreate {
	dc.mutation.SetTrackingNumber(s)
	return dc
}

// SetNillableTrackingNumber sets the "tracking_number" field if the given value is not nil.
func (dc *DeliveryCreate) SetNillableTrackingNumber(s *string) *DeliveryCreate {
	if s != nil {
		dc.SetTrackingNumber(*s)
	}
	return dc
}

// SetCreatedAt sets the "created_at" field.
func (dc *DeliveryCreate) SetCreatedAt(t time.Time) *DeliveryCreate {
	dc.mutation.SetCreatedAt(t)
	return dc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (dc *DeliveryCreate) SetNillableCreatedAt(t *time.Time) *DeliveryCreate {
	if t != nil {
		dc.SetCreatedAt(*t)
	}
	return dc
}

// Mutation returns the DeliveryMutation object of the builder.
func (dc *DeliveryCreate) Mutation() *DeliveryMutation {
	return dc.mutation
}

// Save creates the Delivery in the database.
func (dc *DeliveryCreate) Save(ctx context.Context) (*Delivery, error) {
	dc.defaults()
	return withHooks(ctx, dc.sqlSave, dc.mutation, dc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DeliveryCreate) SaveX(ctx context.Context) *Delivery {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dc *DeliveryCreate) Exec(ctx context.Context) error {
	_, err := dc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dc *DeliveryCreate) ExecX(ctx context.Context) {
	if err := dc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (dc *DeliveryCreate) defaults() {
	if _, ok := dc.mutation.CreatedAt(); !ok {
		v := delivery.DefaultCreatedAt()
		dc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (dc *DeliveryCreate) check() error {
	if _, ok := dc.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "Delivery.order_id"`)}
	}
	if _, ok := dc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Delivery.user_id"`)}
	}
	if _, ok := dc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Delivery.status"`)}
	}
	if v, ok := dc.mutation.Status(); ok {
		if err := delivery.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Delivery.status": %w`, err)}
		}
	}
	if _, ok := dc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Delivery.created_at"`)}
	}
	return nil
}

func (dc *DeliveryCreate) sqlSave(ctx context.Context) (*Delivery, error) {
	if err := dc.check(); err != nil {
		return nil, err
	}
	_node, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	dc.mutation.id = &_node.ID
	dc.mutation.done = true
	return _node, nil
}

func (dc *DeliveryCreate) createSpec() (*Delivery, *sqlgraph.CreateSpec) {
	var (
		_node = &Delivery{config: dc.config}
		_spec = sqlgraph.NewCreateSpec(delivery.Table, sqlgraph.NewFieldSpec(delivery.FieldID, field.TypeInt))
	)
	if value, ok := dc.mutation.OrderID(); ok {
		_spec.SetField(delivery.FieldOrderID, field.TypeInt, value)
		_node.OrderID = value
	}
	if value, ok := dc.mutation.UserID(); ok {
		_spec.SetField(delivery.FieldUserID, field.TypeInt, value)
		_node.UserID = value
	}
	if value, ok := dc.mutation.Status(); ok {
		_spec.SetField(delivery.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := dc.mutation.TrackingNumber(); ok {
		_spec.SetField(delivery.FieldTrackingNumber, field.TypeString, value)
		_node.TrackingNumber = value
	}
	if value, ok := dc.mutation.CreatedAt(); ok {
		_spec.SetField(delivery.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	return _node, _spec
}

// DeliveryCreateBulk is the builder for creating many Delivery entities in bulk.
type DeliveryCreateBulk struct {
	config
	err      error
	builders []*DeliveryCreate
}

// Save creates the Delivery entities in the database.
func (dcb *DeliveryCreateBulk) Save(ctx context.Context) ([]*Delivery, error) {
	if dcb.err != nil {
		return nil, dcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(dcb.builders))
	nodes := make([]*Delivery, len(dcb.builders))
	mutators := make([]Mutator, len(dcb.builders))
	for i := range dcb.builders {
		func(i int, root context.Context) {
			builder := dcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DeliveryMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, dcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, dcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, dcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (dcb *DeliveryCreateBulk) SaveX(ctx context.Context) []*Delivery {
	v, err := dcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (dcb *DeliveryCreateBulk) Exec(ctx context.Context) error {
	_, err := dcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dcb *DeliveryCreateBulk) ExecX(ctx context.Context) {
	if err := dcb.Exec(ctx); err != nil {
		panic(err)
	}
}
