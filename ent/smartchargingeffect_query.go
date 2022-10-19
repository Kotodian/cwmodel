// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/ent-practice/ent/connector"
	"github.com/Kotodian/ent-practice/ent/equipment"
	"github.com/Kotodian/ent-practice/ent/orderinfo"
	"github.com/Kotodian/ent-practice/ent/predicate"
	"github.com/Kotodian/ent-practice/ent/smartchargingeffect"
	"github.com/Kotodian/gokit/datasource"
)

// SmartChargingEffectQuery is the builder for querying SmartChargingEffect entities.
type SmartChargingEffectQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.SmartChargingEffect
	withEquipment *EquipmentQuery
	withConnector *ConnectorQuery
	withOrderInfo *OrderInfoQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SmartChargingEffectQuery builder.
func (sceq *SmartChargingEffectQuery) Where(ps ...predicate.SmartChargingEffect) *SmartChargingEffectQuery {
	sceq.predicates = append(sceq.predicates, ps...)
	return sceq
}

// Limit adds a limit step to the query.
func (sceq *SmartChargingEffectQuery) Limit(limit int) *SmartChargingEffectQuery {
	sceq.limit = &limit
	return sceq
}

// Offset adds an offset step to the query.
func (sceq *SmartChargingEffectQuery) Offset(offset int) *SmartChargingEffectQuery {
	sceq.offset = &offset
	return sceq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sceq *SmartChargingEffectQuery) Unique(unique bool) *SmartChargingEffectQuery {
	sceq.unique = &unique
	return sceq
}

