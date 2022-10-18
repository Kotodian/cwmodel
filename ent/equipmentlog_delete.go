// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipmentlog"
	"github.com/Kotodian/ent-practice/ent/predicate"
)

// EquipmentLogDelete is the builder for deleting a EquipmentLog entity.
type EquipmentLogDelete struct {
	config
	hooks    []Hook
	mutation *EquipmentLogMutation
}

// Where appends a list predicates to the EquipmentLogDelete builder.
func (eld *EquipmentLogDelete) Where(ps ...predicate.EquipmentLog) *EquipmentLogDelete {
	eld.mutation.Where(ps...)
	return eld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (eld *EquipmentLogDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eld.hooks) == 0 {
		affected, err = eld.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EquipmentLogMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eld.mutation = mutation
			affected, err = eld.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eld.hooks) - 1; i >= 0; i-- {
			if eld.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eld.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eld.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (eld *EquipmentLogDelete) ExecX(ctx context.Context) int {
	n, err := eld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (eld *EquipmentLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: equipmentlog.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentlog.FieldID,
			},
		},
	}
	if ps := eld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, eld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// EquipmentLogDeleteOne is the builder for deleting a single EquipmentLog entity.
type EquipmentLogDeleteOne struct {
	eld *EquipmentLogDelete
}

// Exec executes the deletion query.
func (eldo *EquipmentLogDeleteOne) Exec(ctx context.Context) error {
	n, err := eldo.eld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{equipmentlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (eldo *EquipmentLogDeleteOne) ExecX(ctx context.Context) {
	eldo.eld.ExecX(ctx)
}
