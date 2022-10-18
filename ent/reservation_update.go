// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/connector"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/ent-practice/ent/reservation"
)

// ReservationUpdate is the builder for updating Reservation entities.
type ReservationUpdate struct {
	config
	hooks    []Hook
	mutation *ReservationMutation
}

// Where appends a list predicates to the ReservationUpdate builder.
func (ru *ReservationUpdate) Where(ps ...predicate.Reservation) *ReservationUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetReservationID sets the "reservation_id" field.
func (ru *ReservationUpdate) SetReservationID(i int64) *ReservationUpdate {
	ru.mutation.ResetReservationID()
	ru.mutation.SetReservationID(i)
	return ru
}

// AddReservationID adds i to the "reservation_id" field.
func (ru *ReservationUpdate) AddReservationID(i int64) *ReservationUpdate {
	ru.mutation.AddReservationID(i)
	return ru
}

// SetAuthorizationMode sets the "authorization_mode" field.
func (ru *ReservationUpdate) SetAuthorizationMode(i int) *ReservationUpdate {
	ru.mutation.ResetAuthorizationMode()
	ru.mutation.SetAuthorizationMode(i)
	return ru
}

// AddAuthorizationMode adds i to the "authorization_mode" field.
func (ru *ReservationUpdate) AddAuthorizationMode(i int) *ReservationUpdate {
	ru.mutation.AddAuthorizationMode(i)
	return ru
}

// SetAuthorizationID sets the "authorization_id" field.
func (ru *ReservationUpdate) SetAuthorizationID(s string) *ReservationUpdate {
	ru.mutation.SetAuthorizationID(s)
	return ru
}

// SetAdditional sets the "additional" field.
func (ru *ReservationUpdate) SetAdditional(s string) *ReservationUpdate {
	ru.mutation.SetAdditional(s)
	return ru
}

// SetNillableAdditional sets the "additional" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableAdditional(s *string) *ReservationUpdate {
	if s != nil {
		ru.SetAdditional(*s)
	}
	return ru
}

// ClearAdditional clears the value of the "additional" field.
func (ru *ReservationUpdate) ClearAdditional() *ReservationUpdate {
	ru.mutation.ClearAdditional()
	return ru
}

