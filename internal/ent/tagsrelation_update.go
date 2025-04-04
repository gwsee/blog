// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogs"
	"blog/internal/ent/predicate"
	"blog/internal/ent/tags"
	"blog/internal/ent/tagsrelation"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TagsRelationUpdate is the builder for updating TagsRelation entities.
type TagsRelationUpdate struct {
	config
	hooks    []Hook
	mutation *TagsRelationMutation
}

// Where appends a list predicates to the TagsRelationUpdate builder.
func (tru *TagsRelationUpdate) Where(ps ...predicate.TagsRelation) *TagsRelationUpdate {
	tru.mutation.Where(ps...)
	return tru
}

// SetTagID sets the "tag_id" field.
func (tru *TagsRelationUpdate) SetTagID(i int) *TagsRelationUpdate {
	tru.mutation.SetTagID(i)
	return tru
}

// SetNillableTagID sets the "tag_id" field if the given value is not nil.
func (tru *TagsRelationUpdate) SetNillableTagID(i *int) *TagsRelationUpdate {
	if i != nil {
		tru.SetTagID(*i)
	}
	return tru
}

// SetRelation sets the "relation" field.
func (tru *TagsRelationUpdate) SetRelation(s string) *TagsRelationUpdate {
	tru.mutation.SetRelation(s)
	return tru
}

// SetNillableRelation sets the "relation" field if the given value is not nil.
func (tru *TagsRelationUpdate) SetNillableRelation(s *string) *TagsRelationUpdate {
	if s != nil {
		tru.SetRelation(*s)
	}
	return tru
}

// SetRelationID sets the "relation_id" field.
func (tru *TagsRelationUpdate) SetRelationID(i int) *TagsRelationUpdate {
	tru.mutation.SetRelationID(i)
	return tru
}

// SetNillableRelationID sets the "relation_id" field if the given value is not nil.
func (tru *TagsRelationUpdate) SetNillableRelationID(i *int) *TagsRelationUpdate {
	if i != nil {
		tru.SetRelationID(*i)
	}
	return tru
}

// SetBlogID sets the "blog" edge to the Blogs entity by ID.
func (tru *TagsRelationUpdate) SetBlogID(id int) *TagsRelationUpdate {
	tru.mutation.SetBlogID(id)
	return tru
}

// SetBlog sets the "blog" edge to the Blogs entity.
func (tru *TagsRelationUpdate) SetBlog(b *Blogs) *TagsRelationUpdate {
	return tru.SetBlogID(b.ID)
}

// SetTag sets the "tag" edge to the Tags entity.
func (tru *TagsRelationUpdate) SetTag(t *Tags) *TagsRelationUpdate {
	return tru.SetTagID(t.ID)
}

// Mutation returns the TagsRelationMutation object of the builder.
func (tru *TagsRelationUpdate) Mutation() *TagsRelationMutation {
	return tru.mutation
}

// ClearBlog clears the "blog" edge to the Blogs entity.
func (tru *TagsRelationUpdate) ClearBlog() *TagsRelationUpdate {
	tru.mutation.ClearBlog()
	return tru
}

// ClearTag clears the "tag" edge to the Tags entity.
func (tru *TagsRelationUpdate) ClearTag() *TagsRelationUpdate {
	tru.mutation.ClearTag()
	return tru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tru *TagsRelationUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tru.sqlSave, tru.mutation, tru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tru *TagsRelationUpdate) SaveX(ctx context.Context) int {
	affected, err := tru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tru *TagsRelationUpdate) Exec(ctx context.Context) error {
	_, err := tru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tru *TagsRelationUpdate) ExecX(ctx context.Context) {
	if err := tru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tru *TagsRelationUpdate) check() error {
	if tru.mutation.BlogCleared() && len(tru.mutation.BlogIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "TagsRelation.blog"`)
	}
	if tru.mutation.TagCleared() && len(tru.mutation.TagIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "TagsRelation.tag"`)
	}
	return nil
}

