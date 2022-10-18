// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentiot"
	"github.com/Kotodian/ent-practice/ent/predicate"
)

// EquipmentIotUpdate is the builder for updating EquipmentIot entities.
type EquipmentIotUpdate struct {
	config
	hooks    []Hook
	mutation *EquipmentIotMutation
}

// Where appends a list predicates to the EquipmentIotUpdate builder.
func (eiu *EquipmentIotUpdate) Where(ps ...predicate.EquipmentIot) *EquipmentIotUpdate {
	eiu.mutation.Where(ps...)
	return eiu
}

// SetIccid sets the "iccid" field.
func (eiu *EquipmentIotUpdate) SetIccid(s string) *EquipmentIotUpdate {
	eiu.mutation.SetIccid(s)
	return eiu
}

// SetNillableIccid sets the "iccid" field if the given value is not nil.
func (eiu *EquipmentIotUpdate) SetNillableIccid(s *string) *EquipmentIotUpdate {
	if s != nil {
		eiu.SetIccid(*s)
	}
	return eiu
}

// ClearIccid clears the value of the "iccid" field.
func (eiu *EquipmentIotUpdate) ClearIccid() *EquipmentIotUpdate {
	eiu.mutation.ClearIccid()
	return eiu
}

// SetImei sets the "imei" field.
func (eiu *EquipmentIotUpdate) SetImei(s string) *EquipmentIotUpdate {
	eiu.mutation.SetImei(s)
	return eiu
}

// SetNillableImei sets the "imei" field if the given value is not nil.
func (eiu *EquipmentIotUpdate) SetNillableImei(s *string) *EquipmentIotUpdate {
	if s != nil {
		eiu.SetImei(*s)
	}
	return eiu
}

// ClearImei clears the value of the "imei" field.
func (eiu *EquipmentIotUpdate) ClearImei() *EquipmentIotUpdate {
	eiu.mutation.ClearImei()
	return eiu
}

// SetRemoteAddress sets the "remote_address" field.
func (eiu *EquipmentIotUpdate) SetRemoteAddress(s string) *EquipmentIotUpdate {
	eiu.mutation.SetRemoteAddress(s)
	return eiu
}

// SetNillableRemoteAddress sets the "remote_address" field if the given value is not nil.
func (eiu *EquipmentIotUpdate) SetNillableRemoteAddress(s *string) *EquipmentIotUpdate {
	if s != nil {
		eiu.SetRemoteAddress(*s)
	}
	return eiu
}

