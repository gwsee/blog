// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/accountproject"
	"blog/internal/ent/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountProjectQuery is the builder for querying AccountProject entities.
type AccountProjectQuery struct {
	config
	ctx        *QueryContext
	order      []accountproject.OrderOption
	inters     []Interceptor
	predicates []predicate.AccountProject
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AccountProjectQuery builder.
func (apq *AccountProjectQuery) Where(ps ...predicate.AccountProject) *AccountProjectQuery {
	apq.predicates = append(apq.predicates, ps...)
	return apq
}

// Limit the number of records to be returned by this query.
func (apq *AccountProjectQuery) Limit(limit int) *AccountProjectQuery {
	apq.ctx.Limit = &limit
	return apq
}

// Offset to start from.
func (apq *AccountProjectQuery) Offset(offset int) *AccountProjectQuery {
	apq.ctx.Offset = &offset
	return apq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (apq *AccountProjectQuery) Unique(unique bool) *AccountProjectQuery {
	apq.ctx.Unique = &unique
	return apq
}

// Order specifies how the records should be ordered.
func (apq *AccountProjectQuery) Order(o ...accountproject.OrderOption) *AccountProjectQuery {
	apq.order = append(apq.order, o...)
	return apq
}

// First returns the first AccountProject entity from the query.
// Returns a *NotFoundError when no AccountProject was found.
func (apq *AccountProjectQuery) First(ctx context.Context) (*AccountProject, error) {
	nodes, err := apq.Limit(1).All(setContextOp(ctx, apq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{accountproject.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (apq *AccountProjectQuery) FirstX(ctx context.Context) *AccountProject {
	node, err := apq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AccountProject ID from the query.
// Returns a *NotFoundError when no AccountProject ID was found.
func (apq *AccountProjectQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = apq.Limit(1).IDs(setContextOp(ctx, apq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{accountproject.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (apq *AccountProjectQuery) FirstIDX(ctx context.Context) int {
	id, err := apq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AccountProject entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AccountProject entity is found.
// Returns a *NotFoundError when no AccountProject entities are found.
func (apq *AccountProjectQuery) Only(ctx context.Context) (*AccountProject, error) {
	nodes, err := apq.Limit(2).All(setContextOp(ctx, apq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{accountproject.Label}
	default:
		return nil, &NotSingularError{accountproject.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (apq *AccountProjectQuery) OnlyX(ctx context.Context) *AccountProject {
	node, err := apq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AccountProject ID in the query.
// Returns a *NotSingularError when more than one AccountProject ID is found.
// Returns a *NotFoundError when no entities are found.
func (apq *AccountProjectQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = apq.Limit(2).IDs(setContextOp(ctx, apq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{accountproject.Label}
	default:
		err = &NotSingularError{accountproject.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (apq *AccountProjectQuery) OnlyIDX(ctx context.Context) int {
	id, err := apq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AccountProjects.
func (apq *AccountProjectQuery) All(ctx context.Context) ([]*AccountProject, error) {
	ctx = setContextOp(ctx, apq.ctx, ent.OpQueryAll)
	if err := apq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*AccountProject, *AccountProjectQuery]()
	return withInterceptors[[]*AccountProject](ctx, apq, qr, apq.inters)
}

// AllX is like All, but panics if an error occurs.
func (apq *AccountProjectQuery) AllX(ctx context.Context) []*AccountProject {
	nodes, err := apq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AccountProject IDs.
func (apq *AccountProjectQuery) IDs(ctx context.Context) (ids []int, err error) {
	if apq.ctx.Unique == nil && apq.path != nil {
		apq.Unique(true)
	}
	ctx = setContextOp(ctx, apq.ctx, ent.OpQueryIDs)
	if err = apq.Select(accountproject.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (apq *AccountProjectQuery) IDsX(ctx context.Context) []int {
	ids, err := apq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (apq *AccountProjectQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, apq.ctx, ent.OpQueryCount)
	if err := apq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, apq, querierCount[*AccountProjectQuery](), apq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (apq *AccountProjectQuery) CountX(ctx context.Context) int {
	count, err := apq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (apq *AccountProjectQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, apq.ctx, ent.OpQueryExist)
	switch _, err := apq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (apq *AccountProjectQuery) ExistX(ctx context.Context) bool {
	exist, err := apq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AccountProjectQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (apq *AccountProjectQuery) Clone() *AccountProjectQuery {
	if apq == nil {
		return nil
	}
	return &AccountProjectQuery{
		config:     apq.config,
		ctx:        apq.ctx.Clone(),
		order:      append([]accountproject.OrderOption{}, apq.order...),
		inters:     append([]Interceptor{}, apq.inters...),
		predicates: append([]predicate.AccountProject{}, apq.predicates...),
		// clone intermediate query.
		sql:  apq.sql.Clone(),
		path: apq.path,
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
//	client.AccountProject.Query().
//		GroupBy(accountproject.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (apq *AccountProjectQuery) GroupBy(field string, fields ...string) *AccountProjectGroupBy {
	apq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &AccountProjectGroupBy{build: apq}
	grbuild.flds = &apq.ctx.Fields
	grbuild.label = accountproject.Label
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
//	client.AccountProject.Query().
//		Select(accountproject.FieldCreatedAt).
//		Scan(ctx, &v)
func (apq *AccountProjectQuery) Select(fields ...string) *AccountProjectSelect {
	apq.ctx.Fields = append(apq.ctx.Fields, fields...)
	sbuild := &AccountProjectSelect{AccountProjectQuery: apq}
	sbuild.label = accountproject.Label
	sbuild.flds, sbuild.scan = &apq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a AccountProjectSelect configured with the given aggregations.
func (apq *AccountProjectQuery) Aggregate(fns ...AggregateFunc) *AccountProjectSelect {
	return apq.Select().Aggregate(fns...)
}

func (apq *AccountProjectQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range apq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, apq); err != nil {
				return err
			}
		}
	}
	for _, f := range apq.ctx.Fields {
		if !accountproject.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if apq.path != nil {
		prev, err := apq.path(ctx)
		if err != nil {
			return err
		}
		apq.sql = prev
	}
	return nil
}

func (apq *AccountProjectQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*AccountProject, error) {
	var (
		nodes = []*AccountProject{}
		_spec = apq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*AccountProject).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &AccountProject{config: apq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, apq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (apq *AccountProjectQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := apq.querySpec()
	_spec.Node.Columns = apq.ctx.Fields
	if len(apq.ctx.Fields) > 0 {
		_spec.Unique = apq.ctx.Unique != nil && *apq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, apq.driver, _spec)
}

func (apq *AccountProjectQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(accountproject.Table, accountproject.Columns, sqlgraph.NewFieldSpec(accountproject.FieldID, field.TypeInt))
	_spec.From = apq.sql
	if unique := apq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if apq.path != nil {
		_spec.Unique = true
	}
	if fields := apq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, accountproject.FieldID)
		for i := range fields {
			if fields[i] != accountproject.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := apq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := apq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := apq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := apq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (apq *AccountProjectQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(apq.driver.Dialect())
	t1 := builder.Table(accountproject.Table)
	columns := apq.ctx.Fields
	if len(columns) == 0 {
		columns = accountproject.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if apq.sql != nil {
		selector = apq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if apq.ctx.Unique != nil && *apq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range apq.predicates {
		p(selector)
	}
	for _, p := range apq.order {
		p(selector)
	}
	if offset := apq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := apq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AccountProjectGroupBy is the group-by builder for AccountProject entities.
type AccountProjectGroupBy struct {
	selector
	build *AccountProjectQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (apgb *AccountProjectGroupBy) Aggregate(fns ...AggregateFunc) *AccountProjectGroupBy {
	apgb.fns = append(apgb.fns, fns...)
	return apgb
}

// Scan applies the selector query and scans the result into the given value.
func (apgb *AccountProjectGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, apgb.build.ctx, ent.OpQueryGroupBy)
	if err := apgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountProjectQuery, *AccountProjectGroupBy](ctx, apgb.build, apgb, apgb.build.inters, v)
}

func (apgb *AccountProjectGroupBy) sqlScan(ctx context.Context, root *AccountProjectQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(apgb.fns))
	for _, fn := range apgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*apgb.flds)+len(apgb.fns))
		for _, f := range *apgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*apgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := apgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// AccountProjectSelect is the builder for selecting fields of AccountProject entities.
type AccountProjectSelect struct {
	*AccountProjectQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (aps *AccountProjectSelect) Aggregate(fns ...AggregateFunc) *AccountProjectSelect {
	aps.fns = append(aps.fns, fns...)
	return aps
}

// Scan applies the selector query and scans the result into the given value.
func (aps *AccountProjectSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, aps.ctx, ent.OpQuerySelect)
	if err := aps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*AccountProjectQuery, *AccountProjectSelect](ctx, aps.AccountProjectQuery, aps, aps.inters, v)
}

func (aps *AccountProjectSelect) sqlScan(ctx context.Context, root *AccountProjectQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(aps.fns))
	for _, fn := range aps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*aps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := aps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}