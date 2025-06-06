// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogsextend"
	"blog/internal/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlogsExtendUpdate is the builder for updating BlogsExtend entities.
type BlogsExtendUpdate struct {
	config
	hooks    []Hook
	mutation *BlogsExtendMutation
}

// Where appends a list predicates to the BlogsExtendUpdate builder.
func (beu *BlogsExtendUpdate) Where(ps ...predicate.BlogsExtend) *BlogsExtendUpdate {
	beu.mutation.Where(ps...)
	return beu
}

// SetBlogID sets the "blog_id" field.
func (beu *BlogsExtendUpdate) SetBlogID(i int) *BlogsExtendUpdate {
	beu.mutation.ResetBlogID()
	beu.mutation.SetBlogID(i)
	return beu
}

// SetNillableBlogID sets the "blog_id" field if the given value is not nil.
func (beu *BlogsExtendUpdate) SetNillableBlogID(i *int) *BlogsExtendUpdate {
	if i != nil {
		beu.SetBlogID(*i)
	}
	return beu
}

// AddBlogID adds i to the "blog_id" field.
func (beu *BlogsExtendUpdate) AddBlogID(i int) *BlogsExtendUpdate {
	beu.mutation.AddBlogID(i)
	return beu
}

// SetAccountID sets the "account_id" field.
func (beu *BlogsExtendUpdate) SetAccountID(i int) *BlogsExtendUpdate {
	beu.mutation.ResetAccountID()
	beu.mutation.SetAccountID(i)
	return beu
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (beu *BlogsExtendUpdate) SetNillableAccountID(i *int) *BlogsExtendUpdate {
	if i != nil {
		beu.SetAccountID(*i)
	}
	return beu
}

// AddAccountID adds i to the "account_id" field.
func (beu *BlogsExtendUpdate) AddAccountID(i int) *BlogsExtendUpdate {
	beu.mutation.AddAccountID(i)
	return beu
}

// SetBrowseNum sets the "browse_num" field.
func (beu *BlogsExtendUpdate) SetBrowseNum(i int) *BlogsExtendUpdate {
	beu.mutation.ResetBrowseNum()
	beu.mutation.SetBrowseNum(i)
	return beu
}

// SetNillableBrowseNum sets the "browse_num" field if the given value is not nil.
func (beu *BlogsExtendUpdate) SetNillableBrowseNum(i *int) *BlogsExtendUpdate {
	if i != nil {
		beu.SetBrowseNum(*i)
	}
	return beu
}

// AddBrowseNum adds i to the "browse_num" field.
func (beu *BlogsExtendUpdate) AddBrowseNum(i int) *BlogsExtendUpdate {
	beu.mutation.AddBrowseNum(i)
	return beu
}

// SetBrowseAt sets the "browse_at" field.
func (beu *BlogsExtendUpdate) SetBrowseAt(i int64) *BlogsExtendUpdate {
	beu.mutation.ResetBrowseAt()
	beu.mutation.SetBrowseAt(i)
	return beu
}

// SetNillableBrowseAt sets the "browse_at" field if the given value is not nil.
func (beu *BlogsExtendUpdate) SetNillableBrowseAt(i *int64) *BlogsExtendUpdate {
	if i != nil {
		beu.SetBrowseAt(*i)
	}
	return beu
}

// AddBrowseAt adds i to the "browse_at" field.
func (beu *BlogsExtendUpdate) AddBrowseAt(i int64) *BlogsExtendUpdate {
	beu.mutation.AddBrowseAt(i)
	return beu
}

// SetCollect sets the "collect" field.
func (beu *BlogsExtendUpdate) SetCollect(b bool) *BlogsExtendUpdate {
	beu.mutation.SetCollect(b)
	return beu
}

// SetNillableCollect sets the "collect" field if the given value is not nil.
func (beu *BlogsExtendUpdate) SetNillableCollect(b *bool) *BlogsExtendUpdate {
	if b != nil {
		beu.SetCollect(*b)
	}
	return beu
}

// SetCollectAt sets the "collect_at" field.
func (beu *BlogsExtendUpdate) SetCollectAt(i int64) *BlogsExtendUpdate {
	beu.mutation.ResetCollectAt()
	beu.mutation.SetCollectAt(i)
	return beu
}

// SetNillableCollectAt sets the "collect_at" field if the given value is not nil.
func (beu *BlogsExtendUpdate) SetNillableCollectAt(i *int64) *BlogsExtendUpdate {
	if i != nil {
		beu.SetCollectAt(*i)
	}
	return beu
}

// AddCollectAt adds i to the "collect_at" field.
func (beu *BlogsExtendUpdate) AddCollectAt(i int64) *BlogsExtendUpdate {
	beu.mutation.AddCollectAt(i)
	return beu
}

// SetLove sets the "love" field.
func (beu *BlogsExtendUpdate) SetLove(b bool) *BlogsExtendUpdate {
	beu.mutation.SetLove(b)
	return beu
}

// SetNillableLove sets the "love" field if the given value is not nil.
func (beu *BlogsExtendUpdate) SetNillableLove(b *bool) *BlogsExtendUpdate {
	if b != nil {
		beu.SetLove(*b)
	}
	return beu
}

// SetLoveAt sets the "love_at" field.
func (beu *BlogsExtendUpdate) SetLoveAt(i int64) *BlogsExtendUpdate {
	beu.mutation.ResetLoveAt()
	beu.mutation.SetLoveAt(i)
	return beu
}

// SetNillableLoveAt sets the "love_at" field if the given value is not nil.
func (beu *BlogsExtendUpdate) SetNillableLoveAt(i *int64) *BlogsExtendUpdate {
	if i != nil {
		beu.SetLoveAt(*i)
	}
	return beu
}

// AddLoveAt adds i to the "love_at" field.
func (beu *BlogsExtendUpdate) AddLoveAt(i int64) *BlogsExtendUpdate {
	beu.mutation.AddLoveAt(i)
	return beu
}

// Mutation returns the BlogsExtendMutation object of the builder.
func (beu *BlogsExtendUpdate) Mutation() *BlogsExtendMutation {
	return beu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (beu *BlogsExtendUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, beu.sqlSave, beu.mutation, beu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (beu *BlogsExtendUpdate) SaveX(ctx context.Context) int {
	affected, err := beu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (beu *BlogsExtendUpdate) Exec(ctx context.Context) error {
	_, err := beu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (beu *BlogsExtendUpdate) ExecX(ctx context.Context) {
	if err := beu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (beu *BlogsExtendUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(blogsextend.Table, blogsextend.Columns, sqlgraph.NewFieldSpec(blogsextend.FieldID, field.TypeInt))
	if ps := beu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := beu.mutation.BlogID(); ok {
		_spec.SetField(blogsextend.FieldBlogID, field.TypeInt, value)
	}
	if value, ok := beu.mutation.AddedBlogID(); ok {
		_spec.AddField(blogsextend.FieldBlogID, field.TypeInt, value)
	}
	if value, ok := beu.mutation.AccountID(); ok {
		_spec.SetField(blogsextend.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := beu.mutation.AddedAccountID(); ok {
		_spec.AddField(blogsextend.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := beu.mutation.BrowseNum(); ok {
		_spec.SetField(blogsextend.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := beu.mutation.AddedBrowseNum(); ok {
		_spec.AddField(blogsextend.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := beu.mutation.BrowseAt(); ok {
		_spec.SetField(blogsextend.FieldBrowseAt, field.TypeInt64, value)
	}
	if value, ok := beu.mutation.AddedBrowseAt(); ok {
		_spec.AddField(blogsextend.FieldBrowseAt, field.TypeInt64, value)
	}
	if value, ok := beu.mutation.Collect(); ok {
		_spec.SetField(blogsextend.FieldCollect, field.TypeBool, value)
	}
	if value, ok := beu.mutation.CollectAt(); ok {
		_spec.SetField(blogsextend.FieldCollectAt, field.TypeInt64, value)
	}
	if value, ok := beu.mutation.AddedCollectAt(); ok {
		_spec.AddField(blogsextend.FieldCollectAt, field.TypeInt64, value)
	}
	if value, ok := beu.mutation.Love(); ok {
		_spec.SetField(blogsextend.FieldLove, field.TypeBool, value)
	}
	if value, ok := beu.mutation.LoveAt(); ok {
		_spec.SetField(blogsextend.FieldLoveAt, field.TypeInt64, value)
	}
	if value, ok := beu.mutation.AddedLoveAt(); ok {
		_spec.AddField(blogsextend.FieldLoveAt, field.TypeInt64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, beu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blogsextend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	beu.mutation.done = true
	return n, nil
}

// BlogsExtendUpdateOne is the builder for updating a single BlogsExtend entity.
type BlogsExtendUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BlogsExtendMutation
}

// SetBlogID sets the "blog_id" field.
func (beuo *BlogsExtendUpdateOne) SetBlogID(i int) *BlogsExtendUpdateOne {
	beuo.mutation.ResetBlogID()
	beuo.mutation.SetBlogID(i)
	return beuo
}

// SetNillableBlogID sets the "blog_id" field if the given value is not nil.
func (beuo *BlogsExtendUpdateOne) SetNillableBlogID(i *int) *BlogsExtendUpdateOne {
	if i != nil {
		beuo.SetBlogID(*i)
	}
	return beuo
}

// AddBlogID adds i to the "blog_id" field.
func (beuo *BlogsExtendUpdateOne) AddBlogID(i int) *BlogsExtendUpdateOne {
	beuo.mutation.AddBlogID(i)
	return beuo
}

// SetAccountID sets the "account_id" field.
func (beuo *BlogsExtendUpdateOne) SetAccountID(i int) *BlogsExtendUpdateOne {
	beuo.mutation.ResetAccountID()
	beuo.mutation.SetAccountID(i)
	return beuo
}

// SetNillableAccountID sets the "account_id" field if the given value is not nil.
func (beuo *BlogsExtendUpdateOne) SetNillableAccountID(i *int) *BlogsExtendUpdateOne {
	if i != nil {
		beuo.SetAccountID(*i)
	}
	return beuo
}

// AddAccountID adds i to the "account_id" field.
func (beuo *BlogsExtendUpdateOne) AddAccountID(i int) *BlogsExtendUpdateOne {
	beuo.mutation.AddAccountID(i)
	return beuo
}

// SetBrowseNum sets the "browse_num" field.
func (beuo *BlogsExtendUpdateOne) SetBrowseNum(i int) *BlogsExtendUpdateOne {
	beuo.mutation.ResetBrowseNum()
	beuo.mutation.SetBrowseNum(i)
	return beuo
}

// SetNillableBrowseNum sets the "browse_num" field if the given value is not nil.
func (beuo *BlogsExtendUpdateOne) SetNillableBrowseNum(i *int) *BlogsExtendUpdateOne {
	if i != nil {
		beuo.SetBrowseNum(*i)
	}
	return beuo
}

// AddBrowseNum adds i to the "browse_num" field.
func (beuo *BlogsExtendUpdateOne) AddBrowseNum(i int) *BlogsExtendUpdateOne {
	beuo.mutation.AddBrowseNum(i)
	return beuo
}

// SetBrowseAt sets the "browse_at" field.
func (beuo *BlogsExtendUpdateOne) SetBrowseAt(i int64) *BlogsExtendUpdateOne {
	beuo.mutation.ResetBrowseAt()
	beuo.mutation.SetBrowseAt(i)
	return beuo
}

// SetNillableBrowseAt sets the "browse_at" field if the given value is not nil.
func (beuo *BlogsExtendUpdateOne) SetNillableBrowseAt(i *int64) *BlogsExtendUpdateOne {
	if i != nil {
		beuo.SetBrowseAt(*i)
	}
	return beuo
}

// AddBrowseAt adds i to the "browse_at" field.
func (beuo *BlogsExtendUpdateOne) AddBrowseAt(i int64) *BlogsExtendUpdateOne {
	beuo.mutation.AddBrowseAt(i)
	return beuo
}

// SetCollect sets the "collect" field.
func (beuo *BlogsExtendUpdateOne) SetCollect(b bool) *BlogsExtendUpdateOne {
	beuo.mutation.SetCollect(b)
	return beuo
}

// SetNillableCollect sets the "collect" field if the given value is not nil.
func (beuo *BlogsExtendUpdateOne) SetNillableCollect(b *bool) *BlogsExtendUpdateOne {
	if b != nil {
		beuo.SetCollect(*b)
	}
	return beuo
}

// SetCollectAt sets the "collect_at" field.
func (beuo *BlogsExtendUpdateOne) SetCollectAt(i int64) *BlogsExtendUpdateOne {
	beuo.mutation.ResetCollectAt()
	beuo.mutation.SetCollectAt(i)
	return beuo
}

// SetNillableCollectAt sets the "collect_at" field if the given value is not nil.
func (beuo *BlogsExtendUpdateOne) SetNillableCollectAt(i *int64) *BlogsExtendUpdateOne {
	if i != nil {
		beuo.SetCollectAt(*i)
	}
	return beuo
}

// AddCollectAt adds i to the "collect_at" field.
func (beuo *BlogsExtendUpdateOne) AddCollectAt(i int64) *BlogsExtendUpdateOne {
	beuo.mutation.AddCollectAt(i)
	return beuo
}

// SetLove sets the "love" field.
func (beuo *BlogsExtendUpdateOne) SetLove(b bool) *BlogsExtendUpdateOne {
	beuo.mutation.SetLove(b)
	return beuo
}

// SetNillableLove sets the "love" field if the given value is not nil.
func (beuo *BlogsExtendUpdateOne) SetNillableLove(b *bool) *BlogsExtendUpdateOne {
	if b != nil {
		beuo.SetLove(*b)
	}
	return beuo
}

// SetLoveAt sets the "love_at" field.
func (beuo *BlogsExtendUpdateOne) SetLoveAt(i int64) *BlogsExtendUpdateOne {
	beuo.mutation.ResetLoveAt()
	beuo.mutation.SetLoveAt(i)
	return beuo
}

// SetNillableLoveAt sets the "love_at" field if the given value is not nil.
func (beuo *BlogsExtendUpdateOne) SetNillableLoveAt(i *int64) *BlogsExtendUpdateOne {
	if i != nil {
		beuo.SetLoveAt(*i)
	}
	return beuo
}

// AddLoveAt adds i to the "love_at" field.
func (beuo *BlogsExtendUpdateOne) AddLoveAt(i int64) *BlogsExtendUpdateOne {
	beuo.mutation.AddLoveAt(i)
	return beuo
}

// Mutation returns the BlogsExtendMutation object of the builder.
func (beuo *BlogsExtendUpdateOne) Mutation() *BlogsExtendMutation {
	return beuo.mutation
}

// Where appends a list predicates to the BlogsExtendUpdate builder.
func (beuo *BlogsExtendUpdateOne) Where(ps ...predicate.BlogsExtend) *BlogsExtendUpdateOne {
	beuo.mutation.Where(ps...)
	return beuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (beuo *BlogsExtendUpdateOne) Select(field string, fields ...string) *BlogsExtendUpdateOne {
	beuo.fields = append([]string{field}, fields...)
	return beuo
}

// Save executes the query and returns the updated BlogsExtend entity.
func (beuo *BlogsExtendUpdateOne) Save(ctx context.Context) (*BlogsExtend, error) {
	return withHooks(ctx, beuo.sqlSave, beuo.mutation, beuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (beuo *BlogsExtendUpdateOne) SaveX(ctx context.Context) *BlogsExtend {
	node, err := beuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (beuo *BlogsExtendUpdateOne) Exec(ctx context.Context) error {
	_, err := beuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (beuo *BlogsExtendUpdateOne) ExecX(ctx context.Context) {
	if err := beuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (beuo *BlogsExtendUpdateOne) sqlSave(ctx context.Context) (_node *BlogsExtend, err error) {
	_spec := sqlgraph.NewUpdateSpec(blogsextend.Table, blogsextend.Columns, sqlgraph.NewFieldSpec(blogsextend.FieldID, field.TypeInt))
	id, ok := beuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "BlogsExtend.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := beuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blogsextend.FieldID)
		for _, f := range fields {
			if !blogsextend.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != blogsextend.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := beuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := beuo.mutation.BlogID(); ok {
		_spec.SetField(blogsextend.FieldBlogID, field.TypeInt, value)
	}
	if value, ok := beuo.mutation.AddedBlogID(); ok {
		_spec.AddField(blogsextend.FieldBlogID, field.TypeInt, value)
	}
	if value, ok := beuo.mutation.AccountID(); ok {
		_spec.SetField(blogsextend.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := beuo.mutation.AddedAccountID(); ok {
		_spec.AddField(blogsextend.FieldAccountID, field.TypeInt, value)
	}
	if value, ok := beuo.mutation.BrowseNum(); ok {
		_spec.SetField(blogsextend.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := beuo.mutation.AddedBrowseNum(); ok {
		_spec.AddField(blogsextend.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := beuo.mutation.BrowseAt(); ok {
		_spec.SetField(blogsextend.FieldBrowseAt, field.TypeInt64, value)
	}
	if value, ok := beuo.mutation.AddedBrowseAt(); ok {
		_spec.AddField(blogsextend.FieldBrowseAt, field.TypeInt64, value)
	}
	if value, ok := beuo.mutation.Collect(); ok {
		_spec.SetField(blogsextend.FieldCollect, field.TypeBool, value)
	}
	if value, ok := beuo.mutation.CollectAt(); ok {
		_spec.SetField(blogsextend.FieldCollectAt, field.TypeInt64, value)
	}
	if value, ok := beuo.mutation.AddedCollectAt(); ok {
		_spec.AddField(blogsextend.FieldCollectAt, field.TypeInt64, value)
	}
	if value, ok := beuo.mutation.Love(); ok {
		_spec.SetField(blogsextend.FieldLove, field.TypeBool, value)
	}
	if value, ok := beuo.mutation.LoveAt(); ok {
		_spec.SetField(blogsextend.FieldLoveAt, field.TypeInt64, value)
	}
	if value, ok := beuo.mutation.AddedLoveAt(); ok {
		_spec.AddField(blogsextend.FieldLoveAt, field.TypeInt64, value)
	}
	_node = &BlogsExtend{config: beuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, beuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{blogsextend.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	beuo.mutation.done = true
	return _node, nil
}
