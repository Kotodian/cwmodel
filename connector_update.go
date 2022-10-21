// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/evse"
	"github.com/Kotodian/cwmodel/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// ConnectorUpdate is the builder for updating Connector entities.
type ConnectorUpdate struct {
	config
	hooks    []Hook
	mutation *ConnectorMutation
}

// Where appends a list predicates to the ConnectorUpdate builder.
func (cu *ConnectorUpdate) Where(ps ...predicate.Connector) *ConnectorUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetVersion sets the "version" field.
func (cu *ConnectorUpdate) SetVersion(i int64) *ConnectorUpdate {
	cu.mutation.ResetVersion()
	cu.mutation.SetVersion(i)
	return cu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableVersion(i *int64) *ConnectorUpdate {
	if i != nil {
		cu.SetVersion(*i)
	}
	return cu
}

// AddVersion adds i to the "version" field.
func (cu *ConnectorUpdate) AddVersion(i int64) *ConnectorUpdate {
	cu.mutation.AddVersion(i)
	return cu
}

// SetUpdatedBy sets the "updated_by" field.
func (cu *ConnectorUpdate) SetUpdatedBy(d datasource.UUID) *ConnectorUpdate {
	cu.mutation.ResetUpdatedBy()
	cu.mutation.SetUpdatedBy(d)
	return cu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableUpdatedBy(d *datasource.UUID) *ConnectorUpdate {
	if d != nil {
		cu.SetUpdatedBy(*d)
	}
	return cu
}

// AddUpdatedBy adds d to the "updated_by" field.
func (cu *ConnectorUpdate) AddUpdatedBy(d datasource.UUID) *ConnectorUpdate {
	cu.mutation.AddUpdatedBy(d)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *ConnectorUpdate) SetUpdatedAt(i int64) *ConnectorUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(i)
	return cu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (cu *ConnectorUpdate) AddUpdatedAt(i int64) *ConnectorUpdate {
	cu.mutation.AddUpdatedAt(i)
	return cu
}

// SetEquipmentID sets the "equipment_id" field.
func (cu *ConnectorUpdate) SetEquipmentID(d datasource.UUID) *ConnectorUpdate {
	cu.mutation.SetEquipmentID(d)
	return cu
}

// SetEvseID sets the "evse_id" field.
func (cu *ConnectorUpdate) SetEvseID(d datasource.UUID) *ConnectorUpdate {
	cu.mutation.SetEvseID(d)
	return cu
}

// SetEquipmentSn sets the "equipment_sn" field.
func (cu *ConnectorUpdate) SetEquipmentSn(s string) *ConnectorUpdate {
	cu.mutation.SetEquipmentSn(s)
	return cu
}

// SetEvseSerial sets the "evse_serial" field.
func (cu *ConnectorUpdate) SetEvseSerial(s string) *ConnectorUpdate {
	cu.mutation.SetEvseSerial(s)
	return cu
}

// SetSerial sets the "serial" field.
func (cu *ConnectorUpdate) SetSerial(s string) *ConnectorUpdate {
	cu.mutation.SetSerial(s)
	return cu
}

// SetCurrentState sets the "current_state" field.
func (cu *ConnectorUpdate) SetCurrentState(i int) *ConnectorUpdate {
	cu.mutation.ResetCurrentState()
	cu.mutation.SetCurrentState(i)
	return cu
}

// AddCurrentState adds i to the "current_state" field.
func (cu *ConnectorUpdate) AddCurrentState(i int) *ConnectorUpdate {
	cu.mutation.AddCurrentState(i)
	return cu
}

// SetBeforeState sets the "before_state" field.
func (cu *ConnectorUpdate) SetBeforeState(i int) *ConnectorUpdate {
	cu.mutation.ResetBeforeState()
	cu.mutation.SetBeforeState(i)
	return cu
}

// AddBeforeState adds i to the "before_state" field.
func (cu *ConnectorUpdate) AddBeforeState(i int) *ConnectorUpdate {
	cu.mutation.AddBeforeState(i)
	return cu
}

// SetChargingState sets the "charging_state" field.
func (cu *ConnectorUpdate) SetChargingState(i int) *ConnectorUpdate {
	cu.mutation.ResetChargingState()
	cu.mutation.SetChargingState(i)
	return cu
}

// SetNillableChargingState sets the "charging_state" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableChargingState(i *int) *ConnectorUpdate {
	if i != nil {
		cu.SetChargingState(*i)
	}
	return cu
}

// AddChargingState adds i to the "charging_state" field.
func (cu *ConnectorUpdate) AddChargingState(i int) *ConnectorUpdate {
	cu.mutation.AddChargingState(i)
	return cu
}

// ClearChargingState clears the value of the "charging_state" field.
func (cu *ConnectorUpdate) ClearChargingState() *ConnectorUpdate {
	cu.mutation.ClearChargingState()
	return cu
}

// SetReservationID sets the "reservation_id" field.
func (cu *ConnectorUpdate) SetReservationID(d datasource.UUID) *ConnectorUpdate {
	cu.mutation.ResetReservationID()
	cu.mutation.SetReservationID(d)
	return cu
}

// SetNillableReservationID sets the "reservation_id" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableReservationID(d *datasource.UUID) *ConnectorUpdate {
	if d != nil {
		cu.SetReservationID(*d)
	}
	return cu
}

