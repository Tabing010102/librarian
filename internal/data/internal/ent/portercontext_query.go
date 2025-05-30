// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tuihub/librarian/internal/data/internal/ent/portercontext"
	"github.com/tuihub/librarian/internal/data/internal/ent/predicate"
	"github.com/tuihub/librarian/internal/data/internal/ent/user"
	"github.com/tuihub/librarian/internal/model"
)

// PorterContextQuery is the builder for querying PorterContext entities.
type PorterContextQuery struct {
	config
	ctx        *QueryContext
	order      []portercontext.OrderOption
	inters     []Interceptor
	predicates []predicate.PorterContext
	withOwner  *UserQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PorterContextQuery builder.
func (pcq *PorterContextQuery) Where(ps ...predicate.PorterContext) *PorterContextQuery {
	pcq.predicates = append(pcq.predicates, ps...)
	return pcq
}

// Limit the number of records to be returned by this query.
func (pcq *PorterContextQuery) Limit(limit int) *PorterContextQuery {
	pcq.ctx.Limit = &limit
	return pcq
}

// Offset to start from.
func (pcq *PorterContextQuery) Offset(offset int) *PorterContextQuery {
	pcq.ctx.Offset = &offset
	return pcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pcq *PorterContextQuery) Unique(unique bool) *PorterContextQuery {
	pcq.ctx.Unique = &unique
	return pcq
}

// Order specifies how the records should be ordered.
func (pcq *PorterContextQuery) Order(o ...portercontext.OrderOption) *PorterContextQuery {
	pcq.order = append(pcq.order, o...)
	return pcq
}

// QueryOwner chains the current query on the "owner" edge.
func (pcq *PorterContextQuery) QueryOwner() *UserQuery {
	query := (&UserClient{config: pcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(portercontext.Table, portercontext.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, portercontext.OwnerTable, portercontext.OwnerColumn),
		)
		fromU = sqlgraph.SetNeighbors(pcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first PorterContext entity from the query.
// Returns a *NotFoundError when no PorterContext was found.
func (pcq *PorterContextQuery) First(ctx context.Context) (*PorterContext, error) {
	nodes, err := pcq.Limit(1).All(setContextOp(ctx, pcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{portercontext.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pcq *PorterContextQuery) FirstX(ctx context.Context) *PorterContext {
	node, err := pcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PorterContext ID from the query.
// Returns a *NotFoundError when no PorterContext ID was found.
func (pcq *PorterContextQuery) FirstID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = pcq.Limit(1).IDs(setContextOp(ctx, pcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{portercontext.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pcq *PorterContextQuery) FirstIDX(ctx context.Context) model.InternalID {
	id, err := pcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PorterContext entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PorterContext entity is found.
// Returns a *NotFoundError when no PorterContext entities are found.
func (pcq *PorterContextQuery) Only(ctx context.Context) (*PorterContext, error) {
	nodes, err := pcq.Limit(2).All(setContextOp(ctx, pcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{portercontext.Label}
	default:
		return nil, &NotSingularError{portercontext.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pcq *PorterContextQuery) OnlyX(ctx context.Context) *PorterContext {
	node, err := pcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PorterContext ID in the query.
// Returns a *NotSingularError when more than one PorterContext ID is found.
// Returns a *NotFoundError when no entities are found.
func (pcq *PorterContextQuery) OnlyID(ctx context.Context) (id model.InternalID, err error) {
	var ids []model.InternalID
	if ids, err = pcq.Limit(2).IDs(setContextOp(ctx, pcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{portercontext.Label}
	default:
		err = &NotSingularError{portercontext.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pcq *PorterContextQuery) OnlyIDX(ctx context.Context) model.InternalID {
	id, err := pcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PorterContexts.
func (pcq *PorterContextQuery) All(ctx context.Context) ([]*PorterContext, error) {
	ctx = setContextOp(ctx, pcq.ctx, ent.OpQueryAll)
	if err := pcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PorterContext, *PorterContextQuery]()
	return withInterceptors[[]*PorterContext](ctx, pcq, qr, pcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pcq *PorterContextQuery) AllX(ctx context.Context) []*PorterContext {
	nodes, err := pcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PorterContext IDs.
func (pcq *PorterContextQuery) IDs(ctx context.Context) (ids []model.InternalID, err error) {
	if pcq.ctx.Unique == nil && pcq.path != nil {
		pcq.Unique(true)
	}
	ctx = setContextOp(ctx, pcq.ctx, ent.OpQueryIDs)
	if err = pcq.Select(portercontext.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pcq *PorterContextQuery) IDsX(ctx context.Context) []model.InternalID {
	ids, err := pcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pcq *PorterContextQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pcq.ctx, ent.OpQueryCount)
	if err := pcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pcq, querierCount[*PorterContextQuery](), pcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pcq *PorterContextQuery) CountX(ctx context.Context) int {
	count, err := pcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pcq *PorterContextQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pcq.ctx, ent.OpQueryExist)
	switch _, err := pcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pcq *PorterContextQuery) ExistX(ctx context.Context) bool {
	exist, err := pcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PorterContextQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pcq *PorterContextQuery) Clone() *PorterContextQuery {
	if pcq == nil {
		return nil
	}
	return &PorterContextQuery{
		config:     pcq.config,
		ctx:        pcq.ctx.Clone(),
		order:      append([]portercontext.OrderOption{}, pcq.order...),
		inters:     append([]Interceptor{}, pcq.inters...),
		predicates: append([]predicate.PorterContext{}, pcq.predicates...),
		withOwner:  pcq.withOwner.Clone(),
		// clone intermediate query.
		sql:  pcq.sql.Clone(),
		path: pcq.path,
	}
}

// WithOwner tells the query-builder to eager-load the nodes that are connected to
// the "owner" edge. The optional arguments are used to configure the query builder of the edge.
func (pcq *PorterContextQuery) WithOwner(opts ...func(*UserQuery)) *PorterContextQuery {
	query := (&UserClient{config: pcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pcq.withOwner = query
	return pcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		GlobalName string `json:"global_name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PorterContext.Query().
//		GroupBy(portercontext.FieldGlobalName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pcq *PorterContextQuery) GroupBy(field string, fields ...string) *PorterContextGroupBy {
	pcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PorterContextGroupBy{build: pcq}
	grbuild.flds = &pcq.ctx.Fields
	grbuild.label = portercontext.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		GlobalName string `json:"global_name,omitempty"`
//	}
//
//	client.PorterContext.Query().
//		Select(portercontext.FieldGlobalName).
//		Scan(ctx, &v)
func (pcq *PorterContextQuery) Select(fields ...string) *PorterContextSelect {
	pcq.ctx.Fields = append(pcq.ctx.Fields, fields...)
	sbuild := &PorterContextSelect{PorterContextQuery: pcq}
	sbuild.label = portercontext.Label
	sbuild.flds, sbuild.scan = &pcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PorterContextSelect configured with the given aggregations.
func (pcq *PorterContextQuery) Aggregate(fns ...AggregateFunc) *PorterContextSelect {
	return pcq.Select().Aggregate(fns...)
}

func (pcq *PorterContextQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pcq); err != nil {
				return err
			}
		}
	}
	for _, f := range pcq.ctx.Fields {
		if !portercontext.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pcq.path != nil {
		prev, err := pcq.path(ctx)
		if err != nil {
			return err
		}
		pcq.sql = prev
	}
	return nil
}

func (pcq *PorterContextQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PorterContext, error) {
	var (
		nodes       = []*PorterContext{}
		withFKs     = pcq.withFKs
		_spec       = pcq.querySpec()
		loadedTypes = [1]bool{
			pcq.withOwner != nil,
		}
	)
	if pcq.withOwner != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, portercontext.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PorterContext).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PorterContext{config: pcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pcq.withOwner; query != nil {
		if err := pcq.loadOwner(ctx, query, nodes, nil,
			func(n *PorterContext, e *User) { n.Edges.Owner = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pcq *PorterContextQuery) loadOwner(ctx context.Context, query *UserQuery, nodes []*PorterContext, init func(*PorterContext), assign func(*PorterContext, *User)) error {
	ids := make([]model.InternalID, 0, len(nodes))
	nodeids := make(map[model.InternalID][]*PorterContext)
	for i := range nodes {
		if nodes[i].user_porter_context == nil {
			continue
		}
		fk := *nodes[i].user_porter_context
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_porter_context" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pcq *PorterContextQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pcq.querySpec()
	_spec.Node.Columns = pcq.ctx.Fields
	if len(pcq.ctx.Fields) > 0 {
		_spec.Unique = pcq.ctx.Unique != nil && *pcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pcq.driver, _spec)
}

func (pcq *PorterContextQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(portercontext.Table, portercontext.Columns, sqlgraph.NewFieldSpec(portercontext.FieldID, field.TypeInt64))
	_spec.From = pcq.sql
	if unique := pcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pcq.path != nil {
		_spec.Unique = true
	}
	if fields := pcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, portercontext.FieldID)
		for i := range fields {
			if fields[i] != portercontext.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pcq *PorterContextQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pcq.driver.Dialect())
	t1 := builder.Table(portercontext.Table)
	columns := pcq.ctx.Fields
	if len(columns) == 0 {
		columns = portercontext.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pcq.sql != nil {
		selector = pcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pcq.ctx.Unique != nil && *pcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pcq.predicates {
		p(selector)
	}
	for _, p := range pcq.order {
		p(selector)
	}
	if offset := pcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PorterContextGroupBy is the group-by builder for PorterContext entities.
type PorterContextGroupBy struct {
	selector
	build *PorterContextQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pcgb *PorterContextGroupBy) Aggregate(fns ...AggregateFunc) *PorterContextGroupBy {
	pcgb.fns = append(pcgb.fns, fns...)
	return pcgb
}

// Scan applies the selector query and scans the result into the given value.
func (pcgb *PorterContextGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcgb.build.ctx, ent.OpQueryGroupBy)
	if err := pcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PorterContextQuery, *PorterContextGroupBy](ctx, pcgb.build, pcgb, pcgb.build.inters, v)
}

func (pcgb *PorterContextGroupBy) sqlScan(ctx context.Context, root *PorterContextQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pcgb.fns))
	for _, fn := range pcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pcgb.flds)+len(pcgb.fns))
		for _, f := range *pcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PorterContextSelect is the builder for selecting fields of PorterContext entities.
type PorterContextSelect struct {
	*PorterContextQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pcs *PorterContextSelect) Aggregate(fns ...AggregateFunc) *PorterContextSelect {
	pcs.fns = append(pcs.fns, fns...)
	return pcs
}

// Scan applies the selector query and scans the result into the given value.
func (pcs *PorterContextSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pcs.ctx, ent.OpQuerySelect)
	if err := pcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PorterContextQuery, *PorterContextSelect](ctx, pcs.PorterContextQuery, pcs, pcs.inters, v)
}

func (pcs *PorterContextSelect) sqlScan(ctx context.Context, root *PorterContextQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pcs.fns))
	for _, fn := range pcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
