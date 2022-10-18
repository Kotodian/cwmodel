// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/equipmentfirmwareeffect"
	"github.com/Kotodian/ent-practice/ent/firmware"
	"github.com/Kotodian/ent-practice/ent/predicate"
)

// EquipmentFirmwareEffectQuery is the builder for querying EquipmentFirmwareEffect entities.
type EquipmentFirmwareEffectQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.EquipmentFirmwareEffect
	// eager-loading edges.
	withEquipment *EquipmentQuery
	withFirmware  *FirmwareQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EquipmentFirmwareEffectQuery builder.
func (efeq *EquipmentFirmwareEffectQuery) Where(ps ...predicate.EquipmentFirmwareEffect) *EquipmentFirmwareEffectQuery {
	efeq.predicates = append(efeq.predicates, ps...)
	return efeq
}

// Limit adds a limit step to the query.
func (efeq *EquipmentFirmwareEffectQuery) Limit(limit int) *EquipmentFirmwareEffectQuery {
	efeq.limit = &limit
	return efeq
}

// Offset adds an offset step to the query.
func (efeq *EquipmentFirmwareEffectQuery) Offset(offset int) *EquipmentFirmwareEffectQuery {
	efeq.offset = &offset
	return efeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (efeq *EquipmentFirmwareEffectQuery) Unique(unique bool) *EquipmentFirmwareEffectQuery {
	efeq.unique = &unique
	return efeq
}

// Order adds an order step to the query.
func (efeq *EquipmentFirmwareEffectQuery) Order(o ...OrderFunc) *EquipmentFirmwareEffectQuery {
	efeq.order = append(efeq.order, o...)
	return efeq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (efeq *EquipmentFirmwareEffectQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: efeq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := efeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := efeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipmentfirmwareeffect.Table, equipmentfirmwareeffect.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, equipmentfirmwareeffect.EquipmentTable, equipmentfirmwareeffect.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(efeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFirmware chains the current query on the "firmware" edge.
func (efeq *EquipmentFirmwareEffectQuery) QueryFirmware() *FirmwareQuery {
	query := &FirmwareQuery{config: efeq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := efeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := efeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipmentfirmwareeffect.Table, equipmentfirmwareeffect.FieldID, selector),
			sqlgraph.To(firmware.Table, firmware.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, equipmentfirmwareeffect.FirmwareTable, equipmentfirmwareeffect.FirmwareColumn),
		)
		fromU = sqlgraph.SetNeighbors(efeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EquipmentFirmwareEffect entity from the query.
// Returns a *NotFoundError when no EquipmentFirmwareEffect was found.
func (efeq *EquipmentFirmwareEffectQuery) First(ctx context.Context) (*EquipmentFirmwareEffect, error) {
	nodes, err := efeq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{equipmentfirmwareeffect.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (efeq *EquipmentFirmwareEffectQuery) FirstX(ctx context.Context) *EquipmentFirmwareEffect {
	node, err := efeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EquipmentFirmwareEffect ID from the query.
// Returns a *NotFoundError when no EquipmentFirmwareEffect ID was found.
func (efeq *EquipmentFirmwareEffectQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = efeq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{equipmentfirmwareeffect.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (efeq *EquipmentFirmwareEffectQuery) FirstIDX(ctx context.Context) int {
	id, err := efeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EquipmentFirmwareEffect entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EquipmentFirmwareEffect entity is found.
// Returns a *NotFoundError when no EquipmentFirmwareEffect entities are found.
func (efeq *EquipmentFirmwareEffectQuery) Only(ctx context.Context) (*EquipmentFirmwareEffect, error) {
	nodes, err := efeq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		return nil, &NotSingularError{equipmentfirmwareeffect.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (efeq *EquipmentFirmwareEffectQuery) OnlyX(ctx context.Context) *EquipmentFirmwareEffect {
	node, err := efeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EquipmentFirmwareEffect ID in the query.
// Returns a *NotSingularError when more than one EquipmentFirmwareEffect ID is found.
// Returns a *NotFoundError when no entities are found.
func (efeq *EquipmentFirmwareEffectQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = efeq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = &NotSingularError{equipmentfirmwareeffect.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (efeq *EquipmentFirmwareEffectQuery) OnlyIDX(ctx context.Context) int {
	id, err := efeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentFirmwareEffects.
func (efeq *EquipmentFirmwareEffectQuery) All(ctx context.Context) ([]*EquipmentFirmwareEffect, error) {
	if err := efeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return efeq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (efeq *EquipmentFirmwareEffectQuery) AllX(ctx context.Context) []*EquipmentFirmwareEffect {
	nodes, err := efeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EquipmentFirmwareEffect IDs.
func (efeq *EquipmentFirmwareEffectQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := efeq.Select(equipmentfirmwareeffect.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (efeq *EquipmentFirmwareEffectQuery) IDsX(ctx context.Context) []int {
	ids, err := efeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (efeq *EquipmentFirmwareEffectQuery) Count(ctx context.Context) (int, error) {
	if err := efeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return efeq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (efeq *EquipmentFirmwareEffectQuery) CountX(ctx context.Context) int {
	count, err := efeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (efeq *EquipmentFirmwareEffectQuery) Exist(ctx context.Context) (bool, error) {
	if err := efeq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return efeq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (efeq *EquipmentFirmwareEffectQuery) ExistX(ctx context.Context) bool {
	exist, err := efeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EquipmentFirmwareEffectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (efeq *EquipmentFirmwareEffectQuery) Clone() *EquipmentFirmwareEffectQuery {
	if efeq == nil {
		return nil
	}
	return &EquipmentFirmwareEffectQuery{
		config:        efeq.config,
		limit:         efeq.limit,
		offset:        efeq.offset,
		order:         append([]OrderFunc{}, efeq.order...),
		predicates:    append([]predicate.EquipmentFirmwareEffect{}, efeq.predicates...),
		withEquipment: efeq.withEquipment.Clone(),
		withFirmware:  efeq.withFirmware.Clone(),
		// clone intermediate query.
		sql:    efeq.sql.Clone(),
		path:   efeq.path,
		unique: efeq.unique,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (efeq *EquipmentFirmwareEffectQuery) WithEquipment(opts ...func(*EquipmentQuery)) *EquipmentFirmwareEffectQuery {
	query := &EquipmentQuery{config: efeq.config}
	for _, opt := range opts {
		opt(query)
	}
	efeq.withEquipment = query
	return efeq
}

// WithFirmware tells the query-builder to eager-load the nodes that are connected to
// the "firmware" edge. The optional arguments are used to configure the query builder of the edge.
func (efeq *EquipmentFirmwareEffectQuery) WithFirmware(opts ...func(*FirmwareQuery)) *EquipmentFirmwareEffectQuery {
	query := &FirmwareQuery{config: efeq.config}
	for _, opt := range opts {
		opt(query)
	}
	efeq.withFirmware = query
	return efeq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		RequestID int64 `json:"request_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EquipmentFirmwareEffect.Query().
//		GroupBy(equipmentfirmwareeffect.FieldRequestID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (efeq *EquipmentFirmwareEffectQuery) GroupBy(field string, fields ...string) *EquipmentFirmwareEffectGroupBy {
	group := &EquipmentFirmwareEffectGroupBy{config: efeq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := efeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return efeq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		RequestID int64 `json:"request_id,omitempty"`
//	}
//
//	client.EquipmentFirmwareEffect.Query().
//		Select(equipmentfirmwareeffect.FieldRequestID).
//		Scan(ctx, &v)
func (efeq *EquipmentFirmwareEffectQuery) Select(fields ...string) *EquipmentFirmwareEffectSelect {
	efeq.fields = append(efeq.fields, fields...)
	return &EquipmentFirmwareEffectSelect{EquipmentFirmwareEffectQuery: efeq}
}

func (efeq *EquipmentFirmwareEffectQuery) prepareQuery(ctx context.Context) error {
	for _, f := range efeq.fields {
		if !equipmentfirmwareeffect.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if efeq.path != nil {
		prev, err := efeq.path(ctx)
		if err != nil {
			return err
		}
		efeq.sql = prev
	}
	return nil
}

func (efeq *EquipmentFirmwareEffectQuery) sqlAll(ctx context.Context) ([]*EquipmentFirmwareEffect, error) {
	var (
		nodes       = []*EquipmentFirmwareEffect{}
		withFKs     = efeq.withFKs
		_spec       = efeq.querySpec()
		loadedTypes = [2]bool{
			efeq.withEquipment != nil,
			efeq.withFirmware != nil,
		}
	)
	if efeq.withEquipment != nil || efeq.withFirmware != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentfirmwareeffect.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &EquipmentFirmwareEffect{config: efeq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, efeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := efeq.withEquipment; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*EquipmentFirmwareEffect)
		for i := range nodes {
			if nodes[i].equipment_equipment_firmware_effect == nil {
				continue
			}
			fk := *nodes[i].equipment_equipment_firmware_effect
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(equipment.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "equipment_equipment_firmware_effect" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Equipment = n
			}
		}
	}

	if query := efeq.withFirmware; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*EquipmentFirmwareEffect)
		for i := range nodes {
			if nodes[i].firmware_equipment_firmware_effect == nil {
				continue
			}
			fk := *nodes[i].firmware_equipment_firmware_effect
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(firmware.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "firmware_equipment_firmware_effect" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Firmware = n
			}
		}
	}

	return nodes, nil
}

func (efeq *EquipmentFirmwareEffectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := efeq.querySpec()
	_spec.Node.Columns = efeq.fields
	if len(efeq.fields) > 0 {
		_spec.Unique = efeq.unique != nil && *efeq.unique
	}
	return sqlgraph.CountNodes(ctx, efeq.driver, _spec)
}

func (efeq *EquipmentFirmwareEffectQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := efeq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (efeq *EquipmentFirmwareEffectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentfirmwareeffect.Table,
			Columns: equipmentfirmwareeffect.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: equipmentfirmwareeffect.FieldID,
			},
		},
		From:   efeq.sql,
		Unique: true,
	}
	if unique := efeq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := efeq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentfirmwareeffect.FieldID)
		for i := range fields {
			if fields[i] != equipmentfirmwareeffect.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := efeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := efeq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := efeq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := efeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (efeq *EquipmentFirmwareEffectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(efeq.driver.Dialect())
	t1 := builder.Table(equipmentfirmwareeffect.Table)
	columns := efeq.fields
	if len(columns) == 0 {
		columns = equipmentfirmwareeffect.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if efeq.sql != nil {
		selector = efeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if efeq.unique != nil && *efeq.unique {
		selector.Distinct()
	}
	for _, p := range efeq.predicates {
		p(selector)
	}
	for _, p := range efeq.order {
		p(selector)
	}
	if offset := efeq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := efeq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentFirmwareEffectGroupBy is the group-by builder for EquipmentFirmwareEffect entities.
type EquipmentFirmwareEffectGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (efegb *EquipmentFirmwareEffectGroupBy) Aggregate(fns ...AggregateFunc) *EquipmentFirmwareEffectGroupBy {
	efegb.fns = append(efegb.fns, fns...)
	return efegb
}

// Scan applies the group-by query and scans the result into the given value.
func (efegb *EquipmentFirmwareEffectGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := efegb.path(ctx)
	if err != nil {
		return err
	}
	efegb.sql = query
	return efegb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := efegb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (efegb *EquipmentFirmwareEffectGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(efegb.fields) > 1 {
		return nil, errors.New("ent: EquipmentFirmwareEffectGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := efegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) StringsX(ctx context.Context) []string {
	v, err := efegb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (efegb *EquipmentFirmwareEffectGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = efegb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = fmt.Errorf("ent: EquipmentFirmwareEffectGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) StringX(ctx context.Context) string {
	v, err := efegb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (efegb *EquipmentFirmwareEffectGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(efegb.fields) > 1 {
		return nil, errors.New("ent: EquipmentFirmwareEffectGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := efegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) IntsX(ctx context.Context) []int {
	v, err := efegb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (efegb *EquipmentFirmwareEffectGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = efegb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = fmt.Errorf("ent: EquipmentFirmwareEffectGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) IntX(ctx context.Context) int {
	v, err := efegb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (efegb *EquipmentFirmwareEffectGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(efegb.fields) > 1 {
		return nil, errors.New("ent: EquipmentFirmwareEffectGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := efegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := efegb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (efegb *EquipmentFirmwareEffectGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = efegb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = fmt.Errorf("ent: EquipmentFirmwareEffectGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) Float64X(ctx context.Context) float64 {
	v, err := efegb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (efegb *EquipmentFirmwareEffectGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(efegb.fields) > 1 {
		return nil, errors.New("ent: EquipmentFirmwareEffectGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := efegb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := efegb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (efegb *EquipmentFirmwareEffectGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = efegb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = fmt.Errorf("ent: EquipmentFirmwareEffectGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (efegb *EquipmentFirmwareEffectGroupBy) BoolX(ctx context.Context) bool {
	v, err := efegb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (efegb *EquipmentFirmwareEffectGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range efegb.fields {
		if !equipmentfirmwareeffect.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := efegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := efegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (efegb *EquipmentFirmwareEffectGroupBy) sqlQuery() *sql.Selector {
	selector := efegb.sql.Select()
	aggregation := make([]string, 0, len(efegb.fns))
	for _, fn := range efegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(efegb.fields)+len(efegb.fns))
		for _, f := range efegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(efegb.fields...)...)
}

// EquipmentFirmwareEffectSelect is the builder for selecting fields of EquipmentFirmwareEffect entities.
type EquipmentFirmwareEffectSelect struct {
	*EquipmentFirmwareEffectQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (efes *EquipmentFirmwareEffectSelect) Scan(ctx context.Context, v interface{}) error {
	if err := efes.prepareQuery(ctx); err != nil {
		return err
	}
	efes.sql = efes.EquipmentFirmwareEffectQuery.sqlQuery(ctx)
	return efes.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) ScanX(ctx context.Context, v interface{}) {
	if err := efes.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (efes *EquipmentFirmwareEffectSelect) Strings(ctx context.Context) ([]string, error) {
	if len(efes.fields) > 1 {
		return nil, errors.New("ent: EquipmentFirmwareEffectSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := efes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) StringsX(ctx context.Context) []string {
	v, err := efes.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (efes *EquipmentFirmwareEffectSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = efes.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = fmt.Errorf("ent: EquipmentFirmwareEffectSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) StringX(ctx context.Context) string {
	v, err := efes.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (efes *EquipmentFirmwareEffectSelect) Ints(ctx context.Context) ([]int, error) {
	if len(efes.fields) > 1 {
		return nil, errors.New("ent: EquipmentFirmwareEffectSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := efes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) IntsX(ctx context.Context) []int {
	v, err := efes.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (efes *EquipmentFirmwareEffectSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = efes.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = fmt.Errorf("ent: EquipmentFirmwareEffectSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) IntX(ctx context.Context) int {
	v, err := efes.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (efes *EquipmentFirmwareEffectSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(efes.fields) > 1 {
		return nil, errors.New("ent: EquipmentFirmwareEffectSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := efes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) Float64sX(ctx context.Context) []float64 {
	v, err := efes.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (efes *EquipmentFirmwareEffectSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = efes.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = fmt.Errorf("ent: EquipmentFirmwareEffectSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) Float64X(ctx context.Context) float64 {
	v, err := efes.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (efes *EquipmentFirmwareEffectSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(efes.fields) > 1 {
		return nil, errors.New("ent: EquipmentFirmwareEffectSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := efes.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) BoolsX(ctx context.Context) []bool {
	v, err := efes.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (efes *EquipmentFirmwareEffectSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = efes.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{equipmentfirmwareeffect.Label}
	default:
		err = fmt.Errorf("ent: EquipmentFirmwareEffectSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (efes *EquipmentFirmwareEffectSelect) BoolX(ctx context.Context) bool {
	v, err := efes.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (efes *EquipmentFirmwareEffectSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := efes.sql.Query()
	if err := efes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
