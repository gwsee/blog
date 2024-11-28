// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/files"
	"blog/internal/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FilesUpdate is the builder for updating Files entities.
type FilesUpdate struct {
	config
	hooks    []Hook
	mutation *FilesMutation
}

// Where appends a list predicates to the FilesUpdate builder.
func (fu *FilesUpdate) Where(ps ...predicate.Files) *FilesUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetUpdatedAt sets the "updated_at" field.
func (fu *FilesUpdate) SetUpdatedAt(i int64) *FilesUpdate {
	fu.mutation.ResetUpdatedAt()
	fu.mutation.SetUpdatedAt(i)
	return fu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (fu *FilesUpdate) AddUpdatedAt(i int64) *FilesUpdate {
	fu.mutation.AddUpdatedAt(i)
	return fu
}

// SetUpdatedBy sets the "updated_by" field.
func (fu *FilesUpdate) SetUpdatedBy(i int64) *FilesUpdate {
	fu.mutation.ResetUpdatedBy()
	fu.mutation.SetUpdatedBy(i)
	return fu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fu *FilesUpdate) SetNillableUpdatedBy(i *int64) *FilesUpdate {
	if i != nil {
		fu.SetUpdatedBy(*i)
	}
	return fu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fu *FilesUpdate) AddUpdatedBy(i int64) *FilesUpdate {
	fu.mutation.AddUpdatedBy(i)
	return fu
}

// SetDeletedAt sets the "deleted_at" field.
func (fu *FilesUpdate) SetDeletedAt(i int64) *FilesUpdate {
	fu.mutation.ResetDeletedAt()
	fu.mutation.SetDeletedAt(i)
	return fu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fu *FilesUpdate) SetNillableDeletedAt(i *int64) *FilesUpdate {
	if i != nil {
		fu.SetDeletedAt(*i)
	}
	return fu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (fu *FilesUpdate) AddDeletedAt(i int64) *FilesUpdate {
	fu.mutation.AddDeletedAt(i)
	return fu
}

// SetDeletedBy sets the "deleted_by" field.
func (fu *FilesUpdate) SetDeletedBy(i int64) *FilesUpdate {
	fu.mutation.ResetDeletedBy()
	fu.mutation.SetDeletedBy(i)
	return fu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fu *FilesUpdate) SetNillableDeletedBy(i *int64) *FilesUpdate {
	if i != nil {
		fu.SetDeletedBy(*i)
	}
	return fu
}

// AddDeletedBy adds i to the "deleted_by" field.
func (fu *FilesUpdate) AddDeletedBy(i int64) *FilesUpdate {
	fu.mutation.AddDeletedBy(i)
	return fu
}

// SetType sets the "type" field.
func (fu *FilesUpdate) SetType(s string) *FilesUpdate {
	fu.mutation.SetType(s)
	return fu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (fu *FilesUpdate) SetNillableType(s *string) *FilesUpdate {
	if s != nil {
		fu.SetType(*s)
	}
	return fu
}

// SetSize sets the "size" field.
func (fu *FilesUpdate) SetSize(i int64) *FilesUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fu *FilesUpdate) SetNillableSize(i *int64) *FilesUpdate {
	if i != nil {
		fu.SetSize(*i)
	}
	return fu
}

// AddSize adds i to the "size" field.
func (fu *FilesUpdate) AddSize(i int64) *FilesUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// SetName sets the "name" field.
func (fu *FilesUpdate) SetName(s string) *FilesUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fu *FilesUpdate) SetNillableName(s *string) *FilesUpdate {
	if s != nil {
		fu.SetName(*s)
	}
	return fu
}

// SetPath sets the "path" field.
func (fu *FilesUpdate) SetPath(s string) *FilesUpdate {
	fu.mutation.SetPath(s)
	return fu
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (fu *FilesUpdate) SetNillablePath(s *string) *FilesUpdate {
	if s != nil {
		fu.SetPath(*s)
	}
	return fu
}

// Mutation returns the FilesMutation object of the builder.
func (fu *FilesUpdate) Mutation() *FilesMutation {
	return fu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FilesUpdate) Save(ctx context.Context) (int, error) {
	if err := fu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FilesUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FilesUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FilesUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fu *FilesUpdate) defaults() error {
	if _, ok := fu.mutation.UpdatedAt(); !ok {
		if files.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized files.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := files.UpdateDefaultUpdatedAt()
		fu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (fu *FilesUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(files.Table, files.Columns, sqlgraph.NewFieldSpec(files.FieldID, field.TypeString))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.UpdatedAt(); ok {
		_spec.SetField(files.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(files.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.UpdatedBy(); ok {
		_spec.SetField(files.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(files.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.DeletedAt(); ok {
		_spec.SetField(files.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(files.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.DeletedBy(); ok {
		_spec.SetField(files.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedDeletedBy(); ok {
		_spec.AddField(files.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.GetType(); ok {
		_spec.SetField(files.FieldType, field.TypeString, value)
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.SetField(files.FieldSize, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.AddField(files.FieldSize, field.TypeInt64, value)
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(files.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.Path(); ok {
		_spec.SetField(files.FieldPath, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{files.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FilesUpdateOne is the builder for updating a single Files entity.
type FilesUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FilesMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (fuo *FilesUpdateOne) SetUpdatedAt(i int64) *FilesUpdateOne {
	fuo.mutation.ResetUpdatedAt()
	fuo.mutation.SetUpdatedAt(i)
	return fuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (fuo *FilesUpdateOne) AddUpdatedAt(i int64) *FilesUpdateOne {
	fuo.mutation.AddUpdatedAt(i)
	return fuo
}

// SetUpdatedBy sets the "updated_by" field.
func (fuo *FilesUpdateOne) SetUpdatedBy(i int64) *FilesUpdateOne {
	fuo.mutation.ResetUpdatedBy()
	fuo.mutation.SetUpdatedBy(i)
	return fuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableUpdatedBy(i *int64) *FilesUpdateOne {
	if i != nil {
		fuo.SetUpdatedBy(*i)
	}
	return fuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (fuo *FilesUpdateOne) AddUpdatedBy(i int64) *FilesUpdateOne {
	fuo.mutation.AddUpdatedBy(i)
	return fuo
}

// SetDeletedAt sets the "deleted_at" field.
func (fuo *FilesUpdateOne) SetDeletedAt(i int64) *FilesUpdateOne {
	fuo.mutation.ResetDeletedAt()
	fuo.mutation.SetDeletedAt(i)
	return fuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableDeletedAt(i *int64) *FilesUpdateOne {
	if i != nil {
		fuo.SetDeletedAt(*i)
	}
	return fuo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (fuo *FilesUpdateOne) AddDeletedAt(i int64) *FilesUpdateOne {
	fuo.mutation.AddDeletedAt(i)
	return fuo
}

// SetDeletedBy sets the "deleted_by" field.
func (fuo *FilesUpdateOne) SetDeletedBy(i int64) *FilesUpdateOne {
	fuo.mutation.ResetDeletedBy()
	fuo.mutation.SetDeletedBy(i)
	return fuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableDeletedBy(i *int64) *FilesUpdateOne {
	if i != nil {
		fuo.SetDeletedBy(*i)
	}
	return fuo
}

// AddDeletedBy adds i to the "deleted_by" field.
func (fuo *FilesUpdateOne) AddDeletedBy(i int64) *FilesUpdateOne {
	fuo.mutation.AddDeletedBy(i)
	return fuo
}

// SetType sets the "type" field.
func (fuo *FilesUpdateOne) SetType(s string) *FilesUpdateOne {
	fuo.mutation.SetType(s)
	return fuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableType(s *string) *FilesUpdateOne {
	if s != nil {
		fuo.SetType(*s)
	}
	return fuo
}

// SetSize sets the "size" field.
func (fuo *FilesUpdateOne) SetSize(i int64) *FilesUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableSize(i *int64) *FilesUpdateOne {
	if i != nil {
		fuo.SetSize(*i)
	}
	return fuo
}

// AddSize adds i to the "size" field.
func (fuo *FilesUpdateOne) AddSize(i int64) *FilesUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// SetName sets the "name" field.
func (fuo *FilesUpdateOne) SetName(s string) *FilesUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillableName(s *string) *FilesUpdateOne {
	if s != nil {
		fuo.SetName(*s)
	}
	return fuo
}

// SetPath sets the "path" field.
func (fuo *FilesUpdateOne) SetPath(s string) *FilesUpdateOne {
	fuo.mutation.SetPath(s)
	return fuo
}

// SetNillablePath sets the "path" field if the given value is not nil.
func (fuo *FilesUpdateOne) SetNillablePath(s *string) *FilesUpdateOne {
	if s != nil {
		fuo.SetPath(*s)
	}
	return fuo
}

// Mutation returns the FilesMutation object of the builder.
func (fuo *FilesUpdateOne) Mutation() *FilesMutation {
	return fuo.mutation
}

// Where appends a list predicates to the FilesUpdate builder.
func (fuo *FilesUpdateOne) Where(ps ...predicate.Files) *FilesUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FilesUpdateOne) Select(field string, fields ...string) *FilesUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated Files entity.
func (fuo *FilesUpdateOne) Save(ctx context.Context) (*Files, error) {
	if err := fuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FilesUpdateOne) SaveX(ctx context.Context) *Files {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FilesUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FilesUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fuo *FilesUpdateOne) defaults() error {
	if _, ok := fuo.mutation.UpdatedAt(); !ok {
		if files.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized files.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := files.UpdateDefaultUpdatedAt()
		fuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (fuo *FilesUpdateOne) sqlSave(ctx context.Context) (_node *Files, err error) {
	_spec := sqlgraph.NewUpdateSpec(files.Table, files.Columns, sqlgraph.NewFieldSpec(files.FieldID, field.TypeString))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Files.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, files.FieldID)
		for _, f := range fields {
			if !files.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != files.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.UpdatedAt(); ok {
		_spec.SetField(files.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(files.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.UpdatedBy(); ok {
		_spec.SetField(files.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(files.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.DeletedAt(); ok {
		_spec.SetField(files.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(files.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.DeletedBy(); ok {
		_spec.SetField(files.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(files.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.GetType(); ok {
		_spec.SetField(files.FieldType, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Size(); ok {
		_spec.SetField(files.FieldSize, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.AddField(files.FieldSize, field.TypeInt64, value)
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(files.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.Path(); ok {
		_spec.SetField(files.FieldPath, field.TypeString, value)
	}
	_node = &Files{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{files.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}