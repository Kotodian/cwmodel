// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentfirmwareeffect"
	"github.com/Kotodian/ent-practice/ent/firmware"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentFirmwareEffectUpdate is the builder for updating EquipmentFirmwareEffect entities.
type EquipmentFirmwareEffectUpdate struct {
	config
	hooks    []Hook
	mutation *EquipmentFirmwareEffectMutation
}

// Where appends a list predicates to the EquipmentFirmwareEffectUpdate builder.
func (efeu *EquipmentFirmwareEffectUpdate) Where(ps ...predicate.EquipmentFirmwareEffect) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.Where(ps...)
	return efeu
}

// SetVersion sets the "version" field.
func (efeu *EquipmentFirmwareEffectUpdate) SetVersion(i int64) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.ResetVersion()
	efeu.mutation.SetVersion(i)
	return efeu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (efeu *EquipmentFirmwareEffectUpdate) SetNillableVersion(i *int64) *EquipmentFirmwareEffectUpdate {
	if i != nil {
		efeu.SetVersion(*i)
	}
	return efeu
}

// AddVersion adds i to the "version" field.
func (efeu *EquipmentFirmwareEffectUpdate) AddVersion(i int64) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.AddVersion(i)
	return efeu
}

// SetUpdatedBy sets the "updated_by" field.
func (efeu *EquipmentFirmwareEffectUpdate) SetUpdatedBy(d datasource.UUID) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.ResetUpdatedBy()
	efeu.mutation.SetUpdatedBy(d)
	return efeu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (efeu *EquipmentFirmwareEffectUpdate) SetNillableUpdatedBy(d *datasource.UUID) *EquipmentFirmwareEffectUpdate {
	if d != nil {
		efeu.SetUpdatedBy(*d)
	}
	return efeu
}

// AddUpdatedBy adds d to the "updated_by" field.
func (efeu *EquipmentFirmwareEffectUpdate) AddUpdatedBy(d datasource.UUID) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.AddUpdatedBy(d)
	return efeu
}

// SetUpdatedAt sets the "updated_at" field.
func (efeu *EquipmentFirmwareEffectUpdate) SetUpdatedAt(i int64) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.ResetUpdatedAt()
	efeu.mutation.SetUpdatedAt(i)
	return efeu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (efeu *EquipmentFirmwareEffectUpdate) SetNillableUpdatedAt(i *int64) *EquipmentFirmwareEffectUpdate {
	if i != nil {
		efeu.SetUpdatedAt(*i)
	}
	return efeu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (efeu *EquipmentFirmwareEffectUpdate) AddUpdatedAt(i int64) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.AddUpdatedAt(i)
	return efeu
}

// SetRequestID sets the "request_id" field.
func (efeu *EquipmentFirmwareEffectUpdate) SetRequestID(i int64) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.ResetRequestID()
	efeu.mutation.SetRequestID(i)
	return efeu
}

// AddRequestID adds i to the "request_id" field.
func (efeu *EquipmentFirmwareEffectUpdate) AddRequestID(i int64) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.AddRequestID(i)
	return efeu
}

// SetState sets the "state" field.
func (efeu *EquipmentFirmwareEffectUpdate) SetState(i int) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.ResetState()
	efeu.mutation.SetState(i)
	return efeu
}

