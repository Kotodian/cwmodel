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
	"github.com/Kotodian/cwmodel/orderinfo"
	"github.com/Kotodian/cwmodel/predicate"
	"github.com/Kotodian/cwmodel/reservation"
	"github.com/Kotodian/cwmodel/smartchargingeffect"
	"github.com/Kotodian/gokit/datasource"
)

// ConnectorQuery is the builder for querying Connector entities.
type ConnectorQuery struct {
	config
	limit                   *int
	offset                  *int
	unique                  *bool
	order                   []OrderFunc
	fields                  []string
	predicates              []predicate.Connector
	withEvse                *EvseQuery
	withEquipment           *EquipmentQuery
	withOrderInfo           *OrderInfoQuery
	withReservation         *ReservationQuery
	withSmartChargingEffect *SmartChargingEffectQuery
	withFKs                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ConnectorQuery builder.
func (cq *ConnectorQuery) Where(ps ...predicate.Connector) *ConnectorQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit adds a limit step to the query.
func (cq *ConnectorQuery) Limit(limit int) *ConnectorQuery {
	cq.limit = &limit
	return cq
}

// Offset adds an offset step to the query.
func (cq *ConnectorQuery) Offset(offset int) *ConnectorQuery {
	cq.offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *ConnectorQuery) Unique(unique bool) *ConnectorQuery {
	cq.unique = &unique
	return cq
}

// Order adds an order step to the query.
func (cq *ConnectorQuery) Order(o ...OrderFunc) *ConnectorQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryEvse chains the current query on the "evse" edge.
func (cq *ConnectorQuery) QueryEvse() *EvseQuery {
	query := &EvseQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(evse.Table, evse.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, connector.EvseTable, connector.EvseColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEquipment chains the current query on the "equipment" edge.
func (cq *ConnectorQuery) QueryEquipment() *EquipmentQuery {
	query := &EquipmentQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(equipment.Table, equipment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, connector.EquipmentTable, connector.EquipmentColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryOrderInfo chains the current query on the "order_info" edge.
func (cq *ConnectorQuery) QueryOrderInfo() *OrderInfoQuery {
	query := &OrderInfoQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(orderinfo.Table, orderinfo.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, connector.OrderInfoTable, connector.OrderInfoColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryReservation chains the current query on the "reservation" edge.
func (cq *ConnectorQuery) QueryReservation() *ReservationQuery {
	query := &ReservationQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(reservation.Table, reservation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, connector.ReservationTable, connector.ReservationColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySmartChargingEffect chains the current query on the "smart_charging_effect" edge.
func (cq *ConnectorQuery) QuerySmartChargingEffect() *SmartChargingEffectQuery {
	query := &SmartChargingEffectQuery{config: cq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(connector.Table, connector.FieldID, selector),
			sqlgraph.To(smartchargingeffect.Table, smartchargingeffect.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, connector.SmartChargingEffectTable, connector.SmartChargingEffectColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Connector entity from the query.
// Returns a *NotFoundError when no Connector was found.
func (cq *ConnectorQuery) First(ctx context.Context) (*Connector, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{connector.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *ConnectorQuery) FirstX(ctx context.Context) *Connector {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Connector ID from the query.
// Returns a *NotFoundError when no Connector ID was found.
func (cq *ConnectorQuery) FirstID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = cq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{connector.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *ConnectorQuery) FirstIDX(ctx context.Context) datasource.UUID {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Connector entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Connector entity is found.
// Returns a *NotFoundError when no Connector entities are found.
func (cq *ConnectorQuery) Only(ctx context.Context) (*Connector, error) {
	nodes, err := cq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{connector.Label}
	default:
		return nil, &NotSingularError{connector.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *ConnectorQuery) OnlyX(ctx context.Context) *Connector {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Connector ID in the query.
// Returns a *NotSingularError when more than one Connector ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *ConnectorQuery) OnlyID(ctx context.Context) (id datasource.UUID, err error) {
	var ids []datasource.UUID
	if ids, err = cq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{connector.Label}
	default:
		err = &NotSingularError{connector.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *ConnectorQuery) OnlyIDX(ctx context.Context) datasource.UUID {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Connectors.
func (cq *ConnectorQuery) All(ctx context.Context) ([]*Connector, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return cq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (cq *ConnectorQuery) AllX(ctx context.Context) []*Connector {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Connector IDs.
func (cq *ConnectorQuery) IDs(ctx context.Context) ([]datasource.UUID, error) {
	var ids []datasource.UUID
	if err := cq.Select(connector.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *ConnectorQuery) IDsX(ctx context.Context) []datasource.UUID {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *ConnectorQuery) Count(ctx context.Context) (int, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return cq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (cq *ConnectorQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *ConnectorQuery) Exist(ctx context.Context) (bool, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return cq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *ConnectorQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ConnectorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *ConnectorQuery) Clone() *ConnectorQuery {
	if cq == nil {
		return nil
	}
	return &ConnectorQuery{
		config:                  cq.config,
		limit:                   cq.limit,
		offset:                  cq.offset,
		order:                   append([]OrderFunc{}, cq.order...),
		predicates:              append([]predicate.Connector{}, cq.predicates...),
		withEvse:                cq.withEvse.Clone(),
		withEquipment:           cq.withEquipment.Clone(),
		withOrderInfo:           cq.withOrderInfo.Clone(),
		withReservation:         cq.withReservation.Clone(),
		withSmartChargingEffect: cq.withSmartChargingEffect.Clone(),
		// clone intermediate query.
		sql:    cq.sql.Clone(),
		path:   cq.path,
		unique: cq.unique,
	}
}

// WithEvse tells the query-builder to eager-load the nodes that are connected to
// the "evse" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithEvse(opts ...func(*EvseQuery)) *ConnectorQuery {
	query := &EvseQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withEvse = query
	return cq
}

// WithEquipment tells the query-builder to eager-load the nodes that are connected to
// the "equipment" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithEquipment(opts ...func(*EquipmentQuery)) *ConnectorQuery {
	query := &EquipmentQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withEquipment = query
	return cq
}

// WithOrderInfo tells the query-builder to eager-load the nodes that are connected to
// the "order_info" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithOrderInfo(opts ...func(*OrderInfoQuery)) *ConnectorQuery {
	query := &OrderInfoQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withOrderInfo = query
	return cq
}

// WithReservation tells the query-builder to eager-load the nodes that are connected to
// the "reservation" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithReservation(opts ...func(*ReservationQuery)) *ConnectorQuery {
	query := &ReservationQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withReservation = query
	return cq
}

// WithSmartChargingEffect tells the query-builder to eager-load the nodes that are connected to
// the "smart_charging_effect" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *ConnectorQuery) WithSmartChargingEffect(opts ...func(*SmartChargingEffectQuery)) *ConnectorQuery {
	query := &SmartChargingEffectQuery{config: cq.config}
	for _, opt := range opts {
		opt(query)
	}
	cq.withSmartChargingEffect = query
	return cq
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
//	client.Connector.Query().
//		GroupBy(connector.FieldVersion).
//		Aggregate(cwmodel.Count()).
//		Scan(ctx, &v)
func (cq *ConnectorQuery) GroupBy(field string, fields ...string) *ConnectorGroupBy {
	grbuild := &ConnectorGroupBy{config: cq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return cq.sqlQuery(ctx), nil
	}
	grbuild.label = connector.Label
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
//	client.Connector.Query().
//		Select(connector.FieldVersion).
//		Scan(ctx, &v)
func (cq *ConnectorQuery) Select(fields ...string) *ConnectorSelect {
	cq.fields = append(cq.fields, fields...)
	selbuild := &ConnectorSelect{ConnectorQuery: cq}
	selbuild.label = connector.Label
	selbuild.flds, selbuild.scan = &cq.fields, selbuild.Scan
	return selbuild
}

func (cq *ConnectorQuery) prepareQuery(ctx context.Context) error {
	for _, f := range cq.fields {
		if !connector.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("cwmodel: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *ConnectorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Connector, error) {
	var (
		nodes       = []*Connector{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [5]bool{
			cq.withEvse != nil,
			cq.withEquipment != nil,
			cq.withOrderInfo != nil,
			cq.withReservation != nil,
			cq.withSmartChargingEffect != nil,
		}
	)
	if cq.withEvse != nil || cq.withEquipment != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, connector.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Connector).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Connector{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withEvse; query != nil {
		if err := cq.loadEvse(ctx, query, nodes, nil,
			func(n *Connector, e *Evse) { n.Edges.Evse = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withEquipment; query != nil {
		if err := cq.loadEquipment(ctx, query, nodes, nil,
			func(n *Connector, e *Equipment) { n.Edges.Equipment = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withOrderInfo; query != nil {
		if err := cq.loadOrderInfo(ctx, query, nodes,
			func(n *Connector) { n.Edges.OrderInfo = []*OrderInfo{} },
			func(n *Connector, e *OrderInfo) { n.Edges.OrderInfo = append(n.Edges.OrderInfo, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withReservation; query != nil {
		if err := cq.loadReservation(ctx, query, nodes,
			func(n *Connector) { n.Edges.Reservation = []*Reservation{} },
			func(n *Connector, e *Reservation) { n.Edges.Reservation = append(n.Edges.Reservation, e) }); err != nil {
			return nil, err
		}
	}
	if query := cq.withSmartChargingEffect; query != nil {
		if err := cq.loadSmartChargingEffect(ctx, query, nodes,
			func(n *Connector) { n.Edges.SmartChargingEffect = []*SmartChargingEffect{} },
			func(n *Connector, e *SmartChargingEffect) {
				n.Edges.SmartChargingEffect = append(n.Edges.SmartChargingEffect, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *ConnectorQuery) loadEvse(ctx context.Context, query *EvseQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *Evse)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*Connector)
	for i := range nodes {
		if nodes[i].evse_id == nil {
			continue
		}
		fk := *nodes[i].evse_id
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(evse.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "evse_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *ConnectorQuery) loadEquipment(ctx context.Context, query *EquipmentQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *Equipment)) error {
	ids := make([]datasource.UUID, 0, len(nodes))
	nodeids := make(map[datasource.UUID][]*Connector)
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
func (cq *ConnectorQuery) loadOrderInfo(ctx context.Context, query *OrderInfoQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *OrderInfo)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[datasource.UUID]*Connector)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.OrderInfo(func(s *sql.Selector) {
		s.Where(sql.InValues(connector.OrderInfoColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.connector_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "connector_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "connector_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ConnectorQuery) loadReservation(ctx context.Context, query *ReservationQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *Reservation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[datasource.UUID]*Connector)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Reservation(func(s *sql.Selector) {
		s.Where(sql.InValues(connector.ReservationColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.connector_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "connector_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "connector_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (cq *ConnectorQuery) loadSmartChargingEffect(ctx context.Context, query *SmartChargingEffectQuery, nodes []*Connector, init func(*Connector), assign func(*Connector, *SmartChargingEffect)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[datasource.UUID]*Connector)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.SmartChargingEffect(func(s *sql.Selector) {
		s.Where(sql.InValues(connector.SmartChargingEffectColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.connector_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "connector_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "connector_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *ConnectorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	_spec.Node.Columns = cq.fields
	if len(cq.fields) > 0 {
		_spec.Unique = cq.unique != nil && *cq.unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *ConnectorQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("cwmodel: check existence: %w", err)
	default:
		return true, nil
	}
}

func (cq *ConnectorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   connector.Table,
			Columns: connector.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: connector.FieldID,
			},
		},
		From:   cq.sql,
		Unique: true,
	}
	if unique := cq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := cq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, connector.FieldID)
		for i := range fields {
			if fields[i] != connector.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *ConnectorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(connector.Table)
	columns := cq.fields
	if len(columns) == 0 {
		columns = connector.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.unique != nil && *cq.unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ConnectorGroupBy is the group-by builder for Connector entities.
type ConnectorGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *ConnectorGroupBy) Aggregate(fns ...AggregateFunc) *ConnectorGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the group-by query and scans the result into the given value.
func (cgb *ConnectorGroupBy) Scan(ctx context.Context, v any) error {
	query, err := cgb.path(ctx)
	if err != nil {
		return err
	}
	cgb.sql = query
	return cgb.sqlScan(ctx, v)
}

func (cgb *ConnectorGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range cgb.fields {
		if !connector.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := cgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (cgb *ConnectorGroupBy) sqlQuery() *sql.Selector {
	selector := cgb.sql.Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(cgb.fields)+len(cgb.fns))
		for _, f := range cgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(cgb.fields...)...)
}

// ConnectorSelect is the builder for selecting fields of Connector entities.
type ConnectorSelect struct {
	*ConnectorQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (cs *ConnectorSelect) Scan(ctx context.Context, v any) error {
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	cs.sql = cs.ConnectorQuery.sqlQuery(ctx)
	return cs.sqlScan(ctx, v)
}

func (cs *ConnectorSelect) sqlScan(ctx context.Context, v any) error {
	rows := &sql.Rows{}
	query, args := cs.sql.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}