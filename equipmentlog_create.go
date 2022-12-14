// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/equipmentlog"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentLogCreate is the builder for creating a EquipmentLog entity.
type EquipmentLogCreate struct {
	config
	mutation *EquipmentLogMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (elc *EquipmentLogCreate) SetVersion(i int64) *EquipmentLogCreate {
	elc.mutation.SetVersion(i)
	return elc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (elc *EquipmentLogCreate) SetNillableVersion(i *int64) *EquipmentLogCreate {
	if i != nil {
		elc.SetVersion(*i)
	}
	return elc
}

// SetCreatedBy sets the "created_by" field.
func (elc *EquipmentLogCreate) SetCreatedBy(d datasource.UUID) *EquipmentLogCreate {
	elc.mutation.SetCreatedBy(d)
	return elc
}

// SetCreatedAt sets the "created_at" field.
func (elc *EquipmentLogCreate) SetCreatedAt(i int64) *EquipmentLogCreate {
	elc.mutation.SetCreatedAt(i)
	return elc
}

// SetUpdatedBy sets the "updated_by" field.
func (elc *EquipmentLogCreate) SetUpdatedBy(d datasource.UUID) *EquipmentLogCreate {
	elc.mutation.SetUpdatedBy(d)
	return elc
}

// SetUpdatedAt sets the "updated_at" field.
func (elc *EquipmentLogCreate) SetUpdatedAt(i int64) *EquipmentLogCreate {
	elc.mutation.SetUpdatedAt(i)
	return elc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (elc *EquipmentLogCreate) SetNillableUpdatedAt(i *int64) *EquipmentLogCreate {
	if i != nil {
		elc.SetUpdatedAt(*i)
	}
	return elc
}

// SetEquipmentID sets the "equipment_id" field.
func (elc *EquipmentLogCreate) SetEquipmentID(d datasource.UUID) *EquipmentLogCreate {
	elc.mutation.SetEquipmentID(d)
	return elc
}

// SetRequestID sets the "request_id" field.
func (elc *EquipmentLogCreate) SetRequestID(i int64) *EquipmentLogCreate {
	elc.mutation.SetRequestID(i)
	return elc
}

// SetState sets the "state" field.
func (elc *EquipmentLogCreate) SetState(i int) *EquipmentLogCreate {
	elc.mutation.SetState(i)
	return elc
}

// SetDataLink sets the "data_link" field.
func (elc *EquipmentLogCreate) SetDataLink(d datasource.UUID) *EquipmentLogCreate {
	elc.mutation.SetDataLink(d)
	return elc
}

// SetID sets the "id" field.
func (elc *EquipmentLogCreate) SetID(d datasource.UUID) *EquipmentLogCreate {
	elc.mutation.SetID(d)
	return elc
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (elc *EquipmentLogCreate) SetEquipment(e *Equipment) *EquipmentLogCreate {
	return elc.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentLogMutation object of the builder.
func (elc *EquipmentLogCreate) Mutation() *EquipmentLogMutation {
	return elc.mutation
}

// Save creates the EquipmentLog in the database.
func (elc *EquipmentLogCreate) Save(ctx context.Context) (*EquipmentLog, error) {
	var (
		err  error
		node *EquipmentLog
	)
	elc.defaults()
	if len(elc.hooks) == 0 {
		if err = elc.check(); err != nil {
			return nil, err
		}
		node, err = elc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = elc.check(); err != nil {
				return nil, err
			}
			elc.mutation = mutation
			if node, err = elc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(elc.hooks) - 1; i >= 0; i-- {
			if elc.hooks[i] == nil {
				return nil, fmt.Errorf("cwmodel: uninitialized hook (forgotten import cwmodel/runtime?)")
			}
			mut = elc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, elc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EquipmentLog)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EquipmentLogMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (elc *EquipmentLogCreate) SaveX(ctx context.Context) *EquipmentLog {
	v, err := elc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (elc *EquipmentLogCreate) Exec(ctx context.Context) error {
	_, err := elc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elc *EquipmentLogCreate) ExecX(ctx context.Context) {
	if err := elc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (elc *EquipmentLogCreate) defaults() {
	if _, ok := elc.mutation.Version(); !ok {
		v := equipmentlog.DefaultVersion
		elc.mutation.SetVersion(v)
	}
	if _, ok := elc.mutation.UpdatedAt(); !ok {
		v := equipmentlog.DefaultUpdatedAt
		elc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (elc *EquipmentLogCreate) check() error {
	if _, ok := elc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`cwmodel: missing required field "EquipmentLog.version"`)}
	}
	if _, ok := elc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cwmodel: missing required field "EquipmentLog.created_by"`)}
	}
	if _, ok := elc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cwmodel: missing required field "EquipmentLog.created_at"`)}
	}
	if _, ok := elc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cwmodel: missing required field "EquipmentLog.updated_by"`)}
	}
	if _, ok := elc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cwmodel: missing required field "EquipmentLog.updated_at"`)}
	}
	if _, ok := elc.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment_id", err: errors.New(`cwmodel: missing required field "EquipmentLog.equipment_id"`)}
	}
	if _, ok := elc.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request_id", err: errors.New(`cwmodel: missing required field "EquipmentLog.request_id"`)}
	}
	if _, ok := elc.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`cwmodel: missing required field "EquipmentLog.state"`)}
	}
	if _, ok := elc.mutation.DataLink(); !ok {
		return &ValidationError{Name: "data_link", err: errors.New(`cwmodel: missing required field "EquipmentLog.data_link"`)}
	}
	if _, ok := elc.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment", err: errors.New(`cwmodel: missing required edge "EquipmentLog.equipment"`)}
	}
	return nil
}

