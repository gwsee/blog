// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/palacestodo"
	"blog/internal/ent/palacestododone"
	"blog/internal/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PalacesTodoDoneUpdate is the builder for updating PalacesTodoDone entities.
type PalacesTodoDoneUpdate struct {
	config
	hooks    []Hook
	mutation *PalacesTodoDoneMutation
}

// Where appends a list predicates to the PalacesTodoDoneUpdate builder.
func (ptdu *PalacesTodoDoneUpdate) Where(ps ...predicate.PalacesTodoDone) *PalacesTodoDoneUpdate {
	ptdu.mutation.Where(ps...)
	return ptdu
}

// SetUpdatedAt sets the "updated_at" field.
func (ptdu *PalacesTodoDoneUpdate) SetUpdatedAt(i int64) *PalacesTodoDoneUpdate {
	ptdu.mutation.ResetUpdatedAt()
	ptdu.mutation.SetUpdatedAt(i)
	return ptdu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ptdu *PalacesTodoDoneUpdate) AddUpdatedAt(i int64) *PalacesTodoDoneUpdate {
	ptdu.mutation.AddUpdatedAt(i)
	return ptdu
}

// SetUpdatedBy sets the "updated_by" field.
func (ptdu *PalacesTodoDoneUpdate) SetUpdatedBy(i int64) *PalacesTodoDoneUpdate {
	ptdu.mutation.ResetUpdatedBy()
	ptdu.mutation.SetUpdatedBy(i)
	return ptdu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ptdu *PalacesTodoDoneUpdate) SetNillableUpdatedBy(i *int64) *PalacesTodoDoneUpdate {
	if i != nil {
		ptdu.SetUpdatedBy(*i)
	}
	return ptdu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ptdu *PalacesTodoDoneUpdate) AddUpdatedBy(i int64) *PalacesTodoDoneUpdate {
	ptdu.mutation.AddUpdatedBy(i)
	return ptdu
}

// SetDeletedAt sets the "deleted_at" field.
func (ptdu *PalacesTodoDoneUpdate) SetDeletedAt(i int64) *PalacesTodoDoneUpdate {
	ptdu.mutation.ResetDeletedAt()
	ptdu.mutation.SetDeletedAt(i)
	return ptdu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ptdu *PalacesTodoDoneUpdate) SetNillableDeletedAt(i *int64) *PalacesTodoDoneUpdate {
	if i != nil {
		ptdu.SetDeletedAt(*i)
	}
	return ptdu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ptdu *PalacesTodoDoneUpdate) AddDeletedAt(i int64) *PalacesTodoDoneUpdate {
	ptdu.mutation.AddDeletedAt(i)
	return ptdu
}

// SetDeletedBy sets the "deleted_by" field.
func (ptdu *PalacesTodoDoneUpdate) SetDeletedBy(i int64) *PalacesTodoDoneUpdate {
	ptdu.mutation.ResetDeletedBy()
	ptdu.mutation.SetDeletedBy(i)
	return ptdu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ptdu *PalacesTodoDoneUpdate) SetNillableDeletedBy(i *int64) *PalacesTodoDoneUpdate {
	if i != nil {
		ptdu.SetDeletedBy(*i)
	}
	return ptdu
}

// AddDeletedBy adds i to the "deleted_by" field.
func (ptdu *PalacesTodoDoneUpdate) AddDeletedBy(i int64) *PalacesTodoDoneUpdate {
	ptdu.mutation.AddDeletedBy(i)
	return ptdu
}

// SetTodoID sets the "todo_id" field.
func (ptdu *PalacesTodoDoneUpdate) SetTodoID(i int) *PalacesTodoDoneUpdate {
	ptdu.mutation.SetTodoID(i)
	return ptdu
}

// SetNillableTodoID sets the "todo_id" field if the given value is not nil.
func (ptdu *PalacesTodoDoneUpdate) SetNillableTodoID(i *int) *PalacesTodoDoneUpdate {
	if i != nil {
		ptdu.SetTodoID(*i)
	}
	return ptdu
}

