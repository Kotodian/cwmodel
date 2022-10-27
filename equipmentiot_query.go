// Code generated by ent, DO NOT EDIT.

package cwmodel

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Kotodian/cwmodel/equipment"
	"github.com/Kotodian/cwmodel/equipmentiot"
	"github.com/Kotodian/cwmodel/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// EquipmentIotQuery is the builder for querying EquipmentIot entities.
type EquipmentIotQuery struct {
	config
	limit         *int
	offset        *int
	unique        *bool
	order         []OrderFunc
	fields        []string
	predicates    []predicate.EquipmentIot
	withEquipment *EquipmentQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EquipmentIotQuery builder.
func (eiq *EquipmentIotQuery) Where(ps ...predicate.EquipmentIot) *EquipmentIotQuery {
	eiq.predicates = append(eiq.predicates, ps...)
	return eiq
}

// Limit adds a limit step to the query.
func (eiq *EquipmentIotQuery) Limit(limit int) *EquipmentIotQuery {
	eiq.limit = &limit
	return eiq
}

// Offset adds an offset step to the query.
func (eiq *EquipmentIotQuery) Offset(offset int) *EquipmentIotQuery {
	eiq.offset = &offset
	return eiq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (eiq *EquipmentIotQuery) Unique(unique bool) *EquipmentIotQuery {
	eiq.unique = &unique
	return eiq
}

// Order adds an order step to the query.
func (eiq *EquipmentIotQuery) Order(o ...OrderFunc) *EquipmentIotQuery {
	eiq.order = append(eiq.order, o...)
	return eiq
}

// QueryEquipment chains the current query on the "equipment" edge.
func (eiq *EquipmentIotQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: eiq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := eiq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(equipmentiot.Table, equipmentiot.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, equipmentiot.EquipmentTable, equipmentiot.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(eiq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first EquipmentIot entity from the query.
// Returns a *NotFoundError when no EquipmentIot was found.
func (eiq *EquipmentIotQuery) First(ctx context.Context) (*EquipmentIot, error) {
	nodes, err := eiq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{equipmentiot.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (eiq *EquipmentIotQuery) FirstX(ctx context.Context) *EquipmentIot {
	node, err := eiq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EquipmentIot ID from the query.
// Returns a *NotFoundError when no EquipmentIot ID was found.
func (eiq *EquipmentIotQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = eiq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{equipmentiot.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (eiq *EquipmentIotQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := eiq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EquipmentIot entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EquipmentIot entity is found.
// Returns a *NotFoundError when no EquipmentIot entities are found.
func (eiq *EquipmentIotQuery) Only(ctx context.Context) (*EquipmentIot, error) {
	nodes, err := eiq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{equipmentiot.Label}
	default:
		return nil, &NotSingularError{equipmentiot.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (eiq *EquipmentIotQuery) OnlyX(ctx context.Context) *EquipmentIot {
	node, err := eiq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EquipmentIot ID in the query.
// Returns a *NotSingularError when more than one EquipmentIot ID is found.
// Returns a *NotFoundError when no entities are found.
func (eiq *EquipmentIotQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = eiq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{equipmentiot.Label}
	default:
		err = &NotSingularError{equipmentiot.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (eiq *EquipmentIotQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := eiq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EquipmentIots.
func (eiq *EquipmentIotQuery) All(ctx context.Context) ([]*EquipmentIot, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return eiq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (eiq *EquipmentIotQuery) AllX(ctx context.Context) []*EquipmentIot {
	nodes, err := eiq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EquipmentIot IDs.
func (eiq *EquipmentIotQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := eiq.Select(equipmentiot.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (eiq *EquipmentIotQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := eiq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (eiq *EquipmentIotQuery) Count(ctx context.Context) (int, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return eiq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (eiq *EquipmentIotQuery) CountX(ctx context.Context) int {
	count, err := eiq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (eiq *EquipmentIotQuery) Exist(ctx context.Context) (bool, error) {
	if err := eiq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return eiq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (eiq *EquipmentIotQuery) ExistX(ctx context.Context) bool {
	exist, err := eiq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EquipmentIotQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (eiq *EquipmentIotQuery) Clone() *EquipmentIotQuery {
	if eiq == nil {
		return nil
	}
	return &EquipmentIotQuery{
		config:        eiq.config,
		limit:         eiq.limit,
		offset:        eiq.offset,
		order:         append([]OrderFunc{}, eiq.order...),
		predicates:    append([]predicate.EquipmentIot{}, eiq.predicates...),
		withEquipment: eiq.withEquipment.Clone(),
		// clone intermediate query.
		sql:    eiq.sql.Clone(),
		path:   eiq.path,
		unique: eiq.unique,
	}
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (eiq *EquipmentIotQuery) WithEquipment(opts ...func(*EquipmentQuery)) *EquipmentIotQuery {
	query := &EquipmentQuery{config: eiq.config}
	for _, opt := range opts {
		opt(query)
	}
	eiq.withEquipment = query
	return eiq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Version int64 `json:"version"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EquipmentIot.Query().
//		GroupBy(equipmentiot.FieldVersion).
//		Aggregate(cwmodel.Count()).
//		Scan(ctx, &v)
func (eiq *EquipmentIotQuery) GroupBy(field string, fields ...string) *EquipmentIotGroupBy {
	grbuild := &EquipmentIotGroupBy{config: eiq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := eiq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return eiq.sqlQuery(ctx), nil
	}
	grbuild.label = equipmentiot.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Version int64 `json:"version"`
//	}
//
//	client.EquipmentIot.Query().
//		Select(equipmentiot.FieldVersion).
//		Scan(ctx, &v)
func (eiq *EquipmentIotQuery) Select(fields ...string) *EquipmentIotSelect {
	eiq.fields = append(eiq.fields, fields...)
	selbuild := &EquipmentIotSelect{EquipmentIotQuery: eiq}
	selbuild.label = equipmentiot.Label
	selbuild.flds, selbuild.scan = &eiq.fields, selbuild.Scan
	return selbuild
}

func (eiq *EquipmentIotQuery) prepareQuery(ctx context.Context) error {
	for _, f := range eiq.fields {
		if !equipmentiot.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cwmodel: invalid field %q for query", f)}
		}
	}
	if eiq.path != nil {
		prev, err := eiq.path(ctx)
		if err != nil {
			return err
		}
		eiq.sql = prev
	}
	return nil
}

func (eiq *EquipmentIotQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EquipmentIot, error) {
	var (
		nodes       = []*EquipmentIot{}
		_spec       = eiq.querySpec()
		loadedTypes = [1]bool{
			eiq.withEquipment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EquipmentIot).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EquipmentIot{config: eiq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, eiq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := eiq.withEquipment; query != nil {
		if err := eiq.loadEquipment(ctx, query, nodes, nil,
			func(n *EquipmentIot, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (eiq *EquipmentIotQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*EquipmentIot, init func(*EquipmentIot), assign func(*EquipmentIot, *Equipment)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*EquipmentIot)
	for i := range nodes {
		fk := nodes[i].EquipmentID
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

func (eiq *EquipmentIotQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := eiq.querySpec()
	_spec.Node.Columns = eiq.fields
	if len(eiq.fields) > 0 {
		_spec.Unique = eiq.unique != nil && *eiq.unique
	}
	return sqlgraph.CountNodes(ctx, eiq.driver, _spec)
}

func (eiq *EquipmentIotQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := eiq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cwmodel: check existence: %w", err)
	default:
		return true, nil
	}
}

func (eiq *EquipmentIotQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   equipmentiot.Table,
			Columns: equipmentiot.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: equipmentiot.FieldID,
			},
		},
		From:   eiq.sql,
		Unique: true,
	}
	if unique := eiq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := eiq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, equipmentiot.FieldID)
		for i := range fields {
			if fields[i] != equipmentiot.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := eiq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := eiq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := eiq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := eiq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (eiq *EquipmentIotQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(eiq.driver.Dialect())
	t1 := builder.Table(equipmentiot.Table)
	columns := eiq.fields
	if len(columns) == 0 {
		columns = equipmentiot.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if eiq.sql != nil {
		selector = eiq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if eiq.unique != nil && *eiq.unique {
		selector.Distinct()
	}
	for _, p := range eiq.predicates {
		p(selector)
	}
	for _, p := range eiq.order {
		p(selector)
	}
	if offset := eiq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := eiq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EquipmentIotGroupBy is the group-by builder for EquipmentIot entities.
type EquipmentIotGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (eigb *EquipmentIotGroupBy) Aggregate(fns ...AggregateFunc) *EquipmentIotGroupBy {
	eigb.fns = append(eigb.fns, fns...)
	return eigb
}

// Scan applies the group-by query and scans the result into the given value.
func (eigb *EquipmentIotGroupBy) Scan(ctx context.Context, v any) error {
	query, err := eigb.path(ctx)
	if err != nil {
		return err
	}
	eigb.sql = query
	return eigb.sqlScan(ctx, v)
}

func (eigb *EquipmentIotGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range eigb.fields {
		if !equipmentiot.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := eigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := eigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (eigb *EquipmentIotGroupBy) sqlQuery() *sql.Selector {
	selector := eigb.sql.Select()
	aggregation := make([]string, 0, len(eigb.fns))
	for _, fn := range eigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(eigb.fields)+len(eigb.fns))
		for _, f := range eigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(eigb.fields...)...)
}

// EquipmentIotSelect is the builder for selecting fields of EquipmentIot entities.
type EquipmentIotSelect struct {
	*EquipmentIotQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (eis *EquipmentIotSelect) Scan(ctx context.Context, v any) error {
	if err := eis.prepareQuery(ctx); err != nil {
		return err
	}
	eis.sql = eis.EquipmentIotQuery.sqlQuery(ctx)
	return eis.sqlScan(ctx, v)
}

func (eis *EquipmentIotSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := eis.sql.Query()
	if err := eis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
