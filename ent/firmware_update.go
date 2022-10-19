// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipmentfirmwareeffect"
	"github.com/Kotodian/ent-practice/ent/firmware"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// FirmwareUpdate is the builder for updating Firmware entities.
type FirmwareUpdate struct {
	config
	hooks    []Hook
	mutation *FirmwareMutation
}

// Where appends a list predicates to the FirmwareUpdate builder.
func (fu *FirmwareUpdate) Where(ps ...predicate.Firmware) *FirmwareUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetVersion sets the "version" field.
func (fu *FirmwareUpdate) SetVersion(i int64) *FirmwareUpdate {
	fu.mutation.ResetVersion()
	fu.mutation.SetVersion(i)
	return fu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (fu *FirmwareUpdate) SetNillableVersion(i *int64) *FirmwareUpdate {
	if i != nil {
		fu.SetVersion(*i)
	}
	return fu
}

// AddVersion adds i to the "version" field.
func (fu *FirmwareUpdate) AddVersion(i int64) *FirmwareUpdate {
	fu.mutation.AddVersion(i)
	return fu
}

// SetUpdatedBy sets the "updated_by" field.
func (fu *FirmwareUpdate) SetUpdatedBy(d datasource.UUID) *FirmwareUpdate {
	fu.mutation.ResetUpdatedBy()
	fu.mutation.SetUpdatedBy(d)
	return fu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fu *FirmwareUpdate) SetNillableUpdatedBy(d *datasource.UUID) *FirmwareUpdate {
	if d != nil {
		fu.SetUpdatedBy(*d)
	}
	return fu
}

// AddUpdatedBy adds d to the "updated_by" field.
func (fu *FirmwareUpdate) AddUpdatedBy(d datasource.UUID) *FirmwareUpdate {
	fu.mutation.AddUpdatedBy(d)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FirmwareUpdate) SetUpdatedAt(i int64) *FirmwareUpdate {
	fu.mutation.ResetUpdatedAt()
	fu.mutation.SetUpdatedAt(i)
	return fu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fu *FirmwareUpdate) SetNillableUpdatedAt(i *int64) *FirmwareUpdate {
	if i != nil {
		fu.SetUpdatedAt(*i)
	}
	return fu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (fu *FirmwareUpdate) AddUpdatedAt(i int64) *FirmwareUpdate {
	fu.mutation.AddUpdatedAt(i)
	return fu
}

// SetEquipVersion sets the "equip_version" field.
func (fu *FirmwareUpdate) SetEquipVersion(s string) *FirmwareUpdate {
	fu.mutation.SetEquipVersion(s)
	return fu
}

// AddEquipmentFirmwareEffectIDs adds the "equipment_firmware_effect" edge to the EquipmentFirmwareEffect entity by IDs.
func (fu *FirmwareUpdate) AddEquipmentFirmwareEffectIDs(ids ...datasource.UUID) *FirmwareUpdate {
	fu.mutation.AddEquipmentFirmwareEffectIDs(ids...)
	return fu
}

// AddEquipmentFirmwareEffect adds the "equipment_firmware_effect" edges to the EquipmentFirmwareEffect entity.
func (fu *FirmwareUpdate) AddEquipmentFirmwareEffect(e ...*EquipmentFirmwareEffect) *FirmwareUpdate {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fu.AddEquipmentFirmwareEffectIDs(ids...)
}

// Mutation returns the FirmwareMutation object of the builder.
func (fu *FirmwareUpdate) Mutation() *FirmwareMutation {
	return fu.mutation
}

// ClearEquipmentFirmwareEffect clears all "equipment_firmware_effect" edges to the EquipmentFirmwareEffect entity.
func (fu *FirmwareUpdate) ClearEquipmentFirmwareEffect() *FirmwareUpdate {
	fu.mutation.ClearEquipmentFirmwareEffect()
	return fu
}

// RemoveEquipmentFirmwareEffectIDs removes the "equipment_firmware_effect" edge to EquipmentFirmwareEffect entities by IDs.
func (fu *FirmwareUpdate) RemoveEquipmentFirmwareEffectIDs(ids ...datasource.UUID) *FirmwareUpdate {
	fu.mutation.RemoveEquipmentFirmwareEffectIDs(ids...)
	return fu
}

// RemoveEquipmentFirmwareEffect removes "equipment_firmware_effect" edges to EquipmentFirmwareEffect entities.
func (fu *FirmwareUpdate) RemoveEquipmentFirmwareEffect(e ...*EquipmentFirmwareEffect) *FirmwareUpdate {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fu.RemoveEquipmentFirmwareEffectIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FirmwareUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fu.hooks) == 0 {
		affected, err = fu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FirmwareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fu.mutation = mutation
			affected, err = fu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fu.hooks) - 1; i >= 0; i-- {
			if fu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FirmwareUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FirmwareUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FirmwareUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fu *FirmwareUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   firmware.Table,
			Columns: firmware.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: firmware.FieldID,
			},
		},
	}
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Version(); ok {
		_spec.SetField(firmware.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedVersion(); ok {
		_spec.AddField(firmware.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.UpdatedBy(); ok {
		_spec.SetField(firmware.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := fu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(firmware.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(firmware.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(firmware.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.EquipVersion(); ok {
		_spec.SetField(firmware.FieldEquipVersion, field.TypeString, value)
	}
	if fu.mutation.EquipmentFirmwareEffectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedEquipmentFirmwareEffectIDs(); len(nodes) > 0 && !fu.mutation.EquipmentFirmwareEffectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.EquipmentFirmwareEffectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{firmware.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// FirmwareUpdateOne is the builder for updating a single Firmware entity.
type FirmwareUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FirmwareMutation
}

// SetVersion sets the "version" field.
func (fuo *FirmwareUpdateOne) SetVersion(i int64) *FirmwareUpdateOne {
	fuo.mutation.ResetVersion()
	fuo.mutation.SetVersion(i)
	return fuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (fuo *FirmwareUpdateOne) SetNillableVersion(i *int64) *FirmwareUpdateOne {
	if i != nil {
		fuo.SetVersion(*i)
	}
	return fuo
}

// AddVersion adds i to the "version" field.
func (fuo *FirmwareUpdateOne) AddVersion(i int64) *FirmwareUpdateOne {
	fuo.mutation.AddVersion(i)
	return fuo
}

// SetUpdatedBy sets the "updated_by" field.
func (fuo *FirmwareUpdateOne) SetUpdatedBy(d datasource.UUID) *FirmwareUpdateOne {
	fuo.mutation.ResetUpdatedBy()
	fuo.mutation.SetUpdatedBy(d)
	return fuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fuo *FirmwareUpdateOne) SetNillableUpdatedBy(d *datasource.UUID) *FirmwareUpdateOne {
	if d != nil {
		fuo.SetUpdatedBy(*d)
	}
	return fuo
}

// AddUpdatedBy adds d to the "updated_by" field.
func (fuo *FirmwareUpdateOne) AddUpdatedBy(d datasource.UUID) *FirmwareUpdateOne {
	fuo.mutation.AddUpdatedBy(d)
	return fuo
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FirmwareUpdateOne) SetUpdatedAt(i int64) *FirmwareUpdateOne {
	fuo.mutation.ResetUpdatedAt()
	fuo.mutation.SetUpdatedAt(i)
	return fuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fuo *FirmwareUpdateOne) SetNillableUpdatedAt(i *int64) *FirmwareUpdateOne {
	if i != nil {
		fuo.SetUpdatedAt(*i)
	}
	return fuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (fuo *FirmwareUpdateOne) AddUpdatedAt(i int64) *FirmwareUpdateOne {
	fuo.mutation.AddUpdatedAt(i)
	return fuo
}

// SetEquipVersion sets the "equip_version" field.
func (fuo *FirmwareUpdateOne) SetEquipVersion(s string) *FirmwareUpdateOne {
	fuo.mutation.SetEquipVersion(s)
	return fuo
}

// AddEquipmentFirmwareEffectIDs adds the "equipment_firmware_effect" edge to the EquipmentFirmwareEffect entity by IDs.
func (fuo *FirmwareUpdateOne) AddEquipmentFirmwareEffectIDs(ids ...datasource.UUID) *FirmwareUpdateOne {
	fuo.mutation.AddEquipmentFirmwareEffectIDs(ids...)
	return fuo
}

// AddEquipmentFirmwareEffect adds the "equipment_firmware_effect" edges to the EquipmentFirmwareEffect entity.
func (fuo *FirmwareUpdateOne) AddEquipmentFirmwareEffect(e ...*EquipmentFirmwareEffect) *FirmwareUpdateOne {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fuo.AddEquipmentFirmwareEffectIDs(ids...)
}

// Mutation returns the FirmwareMutation object of the builder.
func (fuo *FirmwareUpdateOne) Mutation() *FirmwareMutation {
	return fuo.mutation
}

// ClearEquipmentFirmwareEffect clears all "equipment_firmware_effect" edges to the EquipmentFirmwareEffect entity.
func (fuo *FirmwareUpdateOne) ClearEquipmentFirmwareEffect() *FirmwareUpdateOne {
	fuo.mutation.ClearEquipmentFirmwareEffect()
	return fuo
}

// RemoveEquipmentFirmwareEffectIDs removes the "equipment_firmware_effect" edge to EquipmentFirmwareEffect entities by IDs.
func (fuo *FirmwareUpdateOne) RemoveEquipmentFirmwareEffectIDs(ids ...datasource.UUID) *FirmwareUpdateOne {
	fuo.mutation.RemoveEquipmentFirmwareEffectIDs(ids...)
	return fuo
}

// RemoveEquipmentFirmwareEffect removes "equipment_firmware_effect" edges to EquipmentFirmwareEffect entities.
func (fuo *FirmwareUpdateOne) RemoveEquipmentFirmwareEffect(e ...*EquipmentFirmwareEffect) *FirmwareUpdateOne {
	ids := make([]datasource.UUID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return fuo.RemoveEquipmentFirmwareEffectIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FirmwareUpdateOne) Select(field string, fields ...string) *FirmwareUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Firmware entity.
func (fuo *FirmwareUpdateOne) Save(ctx context.Context) (*Firmware, error) {
	var (
		err  error
		node *Firmware
	)
	if len(fuo.hooks) == 0 {
		node, err = fuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FirmwareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fuo.mutation = mutation
			node, err = fuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(fuo.hooks) - 1; i >= 0; i-- {
			if fuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (fuo *FirmwareUpdateOne) SaveX(ctx context.Context) *Firmware {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FirmwareUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FirmwareUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fuo *FirmwareUpdateOne) sqlSave(ctx context.Context) (_node *Firmware, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   firmware.Table,
			Columns: firmware.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: firmware.FieldID,
			},
		},
	}
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Firmware.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, firmware.FieldID)
		for _, f := range fields {
			if !firmware.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != firmware.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Version(); ok {
		_spec.SetField(firmware.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedVersion(); ok {
		_spec.AddField(firmware.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.UpdatedBy(); ok {
		_spec.SetField(firmware.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := fuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(firmware.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(firmware.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(firmware.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.EquipVersion(); ok {
		_spec.SetField(firmware.FieldEquipVersion, field.TypeString, value)
	}
	if fuo.mutation.EquipmentFirmwareEffectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedEquipmentFirmwareEffectIDs(); len(nodes) > 0 && !fuo.mutation.EquipmentFirmwareEffectCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.EquipmentFirmwareEffectIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Firmware{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{firmware.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