// ClearTodoID clears the value of the "todo_id" field.
func (ptdu *PalacesTodoDoneUpdate) ClearTodoID() *PalacesTodoDoneUpdate {
	ptdu.mutation.ClearTodoID()
	return ptdu
}

// SetOwnerID sets the "owner" edge to the PalacesTodo entity by ID.
func (ptdu *PalacesTodoDoneUpdate) SetOwnerID(id int) *PalacesTodoDoneUpdate {
	ptdu.mutation.SetOwnerID(id)
	return ptdu
}

// SetNillableOwnerID sets the "owner" edge to the PalacesTodo entity by ID if the given value is not nil.
func (ptdu *PalacesTodoDoneUpdate) SetNillableOwnerID(id *int) *PalacesTodoDoneUpdate {
	if id != nil {
		ptdu = ptdu.SetOwnerID(*id)
	}
	return ptdu
}

// SetOwner sets the "owner" edge to the PalacesTodo entity.
func (ptdu *PalacesTodoDoneUpdate) SetOwner(p *PalacesTodo) *PalacesTodoDoneUpdate {
	return ptdu.SetOwnerID(p.ID)
}

// Mutation returns the PalacesTodoDoneMutation object of the builder.
func (ptdu *PalacesTodoDoneUpdate) Mutation() *PalacesTodoDoneMutation {
	return ptdu.mutation
}

