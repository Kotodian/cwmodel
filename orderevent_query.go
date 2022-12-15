// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/orderevent"
	"github.com/Kotodian/cwmodel/orderinfo"
	"github.com/Kotodian/cwmodel/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// OrderEventQuery is the builder for querying OrderEvent entities.
type OrderEventQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.OrderEvent
	withOrderInfo *OrderInfoQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderEventQuery builder.
func (oeq *OrderEventQuery) Where(ps ...predicate.OrderEvent) *OrderEventQuery {
	oeq.predicates = append(oeq.predicates, ps...)
	return oeq
}

// Limit adds a limit step to the query.
func (oeq *OrderEventQuery) Limit(limit int) *OrderEventQuery {
	oeq.limit = &limit
	return oeq
}

// Offset adds an offset step to the query.
func (oeq *OrderEventQuery) Offset(offset int) *OrderEventQuery {
	oeq.offset = &offset
	return oeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oeq *OrderEventQuery) Unique(unique bool) *OrderEventQuery {
	oeq.unique = &unique
	return oeq
}

// Order adds an order step to the query.
func (oeq *OrderEventQuery) Order(o ...OrderFunc) *OrderEventQuery {
	oeq.order = append(oeq.order, o...)
	return oeq
}

// QueryOrderInfo chains the current query on the "order_info" edge.
func (oeq *OrderEventQuery) QueryOrderInfo() *OrderInfoQuery {
	query := &OrderInfoQuery{config: oeq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oeq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderevent.Table, orderevent.FieldID, selector),
			sqlgraph.To(orderinfo.Table, orderinfo.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderevent.OrderInfoTable, orderevent.OrderInfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(oeq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderEvent entity from the query.
// Returns a *NotFoundError when no OrderEvent was found.
func (oeq *OrderEventQuery) First(ctx context.Context) (*OrderEvent, error) {
	nodes, err := oeq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oeq *OrderEventQuery) FirstX(ctx context.Context) *OrderEvent {
	node, err := oeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderEvent ID from the query.
// Returns a *NotFoundError when no OrderEvent ID was found.
func (oeq *OrderEventQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = oeq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oeq *OrderEventQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := oeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderEvent entity is found.
// Returns a *NotFoundError when no OrderEvent entities are found.
func (oeq *OrderEventQuery) Only(ctx context.Context) (*OrderEvent, error) {
	nodes, err := oeq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderevent.Label}
	default:
		return nil, &NotSingularError{orderevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oeq *OrderEventQuery) OnlyX(ctx context.Context) *OrderEvent {
	node, err := oeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderEvent ID in the query.
// Returns a *NotSingularError when more than one OrderEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (oeq *OrderEventQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = oeq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderevent.Label}
	default:
		err = &NotSingularError{orderevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oeq *OrderEventQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := oeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderEvents.
func (oeq *OrderEventQuery) All(ctx context.Context) ([]*OrderEvent, error) {
	if err := oeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oeq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oeq *OrderEventQuery) AllX(ctx context.Context) []*OrderEvent {
	nodes, err := oeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderEvent IDs.
func (oeq *OrderEventQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := oeq.Select(orderevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oeq *OrderEventQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := oeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oeq *OrderEventQuery) Count(ctx context.Context) (int, error) {
	if err := oeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oeq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oeq *OrderEventQuery) CountX(ctx context.Context) int {
	count, err := oeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oeq *OrderEventQuery) Exist(ctx context.Context) (bool, error) {
	switch _, err := oeq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cwmodel: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oeq *OrderEventQuery) ExistX(ctx context.Context) bool {
	exist, err := oeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oeq *OrderEventQuery) Clone() *OrderEventQuery {
	if oeq == nil {
		return nil
	}
	return &OrderEventQuery{
		config:        oeq.config,
		limit:         oeq.limit,
		offset:        oeq.offset,
		order:         append([]OrderFunc{}, oeq.order...),
		predicates:    append([]predicate.OrderEvent{}, oeq.predicates...),
		withOrderInfo: oeq.withOrderInfo.Clone(),
		// clone intermediate query.
		sql:    oeq.sql.Clone(),
		path:   oeq.path,
		unique: oeq.unique,
	}
}

// WithOrderInfo tells the query-builder to eager-load the nodes that are connected to
// the "order_info" edge. The optional arguments are used to configure the query builder of the edge.
func (oeq *OrderEventQuery) WithOrderInfo(opts ...func(*OrderInfoQuery)) *OrderEventQuery {
	query := &OrderInfoQuery{config: oeq.config}
	for _, opt := range opts {
		opt(query)
	}
	oeq.withOrderInfo = query
	return oeq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		OrderID datasource.UUID `json:"order_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrderEvent.Query().
//		GroupBy(orderevent.FieldOrderID).
//		Aggregate(cwmodel.Count()).
//		Scan(ctx, &v)
func (oeq *OrderEventQuery) GroupBy(field string, fields ...string) *OrderEventGroupBy {
	grbuild := &OrderEventGroupBy{config: oeq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oeq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oeq.sqlQuery(ctx), nil
	}
	grbuild.label = orderevent.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		OrderID datasource.UUID `json:"order_id,omitempty"`
//	}
//
//	client.OrderEvent.Query().
//		Select(orderevent.FieldOrderID).
//		Scan(ctx, &v)
func (oeq *OrderEventQuery) Select(fields ...string) *OrderEventSelect {
	oeq.fields = append(oeq.fields, fields...)
	selbuild := &OrderEventSelect{OrderEventQuery: oeq}
	selbuild.label = orderevent.Label
	selbuild.flds, selbuild.scan = &oeq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a OrderEventSelect configured with the given aggregations.
func (oeq *OrderEventQuery) Aggregate(fns ...AggregateFunc) *OrderEventSelect {
	return oeq.Select().Aggregate(fns...)
}

func (oeq *OrderEventQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oeq.fields {
		if !orderevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cwmodel: invalid field %q for query", f)}
		}
	}
	if oeq.path != nil {
		prev, err := oeq.path(ctx)
		if err != nil {
			return err
		}
		oeq.sql = prev
	}
	return nil
}

func (oeq *OrderEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderEvent, error) {
	var (
		nodes       = []*OrderEvent{}
		_spec       = oeq.querySpec()
		loadedTypes = [1]bool{
			oeq.withOrderInfo != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderEvent{config: oeq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oeq.withOrderInfo; query != nil {
		if err := oeq.loadOrderInfo(ctx, query, nodes, nil,
			func(n *OrderEvent, e *OrderInfo) { n.Edges.OrderInfo = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oeq *OrderEventQuery) loadOrderInfo(ctx context.Context, query *OrderInfoQuery, nodes []*OrderEvent, init func(*OrderEvent), assign func(*OrderEvent, *OrderInfo)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*OrderEvent)
	for i := range nodes {
		fk := nodes[i].OrderID
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

func (oeq *OrderEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oeq.querySpec()
	_spec.Node.Columns = oeq.fields
	if len(oeq.fields) > 0 {
		_spec.Unique = oeq.unique != nil && *oeq.unique
	}
	return sqlgraph.CountNodes(ctx, oeq.driver, _spec)
}

func (oeq *OrderEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderevent.Table,
			Columns: orderevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: orderevent.FieldID,
			},
		},
		From:   oeq.sql,
		Unique: true,
	}
	if unique := oeq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oeq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderevent.FieldID)
		for i := range fields {
			if fields[i] != orderevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oeq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oeq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oeq *OrderEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oeq.driver.Dialect())
	t1 := builder.Table(orderevent.Table)
	columns := oeq.fields
	if len(columns) == 0 {
		columns = orderevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oeq.sql != nil {
		selector = oeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oeq.unique != nil && *oeq.unique {
		selector.Distinct()
	}
	for _, p := range oeq.predicates {
		p(selector)
	}
	for _, p := range oeq.order {
		p(selector)
	}
	if offset := oeq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oeq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderEventGroupBy is the group-by builder for OrderEvent entities.
type OrderEventGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oegb *OrderEventGroupBy) Aggregate(fns ...AggregateFunc) *OrderEventGroupBy {
	oegb.fns = append(oegb.fns, fns...)
	return oegb
}

// Scan applies the group-by query and scans the result into the given value.
func (oegb *OrderEventGroupBy) Scan(ctx context.Context, v any) error {
	query, err := oegb.path(ctx)
	if err != nil {
		return err
	}
	oegb.sql = query
	return oegb.sqlScan(ctx, v)
}

func (oegb *OrderEventGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range oegb.fields {
		if !orderevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oegb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oegb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oegb *OrderEventGroupBy) sqlQuery() *sql.Selector {
	selector := oegb.sql.Select()
	aggregation := make([]string, 0, len(oegb.fns))
	for _, fn := range oegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oegb.fields)+len(oegb.fns))
		for _, f := range oegb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oegb.fields...)...)
}

// OrderEventSelect is the builder for selecting fields of OrderEvent entities.
type OrderEventSelect struct {
	*OrderEventQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oes *OrderEventSelect) Aggregate(fns ...AggregateFunc) *OrderEventSelect {
	oes.fns = append(oes.fns, fns...)
	return oes
}

// Scan applies the selector query and scans the result into the given value.
func (oes *OrderEventSelect) Scan(ctx context.Context, v any) error {
	if err := oes.prepareQuery(ctx); err != nil {
		return err
	}
	oes.sql = oes.OrderEventQuery.sqlQuery(ctx)
	return oes.sqlScan(ctx, v)
}

func (oes *OrderEventSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(oes.fns))
	for _, fn := range oes.fns {
		aggregation = append(aggregation, fn(oes.sql))
	}
	switch n := len(*oes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		oes.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		oes.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := oes.sql.Query()
	if err := oes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