// AddState adds i to the "state" field.
func (efeu *EquipmentFirmwareEffectUpdate) AddState(i int) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.AddState(i)
	return efeu
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (efeu *EquipmentFirmwareEffectUpdate) SetEquipmentID(id datasource.UUID) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.SetEquipmentID(id)
	return efeu
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (efeu *EquipmentFirmwareEffectUpdate) SetNillableEquipmentID(id *datasource.UUID) *EquipmentFirmwareEffectUpdate {
	if id != nil {
		efeu = efeu.SetEquipmentID(*id)
	}
	return efeu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (efeu *EquipmentFirmwareEffectUpdate) SetEquipment(e *Equipment) *EquipmentFirmwareEffectUpdate {
	return efeu.SetEquipmentID(e.ID)
}

// SetFirmwareID sets the "firmware" edge to the Firmware entity by ID.
func (efeu *EquipmentFirmwareEffectUpdate) SetFirmwareID(id datasource.UUID) *EquipmentFirmwareEffectUpdate {
	efeu.mutation.SetFirmwareID(id)
	return efeu
}

// SetNillableFirmwareID sets the "firmware" edge to the Firmware entity by ID if the given value is not nil.
func (efeu *EquipmentFirmwareEffectUpdate) SetNillableFirmwareID(id *datasource.UUID) *EquipmentFirmwareEffectUpdate {
	if id != nil {
		efeu = efeu.SetFirmwareID(*id)
	}
	return efeu
}

// SetFirmware sets the "firmware" edge to the Firmware entity.
func (efeu *EquipmentFirmwareEffectUpdate) SetFirmware(f *Firmware) *EquipmentFirmwareEffectUpdate {
	return efeu.SetFirmwareID(f.ID)
}

// Mutation returns the EquipmentFirmwareEffectMutation object of the builder.
func (efeu *EquipmentFirmwareEffectUpdate) Mutation() *EquipmentFirmwareEffectMutation {
	return efeu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (efeu *EquipmentFirmwareEffectUpdate) ClearEquipment() *EquipmentFirmwareEffectUpdate {
	efeu.mutation.ClearEquipment()
	return efeu
}

// ClearFirmware clears the "firmware" edge to the Firmware entity.
func (efeu *EquipmentFirmwareEffectUpdate) ClearFirmware() *EquipmentFirmwareEffectUpdate {
	efeu.mutation.ClearFirmware()
	return efeu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (efeu *EquipmentFirmwareEffectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(efeu.hooks) == 0 {
		affected, err = efeu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentFirmwareEffectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			efeu.mutation = mutation
			affected, err = efeu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(efeu.hooks) - 1; i >= 0; i-- {
			if efeu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = efeu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, efeu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (efeu *EquipmentFirmwareEffectUpdate) SaveX(ctx context.Context) int {
	affected, err := efeu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (efeu *EquipmentFirmwareEffectUpdate) Exec(ctx context.Context) error {
	_, err := efeu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efeu *EquipmentFirmwareEffectUpdate) ExecX(ctx context.Context) {
	if err := efeu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (efeu *EquipmentFirmwareEffectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentfirmwareeffect.Table,
			Columns: equipmentfirmwareeffect.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentfirmwareeffect.FieldID,
			},
		},
	}
	if ps := efeu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := efeu.mutation.Version(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := efeu.mutation.AddedVersion(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := efeu.mutation.UpdatedBy(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := efeu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := efeu.mutation.UpdatedAt(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := efeu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := efeu.mutation.RequestID(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldRequestID, field.TypeInt64, value)
	}
	if value, ok := efeu.mutation.AddedRequestID(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldRequestID, field.TypeInt64, value)
	}
	if value, ok := efeu.mutation.State(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldState, field.TypeInt, value)
	}
	if value, ok := efeu.mutation.AddedState(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldState, field.TypeInt, value)
	}
	if efeu.mutation.EquipmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := efeu.mutation.EquipmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if efeu.mutation.FirmwareCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := efeu.mutation.FirmwareIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, efeu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentfirmwareeffect.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EquipmentFirmwareEffectUpdateOne is the builder for updating a single EquipmentFirmwareEffect entity.
type EquipmentFirmwareEffectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EquipmentFirmwareEffectMutation
}

// SetVersion sets the "version" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetVersion(i int64) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.ResetVersion()
	efeuo.mutation.SetVersion(i)
	return efeuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetNillableVersion(i *int64) *EquipmentFirmwareEffectUpdateOne {
	if i != nil {
		efeuo.SetVersion(*i)
	}
	return efeuo
}

// AddVersion adds i to the "version" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) AddVersion(i int64) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.AddVersion(i)
	return efeuo
}

// SetUpdatedBy sets the "updated_by" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetUpdatedBy(d datasource.UUID) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.ResetUpdatedBy()
	efeuo.mutation.SetUpdatedBy(d)
	return efeuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetNillableUpdatedBy(d *datasource.UUID) *EquipmentFirmwareEffectUpdateOne {
	if d != nil {
		efeuo.SetUpdatedBy(*d)
	}
	return efeuo
}

// AddUpdatedBy adds d to the "updated_by" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) AddUpdatedBy(d datasource.UUID) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.AddUpdatedBy(d)
	return efeuo
}

// SetUpdatedAt sets the "updated_at" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetUpdatedAt(i int64) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.ResetUpdatedAt()
	efeuo.mutation.SetUpdatedAt(i)
	return efeuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetNillableUpdatedAt(i *int64) *EquipmentFirmwareEffectUpdateOne {
	if i != nil {
		efeuo.SetUpdatedAt(*i)
	}
	return efeuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) AddUpdatedAt(i int64) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.AddUpdatedAt(i)
	return efeuo
}

// SetRequestID sets the "request_id" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetRequestID(i int64) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.ResetRequestID()
	efeuo.mutation.SetRequestID(i)
	return efeuo
}

// AddRequestID adds i to the "request_id" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) AddRequestID(i int64) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.AddRequestID(i)
	return efeuo
}

// SetState sets the "state" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetState(i int) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.ResetState()
	efeuo.mutation.SetState(i)
	return efeuo
}

