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

// TagsUpdate is the builder for updating Tags entities.
type TagsUpdate struct {
	config
	hooks    []Hook
	mutation *TagsMutation
}

// Where appends a list predicates to the TagsUpdate builder.
func (tu *TagsUpdate) Where(ps ...predicate.Tags) *TagsUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TagsUpdate) SetUpdatedAt(i int64) *TagsUpdate {
	tu.mutation.ResetUpdatedAt()
	tu.mutation.SetUpdatedAt(i)
	return tu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tu *TagsUpdate) AddUpdatedAt(i int64) *TagsUpdate {
	tu.mutation.AddUpdatedAt(i)
	return tu
}

// SetUpdatedBy sets the "updated_by" field.
func (tu *TagsUpdate) SetUpdatedBy(i int64) *TagsUpdate {
	tu.mutation.ResetUpdatedBy()
	tu.mutation.SetUpdatedBy(i)
	return tu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tu *TagsUpdate) SetNillableUpdatedBy(i *int64) *TagsUpdate {
	if i != nil {
		tu.SetUpdatedBy(*i)
	}
	return tu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (tu *TagsUpdate) AddUpdatedBy(i int64) *TagsUpdate {
	tu.mutation.AddUpdatedBy(i)
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TagsUpdate) SetDeletedAt(i int64) *TagsUpdate {
	tu.mutation.ResetDeletedAt()
	tu.mutation.SetDeletedAt(i)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *TagsUpdate) SetNillableDeletedAt(i *int64) *TagsUpdate {
	if i != nil {
		tu.SetDeletedAt(*i)
	}
	return tu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (tu *TagsUpdate) AddDeletedAt(i int64) *TagsUpdate {
	tu.mutation.AddDeletedAt(i)
	return tu
}

// SetDeletedBy sets the "deleted_by" field.
func (tu *TagsUpdate) SetDeletedBy(i int64) *TagsUpdate {
	tu.mutation.ResetDeletedBy()
	tu.mutation.SetDeletedBy(i)
	return tu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tu *TagsUpdate) SetNillableDeletedBy(i *int64) *TagsUpdate {
	if i != nil {
		tu.SetDeletedBy(*i)
	}
	return tu
}

// AddDeletedBy adds i to the "deleted_by" field.
func (tu *TagsUpdate) AddDeletedBy(i int64) *TagsUpdate {
	tu.mutation.AddDeletedBy(i)
	return tu
}

// SetName sets the "name" field.
func (tu *TagsUpdate) SetName(s string) *TagsUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TagsUpdate) SetNillableName(s *string) *TagsUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// AddBlogIDs adds the "blogs" edge to the Blogs entity by IDs.
func (tu *TagsUpdate) AddBlogIDs(ids ...int) *TagsUpdate {
	tu.mutation.AddBlogIDs(ids...)
	return tu
}

// AddBlogs adds the "blogs" edges to the Blogs entity.
func (tu *TagsUpdate) AddBlogs(b ...*Blogs) *TagsUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.AddBlogIDs(ids...)
}

// AddTagRelationIDs adds the "tag_relation" edge to the TagsRelation entity by IDs.
func (tu *TagsUpdate) AddTagRelationIDs(ids ...int) *TagsUpdate {
	tu.mutation.AddTagRelationIDs(ids...)
	return tu
}

// AddTagRelation adds the "tag_relation" edges to the TagsRelation entity.
func (tu *TagsUpdate) AddTagRelation(t ...*TagsRelation) *TagsUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTagRelationIDs(ids...)
}

// Mutation returns the TagsMutation object of the builder.
func (tu *TagsUpdate) Mutation() *TagsMutation {
	return tu.mutation
}

// ClearBlogs clears all "blogs" edges to the Blogs entity.
func (tu *TagsUpdate) ClearBlogs() *TagsUpdate {
	tu.mutation.ClearBlogs()
	return tu
}

// RemoveBlogIDs removes the "blogs" edge to Blogs entities by IDs.
func (tu *TagsUpdate) RemoveBlogIDs(ids ...int) *TagsUpdate {
	tu.mutation.RemoveBlogIDs(ids...)
	return tu
}

// RemoveBlogs removes "blogs" edges to Blogs entities.
func (tu *TagsUpdate) RemoveBlogs(b ...*Blogs) *TagsUpdate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tu.RemoveBlogIDs(ids...)
}

// ClearTagRelation clears all "tag_relation" edges to the TagsRelation entity.
func (tu *TagsUpdate) ClearTagRelation() *TagsUpdate {
	tu.mutation.ClearTagRelation()
	return tu
}

