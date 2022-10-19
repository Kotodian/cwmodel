// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/equipmentfirmwareeffect"
	"github.com/Kotodian/cwmodel/firmware"
	"github.com/Kotodian/gokit/datasource"
)

// FirmwareCreate is the builder for creating a Firmware entity.
type FirmwareCreate struct {
	config
	mutation *FirmwareMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (fc *FirmwareCreate) SetVersion(i int64) *FirmwareCreate {
	fc.mutation.SetVersion(i)
	return fc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (fc *FirmwareCreate) SetNillableVersion(i *int64) *FirmwareCreate {
	if i != nil {
		fc.SetVersion(*i)
	}
	return fc
}

// SetCreatedBy sets the "created_by" field.
func (fc *FirmwareCreate) SetCreatedBy(d datasource.UUID) *FirmwareCreate {
	fc.mutation.SetCreatedBy(d)
	return fc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (fc *FirmwareCreate) SetNillableCreatedBy(d *datasource.UUID) *FirmwareCreate {
	if d != nil {
		fc.SetCreatedBy(*d)
	}
	return fc
}

// SetCreatedAt sets the "created_at" field.
func (fc *FirmwareCreate) SetCreatedAt(i int64) *FirmwareCreate {
	fc.mutation.SetCreatedAt(i)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FirmwareCreate) SetNillableCreatedAt(i *int64) *FirmwareCreate {
	if i != nil {
		fc.SetCreatedAt(*i)
	}
	return fc
}

// SetUpdatedBy sets the "updated_by" field.
func (fc *FirmwareCreate) SetUpdatedBy(d datasource.UUID) *FirmwareCreate {
	fc.mutation.SetUpdatedBy(d)
	return fc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fc *FirmwareCreate) SetNillableUpdatedBy(d *datasource.UUID) *FirmwareCreate {
	if d != nil {
		fc.SetUpdatedBy(*d)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FirmwareCreate) SetUpdatedAt(i int64) *FirmwareCreate {
	fc.mutation.SetUpdatedAt(i)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FirmwareCreate) SetNillableUpdatedAt(i *int64) *FirmwareCreate {
	if i != nil {
		fc.SetUpdatedAt(*i)
	}
	return fc
}

// SetEquipVersion sets the "equip_version" field.
func (fc *FirmwareCreate) SetEquipVersion(s string) *FirmwareCreate {
	fc.mutation.SetEquipVersion(s)
	return fc
}

// SetID sets the "id" field.
func (fc *FirmwareCreate) SetID(d datasource.UUID) *FirmwareCreate {
	fc.mutation.SetID(d)
	return fc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (fc *FirmwareCreate) SetNillableID(d *datasource.UUID) *FirmwareCreate {
	if d != nil {
		fc.SetID(*d)
	}
	return fc
}

// AddEquipmentFirmwareEffectIDs adds the "equipment_firmware_effect" edge to the EquipmentFirmwareEffect entity by IDs.
func (fc *FirmwareCreate) AddEquipmentFirmwareEffectIDs(ids ...datasource.UUID) *FirmwareCreate {
	fc.mutation.AddEquipmentFirmwareEffectIDs(ids...)
	return fc
}

// AddEquipmentFirmwareEffect adds the "equipment_firmware_effect" edges to the EquipmentFirmwareEffect entity.
func (fc *FirmwareCreate) AddEquipmentFirmwareEffect(e ...*EquipmentFirmwareEffect) *FirmwareCreate {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fc.AddEquipmentFirmwareEffectIDs(ids...)
}

// Mutation returns the FirmwareMutation object of the builder.
func (fc *FirmwareCreate) Mutation() *FirmwareMutation {
	return fc.mutation
}

// Save creates the Firmware in the database.
func (fc *FirmwareCreate) Save(ctx context.Context) (*Firmware, error) {
	var (
		err  error
		node *Firmware
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FirmwareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("cwmodel: uninitialized hook (forgotten import cwmodel/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Firmware)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FirmwareMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FirmwareCreate) SaveX(ctx context.Context) *Firmware {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FirmwareCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FirmwareCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FirmwareCreate) defaults() {
	if _, ok := fc.mutation.Version(); !ok {
		v := firmware.DefaultVersion
		fc.mutation.SetVersion(v)
	}
	if _, ok := fc.mutation.CreatedBy(); !ok {
		v := firmware.DefaultCreatedBy
		fc.mutation.SetCreatedBy(v)
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := firmware.DefaultCreatedAt
		fc.mutation.SetCreatedAt(v)
	}
	if _, ok := fc.mutation.UpdatedBy(); !ok {
		v := firmware.DefaultUpdatedBy
		fc.mutation.SetUpdatedBy(v)
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		v := firmware.DefaultUpdatedAt
		fc.mutation.SetUpdatedAt(v)
	}
	if _, ok := fc.mutation.ID(); !ok {
		v := firmware.DefaultID
		fc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FirmwareCreate) check() error {
	if _, ok := fc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`cwmodel: missing required field "Firmware.version"`)}
	}
	if _, ok := fc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cwmodel: missing required field "Firmware.created_by"`)}
	}
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cwmodel: missing required field "Firmware.created_at"`)}
	}
	if _, ok := fc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cwmodel: missing required field "Firmware.updated_by"`)}
	}
	if _, ok := fc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cwmodel: missing required field "Firmware.updated_at"`)}
	}
	if _, ok := fc.mutation.EquipVersion(); !ok {
		return &ValidationError{Name: "equip_version", err: errors.New(`cwmodel: missing required field "Firmware.equip_version"`)}
	}
	return nil
}

func (fc *FirmwareCreate) sqlSave(ctx context.Context) (*Firmware, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
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

func (fc *FirmwareCreate) createSpec() (*Firmware, *sqlgraph.CreateSpec) {
	var (
		_node = &Firmware{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: firmware.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: firmware.FieldID,
			},
		}
	)
	if id, ok := fc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := fc.mutation.Version(); ok {
		_spec.SetField(firmware.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := fc.mutation.CreatedBy(); ok {
		_spec.SetField(firmware.FieldCreatedBy, field.TypeUint64, value)
		_node.CreatedBy = value
	}
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.SetField(firmware.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedBy(); ok {
		_spec.SetField(firmware.FieldUpdatedBy, field.TypeUint64, value)
		_node.UpdatedBy = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.SetField(firmware.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.EquipVersion(); ok {
		_spec.SetField(firmware.FieldEquipVersion, field.TypeString, value)
		_node.EquipVersion = value
	}
	if nodes := fc.mutation.EquipmentFirmwareEffectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   firmware.EquipmentFirmwareEffectTable,
			Columns: []string{firmware.EquipmentFirmwareEffectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: equipmentfirmwareeffect.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FirmwareCreateBulk is the builder for creating many Firmware entities in bulk.
type FirmwareCreateBulk struct {
	config
	builders []*FirmwareCreate
}

// Save creates the Firmware entities in the database.
func (fcb *FirmwareCreateBulk) Save(ctx context.Context) ([]*Firmware, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Firmware, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FirmwareMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FirmwareCreateBulk) SaveX(ctx context.Context) []*Firmware {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FirmwareCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FirmwareCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}