// SetCustomerID sets the "customer_id" field.
func (ru *ReservationUpdate) SetCustomerID(s string) *ReservationUpdate {
	ru.mutation.SetCustomerID(s)
	return ru
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (ru *ReservationUpdate) SetNillableCustomerID(s *string) *ReservationUpdate {
	if s != nil {
		ru.SetCustomerID(*s)
	}
	return ru
}

// ClearCustomerID clears the value of the "customer_id" field.
func (ru *ReservationUpdate) ClearCustomerID() *ReservationUpdate {
	ru.mutation.ClearCustomerID()
	return ru
}

// SetExpired sets the "expired" field.
func (ru *ReservationUpdate) SetExpired(i int64) *ReservationUpdate {
	ru.mutation.ResetExpired()
	ru.mutation.SetExpired(i)
	return ru
}

// AddExpired adds i to the "expired" field.
func (ru *ReservationUpdate) AddExpired(i int64) *ReservationUpdate {
	ru.mutation.AddExpired(i)
	return ru
}

// SetState sets the "state" field.
func (ru *ReservationUpdate) SetState(i int) *ReservationUpdate {
	ru.mutation.ResetState()
	ru.mutation.SetState(i)
	return ru
}

// AddState adds i to the "state" field.
func (ru *ReservationUpdate) AddState(i int) *ReservationUpdate {
	ru.mutation.AddState(i)
	return ru
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (ru *ReservationUpdate) SetEquipmentID(id int) *ReservationUpdate {
	ru.mutation.SetEquipmentID(id)
	return ru
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (ru *ReservationUpdate) SetNillableEquipmentID(id *int) *ReservationUpdate {
	if id != nil {
		ru = ru.SetEquipmentID(*id)
	}
	return ru
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (ru *ReservationUpdate) SetEquipment(e *Equipment) *ReservationUpdate {
	return ru.SetEquipmentID(e.ID)
}

// SetConnectorID sets the "connector" edge to the Connector entity by ID.
func (ru *ReservationUpdate) SetConnectorID(id int) *ReservationUpdate {
	ru.mutation.SetConnectorID(id)
	return ru
}

// SetNillableConnectorID sets the "connector" edge to the Connector entity by ID if the given value is not nil.
func (ru *ReservationUpdate) SetNillableConnectorID(id *int) *ReservationUpdate {
	if id != nil {
		ru = ru.SetConnectorID(*id)
	}
	return ru
}

// SetConnector sets the "connector" edge to the Connector entity.
func (ru *ReservationUpdate) SetConnector(c *Connector) *ReservationUpdate {
	return ru.SetConnectorID(c.ID)
}

// Mutation returns the ReservationMutation object of the builder.
func (ru *ReservationUpdate) Mutation() *ReservationMutation {
	return ru.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (ru *ReservationUpdate) ClearEquipment() *ReservationUpdate {
	ru.mutation.ClearEquipment()
	return ru
}

// ClearConnector clears the "connector" edge to the Connector entity.
func (ru *ReservationUpdate) ClearConnector() *ReservationUpdate {
	ru.mutation.ClearConnector()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *ReservationUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReservationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *ReservationUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *ReservationUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *ReservationUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ru *ReservationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reservation.Table,
			Columns: reservation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reservation.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.ReservationID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: reservation.FieldReservationID,
		})
	}
	if value, ok := ru.mutation.AddedReservationID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: reservation.FieldReservationID,
		})
	}
	if value, ok := ru.mutation.AuthorizationMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldAuthorizationMode,
		})
	}
	if value, ok := ru.mutation.AddedAuthorizationMode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldAuthorizationMode,
		})
	}
	if value, ok := ru.mutation.AuthorizationID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reservation.FieldAuthorizationID,
		})
	}
	if value, ok := ru.mutation.Additional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reservation.FieldAdditional,
		})
	}
	if ru.mutation.AdditionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: reservation.FieldAdditional,
		})
	}
	if value, ok := ru.mutation.CustomerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reservation.FieldCustomerID,
		})
	}
	if ru.mutation.CustomerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: reservation.FieldCustomerID,
		})
	}
	if value, ok := ru.mutation.Expired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: reservation.FieldExpired,
		})
	}
	if value, ok := ru.mutation.AddedExpired(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: reservation.FieldExpired,
		})
	}
	if value, ok := ru.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldState,
		})
	}
	if value, ok := ru.mutation.AddedState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldState,
		})
	}
	if ru.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.EquipmentTable,
			Columns: []string{reservation.EquipmentColumn},
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
	if nodes := ru.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.EquipmentTable,
			Columns: []string{reservation.EquipmentColumn},
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
	if ru.mutation.ConnectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.ConnectorTable,
			Columns: []string{reservation.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connector.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.ConnectorTable,
			Columns: []string{reservation.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connector.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reservation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ReservationUpdateOne is the builder for updating a single Reservation entity.
type ReservationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ReservationMutation
}

// SetReservationID sets the "reservation_id" field.
func (ruo *ReservationUpdateOne) SetReservationID(i int64) *ReservationUpdateOne {
	ruo.mutation.ResetReservationID()
	ruo.mutation.SetReservationID(i)
	return ruo
}

// AddReservationID adds i to the "reservation_id" field.
func (ruo *ReservationUpdateOne) AddReservationID(i int64) *ReservationUpdateOne {
	ruo.mutation.AddReservationID(i)
	return ruo
}

// SetAuthorizationMode sets the "authorization_mode" field.
func (ruo *ReservationUpdateOne) SetAuthorizationMode(i int) *ReservationUpdateOne {
	ruo.mutation.ResetAuthorizationMode()
	ruo.mutation.SetAuthorizationMode(i)
	return ruo
}

// AddAuthorizationMode adds i to the "authorization_mode" field.
func (ruo *ReservationUpdateOne) AddAuthorizationMode(i int) *ReservationUpdateOne {
	ruo.mutation.AddAuthorizationMode(i)
	return ruo
}

// SetAuthorizationID sets the "authorization_id" field.
func (ruo *ReservationUpdateOne) SetAuthorizationID(s string) *ReservationUpdateOne {
	ruo.mutation.SetAuthorizationID(s)
	return ruo
}

// SetAdditional sets the "additional" field.
func (ruo *ReservationUpdateOne) SetAdditional(s string) *ReservationUpdateOne {
	ruo.mutation.SetAdditional(s)
	return ruo
}

// SetNillableAdditional sets the "additional" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableAdditional(s *string) *ReservationUpdateOne {
	if s != nil {
		ruo.SetAdditional(*s)
	}
	return ruo
}

// ClearAdditional clears the value of the "additional" field.
func (ruo *ReservationUpdateOne) ClearAdditional() *ReservationUpdateOne {
	ruo.mutation.ClearAdditional()
	return ruo
}

