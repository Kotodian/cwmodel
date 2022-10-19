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

// EvseUpdate is the builder for updating Evse entities.
type EvseUpdate struct {
	config
	hooks    []Hook
	mutation *EvseMutation
}

// Where appends a list predicates to the EvseUpdate builder.
func (eu *EvseUpdate) Where(ps ...predicate.Evse) *EvseUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetVersion sets the "version" field.
func (eu *EvseUpdate) SetVersion(i int64) *EvseUpdate {
	eu.mutation.ResetVersion()
	eu.mutation.SetVersion(i)
	return eu
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (eu *EvseUpdate) SetNillableVersion(i *int64) *EvseUpdate {
	if i != nil {
		eu.SetVersion(*i)
	}
	return eu
}

// AddVersion adds i to the "version" field.
func (eu *EvseUpdate) AddVersion(i int64) *EvseUpdate {
	eu.mutation.AddVersion(i)
	return eu
}

// SetUpdatedBy sets the "updated_by" field.
func (eu *EvseUpdate) SetUpdatedBy(d datasource.UUID) *EvseUpdate {
	eu.mutation.ResetUpdatedBy()
	eu.mutation.SetUpdatedBy(d)
	return eu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (eu *EvseUpdate) SetNillableUpdatedBy(d *datasource.UUID) *EvseUpdate {
	if d != nil {
		eu.SetUpdatedBy(*d)
	}
	return eu
}

// AddUpdatedBy adds d to the "updated_by" field.
func (eu *EvseUpdate) AddUpdatedBy(d datasource.UUID) *EvseUpdate {
	eu.mutation.AddUpdatedBy(d)
	return eu
}

// SetUpdatedAt sets the "updated_at" field.
func (eu *EvseUpdate) SetUpdatedAt(i int64) *EvseUpdate {
	eu.mutation.ResetUpdatedAt()
	eu.mutation.SetUpdatedAt(i)
	return eu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (eu *EvseUpdate) AddUpdatedAt(i int64) *EvseUpdate {
	eu.mutation.AddUpdatedAt(i)
	return eu
}

// SetSerial sets the "serial" field.
func (eu *EvseUpdate) SetSerial(s string) *EvseUpdate {
	eu.mutation.SetSerial(s)
	return eu
}

// SetConnectorNumber sets the "connector_number" field.
func (eu *EvseUpdate) SetConnectorNumber(i int) *EvseUpdate {
	eu.mutation.ResetConnectorNumber()
	eu.mutation.SetConnectorNumber(i)
	return eu
}

// AddConnectorNumber adds i to the "connector_number" field.
func (eu *EvseUpdate) AddConnectorNumber(i int) *EvseUpdate {
	eu.mutation.AddConnectorNumber(i)
	return eu
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (eu *EvseUpdate) SetEquipmentID(id datasource.UUID) *EvseUpdate {
	eu.mutation.SetEquipmentID(id)
	return eu
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (eu *EvseUpdate) SetEquipment(e *Equipment) *EvseUpdate {
	return eu.SetEquipmentID(e.ID)
}

// AddConnectorIDs adds the "connector" edge to the Connector entity by IDs.
func (eu *EvseUpdate) AddConnectorIDs(ids ...datasource.UUID) *EvseUpdate {
	eu.mutation.AddConnectorIDs(ids...)
	return eu
}

// AddConnector adds the "connector" edges to the Connector entity.
func (eu *EvseUpdate) AddConnector(c ...*Connector) *EvseUpdate {
	ids := make([]datasource.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eu.AddConnectorIDs(ids...)
}

// Mutation returns the EvseMutation object of the builder.
func (eu *EvseUpdate) Mutation() *EvseMutation {
	return eu.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (eu *EvseUpdate) ClearEquipment() *EvseUpdate {
	eu.mutation.ClearEquipment()
	return eu
}

// ClearConnector clears all "connector" edges to the Connector entity.
func (eu *EvseUpdate) ClearConnector() *EvseUpdate {
	eu.mutation.ClearConnector()
	return eu
}

// RemoveConnectorIDs removes the "connector" edge to Connector entities by IDs.
func (eu *EvseUpdate) RemoveConnectorIDs(ids ...datasource.UUID) *EvseUpdate {
	eu.mutation.RemoveConnectorIDs(ids...)
	return eu
}

// RemoveConnector removes "connector" edges to Connector entities.
func (eu *EvseUpdate) RemoveConnector(c ...*Connector) *EvseUpdate {
	ids := make([]datasource.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return eu.RemoveConnectorIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EvseUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	eu.defaults()
	if len(eu.hooks) == 0 {
		if err = eu.check(); err != nil {
			return 0, err
		}
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EvseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = eu.check(); err != nil {
				return 0, err
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("cwmodel: uninitialized hook (forgotten import cwmodel/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EvseUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EvseUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EvseUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (eu *EvseUpdate) defaults() {
	if _, ok := eu.mutation.UpdatedAt(); !ok {
		v := evse.UpdateDefaultUpdatedAt()
		eu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EvseUpdate) check() error {
	if _, ok := eu.mutation.EquipmentID(); eu.mutation.EquipmentCleared() && !ok {
		return errors.New(`cwmodel: clearing a required unique edge "Evse.equipment"`)
	}
	return nil
}

func (eu *EvseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   evse.Table,
			Columns: evse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: evse.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Version(); ok {
		_spec.SetField(evse.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := eu.mutation.AddedVersion(); ok {
		_spec.AddField(evse.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := eu.mutation.UpdatedBy(); ok {
		_spec.SetField(evse.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := eu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(evse.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := eu.mutation.UpdatedAt(); ok {
		_spec.SetField(evse.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := eu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(evse.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := eu.mutation.Serial(); ok {
		_spec.SetField(evse.FieldSerial, field.TypeString, value)
	}
	if value, ok := eu.mutation.ConnectorNumber(); ok {
		_spec.SetField(evse.FieldConnectorNumber, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedConnectorNumber(); ok {
		_spec.AddField(evse.FieldConnectorNumber, field.TypeInt, value)
	}
	if eu.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   evse.EquipmentTable,
			Columns: []string{evse.EquipmentColumn},
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
	if nodes := eu.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   evse.EquipmentTable,
			Columns: []string{evse.EquipmentColumn},
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
	if eu.mutation.ConnectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   evse.ConnectorTable,
			Columns: []string{evse.ConnectorColumn},
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
	if nodes := eu.mutation.RemovedConnectorIDs(); len(nodes) > 0 && !eu.mutation.ConnectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   evse.ConnectorTable,
			Columns: []string{evse.ConnectorColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   evse.ConnectorTable,
			Columns: []string{evse.ConnectorColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{evse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// EvseUpdateOne is the builder for updating a single Evse entity.
type EvseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EvseMutation
}

// SetVersion sets the "version" field.
func (euo *EvseUpdateOne) SetVersion(i int64) *EvseUpdateOne {
	euo.mutation.ResetVersion()
	euo.mutation.SetVersion(i)
	return euo
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (euo *EvseUpdateOne) SetNillableVersion(i *int64) *EvseUpdateOne {
	if i != nil {
		euo.SetVersion(*i)
	}
	return euo
}

// AddVersion adds i to the "version" field.
func (euo *EvseUpdateOne) AddVersion(i int64) *EvseUpdateOne {
	euo.mutation.AddVersion(i)
	return euo
}

// SetUpdatedBy sets the "updated_by" field.
func (euo *EvseUpdateOne) SetUpdatedBy(d datasource.UUID) *EvseUpdateOne {
	euo.mutation.ResetUpdatedBy()
	euo.mutation.SetUpdatedBy(d)
	return euo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (euo *EvseUpdateOne) SetNillableUpdatedBy(d *datasource.UUID) *EvseUpdateOne {
	if d != nil {
		euo.SetUpdatedBy(*d)
	}
	return euo
}

// AddUpdatedBy adds d to the "updated_by" field.
func (euo *EvseUpdateOne) AddUpdatedBy(d datasource.UUID) *EvseUpdateOne {
	euo.mutation.AddUpdatedBy(d)
	return euo
}

// SetUpdatedAt sets the "updated_at" field.
func (euo *EvseUpdateOne) SetUpdatedAt(i int64) *EvseUpdateOne {
	euo.mutation.ResetUpdatedAt()
	euo.mutation.SetUpdatedAt(i)
	return euo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (euo *EvseUpdateOne) AddUpdatedAt(i int64) *EvseUpdateOne {
	euo.mutation.AddUpdatedAt(i)
	return euo
}

// SetSerial sets the "serial" field.
func (euo *EvseUpdateOne) SetSerial(s string) *EvseUpdateOne {
	euo.mutation.SetSerial(s)
	return euo
}

// SetConnectorNumber sets the "connector_number" field.
func (euo *EvseUpdateOne) SetConnectorNumber(i int) *EvseUpdateOne {
	euo.mutation.ResetConnectorNumber()
	euo.mutation.SetConnectorNumber(i)
	return euo
}

// AddConnectorNumber adds i to the "connector_number" field.
func (euo *EvseUpdateOne) AddConnectorNumber(i int) *EvseUpdateOne {
	euo.mutation.AddConnectorNumber(i)
	return euo
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (euo *EvseUpdateOne) SetEquipmentID(id datasource.UUID) *EvseUpdateOne {
	euo.mutation.SetEquipmentID(id)
	return euo
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (euo *EvseUpdateOne) SetEquipment(e *Equipment) *EvseUpdateOne {
	return euo.SetEquipmentID(e.ID)
}

// AddConnectorIDs adds the "connector" edge to the Connector entity by IDs.
func (euo *EvseUpdateOne) AddConnectorIDs(ids ...datasource.UUID) *EvseUpdateOne {
	euo.mutation.AddConnectorIDs(ids...)
	return euo
}

// AddConnector adds the "connector" edges to the Connector entity.
func (euo *EvseUpdateOne) AddConnector(c ...*Connector) *EvseUpdateOne {
	ids := make([]datasource.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return euo.AddConnectorIDs(ids...)
}

// Mutation returns the EvseMutation object of the builder.
func (euo *EvseUpdateOne) Mutation() *EvseMutation {
	return euo.mutation
}

// ClearEquipment clears the "equipment" edge to the Equipment entity.
func (euo *EvseUpdateOne) ClearEquipment() *EvseUpdateOne {
	euo.mutation.ClearEquipment()
	return euo
}

// ClearConnector clears all "connector" edges to the Connector entity.
func (euo *EvseUpdateOne) ClearConnector() *EvseUpdateOne {
	euo.mutation.ClearConnector()
	return euo
}

// RemoveConnectorIDs removes the "connector" edge to Connector entities by IDs.
func (euo *EvseUpdateOne) RemoveConnectorIDs(ids ...datasource.UUID) *EvseUpdateOne {
	euo.mutation.RemoveConnectorIDs(ids...)
	return euo
}

// RemoveConnector removes "connector" edges to Connector entities.
func (euo *EvseUpdateOne) RemoveConnector(c ...*Connector) *EvseUpdateOne {
	ids := make([]datasource.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return euo.RemoveConnectorIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EvseUpdateOne) Select(field string, fields ...string) *EvseUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Evse entity.
func (euo *EvseUpdateOne) Save(ctx context.Context) (*Evse, error) {
	var (
		err  error
		node *Evse
	)
	euo.defaults()
	if len(euo.hooks) == 0 {
		if err = euo.check(); err != nil {
			return nil, err
		}
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EvseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = euo.check(); err != nil {
				return nil, err
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("cwmodel: uninitialized hook (forgotten import cwmodel/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, euo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Evse)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from EvseMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EvseUpdateOne) SaveX(ctx context.Context) *Evse {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EvseUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EvseUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (euo *EvseUpdateOne) defaults() {
	if _, ok := euo.mutation.UpdatedAt(); !ok {
		v := evse.UpdateDefaultUpdatedAt()
		euo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EvseUpdateOne) check() error {
	if _, ok := euo.mutation.EquipmentID(); euo.mutation.EquipmentCleared() && !ok {
		return errors.New(`cwmodel: clearing a required unique edge "Evse.equipment"`)
	}
	return nil
}

func (euo *EvseUpdateOne) sqlSave(ctx context.Context) (_node *Evse, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   evse.Table,
			Columns: evse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: evse.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`cwmodel: missing "Evse.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, evse.FieldID)
		for _, f := range fields {
			if !evse.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("cwmodel: invalid field %q for query", f)}
			}
			if f != evse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Version(); ok {
		_spec.SetField(evse.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := euo.mutation.AddedVersion(); ok {
		_spec.AddField(evse.FieldVersion, field.TypeInt64, value)
	}
	if value, ok := euo.mutation.UpdatedBy(); ok {
		_spec.SetField(evse.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := euo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(evse.FieldUpdatedBy, field.TypeUint64, value)
	}
	if value, ok := euo.mutation.UpdatedAt(); ok {
		_spec.SetField(evse.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := euo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(evse.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := euo.mutation.Serial(); ok {
		_spec.SetField(evse.FieldSerial, field.TypeString, value)
	}
	if value, ok := euo.mutation.ConnectorNumber(); ok {
		_spec.SetField(evse.FieldConnectorNumber, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedConnectorNumber(); ok {
		_spec.AddField(evse.FieldConnectorNumber, field.TypeInt, value)
	}
	if euo.mutation.EquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   evse.EquipmentTable,
			Columns: []string{evse.EquipmentColumn},
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
	if nodes := euo.mutation.EquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   evse.EquipmentTable,
			Columns: []string{evse.EquipmentColumn},
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
	if euo.mutation.ConnectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   evse.ConnectorTable,
			Columns: []string{evse.ConnectorColumn},
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
	if nodes := euo.mutation.RemovedConnectorIDs(); len(nodes) > 0 && !euo.mutation.ConnectorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   evse.ConnectorTable,
			Columns: []string{evse.ConnectorColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ConnectorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   evse.ConnectorTable,
			Columns: []string{evse.ConnectorColumn},
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
	_node = &Evse{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{evse.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
