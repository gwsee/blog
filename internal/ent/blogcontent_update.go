// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogcontent"
	"blog/internal/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogContentUpdate is the builder for updating BlogContent entities.
type BlogContentUpdate struct {
	config
	hooks    []Hook
	mutation *BlogContentMutation
}

// Where appends a list predicates to the BlogContentUpdate builder.
func (bcu *BlogContentUpdate) Where(ps ...predicate.BlogContent) *BlogContentUpdate {
	bcu.mutation.Where(ps...)
	return bcu
}

// SetContent sets the "content" field.
func (bcu *BlogContentUpdate) SetContent(s string) *BlogContentUpdate {
	bcu.mutation.SetContent(s)
	return bcu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (bcu *BlogContentUpdate) SetNillableContent(s *string) *BlogContentUpdate {
	if s != nil {
		bcu.SetContent(*s)
	}
	return bcu
}

// Mutation returns the BlogContentMutation object of the builder.
func (bcu *BlogContentUpdate) Mutation() *BlogContentMutation {
	return bcu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bcu *BlogContentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, bcu.sqlSave, bcu.mutation, bcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bcu *BlogContentUpdate) SaveX(ctx context.Context) int {
	affected, err := bcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bcu *BlogContentUpdate) Exec(ctx context.Context) error {
	_, err := bcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcu *BlogContentUpdate) ExecX(ctx context.Context) {
	if err := bcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcu *BlogContentUpdate) check() error {
	if v, ok := bcu.mutation.Content(); ok {
		if err := blogcontent.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "BlogContent.content": %w`, err)}
		}
	}
	return nil
}

func (bcu *BlogContentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bcu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(blogcontent.Table, blogcontent.Columns, sqlgraph.NewFieldSpec(blogcontent.FieldID, field.TypeInt))
	if ps := bcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcu.mutation.Content(); ok {
		_spec.SetField(blogcontent.FieldContent, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blogcontent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bcu.mutation.done = true
	return n, nil
}

// BlogContentUpdateOne is the builder for updating a single BlogContent entity.
type BlogContentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogContentMutation
}

// SetContent sets the "content" field.
func (bcuo *BlogContentUpdateOne) SetContent(s string) *BlogContentUpdateOne {
	bcuo.mutation.SetContent(s)
	return bcuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (bcuo *BlogContentUpdateOne) SetNillableContent(s *string) *BlogContentUpdateOne {
	if s != nil {
		bcuo.SetContent(*s)
	}
	return bcuo
}

// Mutation returns the BlogContentMutation object of the builder.
func (bcuo *BlogContentUpdateOne) Mutation() *BlogContentMutation {
	return bcuo.mutation
}

// Where appends a list predicates to the BlogContentUpdate builder.
func (bcuo *BlogContentUpdateOne) Where(ps ...predicate.BlogContent) *BlogContentUpdateOne {
	bcuo.mutation.Where(ps...)
	return bcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (bcuo *BlogContentUpdateOne) Select(field string, fields ...string) *BlogContentUpdateOne {
	bcuo.fields = append([]string{field}, fields...)
	return bcuo
}

// Save executes the query and returns the updated BlogContent entity.
func (bcuo *BlogContentUpdateOne) Save(ctx context.Context) (*BlogContent, error) {
	return withHooks(ctx, bcuo.sqlSave, bcuo.mutation, bcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bcuo *BlogContentUpdateOne) SaveX(ctx context.Context) *BlogContent {
	node, err := bcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (bcuo *BlogContentUpdateOne) Exec(ctx context.Context) error {
	_, err := bcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcuo *BlogContentUpdateOne) ExecX(ctx context.Context) {
	if err := bcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bcuo *BlogContentUpdateOne) check() error {
	if v, ok := bcuo.mutation.Content(); ok {
		if err := blogcontent.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "BlogContent.content": %w`, err)}
		}
	}
	return nil
}

func (bcuo *BlogContentUpdateOne) sqlSave(ctx context.Context) (_node *BlogContent, err error) {
	if err := bcuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(blogcontent.Table, blogcontent.Columns, sqlgraph.NewFieldSpec(blogcontent.FieldID, field.TypeInt))
	id, ok := bcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BlogContent.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := bcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blogcontent.FieldID)
		for _, f := range fields {
			if !blogcontent.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blogcontent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := bcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bcuo.mutation.Content(); ok {
		_spec.SetField(blogcontent.FieldContent, field.TypeString, value)
	}
	_node = &BlogContent{config: bcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, bcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blogcontent.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	bcuo.mutation.done = true
	return _node, nil
}
