// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentfirmwareeffect"
	"github.com/Kotodian/ent-practice/ent/firmware"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentFirmwareEffectCreate is the builder for creating a EquipmentFirmwareEffect entity.
type EquipmentFirmwareEffectCreate struct {
	config
	mutation *EquipmentFirmwareEffectMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (efec *EquipmentFirmwareEffectCreate) SetVersion(i int64) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetVersion(i)
	return efec
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (efec *EquipmentFirmwareEffectCreate) SetNillableVersion(i *int64) *EquipmentFirmwareEffectCreate {
	if i != nil {
		efec.SetVersion(*i)
	}
	return efec
}

// SetCreatedBy sets the "created_by" field.
func (efec *EquipmentFirmwareEffectCreate) SetCreatedBy(d datasource.UUID) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetCreatedBy(d)
	return efec
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (efec *EquipmentFirmwareEffectCreate) SetNillableCreatedBy(d *datasource.UUID) *EquipmentFirmwareEffectCreate {
	if d != nil {
		efec.SetCreatedBy(*d)
	}
	return efec
}

// SetCreatedAt sets the "created_at" field.
func (efec *EquipmentFirmwareEffectCreate) SetCreatedAt(i int64) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetCreatedAt(i)
	return efec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (efec *EquipmentFirmwareEffectCreate) SetNillableCreatedAt(i *int64) *EquipmentFirmwareEffectCreate {
	if i != nil {
		efec.SetCreatedAt(*i)
	}
	return efec
}

// SetUpdatedBy sets the "updated_by" field.
func (efec *EquipmentFirmwareEffectCreate) SetUpdatedBy(d datasource.UUID) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetUpdatedBy(d)
	return efec
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (efec *EquipmentFirmwareEffectCreate) SetNillableUpdatedBy(d *datasource.UUID) *EquipmentFirmwareEffectCreate {
	if d != nil {
		efec.SetUpdatedBy(*d)
	}
	return efec
}

// SetUpdatedAt sets the "updated_at" field.
func (efec *EquipmentFirmwareEffectCreate) SetUpdatedAt(i int64) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetUpdatedAt(i)
	return efec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (efec *EquipmentFirmwareEffectCreate) SetNillableUpdatedAt(i *int64) *EquipmentFirmwareEffectCreate {
	if i != nil {
		efec.SetUpdatedAt(*i)
	}
	return efec
}

// SetRequestID sets the "request_id" field.
func (efec *EquipmentFirmwareEffectCreate) SetRequestID(i int64) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetRequestID(i)
	return efec
}

// SetState sets the "state" field.
func (efec *EquipmentFirmwareEffectCreate) SetState(i int) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetState(i)
	return efec
}

