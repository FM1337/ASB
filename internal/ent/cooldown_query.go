// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/FM1337/ASB/internal/ent/cooldown"
	"github.com/FM1337/ASB/internal/ent/predicate"
	"github.com/FM1337/ASB/internal/ent/server"
)

// CooldownQuery is the builder for querying Cooldown entities.
type CooldownQuery struct {
	config
	ctx        *QueryContext
	order      []cooldown.OrderOption
	inters     []Interceptor
	predicates []predicate.Cooldown
	withServer *ServerQuery
	withFKs    bool
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CooldownQuery builder.
func (cq *CooldownQuery) Where(ps ...predicate.Cooldown) *CooldownQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CooldownQuery) Limit(limit int) *CooldownQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CooldownQuery) Offset(offset int) *CooldownQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CooldownQuery) Unique(unique bool) *CooldownQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CooldownQuery) Order(o ...cooldown.OrderOption) *CooldownQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryServer chains the current query on the "server" edge.
func (cq *CooldownQuery) QueryServer() *ServerQuery {
	query := (&ServerClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(cooldown.Table, cooldown.FieldID, selector),
			sqlgraph.To(server.Table, server.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, cooldown.ServerTable, cooldown.ServerColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Cooldown entity from the query.
// Returns a *NotFoundError when no Cooldown was found.
func (cq *CooldownQuery) First(ctx context.Context) (*Cooldown, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{cooldown.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CooldownQuery) FirstX(ctx context.Context) *Cooldown {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Cooldown ID from the query.
// Returns a *NotFoundError when no Cooldown ID was found.
func (cq *CooldownQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{cooldown.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CooldownQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Cooldown entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Cooldown entity is found.
// Returns a *NotFoundError when no Cooldown entities are found.
func (cq *CooldownQuery) Only(ctx context.Context) (*Cooldown, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{cooldown.Label}
	default:
		return nil, &NotSingularError{cooldown.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CooldownQuery) OnlyX(ctx context.Context) *Cooldown {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Cooldown ID in the query.
// Returns a *NotSingularError when more than one Cooldown ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CooldownQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{cooldown.Label}
	default:
		err = &NotSingularError{cooldown.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CooldownQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Cooldowns.
func (cq *CooldownQuery) All(ctx context.Context) ([]*Cooldown, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Cooldown, *CooldownQuery]()
	return withInterceptors[[]*Cooldown](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CooldownQuery) AllX(ctx context.Context) []*Cooldown {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Cooldown IDs.
func (cq *CooldownQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(cooldown.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CooldownQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CooldownQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CooldownQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CooldownQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CooldownQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CooldownQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CooldownQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CooldownQuery) Clone() *CooldownQuery {
	if cq == nil {
		return nil
	}
	return &CooldownQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]cooldown.OrderOption{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Cooldown{}, cq.predicates...),
		withServer: cq.withServer.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithServer tells the query-builder to eager-load the nodes that are connected to
// the "server" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CooldownQuery) WithServer(opts ...func(*ServerQuery)) *CooldownQuery {
	query := (&ServerClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withServer = query
	return cq
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
//	client.Cooldown.Query().
//		GroupBy(cooldown.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CooldownQuery) GroupBy(field string, fields ...string) *CooldownGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CooldownGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = cooldown.Label
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
//	client.Cooldown.Query().
//		Select(cooldown.FieldCreateTime).
//		Scan(ctx, &v)
func (cq *CooldownQuery) Select(fields ...string) *CooldownSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CooldownSelect{CooldownQuery: cq}
	sbuild.label = cooldown.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CooldownSelect configured with the given aggregations.
func (cq *CooldownQuery) Aggregate(fns ...AggregateFunc) *CooldownSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CooldownQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !cooldown.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
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

func (cq *CooldownQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Cooldown, error) {
	var (
		nodes       = []*Cooldown{}
		withFKs     = cq.withFKs
		_spec       = cq.querySpec()
		loadedTypes = [1]bool{
			cq.withServer != nil,
		}
	)
	if cq.withServer != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, cooldown.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Cooldown).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Cooldown{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
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
	if query := cq.withServer; query != nil {
		if err := cq.loadServer(ctx, query, nodes, nil,
			func(n *Cooldown, e *Server) { n.Edges.Server = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CooldownQuery) loadServer(ctx context.Context, query *ServerQuery, nodes []*Cooldown, init func(*Cooldown), assign func(*Cooldown, *Server)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Cooldown)
	for i := range nodes {
		if nodes[i].server_cooldown == nil {
			continue
		}
		fk := *nodes[i].server_cooldown
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(server.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "server_cooldown" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cq *CooldownQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CooldownQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(cooldown.Table, cooldown.Columns, sqlgraph.NewFieldSpec(cooldown.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, cooldown.FieldID)
		for i := range fields {
			if fields[i] != cooldown.FieldID {
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
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
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

func (cq *CooldownQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(cooldown.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = cooldown.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range cq.modifiers {
		m(selector)
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cq *CooldownQuery) Modify(modifiers ...func(s *sql.Selector)) *CooldownSelect {
	cq.modifiers = append(cq.modifiers, modifiers...)
	return cq.Select()
}

// CooldownGroupBy is the group-by builder for Cooldown entities.
type CooldownGroupBy struct {
	selector
	build *CooldownQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CooldownGroupBy) Aggregate(fns ...AggregateFunc) *CooldownGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CooldownGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CooldownQuery, *CooldownGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CooldownGroupBy) sqlScan(ctx context.Context, root *CooldownQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CooldownSelect is the builder for selecting fields of Cooldown entities.
type CooldownSelect struct {
	*CooldownQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CooldownSelect) Aggregate(fns ...AggregateFunc) *CooldownSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CooldownSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CooldownQuery, *CooldownSelect](ctx, cs.CooldownQuery, cs, cs.inters, v)
}

func (cs *CooldownSelect) sqlScan(ctx context.Context, root *CooldownQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (cs *CooldownSelect) Modify(modifiers ...func(s *sql.Selector)) *CooldownSelect {
	cs.modifiers = append(cs.modifiers, modifiers...)
	return cs
}