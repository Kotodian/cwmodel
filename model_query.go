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
	"github.com/Kotodian/cwmodel/firmware"
	"github.com/Kotodian/cwmodel/model"
	"github.com/Kotodian/cwmodel/predicate"
	"github.com/Kotodian/gokit/datasource"
)

// ModelQuery is the builder for querying Model entities.
type ModelQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.Model
	withFirmware *FirmwareQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ModelQuery builder.
func (mq *ModelQuery) Where(ps ...predicate.Model) *ModelQuery {
	mq.predicates = append(mq.predicates, ps...)
	return mq
}

// Limit adds a limit step to the query.
func (mq *ModelQuery) Limit(limit int) *ModelQuery {
	mq.limit = &limit
	return mq
}

// Offset adds an offset step to the query.
func (mq *ModelQuery) Offset(offset int) *ModelQuery {
	mq.offset = &offset
	return mq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (mq *ModelQuery) Unique(unique bool) *ModelQuery {
	mq.unique = &unique
	return mq
}

// Order adds an order step to the query.
func (mq *ModelQuery) Order(o ...OrderFunc) *ModelQuery {
	mq.order = append(mq.order, o...)
	return mq
}

// QueryFirmware chains the current query on the "firmware" edge.
func (mq *ModelQuery) QueryFirmware() *FirmwareQuery {
	query := &FirmwareQuery{config: mq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := mq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(model.Table, model.FieldID, selector),
			sqlgraph.To(firmware.Table, firmware.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, model.FirmwareTable, model.FirmwareColumn),
		)
		fromU = sqlgraph.SetNeighbors(mq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Model entity from the query.
// Returns a *NotFoundError when no Model was found.
func (mq *ModelQuery) First(ctx context.Context) (*Model, error) {
	nodes, err := mq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{model.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (mq *ModelQuery) FirstX(ctx context.Context) *Model {
	node, err := mq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Model ID from the query.
// Returns a *NotFoundError when no Model ID was found.
func (mq *ModelQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = mq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{model.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (mq *ModelQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := mq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Model entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Model entity is found.
// Returns a *NotFoundError when no Model entities are found.
func (mq *ModelQuery) Only(ctx context.Context) (*Model, error) {
	nodes, err := mq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{model.Label}
	default:
		return nil, &NotSingularError{model.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (mq *ModelQuery) OnlyX(ctx context.Context) *Model {
	node, err := mq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Model ID in the query.
// Returns a *NotSingularError when more than one Model ID is found.
// Returns a *NotFoundError when no entities are found.
func (mq *ModelQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = mq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{model.Label}
	default:
		err = &NotSingularError{model.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (mq *ModelQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := mq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Models.
func (mq *ModelQuery) All(ctx context.Context) ([]*Model, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return mq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (mq *ModelQuery) AllX(ctx context.Context) []*Model {
	nodes, err := mq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Model IDs.
func (mq *ModelQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := mq.Select(model.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (mq *ModelQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := mq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (mq *ModelQuery) Count(ctx context.Context) (int, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return mq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (mq *ModelQuery) CountX(ctx context.Context) int {
	count, err := mq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (mq *ModelQuery) Exist(ctx context.Context) (bool, error) {
	if err := mq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return mq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (mq *ModelQuery) ExistX(ctx context.Context) bool {
	exist, err := mq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ModelQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (mq *ModelQuery) Clone() *ModelQuery {
	if mq == nil {
		return nil
	}
	return &ModelQuery{
		config:       mq.config,
		limit:        mq.limit,
		offset:       mq.offset,
		order:        append([]OrderFunc{}, mq.order...),
		predicates:   append([]predicate.Model{}, mq.predicates...),
		withFirmware: mq.withFirmware.Clone(),
		// clone intermediate query.
		sql:    mq.sql.Clone(),
		path:   mq.path,
		unique: mq.unique,
	}
}

// WithFirmware tells the query-builder to eager-load the nodes that are connected to
// the "firmware" edge. The optional arguments are used to configure the query builder of the edge.
func (mq *ModelQuery) WithFirmware(opts ...func(*FirmwareQuery)) *ModelQuery {
	query := &FirmwareQuery{config: mq.config}
	for _, opt := range opts {
		opt(query)
	}
	mq.withFirmware = query
	return mq
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
//	client.Model.Query().
//		GroupBy(model.FieldVersion).
//		Aggregate(cwmodel.Count()).
//		Scan(ctx, &v)
func (mq *ModelQuery) GroupBy(field string, fields ...string) *ModelGroupBy {
	grbuild := &ModelGroupBy{config: mq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := mq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return mq.sqlQuery(ctx), nil
	}
	grbuild.label = model.Label
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
//	client.Model.Query().
//		Select(model.FieldVersion).
//		Scan(ctx, &v)
func (mq *ModelQuery) Select(fields ...string) *ModelSelect {
	mq.fields = append(mq.fields, fields...)
	selbuild := &ModelSelect{ModelQuery: mq}
	selbuild.label = model.Label
	selbuild.flds, selbuild.scan = &mq.fields, selbuild.Scan
	return selbuild
}

func (mq *ModelQuery) prepareQuery(ctx context.Context) error {
	for _, f := range mq.fields {
		if !model.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cwmodel: invalid field %q for query", f)}
		}
	}
	if mq.path != nil {
		prev, err := mq.path(ctx)
		if err != nil {
			return err
		}
		mq.sql = prev
	}
	return nil
}

func (mq *ModelQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Model, error) {
	var (
		nodes       = []*Model{}
		_spec       = mq.querySpec()
		loadedTypes = [1]bool{
			mq.withFirmware != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Model).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Model{config: mq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, mq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := mq.withFirmware; query != nil {
		if err := mq.loadFirmware(ctx, query, nodes,
			func(n *Model) { n.Edges.Firmware = []*Firmware{} },
			func(n *Model, e *Firmware) { n.Edges.Firmware = append(n.Edges.Firmware, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (mq *ModelQuery) loadFirmware(ctx context.Context, query *FirmwareQuery, nodes []*Model, init func(*Model), assign func(*Model, *Firmware)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[datasource.UUID]*Model)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.Firmware(func(s *sql.Selector) {
		s.Where(sql.InValues(model.FirmwareColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ModelID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "model_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (mq *ModelQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := mq.querySpec()
	_spec.Node.Columns = mq.fields
	if len(mq.fields) > 0 {
		_spec.Unique = mq.unique != nil && *mq.unique
	}
	return sqlgraph.CountNodes(ctx, mq.driver, _spec)
}

func (mq *ModelQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := mq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cwmodel: check existence: %w", err)
	default:
		return true, nil
	}
}

func (mq *ModelQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   model.Table,
			Columns: model.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: model.FieldID,
			},
		},
		From:   mq.sql,
		Unique: true,
	}
	if unique := mq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := mq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, model.FieldID)
		for i := range fields {
			if fields[i] != model.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := mq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := mq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := mq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := mq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (mq *ModelQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(mq.driver.Dialect())
	t1 := builder.Table(model.Table)
	columns := mq.fields
	if len(columns) == 0 {
		columns = model.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if mq.sql != nil {
		selector = mq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if mq.unique != nil && *mq.unique {
		selector.Distinct()
	}
	for _, p := range mq.predicates {
		p(selector)
	}
	for _, p := range mq.order {
		p(selector)
	}
	if offset := mq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := mq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ModelGroupBy is the group-by builder for Model entities.
type ModelGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (mgb *ModelGroupBy) Aggregate(fns ...AggregateFunc) *ModelGroupBy {
	mgb.fns = append(mgb.fns, fns...)
	return mgb
}

// Scan applies the group-by query and scans the result into the given value.
func (mgb *ModelGroupBy) Scan(ctx context.Context, v any) error {
	query, err := mgb.path(ctx)
	if err != nil {
		return err
	}
	mgb.sql = query
	return mgb.sqlScan(ctx, v)
}

func (mgb *ModelGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range mgb.fields {
		if !model.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := mgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (mgb *ModelGroupBy) sqlQuery() *sql.Selector {
	selector := mgb.sql.Select()
	aggregation := make([]string, 0, len(mgb.fns))
	for _, fn := range mgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(mgb.fields)+len(mgb.fns))
		for _, f := range mgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(mgb.fields...)...)
}

// ModelSelect is the builder for selecting fields of Model entities.
type ModelSelect struct {
	*ModelQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ms *ModelSelect) Scan(ctx context.Context, v any) error {
	if err := ms.prepareQuery(ctx); err != nil {
		return err
	}
	ms.sql = ms.ModelQuery.sqlQuery(ctx)
	return ms.sqlScan(ctx, v)
}

func (ms *ModelSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := ms.sql.Query()
	if err := ms.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