func (tru *TagsRelationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tagsrelation.Table, tagsrelation.Columns, sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt))
	if ps := tru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tru.mutation.Relation(); ok {
		_spec.SetField(tagsrelation.FieldRelation, field.TypeString, value)
	}
	if tru.mutation.BlogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.BlogTable,
			Columns: []string{tagsrelation.BlogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.BlogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.BlogTable,
			Columns: []string{tagsrelation.BlogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tru.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.TagTable,
			Columns: []string{tagsrelation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tru.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.TagTable,
			Columns: []string{tagsrelation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tagsrelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tru.mutation.done = true
	return n, nil
}

// TagsRelationUpdateOne is the builder for updating a single TagsRelation entity.
type TagsRelationUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagsRelationMutation
}

// SetTagID sets the "tag_id" field.
func (truo *TagsRelationUpdateOne) SetTagID(i int) *TagsRelationUpdateOne {
	truo.mutation.SetTagID(i)
	return truo
}

// SetNillableTagID sets the "tag_id" field if the given value is not nil.
func (truo *TagsRelationUpdateOne) SetNillableTagID(i *int) *TagsRelationUpdateOne {
	if i != nil {
		truo.SetTagID(*i)
	}
	return truo
}

// SetRelation sets the "relation" field.
func (truo *TagsRelationUpdateOne) SetRelation(s string) *TagsRelationUpdateOne {
	truo.mutation.SetRelation(s)
	return truo
}

// SetNillableRelation sets the "relation" field if the given value is not nil.
func (truo *TagsRelationUpdateOne) SetNillableRelation(s *string) *TagsRelationUpdateOne {
	if s != nil {
		truo.SetRelation(*s)
	}
	return truo
}

// SetRelationID sets the "relation_id" field.
func (truo *TagsRelationUpdateOne) SetRelationID(i int) *TagsRelationUpdateOne {
	truo.mutation.SetRelationID(i)
	return truo
}

// SetNillableRelationID sets the "relation_id" field if the given value is not nil.
func (truo *TagsRelationUpdateOne) SetNillableRelationID(i *int) *TagsRelationUpdateOne {
	if i != nil {
		truo.SetRelationID(*i)
	}
	return truo
}

// SetBlogID sets the "blog" edge to the Blogs entity by ID.
func (truo *TagsRelationUpdateOne) SetBlogID(id int) *TagsRelationUpdateOne {
	truo.mutation.SetBlogID(id)
	return truo
}

// SetBlog sets the "blog" edge to the Blogs entity.
func (truo *TagsRelationUpdateOne) SetBlog(b *Blogs) *TagsRelationUpdateOne {
	return truo.SetBlogID(b.ID)
}

// SetTag sets the "tag" edge to the Tags entity.
func (truo *TagsRelationUpdateOne) SetTag(t *Tags) *TagsRelationUpdateOne {
	return truo.SetTagID(t.ID)
}

// Mutation returns the TagsRelationMutation object of the builder.
func (truo *TagsRelationUpdateOne) Mutation() *TagsRelationMutation {
	return truo.mutation
}

// ClearBlog clears the "blog" edge to the Blogs entity.
func (truo *TagsRelationUpdateOne) ClearBlog() *TagsRelationUpdateOne {
	truo.mutation.ClearBlog()
	return truo
}

// ClearTag clears the "tag" edge to the Tags entity.
func (truo *TagsRelationUpdateOne) ClearTag() *TagsRelationUpdateOne {
	truo.mutation.ClearTag()
	return truo
}

// Where appends a list predicates to the TagsRelationUpdate builder.
func (truo *TagsRelationUpdateOne) Where(ps ...predicate.TagsRelation) *TagsRelationUpdateOne {
	truo.mutation.Where(ps...)
	return truo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (truo *TagsRelationUpdateOne) Select(field string, fields ...string) *TagsRelationUpdateOne {
	truo.fields = append([]string{field}, fields...)
	return truo
}

// Save executes the query and returns the updated TagsRelation entity.
func (truo *TagsRelationUpdateOne) Save(ctx context.Context) (*TagsRelation, error) {
	return withHooks(ctx, truo.sqlSave, truo.mutation, truo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (truo *TagsRelationUpdateOne) SaveX(ctx context.Context) *TagsRelation {
	node, err := truo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (truo *TagsRelationUpdateOne) Exec(ctx context.Context) error {
	_, err := truo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (truo *TagsRelationUpdateOne) ExecX(ctx context.Context) {
	if err := truo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (truo *TagsRelationUpdateOne) check() error {
	if truo.mutation.BlogCleared() && len(truo.mutation.BlogIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "TagsRelation.blog"`)
	}
	if truo.mutation.TagCleared() && len(truo.mutation.TagIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "TagsRelation.tag"`)
	}
	return nil
}

func (truo *TagsRelationUpdateOne) sqlSave(ctx context.Context) (_node *TagsRelation, err error) {
	if err := truo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tagsrelation.Table, tagsrelation.Columns, sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt))
	id, ok := truo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TagsRelation.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := truo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tagsrelation.FieldID)
		for _, f := range fields {
			if !tagsrelation.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tagsrelation.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := truo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := truo.mutation.Relation(); ok {
		_spec.SetField(tagsrelation.FieldRelation, field.TypeString, value)
	}
	if truo.mutation.BlogCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.BlogTable,
			Columns: []string{tagsrelation.BlogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.BlogIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.BlogTable,
			Columns: []string{tagsrelation.BlogColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if truo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.TagTable,
			Columns: []string{tagsrelation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := truo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tagsrelation.TagTable,
			Columns: []string{tagsrelation.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TagsRelation{config: truo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, truo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tagsrelation.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	truo.mutation.done = true
	return _node, nil
}