// RemoveTagRelationIDs removes the "tag_relation" edge to TagsRelation entities by IDs.
func (tu *TagsUpdate) RemoveTagRelationIDs(ids ...int) *TagsUpdate {
	tu.mutation.RemoveTagRelationIDs(ids...)
	return tu
}

// RemoveTagRelation removes "tag_relation" edges to TagsRelation entities.
func (tu *TagsUpdate) RemoveTagRelation(t ...*TagsRelation) *TagsUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTagRelationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagsUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagsUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagsUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagsUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TagsUpdate) defaults() error {
	if _, ok := tu.mutation.CreatedAt(); !ok {
		if tags.UpdateDefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized tags.UpdateDefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := tags.UpdateDefaultCreatedAt()
		tu.mutation.SetCreatedAt(v)
	}
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		if tags.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tags.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tags.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TagsUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tags.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tags.name": %w`, err)}
		}
	}
	return nil
}

func (tu *TagsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(tags.Table, tags.Columns, sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(tags.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(tags.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(tags.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(tags.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.UpdatedBy(); ok {
		_spec.SetField(tags.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(tags.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.SetField(tags.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(tags.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.DeletedBy(); ok {
		_spec.SetField(tags.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedDeletedBy(); ok {
		_spec.AddField(tags.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tags.FieldName, field.TypeString, value)
	}
	if tu.mutation.BlogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tags.BlogsTable,
			Columns: tags.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedBlogsIDs(); len(nodes) > 0 && !tu.mutation.BlogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tags.BlogsTable,
			Columns: tags.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.BlogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tags.BlogsTable,
			Columns: tags.BlogsPrimaryKey,
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
	if tu.mutation.TagRelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tags.TagRelationTable,
			Columns: []string{tags.TagRelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTagRelationIDs(); len(nodes) > 0 && !tu.mutation.TagRelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tags.TagRelationTable,
			Columns: []string{tags.TagRelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TagRelationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tags.TagRelationTable,
			Columns: []string{tags.TagRelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tags.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TagsUpdateOne is the builder for updating a single Tags entity.
type TagsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagsMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TagsUpdateOne) SetUpdatedAt(i int64) *TagsUpdateOne {
	tuo.mutation.ResetUpdatedAt()
	tuo.mutation.SetUpdatedAt(i)
	return tuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tuo *TagsUpdateOne) AddUpdatedAt(i int64) *TagsUpdateOne {
	tuo.mutation.AddUpdatedAt(i)
	return tuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tuo *TagsUpdateOne) SetUpdatedBy(i int64) *TagsUpdateOne {
	tuo.mutation.ResetUpdatedBy()
	tuo.mutation.SetUpdatedBy(i)
	return tuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tuo *TagsUpdateOne) SetNillableUpdatedBy(i *int64) *TagsUpdateOne {
	if i != nil {
		tuo.SetUpdatedBy(*i)
	}
	return tuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (tuo *TagsUpdateOne) AddUpdatedBy(i int64) *TagsUpdateOne {
	tuo.mutation.AddUpdatedBy(i)
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TagsUpdateOne) SetDeletedAt(i int64) *TagsUpdateOne {
	tuo.mutation.ResetDeletedAt()
	tuo.mutation.SetDeletedAt(i)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *TagsUpdateOne) SetNillableDeletedAt(i *int64) *TagsUpdateOne {
	if i != nil {
		tuo.SetDeletedAt(*i)
	}
	return tuo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (tuo *TagsUpdateOne) AddDeletedAt(i int64) *TagsUpdateOne {
	tuo.mutation.AddDeletedAt(i)
	return tuo
}

// SetDeletedBy sets the "deleted_by" field.
func (tuo *TagsUpdateOne) SetDeletedBy(i int64) *TagsUpdateOne {
	tuo.mutation.ResetDeletedBy()
	tuo.mutation.SetDeletedBy(i)
	return tuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tuo *TagsUpdateOne) SetNillableDeletedBy(i *int64) *TagsUpdateOne {
	if i != nil {
		tuo.SetDeletedBy(*i)
	}
	return tuo
}

// AddDeletedBy adds i to the "deleted_by" field.
func (tuo *TagsUpdateOne) AddDeletedBy(i int64) *TagsUpdateOne {
	tuo.mutation.AddDeletedBy(i)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TagsUpdateOne) SetName(s string) *TagsUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TagsUpdateOne) SetNillableName(s *string) *TagsUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// AddBlogIDs adds the "blogs" edge to the Blogs entity by IDs.
func (tuo *TagsUpdateOne) AddBlogIDs(ids ...int) *TagsUpdateOne {
	tuo.mutation.AddBlogIDs(ids...)
	return tuo
}

// AddBlogs adds the "blogs" edges to the Blogs entity.
func (tuo *TagsUpdateOne) AddBlogs(b ...*Blogs) *TagsUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.AddBlogIDs(ids...)
}

// AddTagRelationIDs adds the "tag_relation" edge to the TagsRelation entity by IDs.
func (tuo *TagsUpdateOne) AddTagRelationIDs(ids ...int) *TagsUpdateOne {
	tuo.mutation.AddTagRelationIDs(ids...)
	return tuo
}

// AddTagRelation adds the "tag_relation" edges to the TagsRelation entity.
func (tuo *TagsUpdateOne) AddTagRelation(t ...*TagsRelation) *TagsUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTagRelationIDs(ids...)
}

// Mutation returns the TagsMutation object of the builder.
func (tuo *TagsUpdateOne) Mutation() *TagsMutation {
	return tuo.mutation
}

// ClearBlogs clears all "blogs" edges to the Blogs entity.
func (tuo *TagsUpdateOne) ClearBlogs() *TagsUpdateOne {
	tuo.mutation.ClearBlogs()
	return tuo
}

// RemoveBlogIDs removes the "blogs" edge to Blogs entities by IDs.
func (tuo *TagsUpdateOne) RemoveBlogIDs(ids ...int) *TagsUpdateOne {
	tuo.mutation.RemoveBlogIDs(ids...)
	return tuo
}

// RemoveBlogs removes "blogs" edges to Blogs entities.
func (tuo *TagsUpdateOne) RemoveBlogs(b ...*Blogs) *TagsUpdateOne {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return tuo.RemoveBlogIDs(ids...)
}

// ClearTagRelation clears all "tag_relation" edges to the TagsRelation entity.
func (tuo *TagsUpdateOne) ClearTagRelation() *TagsUpdateOne {
	tuo.mutation.ClearTagRelation()
	return tuo
}

// RemoveTagRelationIDs removes the "tag_relation" edge to TagsRelation entities by IDs.
func (tuo *TagsUpdateOne) RemoveTagRelationIDs(ids ...int) *TagsUpdateOne {
	tuo.mutation.RemoveTagRelationIDs(ids...)
	return tuo
}

// RemoveTagRelation removes "tag_relation" edges to TagsRelation entities.
func (tuo *TagsUpdateOne) RemoveTagRelation(t ...*TagsRelation) *TagsUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTagRelationIDs(ids...)
}

// Where appends a list predicates to the TagsUpdate builder.
func (tuo *TagsUpdateOne) Where(ps ...predicate.Tags) *TagsUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagsUpdateOne) Select(field string, fields ...string) *TagsUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tags entity.
func (tuo *TagsUpdateOne) Save(ctx context.Context) (*Tags, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagsUpdateOne) SaveX(ctx context.Context) *Tags {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagsUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagsUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TagsUpdateOne) defaults() error {
	if _, ok := tuo.mutation.CreatedAt(); !ok {
		if tags.UpdateDefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized tags.UpdateDefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := tags.UpdateDefaultCreatedAt()
		tuo.mutation.SetCreatedAt(v)
	}
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		if tags.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized tags.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := tags.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TagsUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tags.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Tags.name": %w`, err)}
		}
	}
	return nil
}

func (tuo *TagsUpdateOne) sqlSave(ctx context.Context) (_node *Tags, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(tags.Table, tags.Columns, sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tags.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tags.FieldID)
		for _, f := range fields {
			if !tags.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tags.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(tags.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(tags.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tags.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(tags.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.UpdatedBy(); ok {
		_spec.SetField(tags.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(tags.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.SetField(tags.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(tags.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.DeletedBy(); ok {
		_spec.SetField(tags.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(tags.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tags.FieldName, field.TypeString, value)
	}
	if tuo.mutation.BlogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tags.BlogsTable,
			Columns: tags.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedBlogsIDs(); len(nodes) > 0 && !tuo.mutation.BlogsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tags.BlogsTable,
			Columns: tags.BlogsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.BlogsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tags.BlogsTable,
			Columns: tags.BlogsPrimaryKey,
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
	if tuo.mutation.TagRelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tags.TagRelationTable,
			Columns: []string{tags.TagRelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTagRelationIDs(); len(nodes) > 0 && !tuo.mutation.TagRelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tags.TagRelationTable,
			Columns: []string{tags.TagRelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TagRelationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tags.TagRelationTable,
			Columns: []string{tags.TagRelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tags{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tags.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
