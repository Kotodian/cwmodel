// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/ent/connector"
	"github.com/Kotodian/cwmodel/ent/equipment"
	"github.com/Kotodian/cwmodel/ent/orderinfo"
	"github.com/Kotodian/cwmodel/ent/predicate"
	"github.com/Kotodian/cwmodel/ent/smartchargingeffect"
	"github.com/Kotodian/cwmodel/ent/types"
	"github.com/Kotodian/gokit/datasource"
)

// SmartChargingEffectUpdate is the builder for updating SmartChargingEffect entities.
type SmartChargingEffectUpdate struct {
	config
	hooks    []Hook
	mutation *SmartChargingEffectMutation
}

// Where appends a list predicates to the SmartChargingEffectUpdate builder.
func (sceu *SmartChargingEffectUpdate) Where(ps ...predicate.SmartChargingEffect) *SmartChargingEffectUpdate {
	sceu.mutation.Where(ps...)
	return sceu
}

// SetVersion sets the "version" field.
func (sceu *SmartChargingEffectUpdate) SetVersion(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.ResetVersion()
	sceu.mutation.SetVersion(i)
	return sceu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (sceu *SmartChargingEffectUpdate) SetNillableVersion(i *int64) *SmartChargingEffectUpdate {
	if i != nil {
		sceu.SetVersion(*i)
	}
	return sceu
}

// AddVersion adds i to the "version" field.
func (sceu *SmartChargingEffectUpdate) AddVersion(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.AddVersion(i)
	return sceu
}

// SetUpdatedBy sets the "updated_by" field.
func (sceu *SmartChargingEffectUpdate) SetUpdatedBy(d datasource.UUID) *SmartChargingEffectUpdate {
	sceu.mutation.ResetUpdatedBy()
	sceu.mutation.SetUpdatedBy(d)
	return sceu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sceu *SmartChargingEffectUpdate) SetNillableUpdatedBy(d *datasource.UUID) *SmartChargingEffectUpdate {
	if d != nil {
		sceu.SetUpdatedBy(*d)
	}
	return sceu
}

// AddUpdatedBy adds d to the "updated_by" field.
func (sceu *SmartChargingEffectUpdate) AddUpdatedBy(d datasource.UUID) *SmartChargingEffectUpdate {
	sceu.mutation.AddUpdatedBy(d)
	return sceu
}

// SetUpdatedAt sets the "updated_at" field.
func (sceu *SmartChargingEffectUpdate) SetUpdatedAt(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.ResetUpdatedAt()
	sceu.mutation.SetUpdatedAt(i)
	return sceu
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sceu *SmartChargingEffectUpdate) SetNillableUpdatedAt(i *int64) *SmartChargingEffectUpdate {
	if i != nil {
		sceu.SetUpdatedAt(*i)
	}
	return sceu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (sceu *SmartChargingEffectUpdate) AddUpdatedAt(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.AddUpdatedAt(i)
	return sceu
}

// SetSmartID sets the "smart_id" field.
func (sceu *SmartChargingEffectUpdate) SetSmartID(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.ResetSmartID()
	sceu.mutation.SetSmartID(i)
	return sceu
}

// AddSmartID adds i to the "smart_id" field.
func (sceu *SmartChargingEffectUpdate) AddSmartID(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.AddSmartID(i)
	return sceu
}

// SetStartTime sets the "start_time" field.
func (sceu *SmartChargingEffectUpdate) SetStartTime(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.ResetStartTime()
	sceu.mutation.SetStartTime(i)
	return sceu
}

// AddStartTime adds i to the "start_time" field.
func (sceu *SmartChargingEffectUpdate) AddStartTime(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.AddStartTime(i)
	return sceu
}

// SetPid sets the "pid" field.
func (sceu *SmartChargingEffectUpdate) SetPid(d datasource.UUID) *SmartChargingEffectUpdate {
	sceu.mutation.ResetPid()
	sceu.mutation.SetPid(d)
	return sceu
}

// AddPid adds d to the "pid" field.
func (sceu *SmartChargingEffectUpdate) AddPid(d datasource.UUID) *SmartChargingEffectUpdate {
	sceu.mutation.AddPid(d)
	return sceu
}

// SetUnit sets the "unit" field.
func (sceu *SmartChargingEffectUpdate) SetUnit(s string) *SmartChargingEffectUpdate {
	sceu.mutation.SetUnit(s)
	return sceu
}

// SetEquipmentSn sets the "equipment_sn" field.
func (sceu *SmartChargingEffectUpdate) SetEquipmentSn(s string) *SmartChargingEffectUpdate {
	sceu.mutation.SetEquipmentSn(s)
	return sceu
}

// SetValidFrom sets the "valid_from" field.
func (sceu *SmartChargingEffectUpdate) SetValidFrom(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.ResetValidFrom()
	sceu.mutation.SetValidFrom(i)
	return sceu
}

// SetNillableValidFrom sets the "valid_from" field if the given value is not nil.
func (sceu *SmartChargingEffectUpdate) SetNillableValidFrom(i *int64) *SmartChargingEffectUpdate {
	if i != nil {
		sceu.SetValidFrom(*i)
	}
	return sceu
}

// AddValidFrom adds i to the "valid_from" field.
func (sceu *SmartChargingEffectUpdate) AddValidFrom(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.AddValidFrom(i)
	return sceu
}

// ClearValidFrom clears the value of the "valid_from" field.
func (sceu *SmartChargingEffectUpdate) ClearValidFrom() *SmartChargingEffectUpdate {
	sceu.mutation.ClearValidFrom()
	return sceu
}

// SetValidTo sets the "valid_to" field.
func (sceu *SmartChargingEffectUpdate) SetValidTo(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.ResetValidTo()
	sceu.mutation.SetValidTo(i)
	return sceu
}

// SetNillableValidTo sets the "valid_to" field if the given value is not nil.
func (sceu *SmartChargingEffectUpdate) SetNillableValidTo(i *int64) *SmartChargingEffectUpdate {
	if i != nil {
		sceu.SetValidTo(*i)
	}
	return sceu
}

// AddValidTo adds i to the "valid_to" field.
func (sceu *SmartChargingEffectUpdate) AddValidTo(i int64) *SmartChargingEffectUpdate {
	sceu.mutation.AddValidTo(i)
	return sceu
}

// ClearValidTo clears the value of the "valid_to" field.
func (sceu *SmartChargingEffectUpdate) ClearValidTo() *SmartChargingEffectUpdate {
	sceu.mutation.ClearValidTo()
	return sceu
}

// SetSpec sets the "spec" field.
func (sceu *SmartChargingEffectUpdate) SetSpec(tsp []types.ChargingSchedulePeriod) *SmartChargingEffectUpdate {
	sceu.mutation.SetSpec(tsp)
	return sceu
}

// AppendSpec appends tsp to the "spec" field.
func (sceu *SmartChargingEffectUpdate) AppendSpec(tsp []types.ChargingSchedulePeriod) *SmartChargingEffectUpdate {
	sceu.mutation.AppendSpec(tsp)
	return sceu
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (sceu *SmartChargingEffectUpdate) SetEquipmentID(id datasource.UUID) *SmartChargingEffectUpdate {
	sceu.mutation.SetEquipmentID(id)
	return sceu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (sceu *SmartChargingEffectUpdate) SetEquipment(e *Equipment) *SmartChargingEffectUpdate {
	return sceu.SetEquipmentID(e.ID)
}

// SetConnectorID sets the "connector" edge to the Connector entity by ID.
func (sceu *SmartChargingEffectUpdate) SetConnectorID(id datasource.UUID) *SmartChargingEffectUpdate {
	sceu.mutation.SetConnectorID(id)
	return sceu
}

// SetConnector sets the "connector" edge to the Connector entity.
func (sceu *SmartChargingEffectUpdate) SetConnector(c *Connector) *SmartChargingEffectUpdate {
	return sceu.SetConnectorID(c.ID)
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (sceu *SmartChargingEffectUpdate) SetOrderInfoID(id datasource.UUID) *SmartChargingEffectUpdate {
	sceu.mutation.SetOrderInfoID(id)
	return sceu
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (sceu *SmartChargingEffectUpdate) SetNillableOrderInfoID(id *datasource.UUID) *SmartChargingEffectUpdate {
	if id != nil {
		sceu = sceu.SetOrderInfoID(*id)
	}
	return sceu
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (sceu *SmartChargingEffectUpdate) SetOrderInfo(o *OrderInfo) *SmartChargingEffectUpdate {
	return sceu.SetOrderInfoID(o.ID)
}

// Mutation returns the SmartChargingEffectMutation object of the builder.
func (sceu *SmartChargingEffectUpdate) Mutation() *SmartChargingEffectMutation {
	return sceu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (sceu *SmartChargingEffectUpdate) ClearEquipment() *SmartChargingEffectUpdate {
	sceu.mutation.ClearEquipment()
	return sceu
}

// ClearConnector clears the "connector" edge to the Connector entity.
func (sceu *SmartChargingEffectUpdate) ClearConnector() *SmartChargingEffectUpdate {
	sceu.mutation.ClearConnector()
	return sceu
}

// ClearOrderInfo clears the "order_info" edge to the OrderInfo entity.
func (sceu *SmartChargingEffectUpdate) ClearOrderInfo() *SmartChargingEffectUpdate {
	sceu.mutation.ClearOrderInfo()
	return sceu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (sceu *SmartChargingEffectUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(sceu.hooks) == 0 {
		if err = sceu.check(); err != nil {
			return 0, err
		}
		affected, err = sceu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SmartChargingEffectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sceu.check(); err != nil {
				return 0, err
			}
			sceu.mutation = mutation
			affected, err = sceu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(sceu.hooks) - 1; i >= 0; i-- {
			if sceu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sceu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sceu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (sceu *SmartChargingEffectUpdate) SaveX(ctx context.Context) int {
	affected, err := sceu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (sceu *SmartChargingEffectUpdate) Exec(ctx context.Context) error {
	_, err := sceu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sceu *SmartChargingEffectUpdate) ExecX(ctx context.Context) {
	if err := sceu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sceu *SmartChargingEffectUpdate) check() error {
	if _, ok := sceu.mutation.EquipmentID(); sceu.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SmartChargingEffect.equipment"`)
	}
	if _, ok := sceu.mutation.ConnectorID(); sceu.mutation.ConnectorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SmartChargingEffect.connector"`)
	}
	return nil
}

func (sceu *SmartChargingEffectUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   smartchargingeffect.Table,
			Columns: smartchargingeffect.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: smartchargingeffect.FieldID,
			},
		},
	}
	if ps := sceu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sceu.mutation.Version(); ok {
		_spec.SetField(smartchargingeffect.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.AddedVersion(); ok {
		_spec.AddField(smartchargingeffect.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.UpdatedBy(); ok {
		_spec.SetField(smartchargingeffect.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := sceu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(smartchargingeffect.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := sceu.mutation.UpdatedAt(); ok {
		_spec.SetField(smartchargingeffect.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(smartchargingeffect.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.SmartID(); ok {
		_spec.SetField(smartchargingeffect.FieldSmartID, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.AddedSmartID(); ok {
		_spec.AddField(smartchargingeffect.FieldSmartID, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.StartTime(); ok {
		_spec.SetField(smartchargingeffect.FieldStartTime, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.AddedStartTime(); ok {
		_spec.AddField(smartchargingeffect.FieldStartTime, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.Pid(); ok {
		_spec.SetField(smartchargingeffect.FieldPid, field.TypeUint64, value)
	}
	if value, ok := sceu.mutation.AddedPid(); ok {
		_spec.AddField(smartchargingeffect.FieldPid, field.TypeUint64, value)
	}
	if value, ok := sceu.mutation.Unit(); ok {
		_spec.SetField(smartchargingeffect.FieldUnit, field.TypeString, value)
	}
	if value, ok := sceu.mutation.EquipmentSn(); ok {
		_spec.SetField(smartchargingeffect.FieldEquipmentSn, field.TypeString, value)
	}
	if value, ok := sceu.mutation.ValidFrom(); ok {
		_spec.SetField(smartchargingeffect.FieldValidFrom, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.AddedValidFrom(); ok {
		_spec.AddField(smartchargingeffect.FieldValidFrom, field.TypeInt64, value)
	}
	if sceu.mutation.ValidFromCleared() {
		_spec.ClearField(smartchargingeffect.FieldValidFrom, field.TypeInt64)
	}
	if value, ok := sceu.mutation.ValidTo(); ok {
		_spec.SetField(smartchargingeffect.FieldValidTo, field.TypeInt64, value)
	}
	if value, ok := sceu.mutation.AddedValidTo(); ok {
		_spec.AddField(smartchargingeffect.FieldValidTo, field.TypeInt64, value)
	}
	if sceu.mutation.ValidToCleared() {
		_spec.ClearField(smartchargingeffect.FieldValidTo, field.TypeInt64)
	}
	if value, ok := sceu.mutation.Spec(); ok {
		_spec.SetField(smartchargingeffect.FieldSpec, field.TypeJSON, value)
	}
	if value, ok := sceu.mutation.AppendedSpec(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, smartchargingeffect.FieldSpec, value)
		})
	}
	if sceu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.EquipmentTable,
			Columns: []string{smartchargingeffect.EquipmentColumn},
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
	if nodes := sceu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.EquipmentTable,
			Columns: []string{smartchargingeffect.EquipmentColumn},
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
	if sceu.mutation.ConnectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.ConnectorTable,
			Columns: []string{smartchargingeffect.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: connector.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sceu.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.ConnectorTable,
			Columns: []string{smartchargingeffect.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: connector.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sceu.mutation.OrderInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   smartchargingeffect.OrderInfoTable,
			Columns: []string{smartchargingeffect.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: orderinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sceu.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   smartchargingeffect.OrderInfoTable,
			Columns: []string{smartchargingeffect.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: orderinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, sceu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{smartchargingeffect.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SmartChargingEffectUpdateOne is the builder for updating a single SmartChargingEffect entity.
type SmartChargingEffectUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SmartChargingEffectMutation
}

// SetVersion sets the "version" field.
func (sceuo *SmartChargingEffectUpdateOne) SetVersion(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.ResetVersion()
	sceuo.mutation.SetVersion(i)
	return sceuo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (sceuo *SmartChargingEffectUpdateOne) SetNillableVersion(i *int64) *SmartChargingEffectUpdateOne {
	if i != nil {
		sceuo.SetVersion(*i)
	}
	return sceuo
}

// AddVersion adds i to the "version" field.
func (sceuo *SmartChargingEffectUpdateOne) AddVersion(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AddVersion(i)
	return sceuo
}

// SetUpdatedBy sets the "updated_by" field.
func (sceuo *SmartChargingEffectUpdateOne) SetUpdatedBy(d datasource.UUID) *SmartChargingEffectUpdateOne {
	sceuo.mutation.ResetUpdatedBy()
	sceuo.mutation.SetUpdatedBy(d)
	return sceuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (sceuo *SmartChargingEffectUpdateOne) SetNillableUpdatedBy(d *datasource.UUID) *SmartChargingEffectUpdateOne {
	if d != nil {
		sceuo.SetUpdatedBy(*d)
	}
	return sceuo
}

// AddUpdatedBy adds d to the "updated_by" field.
func (sceuo *SmartChargingEffectUpdateOne) AddUpdatedBy(d datasource.UUID) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AddUpdatedBy(d)
	return sceuo
}

// SetUpdatedAt sets the "updated_at" field.
func (sceuo *SmartChargingEffectUpdateOne) SetUpdatedAt(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.ResetUpdatedAt()
	sceuo.mutation.SetUpdatedAt(i)
	return sceuo
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sceuo *SmartChargingEffectUpdateOne) SetNillableUpdatedAt(i *int64) *SmartChargingEffectUpdateOne {
	if i != nil {
		sceuo.SetUpdatedAt(*i)
	}
	return sceuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (sceuo *SmartChargingEffectUpdateOne) AddUpdatedAt(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AddUpdatedAt(i)
	return sceuo
}

// SetSmartID sets the "smart_id" field.
func (sceuo *SmartChargingEffectUpdateOne) SetSmartID(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.ResetSmartID()
	sceuo.mutation.SetSmartID(i)
	return sceuo
}

// AddSmartID adds i to the "smart_id" field.
func (sceuo *SmartChargingEffectUpdateOne) AddSmartID(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AddSmartID(i)
	return sceuo
}

// SetStartTime sets the "start_time" field.
func (sceuo *SmartChargingEffectUpdateOne) SetStartTime(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.ResetStartTime()
	sceuo.mutation.SetStartTime(i)
	return sceuo
}

// AddStartTime adds i to the "start_time" field.
func (sceuo *SmartChargingEffectUpdateOne) AddStartTime(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AddStartTime(i)
	return sceuo
}

// SetPid sets the "pid" field.
func (sceuo *SmartChargingEffectUpdateOne) SetPid(d datasource.UUID) *SmartChargingEffectUpdateOne {
	sceuo.mutation.ResetPid()
	sceuo.mutation.SetPid(d)
	return sceuo
}

// AddPid adds d to the "pid" field.
func (sceuo *SmartChargingEffectUpdateOne) AddPid(d datasource.UUID) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AddPid(d)
	return sceuo
}

// SetUnit sets the "unit" field.
func (sceuo *SmartChargingEffectUpdateOne) SetUnit(s string) *SmartChargingEffectUpdateOne {
	sceuo.mutation.SetUnit(s)
	return sceuo
}

// SetEquipmentSn sets the "equipment_sn" field.
func (sceuo *SmartChargingEffectUpdateOne) SetEquipmentSn(s string) *SmartChargingEffectUpdateOne {
	sceuo.mutation.SetEquipmentSn(s)
	return sceuo
}

// SetValidFrom sets the "valid_from" field.
func (sceuo *SmartChargingEffectUpdateOne) SetValidFrom(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.ResetValidFrom()
	sceuo.mutation.SetValidFrom(i)
	return sceuo
}

// SetNillableValidFrom sets the "valid_from" field if the given value is not nil.
func (sceuo *SmartChargingEffectUpdateOne) SetNillableValidFrom(i *int64) *SmartChargingEffectUpdateOne {
	if i != nil {
		sceuo.SetValidFrom(*i)
	}
	return sceuo
}

// AddValidFrom adds i to the "valid_from" field.
func (sceuo *SmartChargingEffectUpdateOne) AddValidFrom(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AddValidFrom(i)
	return sceuo
}

// ClearValidFrom clears the value of the "valid_from" field.
func (sceuo *SmartChargingEffectUpdateOne) ClearValidFrom() *SmartChargingEffectUpdateOne {
	sceuo.mutation.ClearValidFrom()
	return sceuo
}

// SetValidTo sets the "valid_to" field.
func (sceuo *SmartChargingEffectUpdateOne) SetValidTo(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.ResetValidTo()
	sceuo.mutation.SetValidTo(i)
	return sceuo
}

// SetNillableValidTo sets the "valid_to" field if the given value is not nil.
func (sceuo *SmartChargingEffectUpdateOne) SetNillableValidTo(i *int64) *SmartChargingEffectUpdateOne {
	if i != nil {
		sceuo.SetValidTo(*i)
	}
	return sceuo
}

// AddValidTo adds i to the "valid_to" field.
func (sceuo *SmartChargingEffectUpdateOne) AddValidTo(i int64) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AddValidTo(i)
	return sceuo
}

// ClearValidTo clears the value of the "valid_to" field.
func (sceuo *SmartChargingEffectUpdateOne) ClearValidTo() *SmartChargingEffectUpdateOne {
	sceuo.mutation.ClearValidTo()
	return sceuo
}

// SetSpec sets the "spec" field.
func (sceuo *SmartChargingEffectUpdateOne) SetSpec(tsp []types.ChargingSchedulePeriod) *SmartChargingEffectUpdateOne {
	sceuo.mutation.SetSpec(tsp)
	return sceuo
}

// AppendSpec appends tsp to the "spec" field.
func (sceuo *SmartChargingEffectUpdateOne) AppendSpec(tsp []types.ChargingSchedulePeriod) *SmartChargingEffectUpdateOne {
	sceuo.mutation.AppendSpec(tsp)
	return sceuo
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (sceuo *SmartChargingEffectUpdateOne) SetEquipmentID(id datasource.UUID) *SmartChargingEffectUpdateOne {
	sceuo.mutation.SetEquipmentID(id)
	return sceuo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (sceuo *SmartChargingEffectUpdateOne) SetEquipment(e *Equipment) *SmartChargingEffectUpdateOne {
	return sceuo.SetEquipmentID(e.ID)
}

// SetConnectorID sets the "connector" edge to the Connector entity by ID.
func (sceuo *SmartChargingEffectUpdateOne) SetConnectorID(id datasource.UUID) *SmartChargingEffectUpdateOne {
	sceuo.mutation.SetConnectorID(id)
	return sceuo
}

// SetConnector sets the "connector" edge to the Connector entity.
func (sceuo *SmartChargingEffectUpdateOne) SetConnector(c *Connector) *SmartChargingEffectUpdateOne {
	return sceuo.SetConnectorID(c.ID)
}

// SetOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID.
func (sceuo *SmartChargingEffectUpdateOne) SetOrderInfoID(id datasource.UUID) *SmartChargingEffectUpdateOne {
	sceuo.mutation.SetOrderInfoID(id)
	return sceuo
}

// SetNillableOrderInfoID sets the "order_info" edge to the OrderInfo entity by ID if the given value is not nil.
func (sceuo *SmartChargingEffectUpdateOne) SetNillableOrderInfoID(id *datasource.UUID) *SmartChargingEffectUpdateOne {
	if id != nil {
		sceuo = sceuo.SetOrderInfoID(*id)
	}
	return sceuo
}

// SetOrderInfo sets the "order_info" edge to the OrderInfo entity.
func (sceuo *SmartChargingEffectUpdateOne) SetOrderInfo(o *OrderInfo) *SmartChargingEffectUpdateOne {
	return sceuo.SetOrderInfoID(o.ID)
}

// Mutation returns the SmartChargingEffectMutation object of the builder.
func (sceuo *SmartChargingEffectUpdateOne) Mutation() *SmartChargingEffectMutation {
	return sceuo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (sceuo *SmartChargingEffectUpdateOne) ClearEquipment() *SmartChargingEffectUpdateOne {
	sceuo.mutation.ClearEquipment()
	return sceuo
}

// ClearConnector clears the "connector" edge to the Connector entity.
func (sceuo *SmartChargingEffectUpdateOne) ClearConnector() *SmartChargingEffectUpdateOne {
	sceuo.mutation.ClearConnector()
	return sceuo
}

// ClearOrderInfo clears the "order_info" edge to the OrderInfo entity.
func (sceuo *SmartChargingEffectUpdateOne) ClearOrderInfo() *SmartChargingEffectUpdateOne {
	sceuo.mutation.ClearOrderInfo()
	return sceuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (sceuo *SmartChargingEffectUpdateOne) Select(field string, fields ...string) *SmartChargingEffectUpdateOne {
	sceuo.fields = append([]string{field}, fields...)
	return sceuo
}

// Save executes the query and returns the updated SmartChargingEffect entity.
func (sceuo *SmartChargingEffectUpdateOne) Save(ctx context.Context) (*SmartChargingEffect, error) {
	var (
		err  error
		node *SmartChargingEffect
	)
	if len(sceuo.hooks) == 0 {
		if err = sceuo.check(); err != nil {
			return nil, err
		}
		node, err = sceuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SmartChargingEffectMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sceuo.check(); err != nil {
				return nil, err
			}
			sceuo.mutation = mutation
			node, err = sceuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(sceuo.hooks) - 1; i >= 0; i-- {
			if sceuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sceuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sceuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*SmartChargingEffect)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SmartChargingEffectMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (sceuo *SmartChargingEffectUpdateOne) SaveX(ctx context.Context) *SmartChargingEffect {
	node, err := sceuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (sceuo *SmartChargingEffectUpdateOne) Exec(ctx context.Context) error {
	_, err := sceuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sceuo *SmartChargingEffectUpdateOne) ExecX(ctx context.Context) {
	if err := sceuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sceuo *SmartChargingEffectUpdateOne) check() error {
	if _, ok := sceuo.mutation.EquipmentID(); sceuo.mutation.EquipmentCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SmartChargingEffect.equipment"`)
	}
	if _, ok := sceuo.mutation.ConnectorID(); sceuo.mutation.ConnectorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "SmartChargingEffect.connector"`)
	}
	return nil
}

func (sceuo *SmartChargingEffectUpdateOne) sqlSave(ctx context.Context) (_node *SmartChargingEffect, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   smartchargingeffect.Table,
			Columns: smartchargingeffect.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: smartchargingeffect.FieldID,
			},
		},
	}
	id, ok := sceuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "SmartChargingEffect.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := sceuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, smartchargingeffect.FieldID)
		for _, f := range fields {
			if !smartchargingeffect.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != smartchargingeffect.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := sceuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := sceuo.mutation.Version(); ok {
		_spec.SetField(smartchargingeffect.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.AddedVersion(); ok {
		_spec.AddField(smartchargingeffect.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.UpdatedBy(); ok {
		_spec.SetField(smartchargingeffect.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := sceuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(smartchargingeffect.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := sceuo.mutation.UpdatedAt(); ok {
		_spec.SetField(smartchargingeffect.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(smartchargingeffect.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.SmartID(); ok {
		_spec.SetField(smartchargingeffect.FieldSmartID, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.AddedSmartID(); ok {
		_spec.AddField(smartchargingeffect.FieldSmartID, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.StartTime(); ok {
		_spec.SetField(smartchargingeffect.FieldStartTime, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.AddedStartTime(); ok {
		_spec.AddField(smartchargingeffect.FieldStartTime, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.Pid(); ok {
		_spec.SetField(smartchargingeffect.FieldPid, field.TypeUint64, value)
	}
	if value, ok := sceuo.mutation.AddedPid(); ok {
		_spec.AddField(smartchargingeffect.FieldPid, field.TypeUint64, value)
	}
	if value, ok := sceuo.mutation.Unit(); ok {
		_spec.SetField(smartchargingeffect.FieldUnit, field.TypeString, value)
	}
	if value, ok := sceuo.mutation.EquipmentSn(); ok {
		_spec.SetField(smartchargingeffect.FieldEquipmentSn, field.TypeString, value)
	}
	if value, ok := sceuo.mutation.ValidFrom(); ok {
		_spec.SetField(smartchargingeffect.FieldValidFrom, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.AddedValidFrom(); ok {
		_spec.AddField(smartchargingeffect.FieldValidFrom, field.TypeInt64, value)
	}
	if sceuo.mutation.ValidFromCleared() {
		_spec.ClearField(smartchargingeffect.FieldValidFrom, field.TypeInt64)
	}
	if value, ok := sceuo.mutation.ValidTo(); ok {
		_spec.SetField(smartchargingeffect.FieldValidTo, field.TypeInt64, value)
	}
	if value, ok := sceuo.mutation.AddedValidTo(); ok {
		_spec.AddField(smartchargingeffect.FieldValidTo, field.TypeInt64, value)
	}
	if sceuo.mutation.ValidToCleared() {
		_spec.ClearField(smartchargingeffect.FieldValidTo, field.TypeInt64)
	}
	if value, ok := sceuo.mutation.Spec(); ok {
		_spec.SetField(smartchargingeffect.FieldSpec, field.TypeJSON, value)
	}
	if value, ok := sceuo.mutation.AppendedSpec(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, smartchargingeffect.FieldSpec, value)
		})
	}
	if sceuo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.EquipmentTable,
			Columns: []string{smartchargingeffect.EquipmentColumn},
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
	if nodes := sceuo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.EquipmentTable,
			Columns: []string{smartchargingeffect.EquipmentColumn},
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
	if sceuo.mutation.ConnectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.ConnectorTable,
			Columns: []string{smartchargingeffect.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: connector.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sceuo.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   smartchargingeffect.ConnectorTable,
			Columns: []string{smartchargingeffect.ConnectorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: connector.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if sceuo.mutation.OrderInfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   smartchargingeffect.OrderInfoTable,
			Columns: []string{smartchargingeffect.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: orderinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := sceuo.mutation.OrderInfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   smartchargingeffect.OrderInfoTable,
			Columns: []string{smartchargingeffect.OrderInfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUint64,
					Column: orderinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &SmartChargingEffect{config: sceuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, sceuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{smartchargingeffect.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
