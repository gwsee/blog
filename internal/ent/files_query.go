// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/files"
	"blog/internal/ent/filesextend"
	"blog/internal/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FilesQuery is the builder for querying Files entities.
type FilesQuery struct {
	config
	ctx         *QueryContext
	order       []files.OrderOption
	inters      []Interceptor
	predicates  []predicate.Files
	withExtends *FilesExtendQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FilesQuery builder.
func (fq *FilesQuery) Where(ps ...predicate.Files) *FilesQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FilesQuery) Limit(limit int) *FilesQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FilesQuery) Offset(offset int) *FilesQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FilesQuery) Unique(unique bool) *FilesQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FilesQuery) Order(o ...files.OrderOption) *FilesQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// QueryExtends chains the current query on the "extends" edge.
func (fq *FilesQuery) QueryExtends() *FilesExtendQuery {
	query := (&FilesExtendClient{config: fq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := fq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := fq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(files.Table, files.FieldID, selector),
			sqlgraph.To(filesextend.Table, filesextend.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, files.ExtendsTable, files.ExtendsColumn),
		)
		fromU = sqlgraph.SetNeighbors(fq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Files entity from the query.
// Returns a *NotFoundError when no Files was found.
func (fq *FilesQuery) First(ctx context.Context) (*Files, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{files.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FilesQuery) FirstX(ctx context.Context) *Files {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Files ID from the query.
// Returns a *NotFoundError when no Files ID was found.
func (fq *FilesQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{files.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FilesQuery) FirstIDX(ctx context.Context) string {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Files entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Files entity is found.
// Returns a *NotFoundError when no Files entities are found.
func (fq *FilesQuery) Only(ctx context.Context) (*Files, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{files.Label}
	default:
		return nil, &NotSingularError{files.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FilesQuery) OnlyX(ctx context.Context) *Files {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Files ID in the query.
// Returns a *NotSingularError when more than one Files ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FilesQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{files.Label}
	default:
		err = &NotSingularError{files.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FilesQuery) OnlyIDX(ctx context.Context) string {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FilesSlice.
func (fq *FilesQuery) All(ctx context.Context) ([]*Files, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryAll)
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Files, *FilesQuery]()
	return withInterceptors[[]*Files](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FilesQuery) AllX(ctx context.Context) []*Files {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Files IDs.
func (fq *FilesQuery) IDs(ctx context.Context) (ids []string, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryIDs)
	if err = fq.Select(files.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FilesQuery) IDsX(ctx context.Context) []string {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FilesQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryCount)
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FilesQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FilesQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FilesQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryExist)
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FilesQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FilesQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FilesQuery) Clone() *FilesQuery {
	if fq == nil {
		return nil
	}
	return &FilesQuery{
		config:      fq.config,
		ctx:         fq.ctx.Clone(),
		order:       append([]files.OrderOption{}, fq.order...),
		inters:      append([]Interceptor{}, fq.inters...),
		predicates:  append([]predicate.Files{}, fq.predicates...),
		withExtends: fq.withExtends.Clone(),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// WithExtends tells the query-builder to eager-load the nodes that are connected to
// the "extends" edge. The optional arguments are used to configure the query builder of the edge.
func (fq *FilesQuery) WithExtends(opts ...func(*FilesExtendQuery)) *FilesQuery {
	query := (&FilesExtendClient{config: fq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	fq.withExtends = query
	return fq
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
//	client.Files.Query().
//		GroupBy(files.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FilesQuery) GroupBy(field string, fields ...string) *FilesGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FilesGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = files.Label
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
//	client.Files.Query().
//		Select(files.FieldCreatedAt).
//		Scan(ctx, &v)
func (fq *FilesQuery) Select(fields ...string) *FilesSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FilesSelect{FilesQuery: fq}
	sbuild.label = files.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FilesSelect configured with the given aggregations.
func (fq *FilesQuery) Aggregate(fns ...AggregateFunc) *FilesSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FilesQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !files.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FilesQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Files, error) {
	var (
		nodes       = []*Files{}
		_spec       = fq.querySpec()
		loadedTypes = [1]bool{
			fq.withExtends != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Files).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Files{config: fq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := fq.withExtends; query != nil {
		if err := fq.loadExtends(ctx, query, nodes,
			func(n *Files) { n.Edges.Extends = []*FilesExtend{} },
			func(n *Files, e *FilesExtend) { n.Edges.Extends = append(n.Edges.Extends, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FilesQuery) loadExtends(ctx context.Context, query *FilesExtendQuery, nodes []*Files, init func(*Files), assign func(*Files, *FilesExtend)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Files)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(filesextend.FieldFileID)
	}
	query.Where(predicate.FilesExtend(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(files.ExtendsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.FileID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "file_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (fq *FilesQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FilesQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(files.Table, files.Columns, sqlgraph.NewFieldSpec(files.FieldID, field.TypeString))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, files.FieldID)
		for i := range fields {
			if fields[i] != files.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FilesQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(files.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = files.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FilesGroupBy is the group-by builder for Files entities.
type FilesGroupBy struct {
	selector
	build *FilesQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FilesGroupBy) Aggregate(fns ...AggregateFunc) *FilesGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FilesGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, ent.OpQueryGroupBy)
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FilesQuery, *FilesGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FilesGroupBy) sqlScan(ctx context.Context, root *FilesQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FilesSelect is the builder for selecting fields of Files entities.
type FilesSelect struct {
	*FilesQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FilesSelect) Aggregate(fns ...AggregateFunc) *FilesSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FilesSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, ent.OpQuerySelect)
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FilesQuery, *FilesSelect](ctx, fs.FilesQuery, fs, fs.inters, v)
}

func (fs *FilesSelect) sqlScan(ctx context.Context, root *FilesQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
