// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/predicate"
	"blog/internal/ent/travel"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// TravelUpdate is the builder for updating Travel entities.
type TravelUpdate struct {
	config
	hooks    []Hook
	mutation *TravelMutation
}

// Where appends a list predicates to the TravelUpdate builder.
func (tu *TravelUpdate) Where(ps ...predicate.Travel) *TravelUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdatedAt sets the "updated_at" field.
func (tu *TravelUpdate) SetUpdatedAt(i int64) *TravelUpdate {
	tu.mutation.ResetUpdatedAt()
	tu.mutation.SetUpdatedAt(i)
	return tu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tu *TravelUpdate) AddUpdatedAt(i int64) *TravelUpdate {
	tu.mutation.AddUpdatedAt(i)
	return tu
}

// SetUpdatedBy sets the "updated_by" field.
func (tu *TravelUpdate) SetUpdatedBy(i int64) *TravelUpdate {
	tu.mutation.ResetUpdatedBy()
	tu.mutation.SetUpdatedBy(i)
	return tu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableUpdatedBy(i *int64) *TravelUpdate {
	if i != nil {
		tu.SetUpdatedBy(*i)
	}
	return tu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (tu *TravelUpdate) AddUpdatedBy(i int64) *TravelUpdate {
	tu.mutation.AddUpdatedBy(i)
	return tu
}

// SetDeletedAt sets the "deleted_at" field.
func (tu *TravelUpdate) SetDeletedAt(i int64) *TravelUpdate {
	tu.mutation.ResetDeletedAt()
	tu.mutation.SetDeletedAt(i)
	return tu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableDeletedAt(i *int64) *TravelUpdate {
	if i != nil {
		tu.SetDeletedAt(*i)
	}
	return tu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (tu *TravelUpdate) AddDeletedAt(i int64) *TravelUpdate {
	tu.mutation.AddDeletedAt(i)
	return tu
}

// SetDeletedBy sets the "deleted_by" field.
func (tu *TravelUpdate) SetDeletedBy(i int64) *TravelUpdate {
	tu.mutation.ResetDeletedBy()
	tu.mutation.SetDeletedBy(i)
	return tu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableDeletedBy(i *int64) *TravelUpdate {
	if i != nil {
		tu.SetDeletedBy(*i)
	}
	return tu
}

// AddDeletedBy adds i to the "deleted_by" field.
func (tu *TravelUpdate) AddDeletedBy(i int64) *TravelUpdate {
	tu.mutation.AddDeletedBy(i)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TravelUpdate) SetTitle(s string) *TravelUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableTitle(s *string) *TravelUpdate {
	if s != nil {
		tu.SetTitle(*s)
	}
	return tu
}

// SetDescription sets the "description" field.
func (tu *TravelUpdate) SetDescription(s string) *TravelUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableDescription(s *string) *TravelUpdate {
	if s != nil {
		tu.SetDescription(*s)
	}
	return tu
}

// SetVideo sets the "video" field.
func (tu *TravelUpdate) SetVideo(s string) *TravelUpdate {
	tu.mutation.SetVideo(s)
	return tu
}

// SetNillableVideo sets the "video" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableVideo(s *string) *TravelUpdate {
	if s != nil {
		tu.SetVideo(*s)
	}
	return tu
}

// SetPhotos sets the "photos" field.
func (tu *TravelUpdate) SetPhotos(s []string) *TravelUpdate {
	tu.mutation.SetPhotos(s)
	return tu
}

// AppendPhotos appends s to the "photos" field.
func (tu *TravelUpdate) AppendPhotos(s []string) *TravelUpdate {
	tu.mutation.AppendPhotos(s)
	return tu
}

// SetBrowseNum sets the "browse_num" field.
func (tu *TravelUpdate) SetBrowseNum(i int) *TravelUpdate {
	tu.mutation.ResetBrowseNum()
	tu.mutation.SetBrowseNum(i)
	return tu
}

// SetNillableBrowseNum sets the "browse_num" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableBrowseNum(i *int) *TravelUpdate {
	if i != nil {
		tu.SetBrowseNum(*i)
	}
	return tu
}

// AddBrowseNum adds i to the "browse_num" field.
func (tu *TravelUpdate) AddBrowseNum(i int) *TravelUpdate {
	tu.mutation.AddBrowseNum(i)
	return tu
}

// SetThumbNum sets the "thumb_num" field.
func (tu *TravelUpdate) SetThumbNum(i int) *TravelUpdate {
	tu.mutation.ResetThumbNum()
	tu.mutation.SetThumbNum(i)
	return tu
}

// SetNillableThumbNum sets the "thumb_num" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableThumbNum(i *int) *TravelUpdate {
	if i != nil {
		tu.SetThumbNum(*i)
	}
	return tu
}

// AddThumbNum adds i to the "thumb_num" field.
func (tu *TravelUpdate) AddThumbNum(i int) *TravelUpdate {
	tu.mutation.AddThumbNum(i)
	return tu
}

// SetCollectNum sets the "collect_num" field.
func (tu *TravelUpdate) SetCollectNum(i int) *TravelUpdate {
	tu.mutation.ResetCollectNum()
	tu.mutation.SetCollectNum(i)
	return tu
}

// SetNillableCollectNum sets the "collect_num" field if the given value is not nil.
func (tu *TravelUpdate) SetNillableCollectNum(i *int) *TravelUpdate {
	if i != nil {
		tu.SetCollectNum(*i)
	}
	return tu
}

// AddCollectNum adds i to the "collect_num" field.
func (tu *TravelUpdate) AddCollectNum(i int) *TravelUpdate {
	tu.mutation.AddCollectNum(i)
	return tu
}

// Mutation returns the TravelMutation object of the builder.
func (tu *TravelUpdate) Mutation() *TravelMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TravelUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TravelUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TravelUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TravelUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TravelUpdate) defaults() error {
	if _, ok := tu.mutation.UpdatedAt(); !ok {
		if travel.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized travel.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := travel.UpdateDefaultUpdatedAt()
		tu.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TravelUpdate) check() error {
	if v, ok := tu.mutation.Description(); ok {
		if err := travel.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Travel.description": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Video(); ok {
		if err := travel.VideoValidator(v); err != nil {
			return &ValidationError{Name: "video", err: fmt.Errorf(`ent: validator failed for field "Travel.video": %w`, err)}
		}
	}
	return nil
}

func (tu *TravelUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(travel.Table, travel.Columns, sqlgraph.NewFieldSpec(travel.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(travel.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(travel.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.UpdatedBy(); ok {
		_spec.SetField(travel.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(travel.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.DeletedAt(); ok {
		_spec.SetField(travel.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(travel.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.DeletedBy(); ok {
		_spec.SetField(travel.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.AddedDeletedBy(); ok {
		_spec.AddField(travel.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(travel.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(travel.FieldDescription, field.TypeString, value)
	}
	if value, ok := tu.mutation.Video(); ok {
		_spec.SetField(travel.FieldVideo, field.TypeString, value)
	}
	if value, ok := tu.mutation.Photos(); ok {
		_spec.SetField(travel.FieldPhotos, field.TypeJSON, value)
	}
	if value, ok := tu.mutation.AppendedPhotos(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, travel.FieldPhotos, value)
		})
	}
	if value, ok := tu.mutation.BrowseNum(); ok {
		_spec.SetField(travel.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedBrowseNum(); ok {
		_spec.AddField(travel.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := tu.mutation.ThumbNum(); ok {
		_spec.SetField(travel.FieldThumbNum, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedThumbNum(); ok {
		_spec.AddField(travel.FieldThumbNum, field.TypeInt, value)
	}
	if value, ok := tu.mutation.CollectNum(); ok {
		_spec.SetField(travel.FieldCollectNum, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedCollectNum(); ok {
		_spec.AddField(travel.FieldCollectNum, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{travel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TravelUpdateOne is the builder for updating a single Travel entity.
type TravelUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TravelMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (tuo *TravelUpdateOne) SetUpdatedAt(i int64) *TravelUpdateOne {
	tuo.mutation.ResetUpdatedAt()
	tuo.mutation.SetUpdatedAt(i)
	return tuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (tuo *TravelUpdateOne) AddUpdatedAt(i int64) *TravelUpdateOne {
	tuo.mutation.AddUpdatedAt(i)
	return tuo
}

// SetUpdatedBy sets the "updated_by" field.
func (tuo *TravelUpdateOne) SetUpdatedBy(i int64) *TravelUpdateOne {
	tuo.mutation.ResetUpdatedBy()
	tuo.mutation.SetUpdatedBy(i)
	return tuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableUpdatedBy(i *int64) *TravelUpdateOne {
	if i != nil {
		tuo.SetUpdatedBy(*i)
	}
	return tuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (tuo *TravelUpdateOne) AddUpdatedBy(i int64) *TravelUpdateOne {
	tuo.mutation.AddUpdatedBy(i)
	return tuo
}

// SetDeletedAt sets the "deleted_at" field.
func (tuo *TravelUpdateOne) SetDeletedAt(i int64) *TravelUpdateOne {
	tuo.mutation.ResetDeletedAt()
	tuo.mutation.SetDeletedAt(i)
	return tuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableDeletedAt(i *int64) *TravelUpdateOne {
	if i != nil {
		tuo.SetDeletedAt(*i)
	}
	return tuo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (tuo *TravelUpdateOne) AddDeletedAt(i int64) *TravelUpdateOne {
	tuo.mutation.AddDeletedAt(i)
	return tuo
}

// SetDeletedBy sets the "deleted_by" field.
func (tuo *TravelUpdateOne) SetDeletedBy(i int64) *TravelUpdateOne {
	tuo.mutation.ResetDeletedBy()
	tuo.mutation.SetDeletedBy(i)
	return tuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableDeletedBy(i *int64) *TravelUpdateOne {
	if i != nil {
		tuo.SetDeletedBy(*i)
	}
	return tuo
}

// AddDeletedBy adds i to the "deleted_by" field.
func (tuo *TravelUpdateOne) AddDeletedBy(i int64) *TravelUpdateOne {
	tuo.mutation.AddDeletedBy(i)
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TravelUpdateOne) SetTitle(s string) *TravelUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableTitle(s *string) *TravelUpdateOne {
	if s != nil {
		tuo.SetTitle(*s)
	}
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TravelUpdateOne) SetDescription(s string) *TravelUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableDescription(s *string) *TravelUpdateOne {
	if s != nil {
		tuo.SetDescription(*s)
	}
	return tuo
}

// SetVideo sets the "video" field.
func (tuo *TravelUpdateOne) SetVideo(s string) *TravelUpdateOne {
	tuo.mutation.SetVideo(s)
	return tuo
}

// SetNillableVideo sets the "video" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableVideo(s *string) *TravelUpdateOne {
	if s != nil {
		tuo.SetVideo(*s)
	}
	return tuo
}

// SetPhotos sets the "photos" field.
func (tuo *TravelUpdateOne) SetPhotos(s []string) *TravelUpdateOne {
	tuo.mutation.SetPhotos(s)
	return tuo
}

// AppendPhotos appends s to the "photos" field.
func (tuo *TravelUpdateOne) AppendPhotos(s []string) *TravelUpdateOne {
	tuo.mutation.AppendPhotos(s)
	return tuo
}

// SetBrowseNum sets the "browse_num" field.
func (tuo *TravelUpdateOne) SetBrowseNum(i int) *TravelUpdateOne {
	tuo.mutation.ResetBrowseNum()
	tuo.mutation.SetBrowseNum(i)
	return tuo
}

// SetNillableBrowseNum sets the "browse_num" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableBrowseNum(i *int) *TravelUpdateOne {
	if i != nil {
		tuo.SetBrowseNum(*i)
	}
	return tuo
}

// AddBrowseNum adds i to the "browse_num" field.
func (tuo *TravelUpdateOne) AddBrowseNum(i int) *TravelUpdateOne {
	tuo.mutation.AddBrowseNum(i)
	return tuo
}

// SetThumbNum sets the "thumb_num" field.
func (tuo *TravelUpdateOne) SetThumbNum(i int) *TravelUpdateOne {
	tuo.mutation.ResetThumbNum()
	tuo.mutation.SetThumbNum(i)
	return tuo
}

// SetNillableThumbNum sets the "thumb_num" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableThumbNum(i *int) *TravelUpdateOne {
	if i != nil {
		tuo.SetThumbNum(*i)
	}
	return tuo
}

// AddThumbNum adds i to the "thumb_num" field.
func (tuo *TravelUpdateOne) AddThumbNum(i int) *TravelUpdateOne {
	tuo.mutation.AddThumbNum(i)
	return tuo
}

// SetCollectNum sets the "collect_num" field.
func (tuo *TravelUpdateOne) SetCollectNum(i int) *TravelUpdateOne {
	tuo.mutation.ResetCollectNum()
	tuo.mutation.SetCollectNum(i)
	return tuo
}

// SetNillableCollectNum sets the "collect_num" field if the given value is not nil.
func (tuo *TravelUpdateOne) SetNillableCollectNum(i *int) *TravelUpdateOne {
	if i != nil {
		tuo.SetCollectNum(*i)
	}
	return tuo
}

// AddCollectNum adds i to the "collect_num" field.
func (tuo *TravelUpdateOne) AddCollectNum(i int) *TravelUpdateOne {
	tuo.mutation.AddCollectNum(i)
	return tuo
}

// Mutation returns the TravelMutation object of the builder.
func (tuo *TravelUpdateOne) Mutation() *TravelMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TravelUpdate builder.
func (tuo *TravelUpdateOne) Where(ps ...predicate.Travel) *TravelUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TravelUpdateOne) Select(field string, fields ...string) *TravelUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Travel entity.
func (tuo *TravelUpdateOne) Save(ctx context.Context) (*Travel, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TravelUpdateOne) SaveX(ctx context.Context) *Travel {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TravelUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TravelUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TravelUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdatedAt(); !ok {
		if travel.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized travel.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := travel.UpdateDefaultUpdatedAt()
		tuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TravelUpdateOne) check() error {
	if v, ok := tuo.mutation.Description(); ok {
		if err := travel.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Travel.description": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Video(); ok {
		if err := travel.VideoValidator(v); err != nil {
			return &ValidationError{Name: "video", err: fmt.Errorf(`ent: validator failed for field "Travel.video": %w`, err)}
		}
	}
	return nil
}

func (tuo *TravelUpdateOne) sqlSave(ctx context.Context) (_node *Travel, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(travel.Table, travel.Columns, sqlgraph.NewFieldSpec(travel.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Travel.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, travel.FieldID)
		for _, f := range fields {
			if !travel.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != travel.FieldID {
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
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(travel.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(travel.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.UpdatedBy(); ok {
		_spec.SetField(travel.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(travel.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.DeletedAt(); ok {
		_spec.SetField(travel.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(travel.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.DeletedBy(); ok {
		_spec.SetField(travel.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(travel.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(travel.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(travel.FieldDescription, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Video(); ok {
		_spec.SetField(travel.FieldVideo, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Photos(); ok {
		_spec.SetField(travel.FieldPhotos, field.TypeJSON, value)
	}
	if value, ok := tuo.mutation.AppendedPhotos(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, travel.FieldPhotos, value)
		})
	}
	if value, ok := tuo.mutation.BrowseNum(); ok {
		_spec.SetField(travel.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedBrowseNum(); ok {
		_spec.AddField(travel.FieldBrowseNum, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.ThumbNum(); ok {
		_spec.SetField(travel.FieldThumbNum, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedThumbNum(); ok {
		_spec.AddField(travel.FieldThumbNum, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.CollectNum(); ok {
		_spec.SetField(travel.FieldCollectNum, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedCollectNum(); ok {
		_spec.AddField(travel.FieldCollectNum, field.TypeInt, value)
	}
	_node = &Travel{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{travel.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}