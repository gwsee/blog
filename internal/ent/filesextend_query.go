// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/files"
	"blog/internal/ent/filesextend"
	"blog/internal/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FilesExtendQuery is the builder for querying FilesExtend entities.
type FilesExtendQuery struct {
	config
	ctx        *QueryContext
	order      []filesextend.OrderOption
	inters     []Interceptor
	predicates []predicate.FilesExtend
	withFiles  *FilesQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FilesExtendQuery builder.
func (feq *FilesExtendQuery) Where(ps ...predicate.FilesExtend) *FilesExtendQuery {
	feq.predicates = append(feq.predicates, ps...)
	return feq
}

// Limit the number of records to be returned by this query.
func (feq *FilesExtendQuery) Limit(limit int) *FilesExtendQuery {
	feq.ctx.Limit = &limit
	return feq
}

// Offset to start from.
func (feq *FilesExtendQuery) Offset(offset int) *FilesExtendQuery {
	feq.ctx.Offset = &offset
	return feq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (feq *FilesExtendQuery) Unique(unique bool) *FilesExtendQuery {
	feq.ctx.Unique = &unique
	return feq
}

// Order specifies how the records should be ordered.
func (feq *FilesExtendQuery) Order(o ...filesextend.OrderOption) *FilesExtendQuery {
	feq.order = append(feq.order, o...)
	return feq
}

// QueryFiles chains the current query on the "files" edge.
func (feq *FilesExtendQuery) QueryFiles() *FilesQuery {
	query := (&FilesClient{config: feq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := feq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := feq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(filesextend.Table, filesextend.FieldID, selector),
			sqlgraph.To(files.Table, files.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, filesextend.FilesTable, filesextend.FilesColumn),
		)
		fromU = sqlgraph.SetNeighbors(feq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first FilesExtend entity from the query.
// Returns a *NotFoundError when no FilesExtend was found.
func (feq *FilesExtendQuery) First(ctx context.Context) (*FilesExtend, error) {
	nodes, err := feq.Limit(1).All(setContextOp(ctx, feq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{filesextend.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (feq *FilesExtendQuery) FirstX(ctx context.Context) *FilesExtend {
	node, err := feq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FilesExtend ID from the query.
// Returns a *NotFoundError when no FilesExtend ID was found.
func (feq *FilesExtendQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = feq.Limit(1).IDs(setContextOp(ctx, feq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{filesextend.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (feq *FilesExtendQuery) FirstIDX(ctx context.Context) int {
	id, err := feq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FilesExtend entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FilesExtend entity is found.
// Returns a *NotFoundError when no FilesExtend entities are found.
func (feq *FilesExtendQuery) Only(ctx context.Context) (*FilesExtend, error) {
	nodes, err := feq.Limit(2).All(setContextOp(ctx, feq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{filesextend.Label}
	default:
		return nil, &NotSingularError{filesextend.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (feq *FilesExtendQuery) OnlyX(ctx context.Context) *FilesExtend {
	node, err := feq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FilesExtend ID in the query.
// Returns a *NotSingularError when more than one FilesExtend ID is found.
// Returns a *NotFoundError when no entities are found.
func (feq *FilesExtendQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = feq.Limit(2).IDs(setContextOp(ctx, feq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{filesextend.Label}
	default:
		err = &NotSingularError{filesextend.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (feq *FilesExtendQuery) OnlyIDX(ctx context.Context) int {
	id, err := feq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FilesExtends.
func (feq *FilesExtendQuery) All(ctx context.Context) ([]*FilesExtend, error) {
	ctx = setContextOp(ctx, feq.ctx, ent.OpQueryAll)
	if err := feq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FilesExtend, *FilesExtendQuery]()
	return withInterceptors[[]*FilesExtend](ctx, feq, qr, feq.inters)
}

// AllX is like All, but panics if an error occurs.
func (feq *FilesExtendQuery) AllX(ctx context.Context) []*FilesExtend {
	nodes, err := feq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FilesExtend IDs.
func (feq *FilesExtendQuery) IDs(ctx context.Context) (ids []int, err error) {
	if feq.ctx.Unique == nil && feq.path != nil {
		feq.Unique(true)
	}
	ctx = setContextOp(ctx, feq.ctx, ent.OpQueryIDs)
	if err = feq.Select(filesextend.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (feq *FilesExtendQuery) IDsX(ctx context.Context) []int {
	ids, err := feq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (feq *FilesExtendQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, feq.ctx, ent.OpQueryCount)
	if err := feq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, feq, querierCount[*FilesExtendQuery](), feq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (feq *FilesExtendQuery) CountX(ctx context.Context) int {
	count, err := feq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (feq *FilesExtendQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, feq.ctx, ent.OpQueryExist)
	switch _, err := feq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (feq *FilesExtendQuery) ExistX(ctx context.Context) bool {
	exist, err := feq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FilesExtendQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (feq *FilesExtendQuery) Clone() *FilesExtendQuery {
	if feq == nil {
		return nil
	}
	return &FilesExtendQuery{
		config:     feq.config,
		ctx:        feq.ctx.Clone(),
		order:      append([]filesextend.OrderOption{}, feq.order...),
		inters:     append([]Interceptor{}, feq.inters...),
		predicates: append([]predicate.FilesExtend{}, feq.predicates...),
		withFiles:  feq.withFiles.Clone(),
		// clone intermediate query.
		sql:  feq.sql.Clone(),
		path: feq.path,
	}
}

// WithFiles tells the query-builder to eager-load the nodes that are connected to
// the "files" edge. The optional arguments are used to configure the query builder of the edge.
func (feq *FilesExtendQuery) WithFiles(opts ...func(*FilesQuery)) *FilesExtendQuery {
	query := (&FilesClient{config: feq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	feq.withFiles = query
	return feq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FilesExtend.Query().
//		GroupBy(filesextend.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (feq *FilesExtendQuery) GroupBy(field string, fields ...string) *FilesExtendGroupBy {
	feq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FilesExtendGroupBy{build: feq}
	grbuild.flds = &feq.ctx.Fields
	grbuild.label = filesextend.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt int64 `json:"created_at,omitempty"`
//	}
//
//	client.FilesExtend.Query().
//		Select(filesextend.FieldCreatedAt).
//		Scan(ctx, &v)
func (feq *FilesExtendQuery) Select(fields ...string) *FilesExtendSelect {
	feq.ctx.Fields = append(feq.ctx.Fields, fields...)
	sbuild := &FilesExtendSelect{FilesExtendQuery: feq}
	sbuild.label = filesextend.Label
	sbuild.flds, sbuild.scan = &feq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FilesExtendSelect configured with the given aggregations.
func (feq *FilesExtendQuery) Aggregate(fns ...AggregateFunc) *FilesExtendSelect {
	return feq.Select().Aggregate(fns...)
}

func (feq *FilesExtendQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range feq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, feq); err != nil {
				return err
			}
		}
	}
	for _, f := range feq.ctx.Fields {
		if !filesextend.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if feq.path != nil {
		prev, err := feq.path(ctx)
		if err != nil {
			return err
		}
		feq.sql = prev
	}
	return nil
}

func (feq *FilesExtendQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*FilesExtend, error) {
	var (
		nodes       = []*FilesExtend{}
		_spec       = feq.querySpec()
		loadedTypes = [1]bool{
			feq.withFiles != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*FilesExtend).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &FilesExtend{config: feq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, feq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := feq.withFiles; query != nil {
		if err := feq.loadFiles(ctx, query, nodes, nil,
			func(n *FilesExtend, e *Files) { n.Edges.Files = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (feq *FilesExtendQuery) loadFiles(ctx context.Context, query *FilesQuery, nodes []*FilesExtend, init func(*FilesExtend), assign func(*FilesExtend, *Files)) error {
	ids := make([]string, 0, len(nodes))
	nodeids := make(map[string][]*FilesExtend)
	for i := range nodes {
		fk := nodes[i].FileID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(files.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "file_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (feq *FilesExtendQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := feq.querySpec()
	_spec.Node.Columns = feq.ctx.Fields
	if len(feq.ctx.Fields) > 0 {
		_spec.Unique = feq.ctx.Unique != nil && *feq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, feq.driver, _spec)
}

func (feq *FilesExtendQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(filesextend.Table, filesextend.Columns, sqlgraph.NewFieldSpec(filesextend.FieldID, field.TypeInt))
	_spec.From = feq.sql
	if unique := feq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if feq.path != nil {
		_spec.Unique = true
	}
	if fields := feq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filesextend.FieldID)
		for i := range fields {
			if fields[i] != filesextend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if feq.withFiles != nil {
			_spec.Node.AddColumnOnce(filesextend.FieldFileID)
		}
	}
	if ps := feq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := feq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := feq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := feq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (feq *FilesExtendQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(feq.driver.Dialect())
	t1 := builder.Table(filesextend.Table)
	columns := feq.ctx.Fields
	if len(columns) == 0 {
		columns = filesextend.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if feq.sql != nil {
		selector = feq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if feq.ctx.Unique != nil && *feq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range feq.predicates {
		p(selector)
	}
	for _, p := range feq.order {
		p(selector)
	}
	if offset := feq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := feq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FilesExtendGroupBy is the group-by builder for FilesExtend entities.
type FilesExtendGroupBy struct {
	selector
	build *FilesExtendQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fegb *FilesExtendGroupBy) Aggregate(fns ...AggregateFunc) *FilesExtendGroupBy {
	fegb.fns = append(fegb.fns, fns...)
	return fegb
}

// Scan applies the selector query and scans the result into the given value.
func (fegb *FilesExtendGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fegb.build.ctx, ent.OpQueryGroupBy)
	if err := fegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FilesExtendQuery, *FilesExtendGroupBy](ctx, fegb.build, fegb, fegb.build.inters, v)
}

func (fegb *FilesExtendGroupBy) sqlScan(ctx context.Context, root *FilesExtendQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fegb.fns))
	for _, fn := range fegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fegb.flds)+len(fegb.fns))
		for _, f := range *fegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FilesExtendSelect is the builder for selecting fields of FilesExtend entities.
type FilesExtendSelect struct {
	*FilesExtendQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fes *FilesExtendSelect) Aggregate(fns ...AggregateFunc) *FilesExtendSelect {
	fes.fns = append(fes.fns, fns...)
	return fes
}

// Scan applies the selector query and scans the result into the given value.
func (fes *FilesExtendSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fes.ctx, ent.OpQuerySelect)
	if err := fes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FilesExtendQuery, *FilesExtendSelect](ctx, fes.FilesExtendQuery, fes, fes.inters, v)
}

func (fes *FilesExtendSelect) sqlScan(ctx context.Context, root *FilesExtendQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fes.fns))
	for _, fn := range fes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
