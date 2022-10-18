// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/ent-practice/ent/smartchargingevent"
)

// SmartChargingEventDelete is the builder for deleting a SmartChargingEvent entity.
type SmartChargingEventDelete struct {
	config
	hooks    []Hook
	mutation *SmartChargingEventMutation
}

// Where appends a list predicates to the SmartChargingEventDelete builder.
func (sced *SmartChargingEventDelete) Where(ps ...predicate.SmartChargingEvent) *SmartChargingEventDelete {
	sced.mutation.Where(ps...)
	return sced
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sced *SmartChargingEventDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sced.hooks) == 0 {
		affected, err = sced.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SmartChargingEventMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sced.mutation = mutation
			affected, err = sced.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sced.hooks) - 1; i >= 0; i-- {
			if sced.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sced.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sced.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (sced *SmartChargingEventDelete) ExecX(ctx context.Context) int {
	n, err := sced.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sced *SmartChargingEventDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: smartchargingevent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: smartchargingevent.FieldID,
			},
		},
	}
	if ps := sced.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sced.driver, _spec)
}

// SmartChargingEventDeleteOne is the builder for deleting a single SmartChargingEvent entity.
type SmartChargingEventDeleteOne struct {
	sced *SmartChargingEventDelete
}

// Exec executes the deletion query.
func (scedo *SmartChargingEventDeleteOne) Exec(ctx context.Context) error {
	n, err := scedo.sced.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{smartchargingevent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (scedo *SmartChargingEventDeleteOne) ExecX(ctx context.Context) {
	scedo.sced.ExecX(ctx)
}