// AddReservationID adds d to the "reservation_id" field.
func (cu *ConnectorUpdate) AddReservationID(d datasource.UUID) *ConnectorUpdate {
	cu.mutation.AddReservationID(d)
	return cu
}

// ClearReservationID clears the value of the "reservation_id" field.
func (cu *ConnectorUpdate) ClearReservationID() *ConnectorUpdate {
	cu.mutation.ClearReservationID()
	return cu
}

// SetOrderID sets the "order_id" field.
func (cu *ConnectorUpdate) SetOrderID(d datasource.UUID) *ConnectorUpdate {
	cu.mutation.ResetOrderID()
	cu.mutation.SetOrderID(d)
	return cu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableOrderID(d *datasource.UUID) *ConnectorUpdate {
	if d != nil {
		cu.SetOrderID(*d)
	}
	return cu
}

// AddOrderID adds d to the "order_id" field.
func (cu *ConnectorUpdate) AddOrderID(d datasource.UUID) *ConnectorUpdate {
	cu.mutation.AddOrderID(d)
	return cu
}

// ClearOrderID clears the value of the "order_id" field.
func (cu *ConnectorUpdate) ClearOrderID() *ConnectorUpdate {
	cu.mutation.ClearOrderID()
	return cu
}

// SetParkNo sets the "park_no" field.
func (cu *ConnectorUpdate) SetParkNo(s string) *ConnectorUpdate {
	cu.mutation.SetParkNo(s)
	return cu
}

// SetNillableParkNo sets the "park_no" field if the given value is not nil.
func (cu *ConnectorUpdate) SetNillableParkNo(s *string) *ConnectorUpdate {
	if s != nil {
		cu.SetParkNo(*s)
	}
	return cu
}

// SetEvse sets the "evse" edge to the Evse entity.
func (cu *ConnectorUpdate) SetEvse(e *Evse) *ConnectorUpdate {
	return cu.SetEvseID(e.ID)
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (cu *ConnectorUpdate) SetEquipment(e *Equipment) *ConnectorUpdate {
	return cu.SetEquipmentID(e.ID)
}

// Mutation returns the ConnectorMutation object of the builder.
func (cu *ConnectorUpdate) Mutation() *ConnectorMutation {
	return cu.mutation
}

// ClearEvse clears the "evse" edge to the Evse entity.
func (cu *ConnectorUpdate) ClearEvse() *ConnectorUpdate {
	cu.mutation.ClearEvse()
	return cu
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (cu *ConnectorUpdate) ClearEquipment() *ConnectorUpdate {
	cu.mutation.ClearEquipment()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ConnectorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConnectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("cwmodel: uninitialized hook (forgotten import cwmodel/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ConnectorUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ConnectorUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ConnectorUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ConnectorUpdate) defaults() {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		v := connector.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ConnectorUpdate) check() error {
	if _, ok := cu.mutation.EvseID(); cu.mutation.EvseCleared() && !ok {
		return errors.New(`cwmodel: clearing a required unique edge "Connector.evse"`)
	}
	if _, ok := cu.mutation.EquipmentID(); cu.mutation.EquipmentCleared() && !ok {
		return errors.New(`cwmodel: clearing a required unique edge "Connector.equipment"`)
	}
	return nil
}

func (cu *ConnectorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   connector.Table,
			Columns: connector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: connector.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Version(); ok {
		_spec.SetField(connector.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedVersion(); ok {
		_spec.AddField(connector.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.UpdatedBy(); ok {
		_spec.SetField(connector.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(connector.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(connector.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(connector.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.EquipmentSn(); ok {
		_spec.SetField(connector.FieldEquipmentSn, field.TypeString, value)
	}
	if value, ok := cu.mutation.EvseSerial(); ok {
		_spec.SetField(connector.FieldEvseSerial, field.TypeString, value)
	}
	if value, ok := cu.mutation.Serial(); ok {
		_spec.SetField(connector.FieldSerial, field.TypeString, value)
	}
	if value, ok := cu.mutation.CurrentState(); ok {
		_spec.SetField(connector.FieldCurrentState, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedCurrentState(); ok {
		_spec.AddField(connector.FieldCurrentState, field.TypeInt, value)
	}
	if value, ok := cu.mutation.BeforeState(); ok {
		_spec.SetField(connector.FieldBeforeState, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedBeforeState(); ok {
		_spec.AddField(connector.FieldBeforeState, field.TypeInt, value)
	}
	if value, ok := cu.mutation.ChargingState(); ok {
		_spec.SetField(connector.FieldChargingState, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedChargingState(); ok {
		_spec.AddField(connector.FieldChargingState, field.TypeInt, value)
	}
	if cu.mutation.ChargingStateCleared() {
		_spec.ClearField(connector.FieldChargingState, field.TypeInt)
	}
	if value, ok := cu.mutation.ReservationID(); ok {
		_spec.SetField(connector.FieldReservationID, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedReservationID(); ok {
		_spec.AddField(connector.FieldReservationID, field.TypeUint64, value)
	}
	if cu.mutation.ReservationIDCleared() {
		_spec.ClearField(connector.FieldReservationID, field.TypeUint64)
	}
	if value, ok := cu.mutation.OrderID(); ok {
		_spec.SetField(connector.FieldOrderID, field.TypeUint64, value)
	}
	if value, ok := cu.mutation.AddedOrderID(); ok {
		_spec.AddField(connector.FieldOrderID, field.TypeUint64, value)
	}
	if cu.mutation.OrderIDCleared() {
		_spec.ClearField(connector.FieldOrderID, field.TypeUint64)
	}
	if value, ok := cu.mutation.ParkNo(); ok {
		_spec.SetField(connector.FieldParkNo, field.TypeString, value)
	}
	if cu.mutation.EvseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connector.EvseTable,
			Columns: []string{connector.EvseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: evse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.EvseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connector.EvseTable,
			Columns: []string{connector.EvseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: evse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connector.EquipmentTable,
			Columns: []string{connector.EquipmentColumn},
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
	if nodes := cu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connector.EquipmentTable,
			Columns: []string{connector.EquipmentColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ConnectorUpdateOne is the builder for updating a single Connector entity.
type ConnectorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConnectorMutation
}

// SetVersion sets the "version" field.
func (cuo *ConnectorUpdateOne) SetVersion(i int64) *ConnectorUpdateOne {
	cuo.mutation.ResetVersion()
	cuo.mutation.SetVersion(i)
	return cuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableVersion(i *int64) *ConnectorUpdateOne {
	if i != nil {
		cuo.SetVersion(*i)
	}
	return cuo
}

// AddVersion adds i to the "version" field.
func (cuo *ConnectorUpdateOne) AddVersion(i int64) *ConnectorUpdateOne {
	cuo.mutation.AddVersion(i)
	return cuo
}

// SetUpdatedBy sets the "updated_by" field.
func (cuo *ConnectorUpdateOne) SetUpdatedBy(d datasource.UUID) *ConnectorUpdateOne {
	cuo.mutation.ResetUpdatedBy()
	cuo.mutation.SetUpdatedBy(d)
	return cuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableUpdatedBy(d *datasource.UUID) *ConnectorUpdateOne {
	if d != nil {
		cuo.SetUpdatedBy(*d)
	}
	return cuo
}

// AddUpdatedBy adds d to the "updated_by" field.
func (cuo *ConnectorUpdateOne) AddUpdatedBy(d datasource.UUID) *ConnectorUpdateOne {
	cuo.mutation.AddUpdatedBy(d)
	return cuo
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *ConnectorUpdateOne) SetUpdatedAt(i int64) *ConnectorUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(i)
	return cuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (cuo *ConnectorUpdateOne) AddUpdatedAt(i int64) *ConnectorUpdateOne {
	cuo.mutation.AddUpdatedAt(i)
	return cuo
}

// SetEquipmentID sets the "equipment_id" field.
func (cuo *ConnectorUpdateOne) SetEquipmentID(d datasource.UUID) *ConnectorUpdateOne {
	cuo.mutation.SetEquipmentID(d)
	return cuo
}

// SetEvseID sets the "evse_id" field.
func (cuo *ConnectorUpdateOne) SetEvseID(d datasource.UUID) *ConnectorUpdateOne {
	cuo.mutation.SetEvseID(d)
	return cuo
}

// SetEquipmentSn sets the "equipment_sn" field.
func (cuo *ConnectorUpdateOne) SetEquipmentSn(s string) *ConnectorUpdateOne {
	cuo.mutation.SetEquipmentSn(s)
	return cuo
}

// SetEvseSerial sets the "evse_serial" field.
func (cuo *ConnectorUpdateOne) SetEvseSerial(s string) *ConnectorUpdateOne {
	cuo.mutation.SetEvseSerial(s)
	return cuo
}

// SetSerial sets the "serial" field.
func (cuo *ConnectorUpdateOne) SetSerial(s string) *ConnectorUpdateOne {
	cuo.mutation.SetSerial(s)
	return cuo
}

// SetCurrentState sets the "current_state" field.
func (cuo *ConnectorUpdateOne) SetCurrentState(i int) *ConnectorUpdateOne {
	cuo.mutation.ResetCurrentState()
	cuo.mutation.SetCurrentState(i)
	return cuo
}

// AddCurrentState adds i to the "current_state" field.
func (cuo *ConnectorUpdateOne) AddCurrentState(i int) *ConnectorUpdateOne {
	cuo.mutation.AddCurrentState(i)
	return cuo
}

// SetBeforeState sets the "before_state" field.
func (cuo *ConnectorUpdateOne) SetBeforeState(i int) *ConnectorUpdateOne {
	cuo.mutation.ResetBeforeState()
	cuo.mutation.SetBeforeState(i)
	return cuo
}

// AddBeforeState adds i to the "before_state" field.
func (cuo *ConnectorUpdateOne) AddBeforeState(i int) *ConnectorUpdateOne {
	cuo.mutation.AddBeforeState(i)
	return cuo
}

// SetChargingState sets the "charging_state" field.
func (cuo *ConnectorUpdateOne) SetChargingState(i int) *ConnectorUpdateOne {
	cuo.mutation.ResetChargingState()
	cuo.mutation.SetChargingState(i)
	return cuo
}

// SetNillableChargingState sets the "charging_state" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableChargingState(i *int) *ConnectorUpdateOne {
	if i != nil {
		cuo.SetChargingState(*i)
	}
	return cuo
}

// AddChargingState adds i to the "charging_state" field.
func (cuo *ConnectorUpdateOne) AddChargingState(i int) *ConnectorUpdateOne {
	cuo.mutation.AddChargingState(i)
	return cuo
}

// ClearChargingState clears the value of the "charging_state" field.
func (cuo *ConnectorUpdateOne) ClearChargingState() *ConnectorUpdateOne {
	cuo.mutation.ClearChargingState()
	return cuo
}

// SetReservationID sets the "reservation_id" field.
func (cuo *ConnectorUpdateOne) SetReservationID(d datasource.UUID) *ConnectorUpdateOne {
	cuo.mutation.ResetReservationID()
	cuo.mutation.SetReservationID(d)
	return cuo
}

// SetNillableReservationID sets the "reservation_id" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableReservationID(d *datasource.UUID) *ConnectorUpdateOne {
	if d != nil {
		cuo.SetReservationID(*d)
	}
	return cuo
}

// AddReservationID adds d to the "reservation_id" field.
func (cuo *ConnectorUpdateOne) AddReservationID(d datasource.UUID) *ConnectorUpdateOne {
	cuo.mutation.AddReservationID(d)
	return cuo
}

// ClearReservationID clears the value of the "reservation_id" field.
func (cuo *ConnectorUpdateOne) ClearReservationID() *ConnectorUpdateOne {
	cuo.mutation.ClearReservationID()
	return cuo
}

// SetOrderID sets the "order_id" field.
func (cuo *ConnectorUpdateOne) SetOrderID(d datasource.UUID) *ConnectorUpdateOne {
	cuo.mutation.ResetOrderID()
	cuo.mutation.SetOrderID(d)
	return cuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableOrderID(d *datasource.UUID) *ConnectorUpdateOne {
	if d != nil {
		cuo.SetOrderID(*d)
	}
	return cuo
}

// AddOrderID adds d to the "order_id" field.
func (cuo *ConnectorUpdateOne) AddOrderID(d datasource.UUID) *ConnectorUpdateOne {
	cuo.mutation.AddOrderID(d)
	return cuo
}

// ClearOrderID clears the value of the "order_id" field.
func (cuo *ConnectorUpdateOne) ClearOrderID() *ConnectorUpdateOne {
	cuo.mutation.ClearOrderID()
	return cuo
}

// SetParkNo sets the "park_no" field.
func (cuo *ConnectorUpdateOne) SetParkNo(s string) *ConnectorUpdateOne {
	cuo.mutation.SetParkNo(s)
	return cuo
}

// SetNillableParkNo sets the "park_no" field if the given value is not nil.
func (cuo *ConnectorUpdateOne) SetNillableParkNo(s *string) *ConnectorUpdateOne {
	if s != nil {
		cuo.SetParkNo(*s)
	}
	return cuo
}

// SetEvse sets the "evse" edge to the Evse entity.
func (cuo *ConnectorUpdateOne) SetEvse(e *Evse) *ConnectorUpdateOne {
	return cuo.SetEvseID(e.ID)
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (cuo *ConnectorUpdateOne) SetEquipment(e *Equipment) *ConnectorUpdateOne {
	return cuo.SetEquipmentID(e.ID)
}

// Mutation returns the ConnectorMutation object of the builder.
func (cuo *ConnectorUpdateOne) Mutation() *ConnectorMutation {
	return cuo.mutation
}

// ClearEvse clears the "evse" edge to the Evse entity.
func (cuo *ConnectorUpdateOne) ClearEvse() *ConnectorUpdateOne {
	cuo.mutation.ClearEvse()
	return cuo
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (cuo *ConnectorUpdateOne) ClearEquipment() *ConnectorUpdateOne {
	cuo.mutation.ClearEquipment()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ConnectorUpdateOne) Select(field string, fields ...string) *ConnectorUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Connector entity.
func (cuo *ConnectorUpdateOne) Save(ctx context.Context) (*Connector, error) {
	var (
		err  error
		node *Connector
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConnectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("cwmodel: uninitialized hook (forgotten import cwmodel/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Connector)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ConnectorMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ConnectorUpdateOne) SaveX(ctx context.Context) *Connector {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ConnectorUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ConnectorUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ConnectorUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		v := connector.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ConnectorUpdateOne) check() error {
	if _, ok := cuo.mutation.EvseID(); cuo.mutation.EvseCleared() && !ok {
		return errors.New(`cwmodel: clearing a required unique edge "Connector.evse"`)
	}
	if _, ok := cuo.mutation.EquipmentID(); cuo.mutation.EquipmentCleared() && !ok {
		return errors.New(`cwmodel: clearing a required unique edge "Connector.equipment"`)
	}
	return nil
}

func (cuo *ConnectorUpdateOne) sqlSave(ctx context.Context) (_node *Connector, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   connector.Table,
			Columns: connector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: connector.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cwmodel: missing "Connector.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connector.FieldID)
		for _, f := range fields {
			if !connector.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cwmodel: invalid field %q for query", f)}
			}
			if f != connector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Version(); ok {
		_spec.SetField(connector.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedVersion(); ok {
		_spec.AddField(connector.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.UpdatedBy(); ok {
		_spec.SetField(connector.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(connector.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(connector.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(connector.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.EquipmentSn(); ok {
		_spec.SetField(connector.FieldEquipmentSn, field.TypeString, value)
	}
	if value, ok := cuo.mutation.EvseSerial(); ok {
		_spec.SetField(connector.FieldEvseSerial, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Serial(); ok {
		_spec.SetField(connector.FieldSerial, field.TypeString, value)
	}
	if value, ok := cuo.mutation.CurrentState(); ok {
		_spec.SetField(connector.FieldCurrentState, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedCurrentState(); ok {
		_spec.AddField(connector.FieldCurrentState, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.BeforeState(); ok {
		_spec.SetField(connector.FieldBeforeState, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedBeforeState(); ok {
		_spec.AddField(connector.FieldBeforeState, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.ChargingState(); ok {
		_spec.SetField(connector.FieldChargingState, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedChargingState(); ok {
		_spec.AddField(connector.FieldChargingState, field.TypeInt, value)
	}
	if cuo.mutation.ChargingStateCleared() {
		_spec.ClearField(connector.FieldChargingState, field.TypeInt)
	}
	if value, ok := cuo.mutation.ReservationID(); ok {
		_spec.SetField(connector.FieldReservationID, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedReservationID(); ok {
		_spec.AddField(connector.FieldReservationID, field.TypeUint64, value)
	}
	if cuo.mutation.ReservationIDCleared() {
		_spec.ClearField(connector.FieldReservationID, field.TypeUint64)
	}
	if value, ok := cuo.mutation.OrderID(); ok {
		_spec.SetField(connector.FieldOrderID, field.TypeUint64, value)
	}
	if value, ok := cuo.mutation.AddedOrderID(); ok {
		_spec.AddField(connector.FieldOrderID, field.TypeUint64, value)
	}
	if cuo.mutation.OrderIDCleared() {
		_spec.ClearField(connector.FieldOrderID, field.TypeUint64)
	}
	if value, ok := cuo.mutation.ParkNo(); ok {
		_spec.SetField(connector.FieldParkNo, field.TypeString, value)
	}
	if cuo.mutation.EvseCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connector.EvseTable,
			Columns: []string{connector.EvseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: evse.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.EvseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connector.EvseTable,
			Columns: []string{connector.EvseColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: evse.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connector.EquipmentTable,
			Columns: []string{connector.EquipmentColumn},
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
	if nodes := cuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   connector.EquipmentTable,
			Columns: []string{connector.EquipmentColumn},
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
	_node = &Connector{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{connector.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