// AddState adds i to the "state" field.
func (efeuo *EquipmentFirmwareEffectUpdateOne) AddState(i int) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.AddState(i)
	return efeuo
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetEquipmentID(id datasource.UUID) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.SetEquipmentID(id)
	return efeuo
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetNillableEquipmentID(id *datasource.UUID) *EquipmentFirmwareEffectUpdateOne {
	if id != nil {
		efeuo = efeuo.SetEquipmentID(*id)
	}
	return efeuo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetEquipment(e *Equipment) *EquipmentFirmwareEffectUpdateOne {
	return efeuo.SetEquipmentID(e.ID)
}

// SetFirmwareID sets the "firmware" edge to the Firmware entity by ID.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetFirmwareID(id datasource.UUID) *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.SetFirmwareID(id)
	return efeuo
}

// SetNillableFirmwareID sets the "firmware" edge to the Firmware entity by ID if the given value is not nil.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetNillableFirmwareID(id *datasource.UUID) *EquipmentFirmwareEffectUpdateOne {
	if id != nil {
		efeuo = efeuo.SetFirmwareID(*id)
	}
	return efeuo
}

// SetFirmware sets the "firmware" edge to the Firmware entity.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SetFirmware(f *Firmware) *EquipmentFirmwareEffectUpdateOne {
	return efeuo.SetFirmwareID(f.ID)
}

// Mutation returns the EquipmentFirmwareEffectMutation object of the builder.
func (efeuo *EquipmentFirmwareEffectUpdateOne) Mutation() *EquipmentFirmwareEffectMutation {
	return efeuo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (efeuo *EquipmentFirmwareEffectUpdateOne) ClearEquipment() *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.ClearEquipment()
	return efeuo
}

// ClearFirmware clears the "firmware" edge to the Firmware entity.
func (efeuo *EquipmentFirmwareEffectUpdateOne) ClearFirmware() *EquipmentFirmwareEffectUpdateOne {
	efeuo.mutation.ClearFirmware()
	return efeuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (efeuo *EquipmentFirmwareEffectUpdateOne) Select(field string, fields ...string) *EquipmentFirmwareEffectUpdateOne {
	efeuo.fields = append([]string{field}, fields...)
	return efeuo
}

// Save executes the query and returns the updated EquipmentFirmwareEffect entity.
func (efeuo *EquipmentFirmwareEffectUpdateOne) Save(ctx context.Context) (*EquipmentFirmwareEffect, error) {
	var (
		err  error
		node *EquipmentFirmwareEffect
	)
	if len(efeuo.hooks) == 0 {
		node, err = efeuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentFirmwareEffectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			efeuo.mutation = mutation
			node, err = efeuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(efeuo.hooks) - 1; i >= 0; i-- {
			if efeuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = efeuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, efeuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (efeuo *EquipmentFirmwareEffectUpdateOne) SaveX(ctx context.Context) *EquipmentFirmwareEffect {
	node, err := efeuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (efeuo *EquipmentFirmwareEffectUpdateOne) Exec(ctx context.Context) error {
	_, err := efeuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (efeuo *EquipmentFirmwareEffectUpdateOne) ExecX(ctx context.Context) {
	if err := efeuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (efeuo *EquipmentFirmwareEffectUpdateOne) sqlSave(ctx context.Context) (_node *EquipmentFirmwareEffect, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentfirmwareeffect.Table,
			Columns: equipmentfirmwareeffect.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentfirmwareeffect.FieldID,
			},
		},
	}
	id, ok := efeuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EquipmentFirmwareEffect.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := efeuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentfirmwareeffect.FieldID)
		for _, f := range fields {
			if !equipmentfirmwareeffect.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != equipmentfirmwareeffect.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := efeuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := efeuo.mutation.Version(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := efeuo.mutation.AddedVersion(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := efeuo.mutation.UpdatedBy(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := efeuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := efeuo.mutation.UpdatedAt(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := efeuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := efeuo.mutation.RequestID(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldRequestID, field.TypeInt64, value)
	}
	if value, ok := efeuo.mutation.AddedRequestID(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldRequestID, field.TypeInt64, value)
	}
	if value, ok := efeuo.mutation.State(); ok {
		_spec.SetField(equipmentfirmwareeffect.FieldState, field.TypeInt, value)
	}
	if value, ok := efeuo.mutation.AddedState(); ok {
		_spec.AddField(equipmentfirmwareeffect.FieldState, field.TypeInt, value)
	}
	if efeuo.mutation.EquipmentCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := efeuo.mutation.EquipmentIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if efeuo.mutation.FirmwareCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := efeuo.mutation.FirmwareIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EquipmentFirmwareEffect{config: efeuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, efeuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentfirmwareeffect.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
