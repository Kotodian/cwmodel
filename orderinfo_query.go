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
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/orderevent"
	"github.com/Kotodian/cwmodel/orderinfo"
	"github.com/Kotodian/cwmodel/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// OrderInfoQuery is the builder for querying OrderInfo entities.
type OrderInfoQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.OrderInfo
	withEquipment  *EquipmentQuery
	withOrderEvent *OrderEventQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderInfoQuery builder.
func (oiq *OrderInfoQuery) Where(ps ...predicate.OrderInfo) *OrderInfoQuery {
	oiq.predicates = append(oiq.predicates, ps...)
	return oiq
}

// Limit adds a limit step to the query.
func (oiq *OrderInfoQuery) Limit(limit int) *OrderInfoQuery {
	oiq.limit = &limit
	return oiq
}

// Offset adds an offset step to the query.
func (oiq *OrderInfoQuery) Offset(offset int) *OrderInfoQuery {
	oiq.offset = &offset
	return oiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oiq *OrderInfoQuery) Unique(unique bool) *OrderInfoQuery {
	oiq.unique = &unique
	return oiq
}

// Order adds an order step to the query.
func (oiq *OrderInfoQuery) Order(o ...OrderFunc) *OrderInfoQuery {
	oiq.order = append(oiq.order, o...)
	return oiq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (oiq *OrderInfoQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: oiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderinfo.Table, orderinfo.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, orderinfo.EquipmentTable, orderinfo.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderEvent chains the current query on the "order_event" edge.
func (oiq *OrderInfoQuery) QueryOrderEvent() *OrderEventQuery {
	query := &OrderEventQuery{config: oiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := oiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(orderinfo.Table, orderinfo.FieldID, selector),
			sqlgraph.To(orderevent.Table, orderevent.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, orderinfo.OrderEventTable, orderinfo.OrderEventColumn),
		)
		fromU = sqlgraph.SetNeighbors(oiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first OrderInfo entity from the query.
// Returns a *NotFoundError when no OrderInfo was found.
func (oiq *OrderInfoQuery) First(ctx context.Context) (*OrderInfo, error) {
	nodes, err := oiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderinfo.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oiq *OrderInfoQuery) FirstX(ctx context.Context) *OrderInfo {
	node, err := oiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderInfo ID from the query.
// Returns a *NotFoundError when no OrderInfo ID was found.
func (oiq *OrderInfoQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = oiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oiq *OrderInfoQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := oiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderInfo entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderInfo entity is found.
// Returns a *NotFoundError when no OrderInfo entities are found.
func (oiq *OrderInfoQuery) Only(ctx context.Context) (*OrderInfo, error) {
	nodes, err := oiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderinfo.Label}
	default:
		return nil, &NotSingularError{orderinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oiq *OrderInfoQuery) OnlyX(ctx context.Context) *OrderInfo {
	node, err := oiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderInfo ID in the query.
// Returns a *NotSingularError when more than one OrderInfo ID is found.
// Returns a *NotFoundError when no entities are found.
func (oiq *OrderInfoQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = oiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderinfo.Label}
	default:
		err = &NotSingularError{orderinfo.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oiq *OrderInfoQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := oiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderInfos.
func (oiq *OrderInfoQuery) All(ctx context.Context) ([]*OrderInfo, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return oiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (oiq *OrderInfoQuery) AllX(ctx context.Context) []*OrderInfo {
	nodes, err := oiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderInfo IDs.
func (oiq *OrderInfoQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := oiq.Select(orderinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oiq *OrderInfoQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := oiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oiq *OrderInfoQuery) Count(ctx context.Context) (int, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return oiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (oiq *OrderInfoQuery) CountX(ctx context.Context) int {
	count, err := oiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oiq *OrderInfoQuery) Exist(ctx context.Context) (bool, error) {
	if err := oiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return oiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (oiq *OrderInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := oiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderInfoQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oiq *OrderInfoQuery) Clone() *OrderInfoQuery {
	if oiq == nil {
		return nil
	}
	return &OrderInfoQuery{
		config:         oiq.config,
		limit:          oiq.limit,
		offset:         oiq.offset,
		order:          append([]OrderFunc{}, oiq.order...),
		predicates:     append([]predicate.OrderInfo{}, oiq.predicates...),
		withEquipment:  oiq.withEquipment.Clone(),
		withOrderEvent: oiq.withOrderEvent.Clone(),
		// clone intermediate query.
		sql:    oiq.sql.Clone(),
		path:   oiq.path,
		unique: oiq.unique,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (oiq *OrderInfoQuery) WithEquipment(opts ...func(*EquipmentQuery)) *OrderInfoQuery {
	query := &EquipmentQuery{config: oiq.config}
	for _, opt := range opts {
		opt(query)
	}
	oiq.withEquipment = query
	return oiq
}

// WithOrderEvent tells the query-builder to eager-load the nodes that are connected to
// the "order_event" edge. The optional arguments are used to configure the query builder of the edge.
func (oiq *OrderInfoQuery) WithOrderEvent(opts ...func(*OrderEventQuery)) *OrderInfoQuery {
	query := &OrderEventQuery{config: oiq.config}
	for _, opt := range opts {
		opt(query)
	}
	oiq.withOrderEvent = query
	return oiq
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
//	client.OrderInfo.Query().
//		GroupBy(orderinfo.FieldVersion).
//		Aggregate(cwmodel.Count()).
//		Scan(ctx, &v)
func (oiq *OrderInfoQuery) GroupBy(field string, fields ...string) *OrderInfoGroupBy {
	grbuild := &OrderInfoGroupBy{config: oiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := oiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return oiq.sqlQuery(ctx), nil
	}
	grbuild.label = orderinfo.Label
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
//	client.OrderInfo.Query().
//		Select(orderinfo.FieldVersion).
//		Scan(ctx, &v)
func (oiq *OrderInfoQuery) Select(fields ...string) *OrderInfoSelect {
	oiq.fields = append(oiq.fields, fields...)
	selbuild := &OrderInfoSelect{OrderInfoQuery: oiq}
	selbuild.label = orderinfo.Label
	selbuild.flds, selbuild.scan = &oiq.fields, selbuild.Scan
	return selbuild
}

func (oiq *OrderInfoQuery) prepareQuery(ctx context.Context) error {
	for _, f := range oiq.fields {
		if !orderinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cwmodel: invalid field %q for query", f)}
		}
	}
	if oiq.path != nil {
		prev, err := oiq.path(ctx)
		if err != nil {
			return err
		}
		oiq.sql = prev
	}
	return nil
}

func (oiq *OrderInfoQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderInfo, error) {
	var (
		nodes       = []*OrderInfo{}
		withFKs     = oiq.withFKs
		_spec       = oiq.querySpec()
		loadedTypes = [2]bool{
			oiq.withEquipment != nil,
			oiq.withOrderEvent != nil,
		}
	)
	if oiq.withEquipment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, orderinfo.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderInfo).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderInfo{config: oiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := oiq.withEquipment; query != nil {
		if err := oiq.loadEquipment(ctx, query, nodes, nil,
			func(n *OrderInfo, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	if query := oiq.withOrderEvent; query != nil {
		if err := oiq.loadOrderEvent(ctx, query, nodes,
			func(n *OrderInfo) { n.Edges.OrderEvent = []*OrderEvent{} },
			func(n *OrderInfo, e *OrderEvent) { n.Edges.OrderEvent = append(n.Edges.OrderEvent, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (oiq *OrderInfoQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*OrderInfo, init func(*OrderInfo), assign func(*OrderInfo, *Equipment)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*OrderInfo)
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
func (oiq *OrderInfoQuery) loadOrderEvent(ctx context.Context, query *OrderEventQuery, nodes []*OrderInfo, init func(*OrderInfo), assign func(*OrderInfo, *OrderEvent)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[datasource.UUID]*OrderInfo)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.OrderEvent(func(s *sql.Selector) {
		s.Where(sql.InValues(orderinfo.OrderEventColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.order_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "order_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "order_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (oiq *OrderInfoQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oiq.querySpec()
	_spec.Node.Columns = oiq.fields
	if len(oiq.fields) > 0 {
		_spec.Unique = oiq.unique != nil && *oiq.unique
	}
	return sqlgraph.CountNodes(ctx, oiq.driver, _spec)
}

func (oiq *OrderInfoQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := oiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cwmodel: check existence: %w", err)
	default:
		return true, nil
	}
}

func (oiq *OrderInfoQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   orderinfo.Table,
			Columns: orderinfo.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: orderinfo.FieldID,
			},
		},
		From:   oiq.sql,
		Unique: true,
	}
	if unique := oiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := oiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderinfo.FieldID)
		for i := range fields {
			if fields[i] != orderinfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oiq *OrderInfoQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oiq.driver.Dialect())
	t1 := builder.Table(orderinfo.Table)
	columns := oiq.fields
	if len(columns) == 0 {
		columns = orderinfo.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oiq.sql != nil {
		selector = oiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oiq.unique != nil && *oiq.unique {
		selector.Distinct()
	}
	for _, p := range oiq.predicates {
		p(selector)
	}
	for _, p := range oiq.order {
		p(selector)
	}
	if offset := oiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderInfoGroupBy is the group-by builder for OrderInfo entities.
type OrderInfoGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oigb *OrderInfoGroupBy) Aggregate(fns ...AggregateFunc) *OrderInfoGroupBy {
	oigb.fns = append(oigb.fns, fns...)
	return oigb
}

// Scan applies the group-by query and scans the result into the given value.
func (oigb *OrderInfoGroupBy) Scan(ctx context.Context, v any) error {
	query, err := oigb.path(ctx)
	if err != nil {
		return err
	}
	oigb.sql = query
	return oigb.sqlScan(ctx, v)
}

func (oigb *OrderInfoGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range oigb.fields {
		if !orderinfo.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := oigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (oigb *OrderInfoGroupBy) sqlQuery() *sql.Selector {
	selector := oigb.sql.Select()
	aggregation := make([]string, 0, len(oigb.fns))
	for _, fn := range oigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(oigb.fields)+len(oigb.fns))
		for _, f := range oigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(oigb.fields...)...)
}

// OrderInfoSelect is the builder for selecting fields of OrderInfo entities.
type OrderInfoSelect struct {
	*OrderInfoQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ois *OrderInfoSelect) Scan(ctx context.Context, v any) error {
	if err := ois.prepareQuery(ctx); err != nil {
		return err
	}
	ois.sql = ois.OrderInfoQuery.sqlQuery(ctx)
	return ois.sqlScan(ctx, v)
}

func (ois *OrderInfoSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ois.sql.Query()
	if err := ois.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