// ClearOwner clears the "owner" edge to the PalacesTodo entity.
func (ptdu *PalacesTodoDoneUpdate) ClearOwner() *PalacesTodoDoneUpdate {
	ptdu.mutation.ClearOwner()
	return ptdu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptdu *PalacesTodoDoneUpdate) Save(ctx context.Context) (int, error) {
	if err := ptdu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ptdu.sqlSave, ptdu.mutation, ptdu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptdu *PalacesTodoDoneUpdate) SaveX(ctx context.Context) int {
	affected, err := ptdu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptdu *PalacesTodoDoneUpdate) Exec(ctx context.Context) error {
	_, err := ptdu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptdu *PalacesTodoDoneUpdate) ExecX(ctx context.Context) {
	if err := ptdu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptdu *PalacesTodoDoneUpdate) defaults() error {
	if _, ok := ptdu.mutation.CreatedAt(); !ok {
		if palacestododone.UpdateDefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized palacestododone.UpdateDefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := palacestododone.UpdateDefaultCreatedAt()
		ptdu.mutation.SetCreatedAt(v)
	}
	if _, ok := ptdu.mutation.UpdatedAt(); !ok {
		if palacestododone.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized palacestododone.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := palacestododone.UpdateDefaultUpdatedAt()
		ptdu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ptdu *PalacesTodoDoneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(palacestododone.Table, palacestododone.Columns, sqlgraph.NewFieldSpec(palacestododone.FieldID, field.TypeInt))
	if ps := ptdu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptdu.mutation.CreatedAt(); ok {
		_spec.SetField(palacestododone.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.AddedCreatedAt(); ok {
		_spec.AddField(palacestododone.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.UpdatedAt(); ok {
		_spec.SetField(palacestododone.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(palacestododone.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.UpdatedBy(); ok {
		_spec.SetField(palacestododone.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(palacestododone.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.DeletedAt(); ok {
		_spec.SetField(palacestododone.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(palacestododone.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.DeletedBy(); ok {
		_spec.SetField(palacestododone.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := ptdu.mutation.AddedDeletedBy(); ok {
		_spec.AddField(palacestododone.FieldDeletedBy, field.TypeInt64, value)
	}
	if ptdu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   palacestododone.OwnerTable,
			Columns: []string{palacestododone.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(palacestodo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptdu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   palacestododone.OwnerTable,
			Columns: []string{palacestododone.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(palacestodo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ptdu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{palacestododone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ptdu.mutation.done = true
	return n, nil
}

// PalacesTodoDoneUpdateOne is the builder for updating a single PalacesTodoDone entity.
type PalacesTodoDoneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PalacesTodoDoneMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ptduo *PalacesTodoDoneUpdateOne) SetUpdatedAt(i int64) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.ResetUpdatedAt()
	ptduo.mutation.SetUpdatedAt(i)
	return ptduo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ptduo *PalacesTodoDoneUpdateOne) AddUpdatedAt(i int64) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.AddUpdatedAt(i)
	return ptduo
}

// SetUpdatedBy sets the "updated_by" field.
func (ptduo *PalacesTodoDoneUpdateOne) SetUpdatedBy(i int64) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.ResetUpdatedBy()
	ptduo.mutation.SetUpdatedBy(i)
	return ptduo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ptduo *PalacesTodoDoneUpdateOne) SetNillableUpdatedBy(i *int64) *PalacesTodoDoneUpdateOne {
	if i != nil {
		ptduo.SetUpdatedBy(*i)
	}
	return ptduo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ptduo *PalacesTodoDoneUpdateOne) AddUpdatedBy(i int64) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.AddUpdatedBy(i)
	return ptduo
}

// SetDeletedAt sets the "deleted_at" field.
func (ptduo *PalacesTodoDoneUpdateOne) SetDeletedAt(i int64) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.ResetDeletedAt()
	ptduo.mutation.SetDeletedAt(i)
	return ptduo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ptduo *PalacesTodoDoneUpdateOne) SetNillableDeletedAt(i *int64) *PalacesTodoDoneUpdateOne {
	if i != nil {
		ptduo.SetDeletedAt(*i)
	}
	return ptduo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ptduo *PalacesTodoDoneUpdateOne) AddDeletedAt(i int64) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.AddDeletedAt(i)
	return ptduo
}

// SetDeletedBy sets the "deleted_by" field.
func (ptduo *PalacesTodoDoneUpdateOne) SetDeletedBy(i int64) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.ResetDeletedBy()
	ptduo.mutation.SetDeletedBy(i)
	return ptduo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ptduo *PalacesTodoDoneUpdateOne) SetNillableDeletedBy(i *int64) *PalacesTodoDoneUpdateOne {
	if i != nil {
		ptduo.SetDeletedBy(*i)
	}
	return ptduo
}

// AddDeletedBy adds i to the "deleted_by" field.
func (ptduo *PalacesTodoDoneUpdateOne) AddDeletedBy(i int64) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.AddDeletedBy(i)
	return ptduo
}

// SetTodoID sets the "todo_id" field.
func (ptduo *PalacesTodoDoneUpdateOne) SetTodoID(i int) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.SetTodoID(i)
	return ptduo
}

// SetNillableTodoID sets the "todo_id" field if the given value is not nil.
func (ptduo *PalacesTodoDoneUpdateOne) SetNillableTodoID(i *int) *PalacesTodoDoneUpdateOne {
	if i != nil {
		ptduo.SetTodoID(*i)
	}
	return ptduo
}

// ClearTodoID clears the value of the "todo_id" field.
func (ptduo *PalacesTodoDoneUpdateOne) ClearTodoID() *PalacesTodoDoneUpdateOne {
	ptduo.mutation.ClearTodoID()
	return ptduo
}

// SetOwnerID sets the "owner" edge to the PalacesTodo entity by ID.
func (ptduo *PalacesTodoDoneUpdateOne) SetOwnerID(id int) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.SetOwnerID(id)
	return ptduo
}

// SetNillableOwnerID sets the "owner" edge to the PalacesTodo entity by ID if the given value is not nil.
func (ptduo *PalacesTodoDoneUpdateOne) SetNillableOwnerID(id *int) *PalacesTodoDoneUpdateOne {
	if id != nil {
		ptduo = ptduo.SetOwnerID(*id)
	}
	return ptduo
}

// SetOwner sets the "owner" edge to the PalacesTodo entity.
func (ptduo *PalacesTodoDoneUpdateOne) SetOwner(p *PalacesTodo) *PalacesTodoDoneUpdateOne {
	return ptduo.SetOwnerID(p.ID)
}

// Mutation returns the PalacesTodoDoneMutation object of the builder.
func (ptduo *PalacesTodoDoneUpdateOne) Mutation() *PalacesTodoDoneMutation {
	return ptduo.mutation
}

// ClearOwner clears the "owner" edge to the PalacesTodo entity.
func (ptduo *PalacesTodoDoneUpdateOne) ClearOwner() *PalacesTodoDoneUpdateOne {
	ptduo.mutation.ClearOwner()
	return ptduo
}

// Where appends a list predicates to the PalacesTodoDoneUpdate builder.
func (ptduo *PalacesTodoDoneUpdateOne) Where(ps ...predicate.PalacesTodoDone) *PalacesTodoDoneUpdateOne {
	ptduo.mutation.Where(ps...)
	return ptduo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptduo *PalacesTodoDoneUpdateOne) Select(field string, fields ...string) *PalacesTodoDoneUpdateOne {
	ptduo.fields = append([]string{field}, fields...)
	return ptduo
}

// Save executes the query and returns the updated PalacesTodoDone entity.
func (ptduo *PalacesTodoDoneUpdateOne) Save(ctx context.Context) (*PalacesTodoDone, error) {
	if err := ptduo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ptduo.sqlSave, ptduo.mutation, ptduo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ptduo *PalacesTodoDoneUpdateOne) SaveX(ctx context.Context) *PalacesTodoDone {
	node, err := ptduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptduo *PalacesTodoDoneUpdateOne) Exec(ctx context.Context) error {
	_, err := ptduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptduo *PalacesTodoDoneUpdateOne) ExecX(ctx context.Context) {
	if err := ptduo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ptduo *PalacesTodoDoneUpdateOne) defaults() error {
	if _, ok := ptduo.mutation.CreatedAt(); !ok {
		if palacestododone.UpdateDefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized palacestododone.UpdateDefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := palacestododone.UpdateDefaultCreatedAt()
		ptduo.mutation.SetCreatedAt(v)
	}
	if _, ok := ptduo.mutation.UpdatedAt(); !ok {
		if palacestododone.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized palacestododone.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := palacestododone.UpdateDefaultUpdatedAt()
		ptduo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ptduo *PalacesTodoDoneUpdateOne) sqlSave(ctx context.Context) (_node *PalacesTodoDone, err error) {
	_spec := sqlgraph.NewUpdateSpec(palacestododone.Table, palacestododone.Columns, sqlgraph.NewFieldSpec(palacestododone.FieldID, field.TypeInt))
	id, ok := ptduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PalacesTodoDone.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptduo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, palacestododone.FieldID)
		for _, f := range fields {
			if !palacestododone.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != palacestododone.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptduo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptduo.mutation.CreatedAt(); ok {
		_spec.SetField(palacestododone.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.AddedCreatedAt(); ok {
		_spec.AddField(palacestododone.FieldCreatedAt, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.UpdatedAt(); ok {
		_spec.SetField(palacestododone.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(palacestododone.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.UpdatedBy(); ok {
		_spec.SetField(palacestododone.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(palacestododone.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.DeletedAt(); ok {
		_spec.SetField(palacestododone.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(palacestododone.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.DeletedBy(); ok {
		_spec.SetField(palacestododone.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := ptduo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(palacestododone.FieldDeletedBy, field.TypeInt64, value)
	}
	if ptduo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   palacestododone.OwnerTable,
			Columns: []string{palacestododone.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(palacestodo.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ptduo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   palacestododone.OwnerTable,
			Columns: []string{palacestododone.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(palacestodo.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PalacesTodoDone{config: ptduo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{palacestododone.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ptduo.mutation.done = true
	return _node, nil
}
