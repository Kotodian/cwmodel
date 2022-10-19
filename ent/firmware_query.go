// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/equipmentfirmwareeffect"
	"github.com/Kotodian/ent-practice/ent/firmware"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// FirmwareQuery is the builder for querying Firmware entities.
type FirmwareQuery struct {
	config
	limit                       *int
	offset                      *int
	unique                      *bool
	order                       []OrderFunc
	fields                      []string
	predicates                  []predicate.Firmware
	withEquipmentFirmwareEffect *EquipmentFirmwareEffectQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FirmwareQuery builder.
func (fq *FirmwareQuery) Where(ps ...predicate.Firmware) *FirmwareQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit adds a limit step to the query.
func (fq *FirmwareQuery) Limit(limit int) *FirmwareQuery {
	fq.limit = &limit
	return fq
}

// Offset adds an offset step to the query.
func (fq *FirmwareQuery) Offset(offset int) *FirmwareQuery {
	fq.offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FirmwareQuery) Unique(unique bool) *FirmwareQuery {
	fq.unique = &unique
	return fq
}

// Order adds an order step to the query.
func (fq *FirmwareQuery) Order(o ...OrderFunc) *FirmwareQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryEquipmentFirmwareEffect chains the current query on the "equipment_firmware_effect" edge.
func (fq *FirmwareQuery) QueryEquipmentFirmwareEffect() *EquipmentFirmwareEffectQuery {
	query := &EquipmentFirmwareEffectQuery{config: fq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(firmware.Table, firmware.FieldID, selector),
			sqlgraph.To(equipmentfirmwareeffect.Table, equipmentfirmwareeffect.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, firmware.EquipmentFirmwareEffectTable, firmware.EquipmentFirmwareEffectColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Firmware entity from the query.
// Returns a *NotFoundError when no Firmware was found.
func (fq *FirmwareQuery) First(ctx context.Context) (*Firmware, error) {
	nodes, err := fq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{firmware.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FirmwareQuery) FirstX(ctx context.Context) *Firmware {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Firmware ID from the query.
// Returns a *NotFoundError when no Firmware ID was found.
func (fq *FirmwareQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = fq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{firmware.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FirmwareQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Firmware entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Firmware entity is found.
// Returns a *NotFoundError when no Firmware entities are found.
func (fq *FirmwareQuery) Only(ctx context.Context) (*Firmware, error) {
	nodes, err := fq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{firmware.Label}
	default:
		return nil, &NotSingularError{firmware.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FirmwareQuery) OnlyX(ctx context.Context) *Firmware {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Firmware ID in the query.
// Returns a *NotSingularError when more than one Firmware ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FirmwareQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = fq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{firmware.Label}
	default:
		err = &NotSingularError{firmware.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FirmwareQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Firmwares.
func (fq *FirmwareQuery) All(ctx context.Context) ([]*Firmware, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return fq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (fq *FirmwareQuery) AllX(ctx context.Context) []*Firmware {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Firmware IDs.
func (fq *FirmwareQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := fq.Select(firmware.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FirmwareQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FirmwareQuery) Count(ctx context.Context) (int, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return fq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FirmwareQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FirmwareQuery) Exist(ctx context.Context) (bool, error) {
	if err := fq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return fq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FirmwareQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FirmwareQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FirmwareQuery) Clone() *FirmwareQuery {
	if fq == nil {
		return nil
	}
	return &FirmwareQuery{
		config:                      fq.config,
		limit:                       fq.limit,
		offset:                      fq.offset,
		order:                       append([]OrderFunc{}, fq.order...),
		predicates:                  append([]predicate.Firmware{}, fq.predicates...),
		withEquipmentFirmwareEffect: fq.withEquipmentFirmwareEffect.Clone(),
		// clone intermediate query.
		sql:    fq.sql.Clone(),
		path:   fq.path,
		unique: fq.unique,
	}
}

// WithEquipmentFirmwareEffect tells the query-builder to eager-load the nodes that are connected to
// the "equipment_firmware_effect" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FirmwareQuery) WithEquipmentFirmwareEffect(opts ...func(*EquipmentFirmwareEffectQuery)) *FirmwareQuery {
	query := &EquipmentFirmwareEffectQuery{config: fq.config}
	for _, opt := range opts {
		opt(query)
	}
	fq.withEquipmentFirmwareEffect = query
	return fq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Version int64 `json:"version,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Firmware.Query().
//		GroupBy(firmware.FieldVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FirmwareQuery) GroupBy(field string, fields ...string) *FirmwareGroupBy {
	grbuild := &FirmwareGroupBy{config: fq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return fq.sqlQuery(ctx), nil
	}
	grbuild.label = firmware.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Version int64 `json:"version,omitempty"`
//	}
//
//	client.Firmware.Query().
//		Select(firmware.FieldVersion).
//		Scan(ctx, &v)
func (fq *FirmwareQuery) Select(fields ...string) *FirmwareSelect {
	fq.fields = append(fq.fields, fields...)
	selbuild := &FirmwareSelect{FirmwareQuery: fq}
	selbuild.label = firmware.Label
	selbuild.flds, selbuild.scan = &fq.fields, selbuild.Scan
	return selbuild
}

func (fq *FirmwareQuery) prepareQuery(ctx context.Context) error {
	for _, f := range fq.fields {
		if !firmware.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FirmwareQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Firmware, error) {
	var (
		nodes       = []*Firmware{}
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withEquipmentFirmwareEffect != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Firmware).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Firmware{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withEquipmentFirmwareEffect; query != nil {
		if err := fq.loadEquipmentFirmwareEffect(ctx, query, nodes,
			func(n *Firmware) { n.Edges.EquipmentFirmwareEffect = []*EquipmentFirmwareEffect{} },
			func(n *Firmware, e *EquipmentFirmwareEffect) {
				n.Edges.EquipmentFirmwareEffect = append(n.Edges.EquipmentFirmwareEffect, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FirmwareQuery) loadEquipmentFirmwareEffect(ctx context.Context, query *EquipmentFirmwareEffectQuery, nodes []*Firmware, init func(*Firmware), assign func(*Firmware, *EquipmentFirmwareEffect)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[datasource.UUID]*Firmware)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.EquipmentFirmwareEffect(func(s *sql.Selector) {
		s.Where(sql.InValues(firmware.EquipmentFirmwareEffectColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.firmware_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "firmware_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "firmware_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fq *FirmwareQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.fields
	if len(fq.fields) > 0 {
		_spec.Unique = fq.unique != nil && *fq.unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FirmwareQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (fq *FirmwareQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   firmware.Table,
			Columns: firmware.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: firmware.FieldID,
			},
		},
		From:   fq.sql,
		Unique: true,
	}
	if unique := fq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := fq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, firmware.FieldID)
		for i := range fields {
			if fields[i] != firmware.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FirmwareQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(firmware.Table)
	columns := fq.fields
	if len(columns) == 0 {
		columns = firmware.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.unique != nil && *fq.unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FirmwareGroupBy is the group-by builder for Firmware entities.
type FirmwareGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FirmwareGroupBy) Aggregate(fns ...AggregateFunc) *FirmwareGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the group-by query and scans the result into the given value.
func (fgb *FirmwareGroupBy) Scan(ctx context.Context, v any) error {
	query, err := fgb.path(ctx)
	if err != nil {
		return err
	}
	fgb.sql = query
	return fgb.sqlScan(ctx, v)
}

func (fgb *FirmwareGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range fgb.fields {
		if !firmware.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := fgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (fgb *FirmwareGroupBy) sqlQuery() *sql.Selector {
	selector := fgb.sql.Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(fgb.fields)+len(fgb.fns))
		for _, f := range fgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(fgb.fields...)...)
}

// FirmwareSelect is the builder for selecting fields of Firmware entities.
type FirmwareSelect struct {
	*FirmwareQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FirmwareSelect) Scan(ctx context.Context, v any) error {
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	fs.sql = fs.FirmwareQuery.sqlQuery(ctx)
	return fs.sqlScan(ctx, v)
}

func (fs *FirmwareSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := fs.sql.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}