func (elc *EquipmentLogCreate) sqlSave(ctx context.Context) (*EquipmentLog, error) {
	_node, _spec := elc.createSpec()
	if err := sqlgraph.CreateNode(ctx, elc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = datasource.UUID(id)
	}
	return _node, nil
}

func (elc *EquipmentLogCreate) createSpec() (*EquipmentLog, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentLog{config: elc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: equipmentlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentlog.FieldID,
			},
		}
	)
	if id, ok := elc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := elc.mutation.Version(); ok {
		_spec.SetField(equipmentlog.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := elc.mutation.CreatedBy(); ok {
		_spec.SetField(equipmentlog.FieldCreatedBy, field.TypeUint64, value)
		_node.CreatedBy = value
	}
	if value, ok := elc.mutation.CreatedAt(); ok {
		_spec.SetField(equipmentlog.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := elc.mutation.UpdatedBy(); ok {
		_spec.SetField(equipmentlog.FieldUpdatedBy, field.TypeUint64, value)
		_node.UpdatedBy = value
	}
	if value, ok := elc.mutation.UpdatedAt(); ok {
		_spec.SetField(equipmentlog.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := elc.mutation.RequestID(); ok {
		_spec.SetField(equipmentlog.FieldRequestID, field.TypeInt64, value)
		_node.RequestID = value
	}
	if value, ok := elc.mutation.State(); ok {
		_spec.SetField(equipmentlog.FieldState, field.TypeInt, value)
		_node.State = value
	}
	if value, ok := elc.mutation.DataLink(); ok {
		_spec.SetField(equipmentlog.FieldDataLink, field.TypeUint64, value)
		_node.DataLink = value
	}
	if nodes := elc.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentlog.EquipmentTable,
			Columns: []string{equipmentlog.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.EquipmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EquipmentLogCreateBulk is the builder for creating many EquipmentLog entities in bulk.
type EquipmentLogCreateBulk struct {
	config
	builders []*EquipmentLogCreate
}

// Save creates the EquipmentLog entities in the database.
func (elcb *EquipmentLogCreateBulk) Save(ctx context.Context) ([]*EquipmentLog, error) {
	specs := make([]*sqlgraph.CreateSpec, len(elcb.builders))
	nodes := make([]*EquipmentLog, len(elcb.builders))
	mutators := make([]Mutator, len(elcb.builders))
	for i := range elcb.builders {
		func(i int, root context.Context) {
			builder := elcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentLogMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, elcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, elcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = datasource.UUID(id)
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
		if _, err := mutators[0].Mutate(ctx, elcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (elcb *EquipmentLogCreateBulk) SaveX(ctx context.Context) []*EquipmentLog {
	v, err := elcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (elcb *EquipmentLogCreateBulk) Exec(ctx context.Context) error {
	_, err := elcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (elcb *EquipmentLogCreateBulk) ExecX(ctx context.Context) {
	if err := elcb.Exec(ctx); err != nil {
		panic(err)
	}
}
