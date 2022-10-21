// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/evse"
	"github.com/Kotodian/gokit/datasource"
)

// ConnectorCreate is the builder for creating a Connector entity.
type ConnectorCreate struct {
	config
	mutation *ConnectorMutation
	hooks    []Hook
}

// SetVersion sets the "version" field.
func (cc *ConnectorCreate) SetVersion(i int64) *ConnectorCreate {
	cc.mutation.SetVersion(i)
	return cc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableVersion(i *int64) *ConnectorCreate {
	if i != nil {
		cc.SetVersion(*i)
	}
	return cc
}

// SetCreatedBy sets the "created_by" field.
func (cc *ConnectorCreate) SetCreatedBy(d datasource.UUID) *ConnectorCreate {
	cc.mutation.SetCreatedBy(d)
	return cc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableCreatedBy(d *datasource.UUID) *ConnectorCreate {
	if d != nil {
		cc.SetCreatedBy(*d)
	}
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *ConnectorCreate) SetCreatedAt(i int64) *ConnectorCreate {
	cc.mutation.SetCreatedAt(i)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableCreatedAt(i *int64) *ConnectorCreate {
	if i != nil {
		cc.SetCreatedAt(*i)
	}
	return cc
}

// SetUpdatedBy sets the "updated_by" field.
func (cc *ConnectorCreate) SetUpdatedBy(d datasource.UUID) *ConnectorCreate {
	cc.mutation.SetUpdatedBy(d)
	return cc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableUpdatedBy(d *datasource.UUID) *ConnectorCreate {
	if d != nil {
		cc.SetUpdatedBy(*d)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *ConnectorCreate) SetUpdatedAt(i int64) *ConnectorCreate {
	cc.mutation.SetUpdatedAt(i)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableUpdatedAt(i *int64) *ConnectorCreate {
	if i != nil {
		cc.SetUpdatedAt(*i)
	}
	return cc
}

// SetEquipmentID sets the "equipment_id" field.
func (cc *ConnectorCreate) SetEquipmentID(d datasource.UUID) *ConnectorCreate {
	cc.mutation.SetEquipmentID(d)
	return cc
}

// SetEvseID sets the "evse_id" field.
func (cc *ConnectorCreate) SetEvseID(d datasource.UUID) *ConnectorCreate {
	cc.mutation.SetEvseID(d)
	return cc
}

// SetEquipmentSn sets the "equipment_sn" field.
func (cc *ConnectorCreate) SetEquipmentSn(s string) *ConnectorCreate {
	cc.mutation.SetEquipmentSn(s)
	return cc
}

// SetEvseSerial sets the "evse_serial" field.
func (cc *ConnectorCreate) SetEvseSerial(s string) *ConnectorCreate {
	cc.mutation.SetEvseSerial(s)
	return cc
}

// SetSerial sets the "serial" field.
func (cc *ConnectorCreate) SetSerial(s string) *ConnectorCreate {
	cc.mutation.SetSerial(s)
	return cc
}

// SetCurrentState sets the "current_state" field.
func (cc *ConnectorCreate) SetCurrentState(i int) *ConnectorCreate {
	cc.mutation.SetCurrentState(i)
	return cc
}

// SetBeforeState sets the "before_state" field.
func (cc *ConnectorCreate) SetBeforeState(i int) *ConnectorCreate {
	cc.mutation.SetBeforeState(i)
	return cc
}

// SetChargingState sets the "charging_state" field.
func (cc *ConnectorCreate) SetChargingState(i int) *ConnectorCreate {
	cc.mutation.SetChargingState(i)
	return cc
}

// SetNillableChargingState sets the "charging_state" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableChargingState(i *int) *ConnectorCreate {
	if i != nil {
		cc.SetChargingState(*i)
	}
	return cc
}

// SetReservationID sets the "reservation_id" field.
func (cc *ConnectorCreate) SetReservationID(d datasource.UUID) *ConnectorCreate {
	cc.mutation.SetReservationID(d)
	return cc
}

// SetNillableReservationID sets the "reservation_id" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableReservationID(d *datasource.UUID) *ConnectorCreate {
	if d != nil {
		cc.SetReservationID(*d)
	}
	return cc
}

// SetOrderID sets the "order_id" field.
func (cc *ConnectorCreate) SetOrderID(d datasource.UUID) *ConnectorCreate {
	cc.mutation.SetOrderID(d)
	return cc
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableOrderID(d *datasource.UUID) *ConnectorCreate {
	if d != nil {
		cc.SetOrderID(*d)
	}
	return cc
}

// SetParkNo sets the "park_no" field.
func (cc *ConnectorCreate) SetParkNo(s string) *ConnectorCreate {
	cc.mutation.SetParkNo(s)
	return cc
}

// SetNillableParkNo sets the "park_no" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableParkNo(s *string) *ConnectorCreate {
	if s != nil {
		cc.SetParkNo(*s)
	}
	return cc
}

// SetID sets the "id" field.
func (cc *ConnectorCreate) SetID(d datasource.UUID) *ConnectorCreate {
	cc.mutation.SetID(d)
	return cc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (cc *ConnectorCreate) SetNillableID(d *datasource.UUID) *ConnectorCreate {
	if d != nil {
		cc.SetID(*d)
	}
	return cc
}

// SetEvse sets the "evse" edge to the Evse entity.
func (cc *ConnectorCreate) SetEvse(e *Evse) *ConnectorCreate {
	return cc.SetEvseID(e.ID)
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (cc *ConnectorCreate) SetEquipment(e *Equipment) *ConnectorCreate {
	return cc.SetEquipmentID(e.ID)
}

// Mutation returns the ConnectorMutation object of the builder.
func (cc *ConnectorCreate) Mutation() *ConnectorMutation {
	return cc.mutation
}

// Save creates the Connector in the database.
func (cc *ConnectorCreate) Save(ctx context.Context) (*Connector, error) {
	var (
		err  error
		node *Connector
	)
	cc.defaults()
	if len(cc.hooks) == 0 {
		if err = cc.check(); err != nil {
			return nil, err
		}
		node, err = cc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConnectorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cc.check(); err != nil {
				return nil, err
			}
			cc.mutation = mutation
			if node, err = cc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(cc.hooks) - 1; i >= 0; i-- {
			if cc.hooks[i] == nil {
				return nil, fmt.Errorf("cwmodel: uninitialized hook (forgotten import cwmodel/runtime?)")
			}
			mut = cc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, cc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (cc *ConnectorCreate) SaveX(ctx context.Context) *Connector {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *ConnectorCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *ConnectorCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *ConnectorCreate) defaults() {
	if _, ok := cc.mutation.Version(); !ok {
		v := connector.DefaultVersion
		cc.mutation.SetVersion(v)
	}
	if _, ok := cc.mutation.CreatedBy(); !ok {
		v := connector.DefaultCreatedBy
		cc.mutation.SetCreatedBy(v)
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := connector.DefaultCreatedAt
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedBy(); !ok {
		v := connector.DefaultUpdatedBy
		cc.mutation.SetUpdatedBy(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := connector.DefaultUpdatedAt
		cc.mutation.SetUpdatedAt(v)
	}
	if _, ok := cc.mutation.ParkNo(); !ok {
		v := connector.DefaultParkNo
		cc.mutation.SetParkNo(v)
	}
	if _, ok := cc.mutation.ID(); !ok {
		v := connector.DefaultID
		cc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *ConnectorCreate) check() error {
	if _, ok := cc.mutation.Version(); !ok {
		return &ValidationError{Name: "version", err: errors.New(`cwmodel: missing required field "Connector.version"`)}
	}
	if _, ok := cc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`cwmodel: missing required field "Connector.created_by"`)}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`cwmodel: missing required field "Connector.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`cwmodel: missing required field "Connector.updated_by"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`cwmodel: missing required field "Connector.updated_at"`)}
	}
	if _, ok := cc.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment_id", err: errors.New(`cwmodel: missing required field "Connector.equipment_id"`)}
	}
	if _, ok := cc.mutation.EvseID(); !ok {
		return &ValidationError{Name: "evse_id", err: errors.New(`cwmodel: missing required field "Connector.evse_id"`)}
	}
	if _, ok := cc.mutation.EquipmentSn(); !ok {
		return &ValidationError{Name: "equipment_sn", err: errors.New(`cwmodel: missing required field "Connector.equipment_sn"`)}
	}
	if _, ok := cc.mutation.EvseSerial(); !ok {
		return &ValidationError{Name: "evse_serial", err: errors.New(`cwmodel: missing required field "Connector.evse_serial"`)}
	}
	if _, ok := cc.mutation.Serial(); !ok {
		return &ValidationError{Name: "serial", err: errors.New(`cwmodel: missing required field "Connector.serial"`)}
	}
	if _, ok := cc.mutation.CurrentState(); !ok {
		return &ValidationError{Name: "current_state", err: errors.New(`cwmodel: missing required field "Connector.current_state"`)}
	}
	if _, ok := cc.mutation.BeforeState(); !ok {
		return &ValidationError{Name: "before_state", err: errors.New(`cwmodel: missing required field "Connector.before_state"`)}
	}
	if _, ok := cc.mutation.ParkNo(); !ok {
		return &ValidationError{Name: "park_no", err: errors.New(`cwmodel: missing required field "Connector.park_no"`)}
	}
	if _, ok := cc.mutation.EvseID(); !ok {
		return &ValidationError{Name: "evse", err: errors.New(`cwmodel: missing required edge "Connector.evse"`)}
	}
	if _, ok := cc.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment", err: errors.New(`cwmodel: missing required edge "Connector.equipment"`)}
	}
	return nil
}

func (cc *ConnectorCreate) sqlSave(ctx context.Context) (*Connector, error) {
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
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

func (cc *ConnectorCreate) createSpec() (*Connector, *sqlgraph.CreateSpec) {
	var (
		_node = &Connector{config: cc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: connector.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: connector.FieldID,
			},
		}
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.Version(); ok {
		_spec.SetField(connector.FieldVersion, field.TypeInt64, value)
		_node.Version = value
	}
	if value, ok := cc.mutation.CreatedBy(); ok {
		_spec.SetField(connector.FieldCreatedBy, field.TypeUint64, value)
		_node.CreatedBy = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(connector.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedBy(); ok {
		_spec.SetField(connector.FieldUpdatedBy, field.TypeUint64, value)
		_node.UpdatedBy = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(connector.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.EquipmentSn(); ok {
		_spec.SetField(connector.FieldEquipmentSn, field.TypeString, value)
		_node.EquipmentSn = value
	}
	if value, ok := cc.mutation.EvseSerial(); ok {
		_spec.SetField(connector.FieldEvseSerial, field.TypeString, value)
		_node.EvseSerial = value
	}
	if value, ok := cc.mutation.Serial(); ok {
		_spec.SetField(connector.FieldSerial, field.TypeString, value)
		_node.Serial = value
	}
	if value, ok := cc.mutation.CurrentState(); ok {
		_spec.SetField(connector.FieldCurrentState, field.TypeInt, value)
		_node.CurrentState = value
	}
	if value, ok := cc.mutation.BeforeState(); ok {
		_spec.SetField(connector.FieldBeforeState, field.TypeInt, value)
		_node.BeforeState = value
	}
	if value, ok := cc.mutation.ChargingState(); ok {
		_spec.SetField(connector.FieldChargingState, field.TypeInt, value)
		_node.ChargingState = value
	}
	if value, ok := cc.mutation.ReservationID(); ok {
		_spec.SetField(connector.FieldReservationID, field.TypeUint64, value)
		_node.ReservationID = &value
	}
	if value, ok := cc.mutation.OrderID(); ok {
		_spec.SetField(connector.FieldOrderID, field.TypeUint64, value)
		_node.OrderID = &value
	}
	if value, ok := cc.mutation.ParkNo(); ok {
		_spec.SetField(connector.FieldParkNo, field.TypeString, value)
		_node.ParkNo = value
	}
	if nodes := cc.mutation.EvseIDs(); len(nodes) > 0 {
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
		_node.EvseID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.EquipmentIDs(); len(nodes) > 0 {
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
		_node.EquipmentID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ConnectorCreateBulk is the builder for creating many Connector entities in bulk.
type ConnectorCreateBulk struct {
	config
	builders []*ConnectorCreate
}

// Save creates the Connector entities in the database.
func (ccb *ConnectorCreateBulk) Save(ctx context.Context) ([]*Connector, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Connector, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ConnectorMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *ConnectorCreateBulk) SaveX(ctx context.Context) []*Connector {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *ConnectorCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *ConnectorCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
