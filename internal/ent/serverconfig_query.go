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
	"github.com/FM1337/ASB/internal/ent/predicate"
	"github.com/FM1337/ASB/internal/ent/server"
	"github.com/FM1337/ASB/internal/ent/serverconfig"
)

// ServerConfigQuery is the builder for querying ServerConfig entities.
type ServerConfigQuery struct {
	config
	ctx        *QueryContext
	order      []serverconfig.OrderOption
	inters     []Interceptor
	predicates []predicate.ServerConfig
	withServer *ServerQuery
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ServerConfigQuery builder.
func (scq *ServerConfigQuery) Where(ps ...predicate.ServerConfig) *ServerConfigQuery {
	scq.predicates = append(scq.predicates, ps...)
	return scq
}

// Limit the number of records to be returned by this query.
func (scq *ServerConfigQuery) Limit(limit int) *ServerConfigQuery {
	scq.ctx.Limit = &limit
	return scq
}

// Offset to start from.
func (scq *ServerConfigQuery) Offset(offset int) *ServerConfigQuery {
	scq.ctx.Offset = &offset
	return scq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (scq *ServerConfigQuery) Unique(unique bool) *ServerConfigQuery {
	scq.ctx.Unique = &unique
	return scq
}

// Order specifies how the records should be ordered.
func (scq *ServerConfigQuery) Order(o ...serverconfig.OrderOption) *ServerConfigQuery {
	scq.order = append(scq.order, o...)
	return scq
}

// QueryServer chains the current query on the "server" edge.
func (scq *ServerConfigQuery) QueryServer() *ServerQuery {
	query := (&ServerClient{config: scq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := scq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := scq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(serverconfig.Table, serverconfig.FieldID, selector),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, serverconfig.ServerTable, serverconfig.ServerColumn),
		)
		fromU = sqlgraph.SetNeighbors(scq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ServerConfig entity from the query.
// Returns a *NotFoundError when no ServerConfig was found.
func (scq *ServerConfigQuery) First(ctx context.Context) (*ServerConfig, error) {
	nodes, err := scq.Limit(1).All(setContextOp(ctx, scq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{serverconfig.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (scq *ServerConfigQuery) FirstX(ctx context.Context) *ServerConfig {
	node, err := scq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ServerConfig ID from the query.
// Returns a *NotFoundError when no ServerConfig ID was found.
func (scq *ServerConfigQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = scq.Limit(1).IDs(setContextOp(ctx, scq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{serverconfig.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (scq *ServerConfigQuery) FirstIDX(ctx context.Context) int {
	id, err := scq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ServerConfig entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ServerConfig entity is found.
// Returns a *NotFoundError when no ServerConfig entities are found.
func (scq *ServerConfigQuery) Only(ctx context.Context) (*ServerConfig, error) {
	nodes, err := scq.Limit(2).All(setContextOp(ctx, scq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{serverconfig.Label}
	default:
		return nil, &NotSingularError{serverconfig.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (scq *ServerConfigQuery) OnlyX(ctx context.Context) *ServerConfig {
	node, err := scq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ServerConfig ID in the query.
// Returns a *NotSingularError when more than one ServerConfig ID is found.
// Returns a *NotFoundError when no entities are found.
func (scq *ServerConfigQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = scq.Limit(2).IDs(setContextOp(ctx, scq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{serverconfig.Label}
	default:
		err = &NotSingularError{serverconfig.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (scq *ServerConfigQuery) OnlyIDX(ctx context.Context) int {
	id, err := scq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ServerConfigs.
func (scq *ServerConfigQuery) All(ctx context.Context) ([]*ServerConfig, error) {
	ctx = setContextOp(ctx, scq.ctx, "All")
	if err := scq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ServerConfig, *ServerConfigQuery]()
	return withInterceptors[[]*ServerConfig](ctx, scq, qr, scq.inters)
}

// AllX is like All, but panics if an error occurs.
func (scq *ServerConfigQuery) AllX(ctx context.Context) []*ServerConfig {
	nodes, err := scq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ServerConfig IDs.
func (scq *ServerConfigQuery) IDs(ctx context.Context) (ids []int, err error) {
	if scq.ctx.Unique == nil && scq.path != nil {
		scq.Unique(true)
	}
	ctx = setContextOp(ctx, scq.ctx, "IDs")
	if err = scq.Select(serverconfig.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (scq *ServerConfigQuery) IDsX(ctx context.Context) []int {
	ids, err := scq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (scq *ServerConfigQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, scq.ctx, "Count")
	if err := scq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, scq, querierCount[*ServerConfigQuery](), scq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (scq *ServerConfigQuery) CountX(ctx context.Context) int {
	count, err := scq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (scq *ServerConfigQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, scq.ctx, "Exist")
	switch _, err := scq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (scq *ServerConfigQuery) ExistX(ctx context.Context) bool {
	exist, err := scq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ServerConfigQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (scq *ServerConfigQuery) Clone() *ServerConfigQuery {
	if scq == nil {
		return nil
	}
	return &ServerConfigQuery{
		config:     scq.config,
		ctx:        scq.ctx.Clone(),
		order:      append([]serverconfig.OrderOption{}, scq.order...),
		inters:     append([]Interceptor{}, scq.inters...),
		predicates: append([]predicate.ServerConfig{}, scq.predicates...),
		withServer: scq.withServer.Clone(),
		// clone intermediate query.
		sql:  scq.sql.Clone(),
		path: scq.path,
	}
}

// WithServer tells the query-builder to eager-load the nodes that are connected to
// the "server" edge. The optional arguments are used to configure the query builder of the edge.
func (scq *ServerConfigQuery) WithServer(opts ...func(*ServerQuery)) *ServerConfigQuery {
	query := (&ServerClient{config: scq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	scq.withServer = query
	return scq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ServerConfig.Query().
//		GroupBy(serverconfig.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (scq *ServerConfigQuery) GroupBy(field string, fields ...string) *ServerConfigGroupBy {
	scq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ServerConfigGroupBy{build: scq}
	grbuild.flds = &scq.ctx.Fields
	grbuild.label = serverconfig.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.ServerConfig.Query().
//		Select(serverconfig.FieldCreateTime).
//		Scan(ctx, &v)
func (scq *ServerConfigQuery) Select(fields ...string) *ServerConfigSelect {
	scq.ctx.Fields = append(scq.ctx.Fields, fields...)
	sbuild := &ServerConfigSelect{ServerConfigQuery: scq}
	sbuild.label = serverconfig.Label
	sbuild.flds, sbuild.scan = &scq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ServerConfigSelect configured with the given aggregations.
func (scq *ServerConfigQuery) Aggregate(fns ...AggregateFunc) *ServerConfigSelect {
	return scq.Select().Aggregate(fns...)
}

func (scq *ServerConfigQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range scq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, scq); err != nil {
				return err
			}
		}
	}
	for _, f := range scq.ctx.Fields {
		if !serverconfig.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if scq.path != nil {
		prev, err := scq.path(ctx)
		if err != nil {
			return err
		}
		scq.sql = prev
	}
	return nil
}

func (scq *ServerConfigQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ServerConfig, error) {
	var (
		nodes       = []*ServerConfig{}
		_spec       = scq.querySpec()
		loadedTypes = [1]bool{
			scq.withServer != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ServerConfig).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ServerConfig{config: scq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, scq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := scq.withServer; query != nil {
		if err := scq.loadServer(ctx, query, nodes,
			func(n *ServerConfig) { n.Edges.Server = []*Server{} },
			func(n *ServerConfig, e *Server) { n.Edges.Server = append(n.Edges.Server, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (scq *ServerConfigQuery) loadServer(ctx context.Context, query *ServerQuery, nodes []*ServerConfig, init func(*ServerConfig), assign func(*ServerConfig, *Server)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ServerConfig)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Server(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(serverconfig.ServerColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.server_configuration
		if fk == nil {
			return fmt.Errorf(`foreign-key "server_configuration" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "server_configuration" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (scq *ServerConfigQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := scq.querySpec()
	if len(scq.modifiers) > 0 {
		_spec.Modifiers = scq.modifiers
	}
	_spec.Node.Columns = scq.ctx.Fields
	if len(scq.ctx.Fields) > 0 {
		_spec.Unique = scq.ctx.Unique != nil && *scq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, scq.driver, _spec)
}

func (scq *ServerConfigQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(serverconfig.Table, serverconfig.Columns, sqlgraph.NewFieldSpec(serverconfig.FieldID, field.TypeInt))
	_spec.From = scq.sql
	if unique := scq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if scq.path != nil {
		_spec.Unique = true
	}
	if fields := scq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serverconfig.FieldID)
		for i := range fields {
			if fields[i] != serverconfig.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := scq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := scq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := scq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := scq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (scq *ServerConfigQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(scq.driver.Dialect())
	t1 := builder.Table(serverconfig.Table)
	columns := scq.ctx.Fields
	if len(columns) == 0 {
		columns = serverconfig.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if scq.sql != nil {
		selector = scq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if scq.ctx.Unique != nil && *scq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range scq.modifiers {
		m(selector)
	}
	for _, p := range scq.predicates {
		p(selector)
	}
	for _, p := range scq.order {
		p(selector)
	}
	if offset := scq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := scq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scq *ServerConfigQuery) Modify(modifiers ...func(s *sql.Selector)) *ServerConfigSelect {
	scq.modifiers = append(scq.modifiers, modifiers...)
	return scq.Select()
}

// ServerConfigGroupBy is the group-by builder for ServerConfig entities.
type ServerConfigGroupBy struct {
	selector
	build *ServerConfigQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (scgb *ServerConfigGroupBy) Aggregate(fns ...AggregateFunc) *ServerConfigGroupBy {
	scgb.fns = append(scgb.fns, fns...)
	return scgb
}

// Scan applies the selector query and scans the result into the given value.
func (scgb *ServerConfigGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scgb.build.ctx, "GroupBy")
	if err := scgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServerConfigQuery, *ServerConfigGroupBy](ctx, scgb.build, scgb, scgb.build.inters, v)
}

func (scgb *ServerConfigGroupBy) sqlScan(ctx context.Context, root *ServerConfigQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(scgb.fns))
	for _, fn := range scgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*scgb.flds)+len(scgb.fns))
		for _, f := range *scgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*scgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ServerConfigSelect is the builder for selecting fields of ServerConfig entities.
type ServerConfigSelect struct {
	*ServerConfigQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (scs *ServerConfigSelect) Aggregate(fns ...AggregateFunc) *ServerConfigSelect {
	scs.fns = append(scs.fns, fns...)
	return scs
}

// Scan applies the selector query and scans the result into the given value.
func (scs *ServerConfigSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, scs.ctx, "Select")
	if err := scs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ServerConfigQuery, *ServerConfigSelect](ctx, scs.ServerConfigQuery, scs, scs.inters, v)
}

func (scs *ServerConfigSelect) sqlScan(ctx context.Context, root *ServerConfigQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(scs.fns))
	for _, fn := range scs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*scs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := scs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (scs *ServerConfigSelect) Modify(modifiers ...func(s *sql.Selector)) *ServerConfigSelect {
	scs.modifiers = append(scs.modifiers, modifiers...)
	return scs
}
