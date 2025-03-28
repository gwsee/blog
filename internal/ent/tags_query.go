// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogs"
	"blog/internal/ent/predicate"
	"blog/internal/ent/tags"
	"blog/internal/ent/tagsrelation"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TagsQuery is the builder for querying Tags entities.
type TagsQuery struct {
	config
	ctx             *QueryContext
	order           []tags.OrderOption
	inters          []Interceptor
	predicates      []predicate.Tags
	withBlogs       *BlogsQuery
	withTagRelation *TagsRelationQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TagsQuery builder.
func (tq *TagsQuery) Where(ps ...predicate.Tags) *TagsQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TagsQuery) Limit(limit int) *TagsQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TagsQuery) Offset(offset int) *TagsQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TagsQuery) Unique(unique bool) *TagsQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TagsQuery) Order(o ...tags.OrderOption) *TagsQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// QueryBlogs chains the current query on the "blogs" edge.
func (tq *TagsQuery) QueryBlogs() *BlogsQuery {
	query := (&BlogsClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tags.Table, tags.FieldID, selector),
			sqlgraph.To(blogs.Table, blogs.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tags.BlogsTable, tags.BlogsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTagRelation chains the current query on the "tag_relation" edge.
func (tq *TagsQuery) QueryTagRelation() *TagsRelationQuery {
	query := (&TagsRelationClient{config: tq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(tags.Table, tags.FieldID, selector),
			sqlgraph.To(tagsrelation.Table, tagsrelation.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, tags.TagRelationTable, tags.TagRelationColumn),
		)
		fromU = sqlgraph.SetNeighbors(tq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Tags entity from the query.
// Returns a *NotFoundError when no Tags was found.
func (tq *TagsQuery) First(ctx context.Context) (*Tags, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{tags.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TagsQuery) FirstX(ctx context.Context) *Tags {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Tags ID from the query.
// Returns a *NotFoundError when no Tags ID was found.
func (tq *TagsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{tags.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TagsQuery) FirstIDX(ctx context.Context) int {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Tags entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Tags entity is found.
// Returns a *NotFoundError when no Tags entities are found.
func (tq *TagsQuery) Only(ctx context.Context) (*Tags, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{tags.Label}
	default:
		return nil, &NotSingularError{tags.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TagsQuery) OnlyX(ctx context.Context) *Tags {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Tags ID in the query.
// Returns a *NotSingularError when more than one Tags ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TagsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{tags.Label}
	default:
		err = &NotSingularError{tags.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TagsQuery) OnlyIDX(ctx context.Context) int {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TagsSlice.
func (tq *TagsQuery) All(ctx context.Context) ([]*Tags, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryAll)
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Tags, *TagsQuery]()
	return withInterceptors[[]*Tags](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TagsQuery) AllX(ctx context.Context) []*Tags {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Tags IDs.
func (tq *TagsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryIDs)
	if err = tq.Select(tags.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TagsQuery) IDsX(ctx context.Context) []int {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TagsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryCount)
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TagsQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TagsQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TagsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, ent.OpQueryExist)
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TagsQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TagsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TagsQuery) Clone() *TagsQuery {
	if tq == nil {
		return nil
	}
	return &TagsQuery{
		config:          tq.config,
		ctx:             tq.ctx.Clone(),
		order:           append([]tags.OrderOption{}, tq.order...),
		inters:          append([]Interceptor{}, tq.inters...),
		predicates:      append([]predicate.Tags{}, tq.predicates...),
		withBlogs:       tq.withBlogs.Clone(),
		withTagRelation: tq.withTagRelation.Clone(),
		// clone intermediate query.
		sql:  tq.sql.Clone(),
		path: tq.path,
	}
}

// WithBlogs tells the query-builder to eager-load the nodes that are connected to
// the "blogs" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TagsQuery) WithBlogs(opts ...func(*BlogsQuery)) *TagsQuery {
	query := (&BlogsClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withBlogs = query
	return tq
}

// WithTagRelation tells the query-builder to eager-load the nodes that are connected to
// the "tag_relation" edge. The optional arguments are used to configure the query builder of the edge.
func (tq *TagsQuery) WithTagRelation(opts ...func(*TagsRelationQuery)) *TagsQuery {
	query := (&TagsRelationClient{config: tq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tq.withTagRelation = query
	return tq
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
//	client.Tags.Query().
//		GroupBy(tags.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TagsQuery) GroupBy(field string, fields ...string) *TagsGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TagsGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = tags.Label
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
//	client.Tags.Query().
//		Select(tags.FieldCreatedAt).
//		Scan(ctx, &v)
func (tq *TagsQuery) Select(fields ...string) *TagsSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TagsSelect{TagsQuery: tq}
	sbuild.label = tags.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TagsSelect configured with the given aggregations.
func (tq *TagsQuery) Aggregate(fns ...AggregateFunc) *TagsSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TagsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	for _, f := range tq.ctx.Fields {
		if !tags.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.sql = prev
	}
	return nil
}

func (tq *TagsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Tags, error) {
	var (
		nodes       = []*Tags{}
		_spec       = tq.querySpec()
		loadedTypes = [2]bool{
			tq.withBlogs != nil,
			tq.withTagRelation != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Tags).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Tags{config: tq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tq.withBlogs; query != nil {
		if err := tq.loadBlogs(ctx, query, nodes,
			func(n *Tags) { n.Edges.Blogs = []*Blogs{} },
			func(n *Tags, e *Blogs) { n.Edges.Blogs = append(n.Edges.Blogs, e) }); err != nil {
			return nil, err
		}
	}
	if query := tq.withTagRelation; query != nil {
		if err := tq.loadTagRelation(ctx, query, nodes,
			func(n *Tags) { n.Edges.TagRelation = []*TagsRelation{} },
			func(n *Tags, e *TagsRelation) { n.Edges.TagRelation = append(n.Edges.TagRelation, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tq *TagsQuery) loadBlogs(ctx context.Context, query *BlogsQuery, nodes []*Tags, init func(*Tags), assign func(*Tags, *Blogs)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Tags)
	nids := make(map[int]map[*Tags]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(tags.BlogsTable)
		s.Join(joinT).On(s.C(blogs.FieldID), joinT.C(tags.BlogsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(tags.BlogsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(tags.BlogsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Tags]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Blogs](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "blogs" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (tq *TagsQuery) loadTagRelation(ctx context.Context, query *TagsRelationQuery, nodes []*Tags, init func(*Tags), assign func(*Tags, *TagsRelation)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Tags)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(tagsrelation.FieldTagID)
	}
	query.Where(predicate.TagsRelation(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(tags.TagRelationColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.TagID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "tag_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (tq *TagsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tq.querySpec()
	_spec.Node.Columns = tq.ctx.Fields
	if len(tq.ctx.Fields) > 0 {
		_spec.Unique = tq.ctx.Unique != nil && *tq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tq.driver, _spec)
}

func (tq *TagsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(tags.Table, tags.Columns, sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt))
	_spec.From = tq.sql
	if unique := tq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tq.path != nil {
		_spec.Unique = true
	}
	if fields := tq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tags.FieldID)
		for i := range fields {
			if fields[i] != tags.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tq *TagsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tq.driver.Dialect())
	t1 := builder.Table(tags.Table)
	columns := tq.ctx.Fields
	if len(columns) == 0 {
		columns = tags.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tq.sql != nil {
		selector = tq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tq.ctx.Unique != nil && *tq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tq.predicates {
		p(selector)
	}
	for _, p := range tq.order {
		p(selector)
	}
	if offset := tq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TagsGroupBy is the group-by builder for Tags entities.
type TagsGroupBy struct {
	selector
	build *TagsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TagsGroupBy) Aggregate(fns ...AggregateFunc) *TagsGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TagsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TagsQuery, *TagsGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TagsGroupBy) sqlScan(ctx context.Context, root *TagsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tgb.fns))
	for _, fn := range tgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tgb.flds)+len(tgb.fns))
		for _, f := range *tgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TagsSelect is the builder for selecting fields of Tags entities.
type TagsSelect struct {
	*TagsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TagsSelect) Aggregate(fns ...AggregateFunc) *TagsSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TagsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TagsQuery, *TagsSelect](ctx, ts.TagsQuery, ts, ts.inters, v)
}

func (ts *TagsSelect) sqlScan(ctx context.Context, root *TagsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ts.fns))
	for _, fn := range ts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
