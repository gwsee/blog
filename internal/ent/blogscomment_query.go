// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogscomment"
	"blog/internal/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogsCommentQuery is the builder for querying BlogsComment entities.
type BlogsCommentQuery struct {
	config
	ctx        *QueryContext
	order      []blogscomment.OrderOption
	inters     []Interceptor
	predicates []predicate.BlogsComment
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlogsCommentQuery builder.
func (bcq *BlogsCommentQuery) Where(ps ...predicate.BlogsComment) *BlogsCommentQuery {
	bcq.predicates = append(bcq.predicates, ps...)
	return bcq
}

// Limit the number of records to be returned by this query.
func (bcq *BlogsCommentQuery) Limit(limit int) *BlogsCommentQuery {
	bcq.ctx.Limit = &limit
	return bcq
}

// Offset to start from.
func (bcq *BlogsCommentQuery) Offset(offset int) *BlogsCommentQuery {
	bcq.ctx.Offset = &offset
	return bcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (bcq *BlogsCommentQuery) Unique(unique bool) *BlogsCommentQuery {
	bcq.ctx.Unique = &unique
	return bcq
}

// Order specifies how the records should be ordered.
func (bcq *BlogsCommentQuery) Order(o ...blogscomment.OrderOption) *BlogsCommentQuery {
	bcq.order = append(bcq.order, o...)
	return bcq
}

// First returns the first BlogsComment entity from the query.
// Returns a *NotFoundError when no BlogsComment was found.
func (bcq *BlogsCommentQuery) First(ctx context.Context) (*BlogsComment, error) {
	nodes, err := bcq.Limit(1).All(setContextOp(ctx, bcq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blogscomment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (bcq *BlogsCommentQuery) FirstX(ctx context.Context) *BlogsComment {
	node, err := bcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BlogsComment ID from the query.
// Returns a *NotFoundError when no BlogsComment ID was found.
func (bcq *BlogsCommentQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcq.Limit(1).IDs(setContextOp(ctx, bcq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blogscomment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (bcq *BlogsCommentQuery) FirstIDX(ctx context.Context) int {
	id, err := bcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BlogsComment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BlogsComment entity is found.
// Returns a *NotFoundError when no BlogsComment entities are found.
func (bcq *BlogsCommentQuery) Only(ctx context.Context) (*BlogsComment, error) {
	nodes, err := bcq.Limit(2).All(setContextOp(ctx, bcq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blogscomment.Label}
	default:
		return nil, &NotSingularError{blogscomment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (bcq *BlogsCommentQuery) OnlyX(ctx context.Context) *BlogsComment {
	node, err := bcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BlogsComment ID in the query.
// Returns a *NotSingularError when more than one BlogsComment ID is found.
// Returns a *NotFoundError when no entities are found.
func (bcq *BlogsCommentQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = bcq.Limit(2).IDs(setContextOp(ctx, bcq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blogscomment.Label}
	default:
		err = &NotSingularError{blogscomment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (bcq *BlogsCommentQuery) OnlyIDX(ctx context.Context) int {
	id, err := bcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlogsComments.
func (bcq *BlogsCommentQuery) All(ctx context.Context) ([]*BlogsComment, error) {
	ctx = setContextOp(ctx, bcq.ctx, ent.OpQueryAll)
	if err := bcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*BlogsComment, *BlogsCommentQuery]()
	return withInterceptors[[]*BlogsComment](ctx, bcq, qr, bcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (bcq *BlogsCommentQuery) AllX(ctx context.Context) []*BlogsComment {
	nodes, err := bcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BlogsComment IDs.
func (bcq *BlogsCommentQuery) IDs(ctx context.Context) (ids []int, err error) {
	if bcq.ctx.Unique == nil && bcq.path != nil {
		bcq.Unique(true)
	}
	ctx = setContextOp(ctx, bcq.ctx, ent.OpQueryIDs)
	if err = bcq.Select(blogscomment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (bcq *BlogsCommentQuery) IDsX(ctx context.Context) []int {
	ids, err := bcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (bcq *BlogsCommentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, bcq.ctx, ent.OpQueryCount)
	if err := bcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, bcq, querierCount[*BlogsCommentQuery](), bcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (bcq *BlogsCommentQuery) CountX(ctx context.Context) int {
	count, err := bcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (bcq *BlogsCommentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, bcq.ctx, ent.OpQueryExist)
	switch _, err := bcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (bcq *BlogsCommentQuery) ExistX(ctx context.Context) bool {
	exist, err := bcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlogsCommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (bcq *BlogsCommentQuery) Clone() *BlogsCommentQuery {
	if bcq == nil {
		return nil
	}
	return &BlogsCommentQuery{
		config:     bcq.config,
		ctx:        bcq.ctx.Clone(),
		order:      append([]blogscomment.OrderOption{}, bcq.order...),
		inters:     append([]Interceptor{}, bcq.inters...),
		predicates: append([]predicate.BlogsComment{}, bcq.predicates...),
		// clone intermediate query.
		sql:  bcq.sql.Clone(),
		path: bcq.path,
	}
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
//	client.BlogsComment.Query().
//		GroupBy(blogscomment.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (bcq *BlogsCommentQuery) GroupBy(field string, fields ...string) *BlogsCommentGroupBy {
	bcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &BlogsCommentGroupBy{build: bcq}
	grbuild.flds = &bcq.ctx.Fields
	grbuild.label = blogscomment.Label
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
//	client.BlogsComment.Query().
//		Select(blogscomment.FieldCreatedAt).
//		Scan(ctx, &v)
func (bcq *BlogsCommentQuery) Select(fields ...string) *BlogsCommentSelect {
	bcq.ctx.Fields = append(bcq.ctx.Fields, fields...)
	sbuild := &BlogsCommentSelect{BlogsCommentQuery: bcq}
	sbuild.label = blogscomment.Label
	sbuild.flds, sbuild.scan = &bcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a BlogsCommentSelect configured with the given aggregations.
func (bcq *BlogsCommentQuery) Aggregate(fns ...AggregateFunc) *BlogsCommentSelect {
	return bcq.Select().Aggregate(fns...)
}

func (bcq *BlogsCommentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range bcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, bcq); err != nil {
				return err
			}
		}
	}
	for _, f := range bcq.ctx.Fields {
		if !blogscomment.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if bcq.path != nil {
		prev, err := bcq.path(ctx)
		if err != nil {
			return err
		}
		bcq.sql = prev
	}
	return nil
}

func (bcq *BlogsCommentQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BlogsComment, error) {
	var (
		nodes = []*BlogsComment{}
		_spec = bcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*BlogsComment).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &BlogsComment{config: bcq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, bcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (bcq *BlogsCommentQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := bcq.querySpec()
	_spec.Node.Columns = bcq.ctx.Fields
	if len(bcq.ctx.Fields) > 0 {
		_spec.Unique = bcq.ctx.Unique != nil && *bcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, bcq.driver, _spec)
}

func (bcq *BlogsCommentQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(blogscomment.Table, blogscomment.Columns, sqlgraph.NewFieldSpec(blogscomment.FieldID, field.TypeInt))
	_spec.From = bcq.sql
	if unique := bcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if bcq.path != nil {
		_spec.Unique = true
	}
	if fields := bcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blogscomment.FieldID)
		for i := range fields {
			if fields[i] != blogscomment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := bcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := bcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := bcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := bcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (bcq *BlogsCommentQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(bcq.driver.Dialect())
	t1 := builder.Table(blogscomment.Table)
	columns := bcq.ctx.Fields
	if len(columns) == 0 {
		columns = blogscomment.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if bcq.sql != nil {
		selector = bcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if bcq.ctx.Unique != nil && *bcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range bcq.predicates {
		p(selector)
	}
	for _, p := range bcq.order {
		p(selector)
	}
	if offset := bcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := bcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlogsCommentGroupBy is the group-by builder for BlogsComment entities.
type BlogsCommentGroupBy struct {
	selector
	build *BlogsCommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bcgb *BlogsCommentGroupBy) Aggregate(fns ...AggregateFunc) *BlogsCommentGroupBy {
	bcgb.fns = append(bcgb.fns, fns...)
	return bcgb
}

// Scan applies the selector query and scans the result into the given value.
func (bcgb *BlogsCommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bcgb.build.ctx, ent.OpQueryGroupBy)
	if err := bcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlogsCommentQuery, *BlogsCommentGroupBy](ctx, bcgb.build, bcgb, bcgb.build.inters, v)
}

func (bcgb *BlogsCommentGroupBy) sqlScan(ctx context.Context, root *BlogsCommentQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(bcgb.fns))
	for _, fn := range bcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*bcgb.flds)+len(bcgb.fns))
		for _, f := range *bcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*bcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// BlogsCommentSelect is the builder for selecting fields of BlogsComment entities.
type BlogsCommentSelect struct {
	*BlogsCommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (bcs *BlogsCommentSelect) Aggregate(fns ...AggregateFunc) *BlogsCommentSelect {
	bcs.fns = append(bcs.fns, fns...)
	return bcs
}

// Scan applies the selector query and scans the result into the given value.
func (bcs *BlogsCommentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, bcs.ctx, ent.OpQuerySelect)
	if err := bcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*BlogsCommentQuery, *BlogsCommentSelect](ctx, bcs.BlogsCommentQuery, bcs, bcs.inters, v)
}

func (bcs *BlogsCommentSelect) sqlScan(ctx context.Context, root *BlogsCommentQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(bcs.fns))
	for _, fn := range bcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*bcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