// SetCustomerID sets the "customer_id" field.
func (ruo *ReservationUpdateOne) SetCustomerID(s string) *ReservationUpdateOne {
	ruo.mutation.SetCustomerID(s)
	return ruo
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableCustomerID(s *string) *ReservationUpdateOne {
	if s != nil {
		ruo.SetCustomerID(*s)
	}
	return ruo
}

// ClearCustomerID clears the value of the "customer_id" field.
func (ruo *ReservationUpdateOne) ClearCustomerID() *ReservationUpdateOne {
	ruo.mutation.ClearCustomerID()
	return ruo
}

// SetExpired sets the "expired" field.
func (ruo *ReservationUpdateOne) SetExpired(i int64) *ReservationUpdateOne {
	ruo.mutation.ResetExpired()
	ruo.mutation.SetExpired(i)
	return ruo
}

// AddExpired adds i to the "expired" field.
func (ruo *ReservationUpdateOne) AddExpired(i int64) *ReservationUpdateOne {
	ruo.mutation.AddExpired(i)
	return ruo
}

// SetState sets the "state" field.
func (ruo *ReservationUpdateOne) SetState(i int) *ReservationUpdateOne {
	ruo.mutation.ResetState()
	ruo.mutation.SetState(i)
	return ruo
}

// AddState adds i to the "state" field.
func (ruo *ReservationUpdateOne) AddState(i int) *ReservationUpdateOne {
	ruo.mutation.AddState(i)
	return ruo
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (ruo *ReservationUpdateOne) SetEquipmentID(id int) *ReservationUpdateOne {
	ruo.mutation.SetEquipmentID(id)
	return ruo
}

// SetNillableEquipmentID sets the "equipment" edge to the Equipment entity by ID if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableEquipmentID(id *int) *ReservationUpdateOne {
	if id != nil {
		ruo = ruo.SetEquipmentID(*id)
	}
	return ruo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (ruo *ReservationUpdateOne) SetEquipment(e *Equipment) *ReservationUpdateOne {
	return ruo.SetEquipmentID(e.ID)
}

// SetConnectorID sets the "connector" edge to the Connector entity by ID.
func (ruo *ReservationUpdateOne) SetConnectorID(id int) *ReservationUpdateOne {
	ruo.mutation.SetConnectorID(id)
	return ruo
}

// SetNillableConnectorID sets the "connector" edge to the Connector entity by ID if the given value is not nil.
func (ruo *ReservationUpdateOne) SetNillableConnectorID(id *int) *ReservationUpdateOne {
	if id != nil {
		ruo = ruo.SetConnectorID(*id)
	}
	return ruo
}

// SetConnector sets the "connector" edge to the Connector entity.
func (ruo *ReservationUpdateOne) SetConnector(c *Connector) *ReservationUpdateOne {
	return ruo.SetConnectorID(c.ID)
}

// Mutation returns the ReservationMutation object of the builder.
func (ruo *ReservationUpdateOne) Mutation() *ReservationMutation {
	return ruo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (ruo *ReservationUpdateOne) ClearEquipment() *ReservationUpdateOne {
	ruo.mutation.ClearEquipment()
	return ruo
}

// ClearConnector clears the "connector" edge to the Connector entity.
func (ruo *ReservationUpdateOne) ClearConnector() *ReservationUpdateOne {
	ruo.mutation.ClearConnector()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *ReservationUpdateOne) Select(field string, fields ...string) *ReservationUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Reservation entity.
func (ruo *ReservationUpdateOne) Save(ctx context.Context) (*Reservation, error) {
	var (
		err  error
		node *Reservation
	)
	if len(ruo.hooks) == 0 {
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReservationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *ReservationUpdateOne) SaveX(ctx context.Context) *Reservation {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *ReservationUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *ReservationUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ruo *ReservationUpdateOne) sqlSave(ctx context.Context) (_node *Reservation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   reservation.Table,
			Columns: reservation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: reservation.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Reservation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, reservation.FieldID)
		for _, f := range fields {
			if !reservation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != reservation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.ReservationID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: reservation.FieldReservationID,
		})
	}
	if value, ok := ruo.mutation.AddedReservationID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: reservation.FieldReservationID,
		})
	}
	if value, ok := ruo.mutation.AuthorizationMode(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldAuthorizationMode,
		})
	}
	if value, ok := ruo.mutation.AddedAuthorizationMode(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldAuthorizationMode,
		})
	}
	if value, ok := ruo.mutation.AuthorizationID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reservation.FieldAuthorizationID,
		})
	}
	if value, ok := ruo.mutation.Additional(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reservation.FieldAdditional,
		})
	}
	if ruo.mutation.AdditionalCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: reservation.FieldAdditional,
		})
	}
	if value, ok := ruo.mutation.CustomerID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: reservation.FieldCustomerID,
		})
	}
	if ruo.mutation.CustomerIDCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: reservation.FieldCustomerID,
		})
	}
	if value, ok := ruo.mutation.Expired(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: reservation.FieldExpired,
		})
	}
	if value, ok := ruo.mutation.AddedExpired(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: reservation.FieldExpired,
		})
	}
	if value, ok := ruo.mutation.State(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldState,
		})
	}
	if value, ok := ruo.mutation.AddedState(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: reservation.FieldState,
		})
	}
	if ruo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.EquipmentTable,
			Columns: []string{reservation.EquipmentColumn},
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
	if nodes := ruo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.EquipmentTable,
			Columns: []string{reservation.EquipmentColumn},
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
	if ruo.mutation.ConnectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.ConnectorTable,
			Columns: []string{reservation.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connector.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   reservation.ConnectorTable,
			Columns: []string{reservation.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: connector.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Reservation{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{reservation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