// SetID sets the "id" field.
func (efec *EquipmentFirmwareEffectCreate) SetID(d datasource.UUID) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetID(d)
	return efec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (efec *EquipmentFirmwareEffectCreate) SetNillableID(d *datasource.UUID) *EquipmentFirmwareEffectCreate {
	if d != nil {
		efec.SetID(*d)
	}
	return efec
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (efec *EquipmentFirmwareEffectCreate) SetEquipmentID(id datasource.UUID) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetEquipmentID(id)
	return efec
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (efec *EquipmentFirmwareEffectCreate) SetNillableEquipmentID(id *datasource.UUID) *EquipmentFirmwareEffectCreate {
	if id != nil {
		efec = efec.SetEquipmentID(*id)
	}
	return efec
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (efec *EquipmentFirmwareEffectCreate) SetEquipment(e *Equipment) *EquipmentFirmwareEffectCreate {
	return efec.SetEquipmentID(e.ID)
}

// SetFirmwareID sets the "firmware" edge to the Firmware entity by ID.
func (efec *EquipmentFirmwareEffectCreate) SetFirmwareID(id datasource.UUID) *EquipmentFirmwareEffectCreate {
	efec.mutation.SetFirmwareID(id)
	return efec
}

// SetNillableFirmwareID sets the "firmware" edge to the Firmware entity by ID if the given value is not nil.
func (efec *EquipmentFirmwareEffectCreate) SetNillableFirmwareID(id *datasource.UUID) *EquipmentFirmwareEffectCreate {
	if id != nil {
		efec = efec.SetFirmwareID(*id)
	}
	return efec
}

// SetFirmware sets the "firmware" edge to the Firmware entity.
func (efec *EquipmentFirmwareEffectCreate) SetFirmware(f *Firmware) *EquipmentFirmwareEffectCreate {
	return efec.SetFirmwareID(f.ID)
}

// Mutation returns the EquipmentFirmwareEffectMutation object of the builder.
func (efec *EquipmentFirmwareEffectCreate) Mutation() *EquipmentFirmwareEffectMutation {
	return efec.mutation
}

// Save creates the EquipmentFirmwareEffect in the database.
func (efec *EquipmentFirmwareEffectCreate) Save(ctx context.Context) (*EquipmentFirmwareEffect, error) {
	var (
		err  error
		node *EquipmentFirmwareEffect
	)
	efec.defaults()
	if len(efec.hooks) == 0 {
		if err = efec.check(); err != nil {
			return nil, err
		}
		node, err = efec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentFirmwareEffectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = efec.check(); err != nil {
				return nil, err
			}
			efec.mutation = mutation
			if node, err = efec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(efec.hooks) - 1; i >= 0; i-- {
			if efec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = efec.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, efec.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*EquipmentFirmwareEffect)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EquipmentFirmwareEffectMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (efec *EquipmentFirmwareEffectCreate) SaveX(ctx context.Context) *EquipmentFirmwareEffect {
	v, err := efec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (efec *EquipmentFirmwareEffectCreate) Exec(ctx context.Context) error {
	_, err := efec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efec *EquipmentFirmwareEffectCreate) ExecX(ctx context.Context) {
	if err := efec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (efec *EquipmentFirmwareEffectCreate) defaults() {
	if _, ok := efec.mutation.Version(); !ok {
		v := equipmentfirmwareeffect.DefaultVersion
		efec.mutation.SetVersion(v)
	}
	if _, ok := efec.mutation.CreatedBy(); !ok {
		v := equipmentfirmwareeffect.DefaultCreatedBy
		efec.mutation.SetCreatedBy(v)
	}
	if _, ok := efec.mutation.CreatedAt(); !ok {
		v := equipmentfirmwareeffect.DefaultCreatedAt
		efec.mutation.SetCreatedAt(v)
	}
	if _, ok := efec.mutation.UpdatedBy(); !ok {
		v := equipmentfirmwareeffect.DefaultUpdatedBy
		efec.mutation.SetUpdatedBy(v)
	}
	if _, ok := efec.mutation.UpdatedAt(); !ok {
		v := equipmentfirmwareeffect.DefaultUpdatedAt
		efec.mutation.SetUpdatedAt(v)
	}
	if _, ok := efec.mutation.ID(); !ok {
		v := equipmentfirmwareeffect.DefaultID
		efec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (efec *EquipmentFirmwareEffectCreate) check() error {
	if _, ok := efec.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`ent: missing required field "EquipmentFirmwareEffect.version"`)}
	}
	if _, ok := efec.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "EquipmentFirmwareEffect.created_by"`)}
	}
	if _, ok := efec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "EquipmentFirmwareEffect.created_at"`)}
	}
	if _, ok := efec.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "EquipmentFirmwareEffect.updated_by"`)}
	}
	if _, ok := efec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "EquipmentFirmwareEffect.updated_at"`)}
	}
	if _, ok := efec.mutation.RequestID(); !ok {
		return &ValidationError{Name: "request_id", err: errors.New(`ent: missing required field "EquipmentFirmwareEffect.request_id"`)}
	}
	if _, ok := efec.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "EquipmentFirmwareEffect.state"`)}
	}
	return nil
}

func (efec *EquipmentFirmwareEffectCreate) sqlSave(ctx context.Context) (*EquipmentFirmwareEffect, error) {
	_node, _spec := efec.createSpec()
	if err := sqlgraph.CreateNode(ctx, efec.driver, _spec); err != nil {
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

func (efec *EquipmentFirmwareEffectCreate) createSpec() (*EquipmentFirmwareEffect, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentFirmwareEffect{config: efec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: equipmentfirmwareeffect.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentfirmwareeffect.FieldID,
			},
		}
	)
	if id, ok := efec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := efec.mutation.Version(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := efec.mutation.CreatedBy(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldCreatedBy, field.TypeUint64, value)
		_node.CreatedBy = value
	}
	if value, ok := efec.mutation.CreatedAt(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := efec.mutation.UpdatedBy(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldUpdatedBy, field.TypeUint64, value)
		_node.UpdatedBy = value
	}
	if value, ok := efec.mutation.UpdatedAt(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := efec.mutation.RequestID(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldRequestID, field.TypeInt64, value)
		_node.RequestID = value
	}
	if value, ok := efec.mutation.State(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldState, field.TypeInt, value)
		_node.State = value
	}
	if nodes := efec.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentfirmwareeffect.EquipmentTable,
			Columns: []string{equipmentfirmwareeffect.EquipmentColumn},
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
		_node.equipment_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := efec.mutation.FirmwareIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   equipmentfirmwareeffect.FirmwareTable,
			Columns: []string{equipmentfirmwareeffect.FirmwareColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: firmware.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.firmware_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EquipmentFirmwareEffectCreateBulk is the builder for creating many EquipmentFirmwareEffect entities in bulk.
type EquipmentFirmwareEffectCreateBulk struct {
	config
	builders []*EquipmentFirmwareEffectCreate
}

// Save creates the EquipmentFirmwareEffect entities in the database.
func (efecb *EquipmentFirmwareEffectCreateBulk) Save(ctx context.Context) ([]*EquipmentFirmwareEffect, error) {
	specs := make([]*sqlgraph.CreateSpec, len(efecb.builders))
	nodes := make([]*EquipmentFirmwareEffect, len(efecb.builders))
	mutators := make([]Mutator, len(efecb.builders))
	for i := range efecb.builders {
		func(i int, root context.Context) {
			builder := efecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentFirmwareEffectMutation)
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
					_, err = mutators[i+1].Mutate(root, efecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, efecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, efecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (efecb *EquipmentFirmwareEffectCreateBulk) SaveX(ctx context.Context) []*EquipmentFirmwareEffect {
	v, err := efecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (efecb *EquipmentFirmwareEffectCreateBulk) Exec(ctx context.Context) error {
	_, err := efecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efecb *EquipmentFirmwareEffectCreateBulk) ExecX(ctx context.Context) {
	if err := efecb.Exec(ctx); err != nil {
		panic(err)
	}
}
