// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/connector"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/evse"
	"github.com/Kotodian/gokit/datasource"
)

// EvseCreate is the builder for creating a Evse entity.
type EvseCreate struct {
	config
	mutation *EvseMutation
	hooks    []Hook
}

// SetSerial sets the "serial" field.
func (ec *EvseCreate) SetSerial(s string) *EvseCreate {
	ec.mutation.SetSerial(s)
	return ec
}

// SetConnectorNumber sets the "connector_number" field.
func (ec *EvseCreate) SetConnectorNumber(i int) *EvseCreate {
	ec.mutation.SetConnectorNumber(i)
	return ec
}

// SetID sets the "id" field.
func (ec *EvseCreate) SetID(d datasource.UUID) *EvseCreate {
	ec.mutation.SetID(d)
	return ec
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ec *EvseCreate) SetNillableID(d *datasource.UUID) *EvseCreate {
	if d != nil {
		ec.SetID(*d)
	}
	return ec
}

// SetEquipmentID sets the "equipment" edge to the Equipment entity by ID.
func (ec *EvseCreate) SetEquipmentID(id datasource.UUID) *EvseCreate {
	ec.mutation.SetEquipmentID(id)
	return ec
}

// SetEquipment sets the "equipment" edge to the Equipment entity.
func (ec *EvseCreate) SetEquipment(e *Equipment) *EvseCreate {
	return ec.SetEquipmentID(e.ID)
}

// AddConnectorIDs adds the "connectors" edge to the Connector entity by IDs.
func (ec *EvseCreate) AddConnectorIDs(ids ...datasource.UUID) *EvseCreate {
	ec.mutation.AddConnectorIDs(ids...)
	return ec
}

// AddConnectors adds the "connectors" edges to the Connector entity.
func (ec *EvseCreate) AddConnectors(c ...*Connector) *EvseCreate {
	ids := make([]datasource.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ec.AddConnectorIDs(ids...)
}

// Mutation returns the EvseMutation object of the builder.
func (ec *EvseCreate) Mutation() *EvseMutation {
	return ec.mutation
}

// Save creates the Evse in the database.
func (ec *EvseCreate) Save(ctx context.Context) (*Evse, error) {
	var (
		err  error
		node *Evse
	)
	ec.defaults()
	if len(ec.hooks) == 0 {
		if err = ec.check(); err != nil {
			return nil, err
		}
		node, err = ec.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EvseMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ec.check(); err != nil {
				return nil, err
			}
			ec.mutation = mutation
			if node, err = ec.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ec.hooks) - 1; i >= 0; i-- {
			if ec.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ec.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ec.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EvseCreate) SaveX(ctx context.Context) *Evse {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EvseCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EvseCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EvseCreate) defaults() {
	if _, ok := ec.mutation.ID(); !ok {
		v := evse.DefaultID
		ec.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EvseCreate) check() error {
	if _, ok := ec.mutation.Serial(); !ok {
		return &ValidationError{Name: "serial", err: errors.New(`ent: missing required field "Evse.serial"`)}
	}
	if _, ok := ec.mutation.ConnectorNumber(); !ok {
		return &ValidationError{Name: "connector_number", err: errors.New(`ent: missing required field "Evse.connector_number"`)}
	}
	if _, ok := ec.mutation.EquipmentID(); !ok {
		return &ValidationError{Name: "equipment", err: errors.New(`ent: missing required edge "Evse.equipment"`)}
	}
	return nil
}

func (ec *EvseCreate) sqlSave(ctx context.Context) (*Evse, error) {
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = datasource.UUID(id)
	}
	return _node, nil
}

func (ec *EvseCreate) createSpec() (*Evse, *sqlgraph.CreateSpec) {
	var (
		_node = &Evse{config: ec.config}
		_spec = &sqlgraph.CreateSpec{
			Table: evse.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: evse.FieldID,
			},
		}
	)
	if id, ok := ec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ec.mutation.Serial(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: evse.FieldSerial,
		})
		_node.Serial = value
	}
	if value, ok := ec.mutation.ConnectorNumber(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: evse.FieldConnectorNumber,
		})
		_node.ConnectorNumber = value
	}
	if nodes := ec.mutation.EquipmentIDs(); len(nodes) > 0 {
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
		_node.equipment_evses = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ConnectorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   evse.ConnectorsTable,
			Columns: []string{evse.ConnectorsColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EvseCreateBulk is the builder for creating many Evse entities in bulk.
type EvseCreateBulk struct {
	config
	builders []*EvseCreate
}

// Save creates the Evse entities in the database.
func (ecb *EvseCreateBulk) Save(ctx context.Context) ([]*Evse, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Evse, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EvseMutation)
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
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
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
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = datasource.UUID(id)
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
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EvseCreateBulk) SaveX(ctx context.Context) []*Evse {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EvseCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EvseCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
