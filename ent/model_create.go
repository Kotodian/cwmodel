// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/ent/model"
	"github.com/Kotodian/gokit/datasource"
)

// ModelCreate is the builder for creating a Model entity.
type ModelCreate struct {
	config
	mutation *ModelMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (mc *ModelCreate) SetVersion(i int64) *ModelCreate {
	mc.mutation.SetVersion(i)
	return mc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (mc *ModelCreate) SetNillableVersion(i *int64) *ModelCreate {
	if i != nil {
		mc.SetVersion(*i)
	}
	return mc
}

// SetCreatedBy sets the "created_by" field.
func (mc *ModelCreate) SetCreatedBy(d datasource.UUID) *ModelCreate {
	mc.mutation.SetCreatedBy(d)
	return mc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (mc *ModelCreate) SetNillableCreatedBy(d *datasource.UUID) *ModelCreate {
	if d != nil {
		mc.SetCreatedBy(*d)
	}
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *ModelCreate) SetCreatedAt(i int64) *ModelCreate {
	mc.mutation.SetCreatedAt(i)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *ModelCreate) SetNillableCreatedAt(i *int64) *ModelCreate {
	if i != nil {
		mc.SetCreatedAt(*i)
	}
	return mc
}

// SetUpdatedBy sets the "updated_by" field.
func (mc *ModelCreate) SetUpdatedBy(d datasource.UUID) *ModelCreate {
	mc.mutation.SetUpdatedBy(d)
	return mc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (mc *ModelCreate) SetNillableUpdatedBy(d *datasource.UUID) *ModelCreate {
	if d != nil {
		mc.SetUpdatedBy(*d)
	}
	return mc
}

// SetUpdatedAt sets the "updated_at" field.
func (mc *ModelCreate) SetUpdatedAt(i int64) *ModelCreate {
	mc.mutation.SetUpdatedAt(i)
	return mc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (mc *ModelCreate) SetNillableUpdatedAt(i *int64) *ModelCreate {
	if i != nil {
		mc.SetUpdatedAt(*i)
	}
	return mc
}

// SetCode sets the "code" field.
func (mc *ModelCreate) SetCode(s string) *ModelCreate {
	mc.mutation.SetCode(s)
	return mc
}

// SetName sets the "name" field.
func (mc *ModelCreate) SetName(s string) *ModelCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetPhaseCategory sets the "phase_category" field.
func (mc *ModelCreate) SetPhaseCategory(s string) *ModelCreate {
	mc.mutation.SetPhaseCategory(s)
	return mc
}

// SetCurrentCategory sets the "current_category" field.
func (mc *ModelCreate) SetCurrentCategory(s string) *ModelCreate {
	mc.mutation.SetCurrentCategory(s)
	return mc
}

// SetID sets the "id" field.
func (mc *ModelCreate) SetID(d datasource.UUID) *ModelCreate {
	mc.mutation.SetID(d)
	return mc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (mc *ModelCreate) SetNillableID(d *datasource.UUID) *ModelCreate {
	if d != nil {
		mc.SetID(*d)
	}
	return mc
}

// Mutation returns the ModelMutation object of the builder.
func (mc *ModelCreate) Mutation() *ModelMutation {
	return mc.mutation
}

// Save creates the Model in the database.
func (mc *ModelCreate) Save(ctx context.Context) (*Model, error) {
	var (
		err  error
		node *Model
	)
	mc.defaults()
	if len(mc.hooks) == 0 {
		if err = mc.check(); err != nil {
			return nil, err
		}
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModelMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mc.check(); err != nil {
				return nil, err
			}
			mc.mutation = mutation
			if node, err = mc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			if mc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, mc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Model)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ModelMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *ModelCreate) SaveX(ctx context.Context) *Model {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *ModelCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *ModelCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *ModelCreate) defaults() {
	if _, ok := mc.mutation.Version(); !ok {
		v := model.DefaultVersion
		mc.mutation.SetVersion(v)
	}
	if _, ok := mc.mutation.CreatedBy(); !ok {
		v := model.DefaultCreatedBy
		mc.mutation.SetCreatedBy(v)
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := model.DefaultCreatedAt
		mc.mutation.SetCreatedAt(v)
	}
	if _, ok := mc.mutation.UpdatedBy(); !ok {
		v := model.DefaultUpdatedBy
		mc.mutation.SetUpdatedBy(v)
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		v := model.DefaultUpdatedAt
		mc.mutation.SetUpdatedAt(v)
	}
	if _, ok := mc.mutation.ID(); !ok {
		v := model.DefaultID
		mc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *ModelCreate) check() error {
	if _, ok := mc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "Model.version"`)}
	}
	if _, ok := mc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Model.created_by"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Model.created_at"`)}
	}
	if _, ok := mc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Model.updated_by"`)}
	}
	if _, ok := mc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Model.updated_at"`)}
	}
	if _, ok := mc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Model.code"`)}
	}
	if _, ok := mc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Model.name"`)}
	}
	if _, ok := mc.mutation.PhaseCategory(); !ok {
		return &ValidationError{Name: "phase_category", err: errors.New(`ent: missing required field "Model.phase_category"`)}
	}
	if _, ok := mc.mutation.CurrentCategory(); !ok {
		return &ValidationError{Name: "current_category", err: errors.New(`ent: missing required field "Model.current_category"`)}
	}
	return nil
}

func (mc *ModelCreate) sqlSave(ctx context.Context) (*Model, error) {
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
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

func (mc *ModelCreate) createSpec() (*Model, *sqlgraph.CreateSpec) {
	var (
		_node = &Model{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: model.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: model.FieldID,
			},
		}
	)
	if id, ok := mc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := mc.mutation.Version(); ok {
		_spec.SetField(model.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := mc.mutation.CreatedBy(); ok {
		_spec.SetField(model.FieldCreatedBy, field.TypeUint64, value)
		_node.CreatedBy = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(model.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.UpdatedBy(); ok {
		_spec.SetField(model.FieldUpdatedBy, field.TypeUint64, value)
		_node.UpdatedBy = value
	}
	if value, ok := mc.mutation.UpdatedAt(); ok {
		_spec.SetField(model.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := mc.mutation.Code(); ok {
		_spec.SetField(model.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := mc.mutation.Name(); ok {
		_spec.SetField(model.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := mc.mutation.PhaseCategory(); ok {
		_spec.SetField(model.FieldPhaseCategory, field.TypeString, value)
		_node.PhaseCategory = value
	}
	if value, ok := mc.mutation.CurrentCategory(); ok {
		_spec.SetField(model.FieldCurrentCategory, field.TypeString, value)
		_node.CurrentCategory = value
	}
	return _node, _spec
}

// ModelCreateBulk is the builder for creating many Model entities in bulk.
type ModelCreateBulk struct {
	config
	builders []*ModelCreate
}

// Save creates the Model entities in the database.
func (mcb *ModelCreateBulk) Save(ctx context.Context) ([]*Model, error) {
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Model, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ModelMutation)
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
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *ModelCreateBulk) SaveX(ctx context.Context) []*Model {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *ModelCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *ModelCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}
