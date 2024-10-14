// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/comment"
	"blog/internal/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdatedAt sets the "updated_at" field.
func (cu *CommentUpdate) SetUpdatedAt(i int64) *CommentUpdate {
	cu.mutation.ResetUpdatedAt()
	cu.mutation.SetUpdatedAt(i)
	return cu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (cu *CommentUpdate) AddUpdatedAt(i int64) *CommentUpdate {
	cu.mutation.AddUpdatedAt(i)
	return cu
}

// SetUpdatedBy sets the "updated_by" field.
func (cu *CommentUpdate) SetUpdatedBy(s string) *CommentUpdate {
	cu.mutation.SetUpdatedBy(s)
	return cu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableUpdatedBy(s *string) *CommentUpdate {
	if s != nil {
		cu.SetUpdatedBy(*s)
	}
	return cu
}

// SetIsDeleted sets the "is_deleted" field.
func (cu *CommentUpdate) SetIsDeleted(u uint8) *CommentUpdate {
	cu.mutation.ResetIsDeleted()
	cu.mutation.SetIsDeleted(u)
	return cu
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableIsDeleted(u *uint8) *CommentUpdate {
	if u != nil {
		cu.SetIsDeleted(*u)
	}
	return cu
}

// AddIsDeleted adds u to the "is_deleted" field.
func (cu *CommentUpdate) AddIsDeleted(u int8) *CommentUpdate {
	cu.mutation.AddIsDeleted(u)
	return cu
}

// SetDeletedAt sets the "deleted_at" field.
func (cu *CommentUpdate) SetDeletedAt(i int64) *CommentUpdate {
	cu.mutation.ResetDeletedAt()
	cu.mutation.SetDeletedAt(i)
	return cu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDeletedAt(i *int64) *CommentUpdate {
	if i != nil {
		cu.SetDeletedAt(*i)
	}
	return cu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (cu *CommentUpdate) AddDeletedAt(i int64) *CommentUpdate {
	cu.mutation.AddDeletedAt(i)
	return cu
}

// SetDeletedBy sets the "deleted_by" field.
func (cu *CommentUpdate) SetDeletedBy(s string) *CommentUpdate {
	cu.mutation.SetDeletedBy(s)
	return cu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableDeletedBy(s *string) *CommentUpdate {
	if s != nil {
		cu.SetDeletedBy(*s)
	}
	return cu
}

// SetAccountID sets the "account_id" field.
func (cu *CommentUpdate) SetAccountID(i int) *CommentUpdate {
	cu.mutation.ResetAccountID()
	cu.mutation.SetAccountID(i)
	return cu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableAccountID(i *int) *CommentUpdate {
	if i != nil {
		cu.SetAccountID(*i)
	}
	return cu
}

// AddAccountID adds i to the "account_id" field.
func (cu *CommentUpdate) AddAccountID(i int) *CommentUpdate {
	cu.mutation.AddAccountID(i)
	return cu
}

// SetBlogID sets the "blog_id" field.
func (cu *CommentUpdate) SetBlogID(i int) *CommentUpdate {
	cu.mutation.ResetBlogID()
	cu.mutation.SetBlogID(i)
	return cu
}

// SetNillableBlogID sets the "blog_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableBlogID(i *int) *CommentUpdate {
	if i != nil {
		cu.SetBlogID(*i)
	}
	return cu
}

// AddBlogID adds i to the "blog_id" field.
func (cu *CommentUpdate) AddBlogID(i int) *CommentUpdate {
	cu.mutation.AddBlogID(i)
	return cu
}

// SetTopID sets the "top_id" field.
func (cu *CommentUpdate) SetTopID(i int) *CommentUpdate {
	cu.mutation.ResetTopID()
	cu.mutation.SetTopID(i)
	return cu
}

// SetNillableTopID sets the "top_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableTopID(i *int) *CommentUpdate {
	if i != nil {
		cu.SetTopID(*i)
	}
	return cu
}

// AddTopID adds i to the "top_id" field.
func (cu *CommentUpdate) AddTopID(i int) *CommentUpdate {
	cu.mutation.AddTopID(i)
	return cu
}

// SetParentID sets the "parent_id" field.
func (cu *CommentUpdate) SetParentID(i int) *CommentUpdate {
	cu.mutation.ResetParentID()
	cu.mutation.SetParentID(i)
	return cu
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableParentID(i *int) *CommentUpdate {
	if i != nil {
		cu.SetParentID(*i)
	}
	return cu
}

// AddParentID adds i to the "parent_id" field.
func (cu *CommentUpdate) AddParentID(i int) *CommentUpdate {
	cu.mutation.AddParentID(i)
	return cu
}

// SetLevel sets the "level" field.
func (cu *CommentUpdate) SetLevel(i int) *CommentUpdate {
	cu.mutation.ResetLevel()
	cu.mutation.SetLevel(i)
	return cu
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableLevel(i *int) *CommentUpdate {
	if i != nil {
		cu.SetLevel(*i)
	}
	return cu
}

// AddLevel adds i to the "level" field.
func (cu *CommentUpdate) AddLevel(i int) *CommentUpdate {
	cu.mutation.AddLevel(i)
	return cu
}

// SetTotal sets the "total" field.
func (cu *CommentUpdate) SetTotal(i int) *CommentUpdate {
	cu.mutation.ResetTotal()
	cu.mutation.SetTotal(i)
	return cu
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableTotal(i *int) *CommentUpdate {
	if i != nil {
		cu.SetTotal(*i)
	}
	return cu
}

// AddTotal adds i to the "total" field.
func (cu *CommentUpdate) AddTotal(i int) *CommentUpdate {
	cu.mutation.AddTotal(i)
	return cu
}

// SetStatus sets the "status" field.
func (cu *CommentUpdate) SetStatus(i int8) *CommentUpdate {
	cu.mutation.ResetStatus()
	cu.mutation.SetStatus(i)
	return cu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableStatus(i *int8) *CommentUpdate {
	if i != nil {
		cu.SetStatus(*i)
	}
	return cu
}

// AddStatus adds i to the "status" field.
func (cu *CommentUpdate) AddStatus(i int8) *CommentUpdate {
	cu.mutation.AddStatus(i)
	return cu
}

// SetContent sets the "content" field.
func (cu *CommentUpdate) SetContent(s string) *CommentUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cu *CommentUpdate) SetNillableContent(s *string) *CommentUpdate {
	if s != nil {
		cu.SetContent(*s)
	}
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	if err := cu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CommentUpdate) defaults() error {
	if _, ok := cu.mutation.UpdatedAt(); !ok {
		if comment.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.UpdateDefaultUpdatedAt()
		cu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cu *CommentUpdate) check() error {
	if v, ok := cu.mutation.AccountID(); ok {
		if err := comment.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`ent: validator failed for field "Comment.account_id": %w`, err)}
		}
	}
	if v, ok := cu.mutation.BlogID(); ok {
		if err := comment.BlogIDValidator(v); err != nil {
			return &ValidationError{Name: "blog_id", err: fmt.Errorf(`ent: validator failed for field "Comment.blog_id": %w`, err)}
		}
	}
	return nil
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(comment.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.UpdatedBy(); ok {
		_spec.SetField(comment.FieldUpdatedBy, field.TypeString, value)
	}
	if value, ok := cu.mutation.IsDeleted(); ok {
		_spec.SetField(comment.FieldIsDeleted, field.TypeUint8, value)
	}
	if value, ok := cu.mutation.AddedIsDeleted(); ok {
		_spec.AddField(comment.FieldIsDeleted, field.TypeUint8, value)
	}
	if value, ok := cu.mutation.DeletedAt(); ok {
		_spec.SetField(comment.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(comment.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := cu.mutation.DeletedBy(); ok {
		_spec.SetField(comment.FieldDeletedBy, field.TypeString, value)
	}
	if value, ok := cu.mutation.AccountID(); ok {
		_spec.SetField(comment.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedAccountID(); ok {
		_spec.AddField(comment.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.BlogID(); ok {
		_spec.SetField(comment.FieldBlogID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedBlogID(); ok {
		_spec.AddField(comment.FieldBlogID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.TopID(); ok {
		_spec.SetField(comment.FieldTopID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedTopID(); ok {
		_spec.AddField(comment.FieldTopID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.ParentID(); ok {
		_spec.SetField(comment.FieldParentID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedParentID(); ok {
		_spec.AddField(comment.FieldParentID, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Level(); ok {
		_spec.SetField(comment.FieldLevel, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedLevel(); ok {
		_spec.AddField(comment.FieldLevel, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Total(); ok {
		_spec.SetField(comment.FieldTotal, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedTotal(); ok {
		_spec.AddField(comment.FieldTotal, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Status(); ok {
		_spec.SetField(comment.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.AddedStatus(); ok {
		_spec.AddField(comment.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (cuo *CommentUpdateOne) SetUpdatedAt(i int64) *CommentUpdateOne {
	cuo.mutation.ResetUpdatedAt()
	cuo.mutation.SetUpdatedAt(i)
	return cuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (cuo *CommentUpdateOne) AddUpdatedAt(i int64) *CommentUpdateOne {
	cuo.mutation.AddUpdatedAt(i)
	return cuo
}

// SetUpdatedBy sets the "updated_by" field.
func (cuo *CommentUpdateOne) SetUpdatedBy(s string) *CommentUpdateOne {
	cuo.mutation.SetUpdatedBy(s)
	return cuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableUpdatedBy(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetUpdatedBy(*s)
	}
	return cuo
}

// SetIsDeleted sets the "is_deleted" field.
func (cuo *CommentUpdateOne) SetIsDeleted(u uint8) *CommentUpdateOne {
	cuo.mutation.ResetIsDeleted()
	cuo.mutation.SetIsDeleted(u)
	return cuo
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableIsDeleted(u *uint8) *CommentUpdateOne {
	if u != nil {
		cuo.SetIsDeleted(*u)
	}
	return cuo
}

// AddIsDeleted adds u to the "is_deleted" field.
func (cuo *CommentUpdateOne) AddIsDeleted(u int8) *CommentUpdateOne {
	cuo.mutation.AddIsDeleted(u)
	return cuo
}

// SetDeletedAt sets the "deleted_at" field.
func (cuo *CommentUpdateOne) SetDeletedAt(i int64) *CommentUpdateOne {
	cuo.mutation.ResetDeletedAt()
	cuo.mutation.SetDeletedAt(i)
	return cuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDeletedAt(i *int64) *CommentUpdateOne {
	if i != nil {
		cuo.SetDeletedAt(*i)
	}
	return cuo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (cuo *CommentUpdateOne) AddDeletedAt(i int64) *CommentUpdateOne {
	cuo.mutation.AddDeletedAt(i)
	return cuo
}

// SetDeletedBy sets the "deleted_by" field.
func (cuo *CommentUpdateOne) SetDeletedBy(s string) *CommentUpdateOne {
	cuo.mutation.SetDeletedBy(s)
	return cuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableDeletedBy(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetDeletedBy(*s)
	}
	return cuo
}

// SetAccountID sets the "account_id" field.
func (cuo *CommentUpdateOne) SetAccountID(i int) *CommentUpdateOne {
	cuo.mutation.ResetAccountID()
	cuo.mutation.SetAccountID(i)
	return cuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableAccountID(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetAccountID(*i)
	}
	return cuo
}

// AddAccountID adds i to the "account_id" field.
func (cuo *CommentUpdateOne) AddAccountID(i int) *CommentUpdateOne {
	cuo.mutation.AddAccountID(i)
	return cuo
}

// SetBlogID sets the "blog_id" field.
func (cuo *CommentUpdateOne) SetBlogID(i int) *CommentUpdateOne {
	cuo.mutation.ResetBlogID()
	cuo.mutation.SetBlogID(i)
	return cuo
}

// SetNillableBlogID sets the "blog_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableBlogID(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetBlogID(*i)
	}
	return cuo
}

// AddBlogID adds i to the "blog_id" field.
func (cuo *CommentUpdateOne) AddBlogID(i int) *CommentUpdateOne {
	cuo.mutation.AddBlogID(i)
	return cuo
}

// SetTopID sets the "top_id" field.
func (cuo *CommentUpdateOne) SetTopID(i int) *CommentUpdateOne {
	cuo.mutation.ResetTopID()
	cuo.mutation.SetTopID(i)
	return cuo
}

// SetNillableTopID sets the "top_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableTopID(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetTopID(*i)
	}
	return cuo
}

// AddTopID adds i to the "top_id" field.
func (cuo *CommentUpdateOne) AddTopID(i int) *CommentUpdateOne {
	cuo.mutation.AddTopID(i)
	return cuo
}

// SetParentID sets the "parent_id" field.
func (cuo *CommentUpdateOne) SetParentID(i int) *CommentUpdateOne {
	cuo.mutation.ResetParentID()
	cuo.mutation.SetParentID(i)
	return cuo
}

// SetNillableParentID sets the "parent_id" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableParentID(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetParentID(*i)
	}
	return cuo
}

// AddParentID adds i to the "parent_id" field.
func (cuo *CommentUpdateOne) AddParentID(i int) *CommentUpdateOne {
	cuo.mutation.AddParentID(i)
	return cuo
}

// SetLevel sets the "level" field.
func (cuo *CommentUpdateOne) SetLevel(i int) *CommentUpdateOne {
	cuo.mutation.ResetLevel()
	cuo.mutation.SetLevel(i)
	return cuo
}

// SetNillableLevel sets the "level" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableLevel(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetLevel(*i)
	}
	return cuo
}

// AddLevel adds i to the "level" field.
func (cuo *CommentUpdateOne) AddLevel(i int) *CommentUpdateOne {
	cuo.mutation.AddLevel(i)
	return cuo
}

// SetTotal sets the "total" field.
func (cuo *CommentUpdateOne) SetTotal(i int) *CommentUpdateOne {
	cuo.mutation.ResetTotal()
	cuo.mutation.SetTotal(i)
	return cuo
}

// SetNillableTotal sets the "total" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableTotal(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetTotal(*i)
	}
	return cuo
}

// AddTotal adds i to the "total" field.
func (cuo *CommentUpdateOne) AddTotal(i int) *CommentUpdateOne {
	cuo.mutation.AddTotal(i)
	return cuo
}

// SetStatus sets the "status" field.
func (cuo *CommentUpdateOne) SetStatus(i int8) *CommentUpdateOne {
	cuo.mutation.ResetStatus()
	cuo.mutation.SetStatus(i)
	return cuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableStatus(i *int8) *CommentUpdateOne {
	if i != nil {
		cuo.SetStatus(*i)
	}
	return cuo
}

// AddStatus adds i to the "status" field.
func (cuo *CommentUpdateOne) AddStatus(i int8) *CommentUpdateOne {
	cuo.mutation.AddStatus(i)
	return cuo
}

// SetContent sets the "content" field.
func (cuo *CommentUpdateOne) SetContent(s string) *CommentUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableContent(s *string) *CommentUpdateOne {
	if s != nil {
		cuo.SetContent(*s)
	}
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (cuo *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Comment entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	if err := cuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CommentUpdateOne) defaults() error {
	if _, ok := cuo.mutation.UpdatedAt(); !ok {
		if comment.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized comment.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := comment.UpdateDefaultUpdatedAt()
		cuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CommentUpdateOne) check() error {
	if v, ok := cuo.mutation.AccountID(); ok {
		if err := comment.AccountIDValidator(v); err != nil {
			return &ValidationError{Name: "account_id", err: fmt.Errorf(`ent: validator failed for field "Comment.account_id": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.BlogID(); ok {
		if err := comment.BlogIDValidator(v); err != nil {
			return &ValidationError{Name: "blog_id", err: fmt.Errorf(`ent: validator failed for field "Comment.blog_id": %w`, err)}
		}
	}
	return nil
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdatedAt(); ok {
		_spec.SetField(comment.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(comment.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.UpdatedBy(); ok {
		_spec.SetField(comment.FieldUpdatedBy, field.TypeString, value)
	}
	if value, ok := cuo.mutation.IsDeleted(); ok {
		_spec.SetField(comment.FieldIsDeleted, field.TypeUint8, value)
	}
	if value, ok := cuo.mutation.AddedIsDeleted(); ok {
		_spec.AddField(comment.FieldIsDeleted, field.TypeUint8, value)
	}
	if value, ok := cuo.mutation.DeletedAt(); ok {
		_spec.SetField(comment.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(comment.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := cuo.mutation.DeletedBy(); ok {
		_spec.SetField(comment.FieldDeletedBy, field.TypeString, value)
	}
	if value, ok := cuo.mutation.AccountID(); ok {
		_spec.SetField(comment.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedAccountID(); ok {
		_spec.AddField(comment.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.BlogID(); ok {
		_spec.SetField(comment.FieldBlogID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedBlogID(); ok {
		_spec.AddField(comment.FieldBlogID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.TopID(); ok {
		_spec.SetField(comment.FieldTopID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedTopID(); ok {
		_spec.AddField(comment.FieldTopID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.ParentID(); ok {
		_spec.SetField(comment.FieldParentID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedParentID(); ok {
		_spec.AddField(comment.FieldParentID, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Level(); ok {
		_spec.SetField(comment.FieldLevel, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedLevel(); ok {
		_spec.AddField(comment.FieldLevel, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Total(); ok {
		_spec.SetField(comment.FieldTotal, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedTotal(); ok {
		_spec.AddField(comment.FieldTotal, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Status(); ok {
		_spec.SetField(comment.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.AddedStatus(); ok {
		_spec.AddField(comment.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.SetField(comment.FieldContent, field.TypeString, value)
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
