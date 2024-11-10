// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"blog/internal/ent"
	"blog/internal/ent/account"
	"blog/internal/ent/blogs"
	"blog/internal/ent/blogscomment"
	"blog/internal/ent/blogscontent"
	"blog/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The AccountFunc type is an adapter to allow the use of ordinary function as a Querier.
type AccountFunc func(context.Context, *ent.AccountQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AccountFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AccountQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AccountQuery", q)
}

// The TraverseAccount type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAccount func(context.Context, *ent.AccountQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAccount) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAccount) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AccountQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AccountQuery", q)
}

// The BlogsFunc type is an adapter to allow the use of ordinary function as a Querier.
type BlogsFunc func(context.Context, *ent.BlogsQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f BlogsFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.BlogsQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.BlogsQuery", q)
}

// The TraverseBlogs type is an adapter to allow the use of ordinary function as Traverser.
type TraverseBlogs func(context.Context, *ent.BlogsQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseBlogs) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseBlogs) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BlogsQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.BlogsQuery", q)
}

// The BlogsCommentFunc type is an adapter to allow the use of ordinary function as a Querier.
type BlogsCommentFunc func(context.Context, *ent.BlogsCommentQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f BlogsCommentFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.BlogsCommentQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.BlogsCommentQuery", q)
}

// The TraverseBlogsComment type is an adapter to allow the use of ordinary function as Traverser.
type TraverseBlogsComment func(context.Context, *ent.BlogsCommentQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseBlogsComment) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseBlogsComment) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BlogsCommentQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.BlogsCommentQuery", q)
}

// The BlogsContentFunc type is an adapter to allow the use of ordinary function as a Querier.
type BlogsContentFunc func(context.Context, *ent.BlogsContentQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f BlogsContentFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.BlogsContentQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.BlogsContentQuery", q)
}

// The TraverseBlogsContent type is an adapter to allow the use of ordinary function as Traverser.
type TraverseBlogsContent func(context.Context, *ent.BlogsContentQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseBlogsContent) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseBlogsContent) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.BlogsContentQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.BlogsContentQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.AccountQuery:
		return &query[*ent.AccountQuery, predicate.Account, account.OrderOption]{typ: ent.TypeAccount, tq: q}, nil
	case *ent.BlogsQuery:
		return &query[*ent.BlogsQuery, predicate.Blogs, blogs.OrderOption]{typ: ent.TypeBlogs, tq: q}, nil
	case *ent.BlogsCommentQuery:
		return &query[*ent.BlogsCommentQuery, predicate.BlogsComment, blogscomment.OrderOption]{typ: ent.TypeBlogsComment, tq: q}, nil
	case *ent.BlogsContentQuery:
		return &query[*ent.BlogsContentQuery, predicate.BlogsContent, blogscontent.OrderOption]{typ: ent.TypeBlogsContent, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}