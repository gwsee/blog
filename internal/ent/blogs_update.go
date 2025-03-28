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
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// BlogsUpdate is the builder for updating Blogs entities.
type BlogsUpdate struct {
	config
	hooks    []Hook
	mutation *BlogsMutation
}

// Where appends a list predicates to the BlogsUpdate builder.
func (bu *BlogsUpdate) Where(ps ...predicate.Blogs) *BlogsUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BlogsUpdate) SetUpdatedAt(i int64) *BlogsUpdate {
	bu.mutation.ResetUpdatedAt()
	bu.mutation.SetUpdatedAt(i)
	return bu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (bu *BlogsUpdate) AddUpdatedAt(i int64) *BlogsUpdate {
	bu.mutation.AddUpdatedAt(i)
	return bu
}

// SetUpdatedBy sets the "updated_by" field.
func (bu *BlogsUpdate) SetUpdatedBy(i int64) *BlogsUpdate {
	bu.mutation.ResetUpdatedBy()
	bu.mutation.SetUpdatedBy(i)
	return bu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableUpdatedBy(i *int64) *BlogsUpdate {
	if i != nil {
		bu.SetUpdatedBy(*i)
	}
	return bu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (bu *BlogsUpdate) AddUpdatedBy(i int64) *BlogsUpdate {
	bu.mutation.AddUpdatedBy(i)
	return bu
}

// SetDeletedAt sets the "deleted_at" field.
func (bu *BlogsUpdate) SetDeletedAt(i int64) *BlogsUpdate {
	bu.mutation.ResetDeletedAt()
	bu.mutation.SetDeletedAt(i)
	return bu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableDeletedAt(i *int64) *BlogsUpdate {
	if i != nil {
		bu.SetDeletedAt(*i)
	}
	return bu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (bu *BlogsUpdate) AddDeletedAt(i int64) *BlogsUpdate {
	bu.mutation.AddDeletedAt(i)
	return bu
}

// SetDeletedBy sets the "deleted_by" field.
func (bu *BlogsUpdate) SetDeletedBy(i int64) *BlogsUpdate {
	bu.mutation.ResetDeletedBy()
	bu.mutation.SetDeletedBy(i)
	return bu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableDeletedBy(i *int64) *BlogsUpdate {
	if i != nil {
		bu.SetDeletedBy(*i)
	}
	return bu
}

// AddDeletedBy adds i to the "deleted_by" field.
func (bu *BlogsUpdate) AddDeletedBy(i int64) *BlogsUpdate {
	bu.mutation.AddDeletedBy(i)
	return bu
}

// SetAccountID sets the "account_id" field.
func (bu *BlogsUpdate) SetAccountID(i int) *BlogsUpdate {
	bu.mutation.ResetAccountID()
	bu.mutation.SetAccountID(i)
	return bu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableAccountID(i *int) *BlogsUpdate {
	if i != nil {
		bu.SetAccountID(*i)
	}
	return bu
}

// AddAccountID adds i to the "account_id" field.
func (bu *BlogsUpdate) AddAccountID(i int) *BlogsUpdate {
	bu.mutation.AddAccountID(i)
	return bu
}

// SetTitle sets the "title" field.
func (bu *BlogsUpdate) SetTitle(s string) *BlogsUpdate {
	bu.mutation.SetTitle(s)
	return bu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableTitle(s *string) *BlogsUpdate {
	if s != nil {
		bu.SetTitle(*s)
	}
	return bu
}

// SetDescription sets the "description" field.
func (bu *BlogsUpdate) SetDescription(s string) *BlogsUpdate {
	bu.mutation.SetDescription(s)
	return bu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableDescription(s *string) *BlogsUpdate {
	if s != nil {
		bu.SetDescription(*s)
	}
	return bu
}

// SetIsHidden sets the "is_hidden" field.
func (bu *BlogsUpdate) SetIsHidden(i int8) *BlogsUpdate {
	bu.mutation.ResetIsHidden()
	bu.mutation.SetIsHidden(i)
	return bu
}

// SetNillableIsHidden sets the "is_hidden" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableIsHidden(i *int8) *BlogsUpdate {
	if i != nil {
		bu.SetIsHidden(*i)
	}
	return bu
}

// AddIsHidden adds i to the "is_hidden" field.
func (bu *BlogsUpdate) AddIsHidden(i int8) *BlogsUpdate {
	bu.mutation.AddIsHidden(i)
	return bu
}

// SetTags sets the "tags" field.
func (bu *BlogsUpdate) SetTags(s []string) *BlogsUpdate {
	bu.mutation.SetTags(s)
	return bu
}

// AppendTags appends s to the "tags" field.
func (bu *BlogsUpdate) AppendTags(s []string) *BlogsUpdate {
	bu.mutation.AppendTags(s)
	return bu
}

// SetCover sets the "cover" field.
func (bu *BlogsUpdate) SetCover(s string) *BlogsUpdate {
	bu.mutation.SetCover(s)
	return bu
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableCover(s *string) *BlogsUpdate {
	if s != nil {
		bu.SetCover(*s)
	}
	return bu
}

// SetBrowseNum sets the "browse_num" field.
func (bu *BlogsUpdate) SetBrowseNum(i int) *BlogsUpdate {
	bu.mutation.ResetBrowseNum()
	bu.mutation.SetBrowseNum(i)
	return bu
}

// SetNillableBrowseNum sets the "browse_num" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableBrowseNum(i *int) *BlogsUpdate {
	if i != nil {
		bu.SetBrowseNum(*i)
	}
	return bu
}

// AddBrowseNum adds i to the "browse_num" field.
func (bu *BlogsUpdate) AddBrowseNum(i int) *BlogsUpdate {
	bu.mutation.AddBrowseNum(i)
	return bu
}

// SetCollectNum sets the "collect_num" field.
func (bu *BlogsUpdate) SetCollectNum(i int) *BlogsUpdate {
	bu.mutation.ResetCollectNum()
	bu.mutation.SetCollectNum(i)
	return bu
}

// SetNillableCollectNum sets the "collect_num" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableCollectNum(i *int) *BlogsUpdate {
	if i != nil {
		bu.SetCollectNum(*i)
	}
	return bu
}

// AddCollectNum adds i to the "collect_num" field.
func (bu *BlogsUpdate) AddCollectNum(i int) *BlogsUpdate {
	bu.mutation.AddCollectNum(i)
	return bu
}

// SetLoveNum sets the "love_num" field.
func (bu *BlogsUpdate) SetLoveNum(i int) *BlogsUpdate {
	bu.mutation.ResetLoveNum()
	bu.mutation.SetLoveNum(i)
	return bu
}

// SetNillableLoveNum sets the "love_num" field if the given value is not nil.
func (bu *BlogsUpdate) SetNillableLoveNum(i *int) *BlogsUpdate {
	if i != nil {
		bu.SetLoveNum(*i)
	}
	return bu
}

// AddLoveNum adds i to the "love_num" field.
func (bu *BlogsUpdate) AddLoveNum(i int) *BlogsUpdate {
	bu.mutation.AddLoveNum(i)
	return bu
}

// AddTagIDs adds the "tag" edge to the Tags entity by IDs.
func (bu *BlogsUpdate) AddTagIDs(ids ...int) *BlogsUpdate {
	bu.mutation.AddTagIDs(ids...)
	return bu
}

// AddTag adds the "tag" edges to the Tags entity.
func (bu *BlogsUpdate) AddTag(t ...*Tags) *BlogsUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTagIDs(ids...)
}

// AddTagRelationIDs adds the "tag_relation" edge to the TagsRelation entity by IDs.
func (bu *BlogsUpdate) AddTagRelationIDs(ids ...int) *BlogsUpdate {
	bu.mutation.AddTagRelationIDs(ids...)
	return bu
}

// AddTagRelation adds the "tag_relation" edges to the TagsRelation entity.
func (bu *BlogsUpdate) AddTagRelation(t ...*TagsRelation) *BlogsUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.AddTagRelationIDs(ids...)
}

// Mutation returns the BlogsMutation object of the builder.
func (bu *BlogsUpdate) Mutation() *BlogsMutation {
	return bu.mutation
}

// ClearTag clears all "tag" edges to the Tags entity.
func (bu *BlogsUpdate) ClearTag() *BlogsUpdate {
	bu.mutation.ClearTag()
	return bu
}

// RemoveTagIDs removes the "tag" edge to Tags entities by IDs.
func (bu *BlogsUpdate) RemoveTagIDs(ids ...int) *BlogsUpdate {
	bu.mutation.RemoveTagIDs(ids...)
	return bu
}

// RemoveTag removes "tag" edges to Tags entities.
func (bu *BlogsUpdate) RemoveTag(t ...*Tags) *BlogsUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTagIDs(ids...)
}

// ClearTagRelation clears all "tag_relation" edges to the TagsRelation entity.
func (bu *BlogsUpdate) ClearTagRelation() *BlogsUpdate {
	bu.mutation.ClearTagRelation()
	return bu
}

// RemoveTagRelationIDs removes the "tag_relation" edge to TagsRelation entities by IDs.
func (bu *BlogsUpdate) RemoveTagRelationIDs(ids ...int) *BlogsUpdate {
	bu.mutation.RemoveTagRelationIDs(ids...)
	return bu
}

// RemoveTagRelation removes "tag_relation" edges to TagsRelation entities.
func (bu *BlogsUpdate) RemoveTagRelation(t ...*TagsRelation) *BlogsUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return bu.RemoveTagRelationIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BlogsUpdate) Save(ctx context.Context) (int, error) {
	if err := bu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BlogsUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BlogsUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BlogsUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BlogsUpdate) defaults() error {
	if _, ok := bu.mutation.CreatedAt(); !ok {
		if blogs.UpdateDefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized blogs.UpdateDefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := blogs.UpdateDefaultCreatedAt()
		bu.mutation.SetCreatedAt(v)
	}
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		if blogs.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized blogs.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := blogs.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (bu *BlogsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(blogs.Table, blogs.Columns, sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.SetField(blogs.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(blogs.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(blogs.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(blogs.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.UpdatedBy(); ok {
		_spec.SetField(blogs.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(blogs.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.DeletedAt(); ok {
		_spec.SetField(blogs.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(blogs.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.DeletedBy(); ok {
		_spec.SetField(blogs.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AddedDeletedBy(); ok {
		_spec.AddField(blogs.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := bu.mutation.AccountID(); ok {
		_spec.SetField(blogs.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedAccountID(); ok {
		_spec.AddField(blogs.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := bu.mutation.Title(); ok {
		_spec.SetField(blogs.FieldTitle, field.TypeString, value)
	}
	if value, ok := bu.mutation.Description(); ok {
		_spec.SetField(blogs.FieldDescription, field.TypeString, value)
	}
	if value, ok := bu.mutation.IsHidden(); ok {
		_spec.SetField(blogs.FieldIsHidden, field.TypeInt8, value)
	}
	if value, ok := bu.mutation.AddedIsHidden(); ok {
		_spec.AddField(blogs.FieldIsHidden, field.TypeInt8, value)
	}
	if value, ok := bu.mutation.Tags(); ok {
		_spec.SetField(blogs.FieldTags, field.TypeJSON, value)
	}
	if value, ok := bu.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, blogs.FieldTags, value)
		})
	}
	if value, ok := bu.mutation.Cover(); ok {
		_spec.SetField(blogs.FieldCover, field.TypeString, value)
	}
	if value, ok := bu.mutation.BrowseNum(); ok {
		_spec.SetField(blogs.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedBrowseNum(); ok {
		_spec.AddField(blogs.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := bu.mutation.CollectNum(); ok {
		_spec.SetField(blogs.FieldCollectNum, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedCollectNum(); ok {
		_spec.AddField(blogs.FieldCollectNum, field.TypeInt, value)
	}
	if value, ok := bu.mutation.LoveNum(); ok {
		_spec.SetField(blogs.FieldLoveNum, field.TypeInt, value)
	}
	if value, ok := bu.mutation.AddedLoveNum(); ok {
		_spec.AddField(blogs.FieldLoveNum, field.TypeInt, value)
	}
	if bu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blogs.TagTable,
			Columns: blogs.TagPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTagIDs(); len(nodes) > 0 && !bu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blogs.TagTable,
			Columns: blogs.TagPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blogs.TagTable,
			Columns: blogs.TagPrimaryKey,
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
	if bu.mutation.TagRelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blogs.TagRelationTable,
			Columns: []string{blogs.TagRelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedTagRelationIDs(); len(nodes) > 0 && !bu.mutation.TagRelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blogs.TagRelationTable,
			Columns: []string{blogs.TagRelationColumn},
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
	if nodes := bu.mutation.TagRelationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blogs.TagRelationTable,
			Columns: []string{blogs.TagRelationColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blogs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BlogsUpdateOne is the builder for updating a single Blogs entity.
type BlogsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogsMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BlogsUpdateOne) SetUpdatedAt(i int64) *BlogsUpdateOne {
	buo.mutation.ResetUpdatedAt()
	buo.mutation.SetUpdatedAt(i)
	return buo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (buo *BlogsUpdateOne) AddUpdatedAt(i int64) *BlogsUpdateOne {
	buo.mutation.AddUpdatedAt(i)
	return buo
}

// SetUpdatedBy sets the "updated_by" field.
func (buo *BlogsUpdateOne) SetUpdatedBy(i int64) *BlogsUpdateOne {
	buo.mutation.ResetUpdatedBy()
	buo.mutation.SetUpdatedBy(i)
	return buo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableUpdatedBy(i *int64) *BlogsUpdateOne {
	if i != nil {
		buo.SetUpdatedBy(*i)
	}
	return buo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (buo *BlogsUpdateOne) AddUpdatedBy(i int64) *BlogsUpdateOne {
	buo.mutation.AddUpdatedBy(i)
	return buo
}

// SetDeletedAt sets the "deleted_at" field.
func (buo *BlogsUpdateOne) SetDeletedAt(i int64) *BlogsUpdateOne {
	buo.mutation.ResetDeletedAt()
	buo.mutation.SetDeletedAt(i)
	return buo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableDeletedAt(i *int64) *BlogsUpdateOne {
	if i != nil {
		buo.SetDeletedAt(*i)
	}
	return buo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (buo *BlogsUpdateOne) AddDeletedAt(i int64) *BlogsUpdateOne {
	buo.mutation.AddDeletedAt(i)
	return buo
}

// SetDeletedBy sets the "deleted_by" field.
func (buo *BlogsUpdateOne) SetDeletedBy(i int64) *BlogsUpdateOne {
	buo.mutation.ResetDeletedBy()
	buo.mutation.SetDeletedBy(i)
	return buo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableDeletedBy(i *int64) *BlogsUpdateOne {
	if i != nil {
		buo.SetDeletedBy(*i)
	}
	return buo
}

// AddDeletedBy adds i to the "deleted_by" field.
func (buo *BlogsUpdateOne) AddDeletedBy(i int64) *BlogsUpdateOne {
	buo.mutation.AddDeletedBy(i)
	return buo
}

// SetAccountID sets the "account_id" field.
func (buo *BlogsUpdateOne) SetAccountID(i int) *BlogsUpdateOne {
	buo.mutation.ResetAccountID()
	buo.mutation.SetAccountID(i)
	return buo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableAccountID(i *int) *BlogsUpdateOne {
	if i != nil {
		buo.SetAccountID(*i)
	}
	return buo
}

// AddAccountID adds i to the "account_id" field.
func (buo *BlogsUpdateOne) AddAccountID(i int) *BlogsUpdateOne {
	buo.mutation.AddAccountID(i)
	return buo
}

// SetTitle sets the "title" field.
func (buo *BlogsUpdateOne) SetTitle(s string) *BlogsUpdateOne {
	buo.mutation.SetTitle(s)
	return buo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableTitle(s *string) *BlogsUpdateOne {
	if s != nil {
		buo.SetTitle(*s)
	}
	return buo
}

// SetDescription sets the "description" field.
func (buo *BlogsUpdateOne) SetDescription(s string) *BlogsUpdateOne {
	buo.mutation.SetDescription(s)
	return buo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableDescription(s *string) *BlogsUpdateOne {
	if s != nil {
		buo.SetDescription(*s)
	}
	return buo
}

// SetIsHidden sets the "is_hidden" field.
func (buo *BlogsUpdateOne) SetIsHidden(i int8) *BlogsUpdateOne {
	buo.mutation.ResetIsHidden()
	buo.mutation.SetIsHidden(i)
	return buo
}

// SetNillableIsHidden sets the "is_hidden" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableIsHidden(i *int8) *BlogsUpdateOne {
	if i != nil {
		buo.SetIsHidden(*i)
	}
	return buo
}

// AddIsHidden adds i to the "is_hidden" field.
func (buo *BlogsUpdateOne) AddIsHidden(i int8) *BlogsUpdateOne {
	buo.mutation.AddIsHidden(i)
	return buo
}

// SetTags sets the "tags" field.
func (buo *BlogsUpdateOne) SetTags(s []string) *BlogsUpdateOne {
	buo.mutation.SetTags(s)
	return buo
}

// AppendTags appends s to the "tags" field.
func (buo *BlogsUpdateOne) AppendTags(s []string) *BlogsUpdateOne {
	buo.mutation.AppendTags(s)
	return buo
}

// SetCover sets the "cover" field.
func (buo *BlogsUpdateOne) SetCover(s string) *BlogsUpdateOne {
	buo.mutation.SetCover(s)
	return buo
}

// SetNillableCover sets the "cover" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableCover(s *string) *BlogsUpdateOne {
	if s != nil {
		buo.SetCover(*s)
	}
	return buo
}

// SetBrowseNum sets the "browse_num" field.
func (buo *BlogsUpdateOne) SetBrowseNum(i int) *BlogsUpdateOne {
	buo.mutation.ResetBrowseNum()
	buo.mutation.SetBrowseNum(i)
	return buo
}

// SetNillableBrowseNum sets the "browse_num" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableBrowseNum(i *int) *BlogsUpdateOne {
	if i != nil {
		buo.SetBrowseNum(*i)
	}
	return buo
}

// AddBrowseNum adds i to the "browse_num" field.
func (buo *BlogsUpdateOne) AddBrowseNum(i int) *BlogsUpdateOne {
	buo.mutation.AddBrowseNum(i)
	return buo
}

// SetCollectNum sets the "collect_num" field.
func (buo *BlogsUpdateOne) SetCollectNum(i int) *BlogsUpdateOne {
	buo.mutation.ResetCollectNum()
	buo.mutation.SetCollectNum(i)
	return buo
}

// SetNillableCollectNum sets the "collect_num" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableCollectNum(i *int) *BlogsUpdateOne {
	if i != nil {
		buo.SetCollectNum(*i)
	}
	return buo
}

// AddCollectNum adds i to the "collect_num" field.
func (buo *BlogsUpdateOne) AddCollectNum(i int) *BlogsUpdateOne {
	buo.mutation.AddCollectNum(i)
	return buo
}

// SetLoveNum sets the "love_num" field.
func (buo *BlogsUpdateOne) SetLoveNum(i int) *BlogsUpdateOne {
	buo.mutation.ResetLoveNum()
	buo.mutation.SetLoveNum(i)
	return buo
}

// SetNillableLoveNum sets the "love_num" field if the given value is not nil.
func (buo *BlogsUpdateOne) SetNillableLoveNum(i *int) *BlogsUpdateOne {
	if i != nil {
		buo.SetLoveNum(*i)
	}
	return buo
}

// AddLoveNum adds i to the "love_num" field.
func (buo *BlogsUpdateOne) AddLoveNum(i int) *BlogsUpdateOne {
	buo.mutation.AddLoveNum(i)
	return buo
}

// AddTagIDs adds the "tag" edge to the Tags entity by IDs.
func (buo *BlogsUpdateOne) AddTagIDs(ids ...int) *BlogsUpdateOne {
	buo.mutation.AddTagIDs(ids...)
	return buo
}

// AddTag adds the "tag" edges to the Tags entity.
func (buo *BlogsUpdateOne) AddTag(t ...*Tags) *BlogsUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTagIDs(ids...)
}

// AddTagRelationIDs adds the "tag_relation" edge to the TagsRelation entity by IDs.
func (buo *BlogsUpdateOne) AddTagRelationIDs(ids ...int) *BlogsUpdateOne {
	buo.mutation.AddTagRelationIDs(ids...)
	return buo
}

// AddTagRelation adds the "tag_relation" edges to the TagsRelation entity.
func (buo *BlogsUpdateOne) AddTagRelation(t ...*TagsRelation) *BlogsUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.AddTagRelationIDs(ids...)
}

// Mutation returns the BlogsMutation object of the builder.
func (buo *BlogsUpdateOne) Mutation() *BlogsMutation {
	return buo.mutation
}

// ClearTag clears all "tag" edges to the Tags entity.
func (buo *BlogsUpdateOne) ClearTag() *BlogsUpdateOne {
	buo.mutation.ClearTag()
	return buo
}

// RemoveTagIDs removes the "tag" edge to Tags entities by IDs.
func (buo *BlogsUpdateOne) RemoveTagIDs(ids ...int) *BlogsUpdateOne {
	buo.mutation.RemoveTagIDs(ids...)
	return buo
}

// RemoveTag removes "tag" edges to Tags entities.
func (buo *BlogsUpdateOne) RemoveTag(t ...*Tags) *BlogsUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTagIDs(ids...)
}

// ClearTagRelation clears all "tag_relation" edges to the TagsRelation entity.
func (buo *BlogsUpdateOne) ClearTagRelation() *BlogsUpdateOne {
	buo.mutation.ClearTagRelation()
	return buo
}

// RemoveTagRelationIDs removes the "tag_relation" edge to TagsRelation entities by IDs.
func (buo *BlogsUpdateOne) RemoveTagRelationIDs(ids ...int) *BlogsUpdateOne {
	buo.mutation.RemoveTagRelationIDs(ids...)
	return buo
}

// RemoveTagRelation removes "tag_relation" edges to TagsRelation entities.
func (buo *BlogsUpdateOne) RemoveTagRelation(t ...*TagsRelation) *BlogsUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return buo.RemoveTagRelationIDs(ids...)
}

// Where appends a list predicates to the BlogsUpdate builder.
func (buo *BlogsUpdateOne) Where(ps ...predicate.Blogs) *BlogsUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BlogsUpdateOne) Select(field string, fields ...string) *BlogsUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Blogs entity.
func (buo *BlogsUpdateOne) Save(ctx context.Context) (*Blogs, error) {
	if err := buo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BlogsUpdateOne) SaveX(ctx context.Context) *Blogs {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BlogsUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BlogsUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BlogsUpdateOne) defaults() error {
	if _, ok := buo.mutation.CreatedAt(); !ok {
		if blogs.UpdateDefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized blogs.UpdateDefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := blogs.UpdateDefaultCreatedAt()
		buo.mutation.SetCreatedAt(v)
	}
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		if blogs.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized blogs.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := blogs.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (buo *BlogsUpdateOne) sqlSave(ctx context.Context) (_node *Blogs, err error) {
	_spec := sqlgraph.NewUpdateSpec(blogs.Table, blogs.Columns, sqlgraph.NewFieldSpec(blogs.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Blogs.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blogs.FieldID)
		for _, f := range fields {
			if !blogs.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blogs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.SetField(blogs.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(blogs.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(blogs.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(blogs.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.UpdatedBy(); ok {
		_spec.SetField(blogs.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(blogs.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.DeletedAt(); ok {
		_spec.SetField(blogs.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(blogs.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.DeletedBy(); ok {
		_spec.SetField(blogs.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(blogs.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := buo.mutation.AccountID(); ok {
		_spec.SetField(blogs.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedAccountID(); ok {
		_spec.AddField(blogs.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := buo.mutation.Title(); ok {
		_spec.SetField(blogs.FieldTitle, field.TypeString, value)
	}
	if value, ok := buo.mutation.Description(); ok {
		_spec.SetField(blogs.FieldDescription, field.TypeString, value)
	}
	if value, ok := buo.mutation.IsHidden(); ok {
		_spec.SetField(blogs.FieldIsHidden, field.TypeInt8, value)
	}
	if value, ok := buo.mutation.AddedIsHidden(); ok {
		_spec.AddField(blogs.FieldIsHidden, field.TypeInt8, value)
	}
	if value, ok := buo.mutation.Tags(); ok {
		_spec.SetField(blogs.FieldTags, field.TypeJSON, value)
	}
	if value, ok := buo.mutation.AppendedTags(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, blogs.FieldTags, value)
		})
	}
	if value, ok := buo.mutation.Cover(); ok {
		_spec.SetField(blogs.FieldCover, field.TypeString, value)
	}
	if value, ok := buo.mutation.BrowseNum(); ok {
		_spec.SetField(blogs.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedBrowseNum(); ok {
		_spec.AddField(blogs.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := buo.mutation.CollectNum(); ok {
		_spec.SetField(blogs.FieldCollectNum, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedCollectNum(); ok {
		_spec.AddField(blogs.FieldCollectNum, field.TypeInt, value)
	}
	if value, ok := buo.mutation.LoveNum(); ok {
		_spec.SetField(blogs.FieldLoveNum, field.TypeInt, value)
	}
	if value, ok := buo.mutation.AddedLoveNum(); ok {
		_spec.AddField(blogs.FieldLoveNum, field.TypeInt, value)
	}
	if buo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blogs.TagTable,
			Columns: blogs.TagPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTagIDs(); len(nodes) > 0 && !buo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blogs.TagTable,
			Columns: blogs.TagPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tags.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   blogs.TagTable,
			Columns: blogs.TagPrimaryKey,
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
	if buo.mutation.TagRelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blogs.TagRelationTable,
			Columns: []string{blogs.TagRelationColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tagsrelation.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedTagRelationIDs(); len(nodes) > 0 && !buo.mutation.TagRelationCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blogs.TagRelationTable,
			Columns: []string{blogs.TagRelationColumn},
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
	if nodes := buo.mutation.TagRelationIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   blogs.TagRelationTable,
			Columns: []string{blogs.TagRelationColumn},
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
	_node = &Blogs{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blogs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
