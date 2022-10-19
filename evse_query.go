// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/connector"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/evse"
	"github.com/Kotodian/cwmodel/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// EvseQuery is the builder for querying Evse entities.
type EvseQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.Evse
	withEquipment *EquipmentQuery
	withConnector *ConnectorQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EvseQuery builder.
func (eq *EvseQuery) Where(ps ...predicate.Evse) *EvseQuery {
	eq.predicates = append(eq.predicates, ps...)
	return eq
}

// Limit adds a limit step to the query.
func (eq *EvseQuery) Limit(limit int) *EvseQuery {
	eq.limit = &limit
	return eq
}

// Offset adds an offset step to the query.
func (eq *EvseQuery) Offset(offset int) *EvseQuery {
	eq.offset = &offset
	return eq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eq *EvseQuery) Unique(unique bool) *EvseQuery {
	eq.unique = &unique
	return eq
}

// Order adds an order step to the query.
func (eq *EvseQuery) Order(o ...OrderFunc) *EvseQuery {
	eq.order = append(eq.order, o...)
	return eq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (eq *EvseQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(evse.Table, evse.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, evse.EquipmentTable, evse.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryConnector chains the current query on the "connector" edge.
func (eq *EvseQuery) QueryConnector() *ConnectorQuery {
	query := &ConnectorQuery{config: eq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(evse.Table, evse.FieldID, selector),
			sqlgraph.To(connector.Table, connector.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, evse.ConnectorTable, evse.ConnectorColumn),
		)
		fromU = sqlgraph.SetNeighbors(eq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Evse entity from the query.
// Returns a *NotFoundError when no Evse was found.
func (eq *EvseQuery) First(ctx context.Context) (*Evse, error) {
	nodes, err := eq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{evse.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eq *EvseQuery) FirstX(ctx context.Context) *Evse {
	node, err := eq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Evse ID from the query.
// Returns a *NotFoundError when no Evse ID was found.
func (eq *EvseQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = eq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{evse.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eq *EvseQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := eq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Evse entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Evse entity is found.
// Returns a *NotFoundError when no Evse entities are found.
func (eq *EvseQuery) Only(ctx context.Context) (*Evse, error) {
	nodes, err := eq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{evse.Label}
	default:
		return nil, &NotSingularError{evse.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eq *EvseQuery) OnlyX(ctx context.Context) *Evse {
	node, err := eq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Evse ID in the query.
// Returns a *NotSingularError when more than one Evse ID is found.
// Returns a *NotFoundError when no entities are found.
func (eq *EvseQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = eq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{evse.Label}
	default:
		err = &NotSingularError{evse.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eq *EvseQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := eq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Evses.
func (eq *EvseQuery) All(ctx context.Context) ([]*Evse, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eq *EvseQuery) AllX(ctx context.Context) []*Evse {
	nodes, err := eq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Evse IDs.
func (eq *EvseQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := eq.Select(evse.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eq *EvseQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := eq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eq *EvseQuery) Count(ctx context.Context) (int, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eq *EvseQuery) CountX(ctx context.Context) int {
	count, err := eq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eq *EvseQuery) Exist(ctx context.Context) (bool, error) {
	if err := eq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eq *EvseQuery) ExistX(ctx context.Context) bool {
	exist, err := eq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EvseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eq *EvseQuery) Clone() *EvseQuery {
	if eq == nil {
		return nil
	}
	return &EvseQuery{
		config:        eq.config,
		limit:         eq.limit,
		offset:        eq.offset,
		order:         append([]OrderFunc{}, eq.order...),
		predicates:    append([]predicate.Evse{}, eq.predicates...),
		withEquipment: eq.withEquipment.Clone(),
		withConnector: eq.withConnector.Clone(),
		// clone intermediate query.
		sql:    eq.sql.Clone(),
		path:   eq.path,
		unique: eq.unique,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EvseQuery) WithEquipment(opts ...func(*EquipmentQuery)) *EvseQuery {
	query := &EquipmentQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withEquipment = query
	return eq
}

// WithConnector tells the query-builder to eager-load the nodes that are connected to
// the "connector" edge. The optional arguments are used to configure the query builder of the edge.
func (eq *EvseQuery) WithConnector(opts ...func(*ConnectorQuery)) *EvseQuery {
	query := &ConnectorQuery{config: eq.config}
	for _, opt := range opts {
		opt(query)
	}
	eq.withConnector = query
	return eq
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
//	client.Evse.Query().
//		GroupBy(evse.FieldVersion).
//		Aggregate(cwmodel.Count()).
//		Scan(ctx, &v)
func (eq *EvseQuery) GroupBy(field string, fields ...string) *EvseGroupBy {
	grbuild := &EvseGroupBy{config: eq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eq.sqlQuery(ctx), nil
	}
	grbuild.label = evse.Label
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
//	client.Evse.Query().
//		Select(evse.FieldVersion).
//		Scan(ctx, &v)
func (eq *EvseQuery) Select(fields ...string) *EvseSelect {
	eq.fields = append(eq.fields, fields...)
	selbuild := &EvseSelect{EvseQuery: eq}
	selbuild.label = evse.Label
	selbuild.flds, selbuild.scan = &eq.fields, selbuild.Scan
	return selbuild
}

func (eq *EvseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eq.fields {
		if !evse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cwmodel: invalid field %q for query", f)}
		}
	}
	if eq.path != nil {
		prev, err := eq.path(ctx)
		if err != nil {
			return err
		}
		eq.sql = prev
	}
	return nil
}

func (eq *EvseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Evse, error) {
	var (
		nodes       = []*Evse{}
		withFKs     = eq.withFKs
		_spec       = eq.querySpec()
		loadedTypes = [2]bool{
			eq.withEquipment != nil,
			eq.withConnector != nil,
		}
	)
	if eq.withEquipment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, evse.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Evse).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Evse{config: eq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eq.withEquipment; query != nil {
		if err := eq.loadEquipment(ctx, query, nodes, nil,
			func(n *Evse, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	if query := eq.withConnector; query != nil {
		if err := eq.loadConnector(ctx, query, nodes,
			func(n *Evse) { n.Edges.Connector = []*Connector{} },
			func(n *Evse, e *Connector) { n.Edges.Connector = append(n.Edges.Connector, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eq *EvseQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*Evse, init func(*Evse), assign func(*Evse, *Equipment)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*Evse)
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
func (eq *EvseQuery) loadConnector(ctx context.Context, query *ConnectorQuery, nodes []*Evse, init func(*Evse), assign func(*Evse, *Connector)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[datasource.UUID]*Evse)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Connector(func(s *sql.Selector) {
		s.Where(sql.InValues(evse.ConnectorColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.evse_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "evse_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "evse_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (eq *EvseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eq.querySpec()
	_spec.Node.Columns = eq.fields
	if len(eq.fields) > 0 {
		_spec.Unique = eq.unique != nil && *eq.unique
	}
	return sqlgraph.CountNodes(ctx, eq.driver, _spec)
}

func (eq *EvseQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := eq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cwmodel: check existence: %w", err)
	default:
		return true, nil
	}
}

func (eq *EvseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   evse.Table,
			Columns: evse.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: evse.FieldID,
			},
		},
		From:   eq.sql,
		Unique: true,
	}
	if unique := eq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := eq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, evse.FieldID)
		for i := range fields {
			if fields[i] != evse.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eq *EvseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eq.driver.Dialect())
	t1 := builder.Table(evse.Table)
	columns := eq.fields
	if len(columns) == 0 {
		columns = evse.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eq.sql != nil {
		selector = eq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eq.unique != nil && *eq.unique {
		selector.Distinct()
	}
	for _, p := range eq.predicates {
		p(selector)
	}
	for _, p := range eq.order {
		p(selector)
	}
	if offset := eq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EvseGroupBy is the group-by builder for Evse entities.
type EvseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (egb *EvseGroupBy) Aggregate(fns ...AggregateFunc) *EvseGroupBy {
	egb.fns = append(egb.fns, fns...)
	return egb
}

// Scan applies the group-by query and scans the result into the given value.
func (egb *EvseGroupBy) Scan(ctx context.Context, v any) error {
	query, err := egb.path(ctx)
	if err != nil {
		return err
	}
	egb.sql = query
	return egb.sqlScan(ctx, v)
}

func (egb *EvseGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range egb.fields {
		if !evse.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := egb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := egb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (egb *EvseGroupBy) sqlQuery() *sql.Selector {
	selector := egb.sql.Select()
	aggregation := make([]string, 0, len(egb.fns))
	for _, fn := range egb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(egb.fields)+len(egb.fns))
		for _, f := range egb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(egb.fields...)...)
}

// EvseSelect is the builder for selecting fields of Evse entities.
type EvseSelect struct {
	*EvseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (es *EvseSelect) Scan(ctx context.Context, v any) error {
	if err := es.prepareQuery(ctx); err != nil {
		return err
	}
	es.sql = es.EvseQuery.sqlQuery(ctx)
	return es.sqlScan(ctx, v)
}

func (es *EvseSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := es.sql.Query()
	if err := es.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}