// Order adds an order step to the query.
func (sceq *SmartChargingEffectQuery) Order(o ...OrderFunc) *SmartChargingEffectQuery {
	sceq.order = append(sceq.order, o...)
	return sceq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (sceq *SmartChargingEffectQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: sceq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sceq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(smartchargingeffect.Table, smartchargingeffect.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, smartchargingeffect.EquipmentTable, smartchargingeffect.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(sceq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryConnector chains the current query on the "connector" edge.
func (sceq *SmartChargingEffectQuery) QueryConnector() *ConnectorQuery {
	query := &ConnectorQuery{config: sceq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sceq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(smartchargingeffect.Table, smartchargingeffect.FieldID, selector),
			sqlgraph.To(connector.Table, connector.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, smartchargingeffect.ConnectorTable, smartchargingeffect.ConnectorColumn),
		)
		fromU = sqlgraph.SetNeighbors(sceq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderInfo chains the current query on the "order_info" edge.
func (sceq *SmartChargingEffectQuery) QueryOrderInfo() *OrderInfoQuery {
	query := &OrderInfoQuery{config: sceq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sceq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(smartchargingeffect.Table, smartchargingeffect.FieldID, selector),
			sqlgraph.To(orderinfo.Table, orderinfo.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, smartchargingeffect.OrderInfoTable, smartchargingeffect.OrderInfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(sceq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first SmartChargingEffect entity from the query.
// Returns a *NotFoundError when no SmartChargingEffect was found.
func (sceq *SmartChargingEffectQuery) First(ctx context.Context) (*SmartChargingEffect, error) {
	nodes, err := sceq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{smartchargingeffect.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sceq *SmartChargingEffectQuery) FirstX(ctx context.Context) *SmartChargingEffect {
	node, err := sceq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SmartChargingEffect ID from the query.
// Returns a *NotFoundError when no SmartChargingEffect ID was found.
func (sceq *SmartChargingEffectQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = sceq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{smartchargingeffect.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sceq *SmartChargingEffectQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := sceq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SmartChargingEffect entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SmartChargingEffect entity is found.
// Returns a *NotFoundError when no SmartChargingEffect entities are found.
func (sceq *SmartChargingEffectQuery) Only(ctx context.Context) (*SmartChargingEffect, error) {
	nodes, err := sceq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{smartchargingeffect.Label}
	default:
		return nil, &NotSingularError{smartchargingeffect.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sceq *SmartChargingEffectQuery) OnlyX(ctx context.Context) *SmartChargingEffect {
	node, err := sceq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SmartChargingEffect ID in the query.
// Returns a *NotSingularError when more than one SmartChargingEffect ID is found.
// Returns a *NotFoundError when no entities are found.
func (sceq *SmartChargingEffectQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = sceq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{smartchargingeffect.Label}
	default:
		err = &NotSingularError{smartchargingeffect.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sceq *SmartChargingEffectQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := sceq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SmartChargingEffects.
func (sceq *SmartChargingEffectQuery) All(ctx context.Context) ([]*SmartChargingEffect, error) {
	if err := sceq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return sceq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sceq *SmartChargingEffectQuery) AllX(ctx context.Context) []*SmartChargingEffect {
	nodes, err := sceq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SmartChargingEffect IDs.
func (sceq *SmartChargingEffectQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := sceq.Select(smartchargingeffect.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sceq *SmartChargingEffectQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := sceq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sceq *SmartChargingEffectQuery) Count(ctx context.Context) (int, error) {
	if err := sceq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return sceq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sceq *SmartChargingEffectQuery) CountX(ctx context.Context) int {
	count, err := sceq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sceq *SmartChargingEffectQuery) Exist(ctx context.Context) (bool, error) {
	if err := sceq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return sceq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sceq *SmartChargingEffectQuery) ExistX(ctx context.Context) bool {
	exist, err := sceq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SmartChargingEffectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sceq *SmartChargingEffectQuery) Clone() *SmartChargingEffectQuery {
	if sceq == nil {
		return nil
	}
	return &SmartChargingEffectQuery{
		config:        sceq.config,
		limit:         sceq.limit,
		offset:        sceq.offset,
		order:         append([]OrderFunc{}, sceq.order...),
		predicates:    append([]predicate.SmartChargingEffect{}, sceq.predicates...),
		withEquipment: sceq.withEquipment.Clone(),
		withConnector: sceq.withConnector.Clone(),
		withOrderInfo: sceq.withOrderInfo.Clone(),
		// clone intermediate query.
		sql:    sceq.sql.Clone(),
		path:   sceq.path,
		unique: sceq.unique,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (sceq *SmartChargingEffectQuery) WithEquipment(opts ...func(*EquipmentQuery)) *SmartChargingEffectQuery {
	query := &EquipmentQuery{config: sceq.config}
	for _, opt := range opts {
		opt(query)
	}
	sceq.withEquipment = query
	return sceq
}

// WithConnector tells the query-builder to eager-load the nodes that are connected to
// the "connector" edge. The optional arguments are used to configure the query builder of the edge.
func (sceq *SmartChargingEffectQuery) WithConnector(opts ...func(*ConnectorQuery)) *SmartChargingEffectQuery {
	query := &ConnectorQuery{config: sceq.config}
	for _, opt := range opts {
		opt(query)
	}
	sceq.withConnector = query
	return sceq
}

// WithOrderInfo tells the query-builder to eager-load the nodes that are connected to
// the "order_info" edge. The optional arguments are used to configure the query builder of the edge.
func (sceq *SmartChargingEffectQuery) WithOrderInfo(opts ...func(*OrderInfoQuery)) *SmartChargingEffectQuery {
	query := &OrderInfoQuery{config: sceq.config}
	for _, opt := range opts {
		opt(query)
	}
	sceq.withOrderInfo = query
	return sceq
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
//	client.SmartChargingEffect.Query().
//		GroupBy(smartchargingeffect.FieldVersion).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sceq *SmartChargingEffectQuery) GroupBy(field string, fields ...string) *SmartChargingEffectGroupBy {
	grbuild := &SmartChargingEffectGroupBy{config: sceq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := sceq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return sceq.sqlQuery(ctx), nil
	}
	grbuild.label = smartchargingeffect.Label
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
//	client.SmartChargingEffect.Query().
//		Select(smartchargingeffect.FieldVersion).
//		Scan(ctx, &v)
func (sceq *SmartChargingEffectQuery) Select(fields ...string) *SmartChargingEffectSelect {
	sceq.fields = append(sceq.fields, fields...)
	selbuild := &SmartChargingEffectSelect{SmartChargingEffectQuery: sceq}
	selbuild.label = smartchargingeffect.Label
	selbuild.flds, selbuild.scan = &sceq.fields, selbuild.Scan
	return selbuild
}

func (sceq *SmartChargingEffectQuery) prepareQuery(ctx context.Context) error {
	for _, f := range sceq.fields {
		if !smartchargingeffect.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sceq.path != nil {
		prev, err := sceq.path(ctx)
		if err != nil {
			return err
		}
		sceq.sql = prev
	}
	return nil
}

func (sceq *SmartChargingEffectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*SmartChargingEffect, error) {
	var (
		nodes       = []*SmartChargingEffect{}
		withFKs     = sceq.withFKs
		_spec       = sceq.querySpec()
		loadedTypes = [3]bool{
			sceq.withEquipment != nil,
			sceq.withConnector != nil,
			sceq.withOrderInfo != nil,
		}
	)
	if sceq.withEquipment != nil || sceq.withConnector != nil || sceq.withOrderInfo != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, smartchargingeffect.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*SmartChargingEffect).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &SmartChargingEffect{config: sceq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sceq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sceq.withEquipment; query != nil {
		if err := sceq.loadEquipment(ctx, query, nodes, nil,
			func(n *SmartChargingEffect, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	if query := sceq.withConnector; query != nil {
		if err := sceq.loadConnector(ctx, query, nodes, nil,
			func(n *SmartChargingEffect, e *Connector) { n.Edges.Connector = e }); err != nil {
			return nil, err
		}
	}
	if query := sceq.withOrderInfo; query != nil {
		if err := sceq.loadOrderInfo(ctx, query, nodes, nil,
			func(n *SmartChargingEffect, e *OrderInfo) { n.Edges.OrderInfo = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sceq *SmartChargingEffectQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*SmartChargingEffect, init func(*SmartChargingEffect), assign func(*SmartChargingEffect, *Equipment)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*SmartChargingEffect)
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
func (sceq *SmartChargingEffectQuery) loadConnector(ctx context.Context, query *ConnectorQuery, nodes []*SmartChargingEffect, init func(*SmartChargingEffect), assign func(*SmartChargingEffect, *Connector)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*SmartChargingEffect)
	for i := range nodes {
		if nodes[i].connector_id == nil {
			continue
		}
		fk := *nodes[i].connector_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(connector.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "connector_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sceq *SmartChargingEffectQuery) loadOrderInfo(ctx context.Context, query *OrderInfoQuery, nodes []*SmartChargingEffect, init func(*SmartChargingEffect), assign func(*SmartChargingEffect, *OrderInfo)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*SmartChargingEffect)
	for i := range nodes {
		if nodes[i].order_id == nil {
			continue
		}
		fk := *nodes[i].order_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(orderinfo.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (sceq *SmartChargingEffectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sceq.querySpec()
	_spec.Node.Columns = sceq.fields
	if len(sceq.fields) > 0 {
		_spec.Unique = sceq.unique != nil && *sceq.unique
	}
	return sqlgraph.CountNodes(ctx, sceq.driver, _spec)
}

func (sceq *SmartChargingEffectQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := sceq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (sceq *SmartChargingEffectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   smartchargingeffect.Table,
			Columns: smartchargingeffect.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: smartchargingeffect.FieldID,
			},
		},
		From:   sceq.sql,
		Unique: true,
	}
	if unique := sceq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sceq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, smartchargingeffect.FieldID)
		for i := range fields {
			if fields[i] != smartchargingeffect.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sceq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sceq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sceq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sceq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sceq *SmartChargingEffectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sceq.driver.Dialect())
	t1 := builder.Table(smartchargingeffect.Table)
	columns := sceq.fields
	if len(columns) == 0 {
		columns = smartchargingeffect.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sceq.sql != nil {
		selector = sceq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sceq.unique != nil && *sceq.unique {
		selector.Distinct()
	}
	for _, p := range sceq.predicates {
		p(selector)
	}
	for _, p := range sceq.order {
		p(selector)
	}
	if offset := sceq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sceq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SmartChargingEffectGroupBy is the group-by builder for SmartChargingEffect entities.
type SmartChargingEffectGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scegb *SmartChargingEffectGroupBy) Aggregate(fns ...AggregateFunc) *SmartChargingEffectGroupBy {
	scegb.fns = append(scegb.fns, fns...)
	return scegb
}

// Scan applies the group-by query and scans the result into the given value.
func (scegb *SmartChargingEffectGroupBy) Scan(ctx context.Context, v any) error {
	query, err := scegb.path(ctx)
	if err != nil {
		return err
	}
	scegb.sql = query
	return scegb.sqlScan(ctx, v)
}

func (scegb *SmartChargingEffectGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range scegb.fields {
		if !smartchargingeffect.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := scegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (scegb *SmartChargingEffectGroupBy) sqlQuery() *sql.Selector {
	selector := scegb.sql.Select()
	aggregation := make([]string, 0, len(scegb.fns))
	for _, fn := range scegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(scegb.fields)+len(scegb.fns))
		for _, f := range scegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(scegb.fields...)...)
}

// SmartChargingEffectSelect is the builder for selecting fields of SmartChargingEffect entities.
type SmartChargingEffectSelect struct {
	*SmartChargingEffectQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (sces *SmartChargingEffectSelect) Scan(ctx context.Context, v any) error {
	if err := sces.prepareQuery(ctx); err != nil {
		return err
	}
	sces.sql = sces.SmartChargingEffectQuery.sqlQuery(ctx)
	return sces.sqlScan(ctx, v)
}

func (sces *SmartChargingEffectSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := sces.sql.Query()
	if err := sces.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}