// ClearRemoteAddress clears the value of the "remote_address" field.
func (eiu *EquipmentIotUpdate) ClearRemoteAddress() *EquipmentIotUpdate {
	eiu.mutation.ClearRemoteAddress()
	return eiu
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eiu *EquipmentIotUpdate) SetEquipmentID(id int) *EquipmentIotUpdate {
	eiu.mutation.SetEquipmentID(id)
	return eiu
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (eiu *EquipmentIotUpdate) SetNillableEquipmentID(id *int) *EquipmentIotUpdate {
	if id != nil {
		eiu = eiu.SetEquipmentID(*id)
	}
	return eiu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eiu *EquipmentIotUpdate) SetEquipment(e *Equipment) *EquipmentIotUpdate {
	return eiu.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentIotMutation object of the builder.
func (eiu *EquipmentIotUpdate) Mutation() *EquipmentIotMutation {
	return eiu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (eiu *EquipmentIotUpdate) ClearEquipment() *EquipmentIotUpdate {
	eiu.mutation.ClearEquipment()
	return eiu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eiu *EquipmentIotUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eiu.hooks) == 0 {
		affected, err = eiu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentIotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eiu.mutation = mutation
			affected, err = eiu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eiu.hooks) - 1; i >= 0; i-- {
			if eiu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eiu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eiu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eiu *EquipmentIotUpdate) SaveX(ctx context.Context) int {
	affected, err := eiu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eiu *EquipmentIotUpdate) Exec(ctx context.Context) error {
	_, err := eiu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiu *EquipmentIotUpdate) ExecX(ctx context.Context) {
	if err := eiu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eiu *EquipmentIotUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentiot.Table,
			Columns: equipmentiot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentiot.FieldID,
			},
		},
	}
	if ps := eiu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiu.mutation.Iccid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldIccid,
		})
	}
	if eiu.mutation.IccidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: equipmentiot.FieldIccid,
		})
	}
	if value, ok := eiu.mutation.Imei(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldImei,
		})
	}
	if eiu.mutation.ImeiCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: equipmentiot.FieldImei,
		})
	}
	if value, ok := eiu.mutation.RemoteAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldRemoteAddress,
		})
	}
	if eiu.mutation.RemoteAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: equipmentiot.FieldRemoteAddress,
		})
	}
	if eiu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentiot.EquipmentTable,
			Columns: []string{equipmentiot.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentiot.EquipmentTable,
			Columns: []string{equipmentiot.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eiu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentiot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EquipmentIotUpdateOne is the builder for updating a single EquipmentIot entity.
type EquipmentIotUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EquipmentIotMutation
}

// SetIccid sets the "iccid" field.
func (eiuo *EquipmentIotUpdateOne) SetIccid(s string) *EquipmentIotUpdateOne {
	eiuo.mutation.SetIccid(s)
	return eiuo
}

// SetNillableIccid sets the "iccid" field if the given value is not nil.
func (eiuo *EquipmentIotUpdateOne) SetNillableIccid(s *string) *EquipmentIotUpdateOne {
	if s != nil {
		eiuo.SetIccid(*s)
	}
	return eiuo
}

// ClearIccid clears the value of the "iccid" field.
func (eiuo *EquipmentIotUpdateOne) ClearIccid() *EquipmentIotUpdateOne {
	eiuo.mutation.ClearIccid()
	return eiuo
}

// SetImei sets the "imei" field.
func (eiuo *EquipmentIotUpdateOne) SetImei(s string) *EquipmentIotUpdateOne {
	eiuo.mutation.SetImei(s)
	return eiuo
}

// SetNillableImei sets the "imei" field if the given value is not nil.
func (eiuo *EquipmentIotUpdateOne) SetNillableImei(s *string) *EquipmentIotUpdateOne {
	if s != nil {
		eiuo.SetImei(*s)
	}
	return eiuo
}

// ClearImei clears the value of the "imei" field.
func (eiuo *EquipmentIotUpdateOne) ClearImei() *EquipmentIotUpdateOne {
	eiuo.mutation.ClearImei()
	return eiuo
}

// SetRemoteAddress sets the "remote_address" field.
func (eiuo *EquipmentIotUpdateOne) SetRemoteAddress(s string) *EquipmentIotUpdateOne {
	eiuo.mutation.SetRemoteAddress(s)
	return eiuo
}

// SetNillableRemoteAddress sets the "remote_address" field if the given value is not nil.
func (eiuo *EquipmentIotUpdateOne) SetNillableRemoteAddress(s *string) *EquipmentIotUpdateOne {
	if s != nil {
		eiuo.SetRemoteAddress(*s)
	}
	return eiuo
}

// ClearRemoteAddress clears the value of the "remote_address" field.
func (eiuo *EquipmentIotUpdateOne) ClearRemoteAddress() *EquipmentIotUpdateOne {
	eiuo.mutation.ClearRemoteAddress()
	return eiuo
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eiuo *EquipmentIotUpdateOne) SetEquipmentID(id int) *EquipmentIotUpdateOne {
	eiuo.mutation.SetEquipmentID(id)
	return eiuo
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (eiuo *EquipmentIotUpdateOne) SetNillableEquipmentID(id *int) *EquipmentIotUpdateOne {
	if id != nil {
		eiuo = eiuo.SetEquipmentID(*id)
	}
	return eiuo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eiuo *EquipmentIotUpdateOne) SetEquipment(e *Equipment) *EquipmentIotUpdateOne {
	return eiuo.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentIotMutation object of the builder.
func (eiuo *EquipmentIotUpdateOne) Mutation() *EquipmentIotMutation {
	return eiuo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (eiuo *EquipmentIotUpdateOne) ClearEquipment() *EquipmentIotUpdateOne {
	eiuo.mutation.ClearEquipment()
	return eiuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (eiuo *EquipmentIotUpdateOne) Select(field string, fields ...string) *EquipmentIotUpdateOne {
	eiuo.fields = append([]string{field}, fields...)
	return eiuo
}

// Save executes the query and returns the updated EquipmentIot entity.
func (eiuo *EquipmentIotUpdateOne) Save(ctx context.Context) (*EquipmentIot, error) {
	var (
		err  error
		node *EquipmentIot
	)
	if len(eiuo.hooks) == 0 {
		node, err = eiuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentIotMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eiuo.mutation = mutation
			node, err = eiuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(eiuo.hooks) - 1; i >= 0; i-- {
			if eiuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eiuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eiuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (eiuo *EquipmentIotUpdateOne) SaveX(ctx context.Context) *EquipmentIot {
	node, err := eiuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (eiuo *EquipmentIotUpdateOne) Exec(ctx context.Context) error {
	_, err := eiuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eiuo *EquipmentIotUpdateOne) ExecX(ctx context.Context) {
	if err := eiuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eiuo *EquipmentIotUpdateOne) sqlSave(ctx context.Context) (_node *EquipmentIot, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentiot.Table,
			Columns: equipmentiot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentiot.FieldID,
			},
		},
	}
	id, ok := eiuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "EquipmentIot.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := eiuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentiot.FieldID)
		for _, f := range fields {
			if !equipmentiot.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != equipmentiot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := eiuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eiuo.mutation.Iccid(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldIccid,
		})
	}
	if eiuo.mutation.IccidCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: equipmentiot.FieldIccid,
		})
	}
	if value, ok := eiuo.mutation.Imei(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldImei,
		})
	}
	if eiuo.mutation.ImeiCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: equipmentiot.FieldImei,
		})
	}
	if value, ok := eiuo.mutation.RemoteAddress(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentiot.FieldRemoteAddress,
		})
	}
	if eiuo.mutation.RemoteAddressCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: equipmentiot.FieldRemoteAddress,
		})
	}
	if eiuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentiot.EquipmentTable,
			Columns: []string{equipmentiot.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eiuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentiot.EquipmentTable,
			Columns: []string{equipmentiot.EquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: equipment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &EquipmentIot{config: eiuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, eiuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{equipmentiot.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
