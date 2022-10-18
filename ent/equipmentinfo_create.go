// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentinfo"
)

// EquipmentInfoCreate is the builder for creating a EquipmentInfo entity.
type EquipmentInfoCreate struct {
	config
	mutation *EquipmentInfoMutation
	hooks    []Hook
}

// SetEquipmentSn sets the "equipment_sn" field.
func (eic *EquipmentInfoCreate) SetEquipmentSn(s string) *EquipmentInfoCreate {
	eic.mutation.SetEquipmentSn(s)
	return eic
}

// SetAccessPod sets the "access_pod" field.
func (eic *EquipmentInfoCreate) SetAccessPod(s string) *EquipmentInfoCreate {
	eic.mutation.SetAccessPod(s)
	return eic
}

// SetState sets the "state" field.
func (eic *EquipmentInfoCreate) SetState(b bool) *EquipmentInfoCreate {
	eic.mutation.SetState(b)
	return eic
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eic *EquipmentInfoCreate) SetEquipmentID(id int) *EquipmentInfoCreate {
	eic.mutation.SetEquipmentID(id)
	return eic
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eic *EquipmentInfoCreate) SetEquipment(e *Equipment) *EquipmentInfoCreate {
	return eic.SetEquipmentID(e.ID)
}

// Mutation returns the EquipmentInfoMutation object of the builder.
func (eic *EquipmentInfoCreate) Mutation() *EquipmentInfoMutation {
	return eic.mutation
}

// Save creates the EquipmentInfo in the database.
func (eic *EquipmentInfoCreate) Save(ctx context.Context) (*EquipmentInfo, error) {
	var (
		err  error
		node *EquipmentInfo
	)
	if len(eic.hooks) == 0 {
		if err = eic.check(); err != nil {
			return nil, err
		}
		node, err = eic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eic.check(); err != nil {
				return nil, err
			}
			eic.mutation = mutation
			if node, err = eic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(eic.hooks) - 1; i >= 0; i-- {
			if eic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (eic *EquipmentInfoCreate) SaveX(ctx context.Context) *EquipmentInfo {
	v, err := eic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eic *EquipmentInfoCreate) Exec(ctx context.Context) error {
	_, err := eic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eic *EquipmentInfoCreate) ExecX(ctx context.Context) {
	if err := eic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eic *EquipmentInfoCreate) check() error {
	if _, ok := eic.mutation.EquipmentSn(); !ok {
		return &ValidationError{Name: "equipment_sn", err: errors.New(`ent: missing required field "EquipmentInfo.equipment_sn"`)}
	}
	if _, ok := eic.mutation.AccessPod(); !ok {
		return &ValidationError{Name: "access_pod", err: errors.New(`ent: missing required field "EquipmentInfo.access_pod"`)}
	}
	if _, ok := eic.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "EquipmentInfo.state"`)}
	}
	if _, ok := eic.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment", err: errors.New(`ent: missing required edge "EquipmentInfo.equipment"`)}
	}
	return nil
}

func (eic *EquipmentInfoCreate) sqlSave(ctx context.Context) (*EquipmentInfo, error) {
	_node, _spec := eic.createSpec()
	if err := sqlgraph.CreateNode(ctx, eic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (eic *EquipmentInfoCreate) createSpec() (*EquipmentInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &EquipmentInfo{config: eic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: equipmentinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentinfo.FieldID,
			},
		}
	)
	if value, ok := eic.mutation.EquipmentSn(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentinfo.FieldEquipmentSn,
		})
		_node.EquipmentSn = value
	}
	if value, ok := eic.mutation.AccessPod(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: equipmentinfo.FieldAccessPod,
		})
		_node.AccessPod = value
	}
	if value, ok := eic.mutation.State(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: equipmentinfo.FieldState,
		})
		_node.State = value
	}
	if nodes := eic.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   equipmentinfo.EquipmentTable,
			Columns: []string{equipmentinfo.EquipmentColumn},
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
		_node.equipment_equipment_info = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EquipmentInfoCreateBulk is the builder for creating many EquipmentInfo entities in bulk.
type EquipmentInfoCreateBulk struct {
	config
	builders []*EquipmentInfoCreate
}

// Save creates the EquipmentInfo entities in the database.
func (eicb *EquipmentInfoCreateBulk) Save(ctx context.Context) ([]*EquipmentInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(eicb.builders))
	nodes := make([]*EquipmentInfo, len(eicb.builders))
	mutators := make([]Mutator, len(eicb.builders))
	for i := range eicb.builders {
		func(i int, root context.Context) {
			builder := eicb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EquipmentInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, eicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, eicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, eicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (eicb *EquipmentInfoCreateBulk) SaveX(ctx context.Context) []*EquipmentInfo {
	v, err := eicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (eicb *EquipmentInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := eicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eicb *EquipmentInfoCreateBulk) ExecX(ctx context.Context) {
	if err := eicb.Exec(ctx); err != nil {
		panic(err)
	}
}
