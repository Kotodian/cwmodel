// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/ent/equipment"
	"github.com/Kotodian/cwmodel/ent/equipmentlog"
	"github.com/Kotodian/cwmodel/ent/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentLogQuery is the builder for querying EquipmentLog entities.
type EquipmentLogQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.EquipmentLog
	withEquipment *EquipmentQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EquipmentLogQuery builder.
func (elq *EquipmentLogQuery) Where(ps ...predicate.EquipmentLog) *EquipmentLogQuery {
	elq.predicates = append(elq.predicates, ps...)
	return elq
}

// Limit adds a limit step to the query.
func (elq *EquipmentLogQuery) Limit(limit int) *EquipmentLogQuery {
	elq.limit = &limit
	return elq
}

// Offset adds an offset step to the query.
func (elq *EquipmentLogQuery) Offset(offset int) *EquipmentLogQuery {
	elq.offset = &offset
	return elq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (elq *EquipmentLogQuery) Unique(unique bool) *EquipmentLogQuery {
	elq.unique = &unique
	return elq
}

// Order adds an order step to the query.
func (elq *EquipmentLogQuery) Order(o ...OrderFunc) *EquipmentLogQuery {
	elq.order = append(elq.order, o...)
	return elq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (elq *EquipmentLogQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: elq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := elq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipmentlog.Table, equipmentlog.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, equipmentlog.EquipmentTable, equipmentlog.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(elq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EquipmentLog entity from the query.
// Returns a *NotFoundError when no EquipmentLog was found.
func (elq *EquipmentLogQuery) First(ctx context.Context) (*EquipmentLog, error) {
	nodes, err := elq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{equipmentlog.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (elq *EquipmentLogQuery) FirstX(ctx context.Context) *EquipmentLog {
	node, err := elq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EquipmentLog ID from the query.
// Returns a *NotFoundError when no EquipmentLog ID was found.
func (elq *EquipmentLogQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = elq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{equipmentlog.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (elq *EquipmentLogQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := elq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EquipmentLog entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EquipmentLog entity is found.
// Returns a *NotFoundError when no EquipmentLog entities are found.
func (elq *EquipmentLogQuery) Only(ctx context.Context) (*EquipmentLog, error) {
	nodes, err := elq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{equipmentlog.Label}
	default:
		return nil, &NotSingularError{equipmentlog.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (elq *EquipmentLogQuery) OnlyX(ctx context.Context) *EquipmentLog {
	node, err := elq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EquipmentLog ID in the query.
// Returns a *NotSingularError when more than one EquipmentLog ID is found.
// Returns a *NotFoundError when no entities are found.
func (elq *EquipmentLogQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = elq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{equipmentlog.Label}
	default:
		err = &NotSingularError{equipmentlog.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (elq *EquipmentLogQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := elq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentLogs.
func (elq *EquipmentLogQuery) All(ctx context.Context) ([]*EquipmentLog, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return elq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (elq *EquipmentLogQuery) AllX(ctx context.Context) []*EquipmentLog {
	nodes, err := elq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EquipmentLog IDs.
func (elq *EquipmentLogQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := elq.Select(equipmentlog.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (elq *EquipmentLogQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := elq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (elq *EquipmentLogQuery) Count(ctx context.Context) (int, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return elq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (elq *EquipmentLogQuery) CountX(ctx context.Context) int {
	count, err := elq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (elq *EquipmentLogQuery) Exist(ctx context.Context) (bool, error) {
	if err := elq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return elq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (elq *EquipmentLogQuery) ExistX(ctx context.Context) bool {
	exist, err := elq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EquipmentLogQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (elq *EquipmentLogQuery) Clone() *EquipmentLogQuery {
	if elq == nil {
		return nil
	}
	return &EquipmentLogQuery{
		config:        elq.config,
		limit:         elq.limit,
		offset:        elq.offset,
		order:         append([]OrderFunc{}, elq.order...),
		predicates:    append([]predicate.EquipmentLog{}, elq.predicates...),
		withEquipment: elq.withEquipment.Clone(),
		// clone intermediate query.
		sql:    elq.sql.Clone(),
		path:   elq.path,
		unique: elq.unique,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (elq *EquipmentLogQuery) WithEquipment(opts ...func(*EquipmentQuery)) *EquipmentLogQuery {
	query := &EquipmentQuery{config: elq.config}
	for _, opt := range opts {
		opt(query)
	}
	elq.withEquipment = query
	return elq
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
//	client.EquipmentLog.Query().
//		GroupBy(equipmentlog.FieldVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (elq *EquipmentLogQuery) GroupBy(field string, fields ...string) *EquipmentLogGroupBy {
	grbuild := &EquipmentLogGroupBy{config: elq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := elq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return elq.sqlQuery(ctx), nil
	}
	grbuild.label = equipmentlog.Label
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
//	client.EquipmentLog.Query().
//		Select(equipmentlog.FieldVersion).
//		Scan(ctx, &v)
func (elq *EquipmentLogQuery) Select(fields ...string) *EquipmentLogSelect {
	elq.fields = append(elq.fields, fields...)
	selbuild := &EquipmentLogSelect{EquipmentLogQuery: elq}
	selbuild.label = equipmentlog.Label
	selbuild.flds, selbuild.scan = &elq.fields, selbuild.Scan
	return selbuild
}

func (elq *EquipmentLogQuery) prepareQuery(ctx context.Context) error {
	for _, f := range elq.fields {
		if !equipmentlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if elq.path != nil {
		prev, err := elq.path(ctx)
		if err != nil {
			return err
		}
		elq.sql = prev
	}
	return nil
}

func (elq *EquipmentLogQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EquipmentLog, error) {
	var (
		nodes       = []*EquipmentLog{}
		withFKs     = elq.withFKs
		_spec       = elq.querySpec()
		loadedTypes = [1]bool{
			elq.withEquipment != nil,
		}
	)
	if elq.withEquipment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentlog.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EquipmentLog).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EquipmentLog{config: elq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, elq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := elq.withEquipment; query != nil {
		if err := elq.loadEquipment(ctx, query, nodes, nil,
			func(n *EquipmentLog, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (elq *EquipmentLogQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*EquipmentLog, init func(*EquipmentLog), assign func(*EquipmentLog, *Equipment)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*EquipmentLog)
	for i := range nodes {
		if nodes[i].equipment_id == nil {
			continue
		}
		fk := *nodes[i].equipment_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(equipment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "equipment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (elq *EquipmentLogQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := elq.querySpec()
	_spec.Node.Columns = elq.fields
	if len(elq.fields) > 0 {
		_spec.Unique = elq.unique != nil && *elq.unique
	}
	return sqlgraph.CountNodes(ctx, elq.driver, _spec)
}

func (elq *EquipmentLogQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := elq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (elq *EquipmentLogQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentlog.Table,
			Columns: equipmentlog.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentlog.FieldID,
			},
		},
		From:   elq.sql,
		Unique: true,
	}
	if unique := elq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := elq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentlog.FieldID)
		for i := range fields {
			if fields[i] != equipmentlog.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := elq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := elq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := elq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := elq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (elq *EquipmentLogQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(elq.driver.Dialect())
	t1 := builder.Table(equipmentlog.Table)
	columns := elq.fields
	if len(columns) == 0 {
		columns = equipmentlog.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if elq.sql != nil {
		selector = elq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if elq.unique != nil && *elq.unique {
		selector.Distinct()
	}
	for _, p := range elq.predicates {
		p(selector)
	}
	for _, p := range elq.order {
		p(selector)
	}
	if offset := elq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := elq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentLogGroupBy is the group-by builder for EquipmentLog entities.
type EquipmentLogGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (elgb *EquipmentLogGroupBy) Aggregate(fns ...AggregateFunc) *EquipmentLogGroupBy {
	elgb.fns = append(elgb.fns, fns...)
	return elgb
}

// Scan applies the group-by query and scans the result into the given value.
func (elgb *EquipmentLogGroupBy) Scan(ctx context.Context, v any) error {
	query, err := elgb.path(ctx)
	if err != nil {
		return err
	}
	elgb.sql = query
	return elgb.sqlScan(ctx, v)
}

func (elgb *EquipmentLogGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range elgb.fields {
		if !equipmentlog.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := elgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := elgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (elgb *EquipmentLogGroupBy) sqlQuery() *sql.Selector {
	selector := elgb.sql.Select()
	aggregation := make([]string, 0, len(elgb.fns))
	for _, fn := range elgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(elgb.fields)+len(elgb.fns))
		for _, f := range elgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(elgb.fields...)...)
}

// EquipmentLogSelect is the builder for selecting fields of EquipmentLog entities.
type EquipmentLogSelect struct {
	*EquipmentLogQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (els *EquipmentLogSelect) Scan(ctx context.Context, v any) error {
	if err := els.prepareQuery(ctx); err != nil {
		return err
	}
	els.sql = els.EquipmentLogQuery.sqlQuery(ctx)
	return els.sqlScan(ctx, v)
}

func (els *EquipmentLogSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := els.sql.Query()
	if err := els.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
