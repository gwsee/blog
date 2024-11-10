// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/account"
	"blog/internal/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountUpdate is the builder for updating Account entities.
type AccountUpdate struct {
	config
	hooks    []Hook
	mutation *AccountMutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (au *AccountUpdate) Where(ps ...predicate.Account) *AccountUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AccountUpdate) SetUpdatedAt(i int64) *AccountUpdate {
	au.mutation.ResetUpdatedAt()
	au.mutation.SetUpdatedAt(i)
	return au
}

// AddUpdatedAt adds i to the "updated_at" field.
func (au *AccountUpdate) AddUpdatedAt(i int64) *AccountUpdate {
	au.mutation.AddUpdatedAt(i)
	return au
}

// SetUpdatedBy sets the "updated_by" field.
func (au *AccountUpdate) SetUpdatedBy(i int64) *AccountUpdate {
	au.mutation.ResetUpdatedBy()
	au.mutation.SetUpdatedBy(i)
	return au
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (au *AccountUpdate) SetNillableUpdatedBy(i *int64) *AccountUpdate {
	if i != nil {
		au.SetUpdatedBy(*i)
	}
	return au
}

// AddUpdatedBy adds i to the "updated_by" field.
func (au *AccountUpdate) AddUpdatedBy(i int64) *AccountUpdate {
	au.mutation.AddUpdatedBy(i)
	return au
}

// SetDeletedAt sets the "deleted_at" field.
func (au *AccountUpdate) SetDeletedAt(i int64) *AccountUpdate {
	au.mutation.ResetDeletedAt()
	au.mutation.SetDeletedAt(i)
	return au
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDeletedAt(i *int64) *AccountUpdate {
	if i != nil {
		au.SetDeletedAt(*i)
	}
	return au
}

// AddDeletedAt adds i to the "deleted_at" field.
func (au *AccountUpdate) AddDeletedAt(i int64) *AccountUpdate {
	au.mutation.AddDeletedAt(i)
	return au
}

// SetDeletedBy sets the "deleted_by" field.
func (au *AccountUpdate) SetDeletedBy(i int64) *AccountUpdate {
	au.mutation.ResetDeletedBy()
	au.mutation.SetDeletedBy(i)
	return au
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (au *AccountUpdate) SetNillableDeletedBy(i *int64) *AccountUpdate {
	if i != nil {
		au.SetDeletedBy(*i)
	}
	return au
}

// AddDeletedBy adds i to the "deleted_by" field.
func (au *AccountUpdate) AddDeletedBy(i int64) *AccountUpdate {
	au.mutation.AddDeletedBy(i)
	return au
}

// SetAccount sets the "account" field.
func (au *AccountUpdate) SetAccount(s string) *AccountUpdate {
	au.mutation.SetAccount(s)
	return au
}

// SetNillableAccount sets the "account" field if the given value is not nil.
func (au *AccountUpdate) SetNillableAccount(s *string) *AccountUpdate {
	if s != nil {
		au.SetAccount(*s)
	}
	return au
}

// SetPassword sets the "password" field.
func (au *AccountUpdate) SetPassword(s string) *AccountUpdate {
	au.mutation.SetPassword(s)
	return au
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (au *AccountUpdate) SetNillablePassword(s *string) *AccountUpdate {
	if s != nil {
		au.SetPassword(*s)
	}
	return au
}

// SetEmail sets the "email" field.
func (au *AccountUpdate) SetEmail(s string) *AccountUpdate {
	au.mutation.SetEmail(s)
	return au
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (au *AccountUpdate) SetNillableEmail(s *string) *AccountUpdate {
	if s != nil {
		au.SetEmail(*s)
	}
	return au
}

// SetStatus sets the "status" field.
func (au *AccountUpdate) SetStatus(i int8) *AccountUpdate {
	au.mutation.ResetStatus()
	au.mutation.SetStatus(i)
	return au
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (au *AccountUpdate) SetNillableStatus(i *int8) *AccountUpdate {
	if i != nil {
		au.SetStatus(*i)
	}
	return au
}

// AddStatus adds i to the "status" field.
func (au *AccountUpdate) AddStatus(i int8) *AccountUpdate {
	au.mutation.AddStatus(i)
	return au
}

// Mutation returns the AccountMutation object of the builder.
func (au *AccountUpdate) Mutation() *AccountMutation {
	return au.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AccountUpdate) Save(ctx context.Context) (int, error) {
	if err := au.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AccountUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AccountUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AccountUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AccountUpdate) defaults() error {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		if account.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized account.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := account.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (au *AccountUpdate) check() error {
	if v, ok := au.mutation.Account(); ok {
		if err := account.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf(`ent: validator failed for field "Account.account": %w`, err)}
		}
	}
	if v, ok := au.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	return nil
}

func (au *AccountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(account.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := au.mutation.UpdatedBy(); ok {
		_spec.SetField(account.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(account.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := au.mutation.DeletedAt(); ok {
		_spec.SetField(account.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedDeletedAt(); ok {
		_spec.AddField(account.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := au.mutation.DeletedBy(); ok {
		_spec.SetField(account.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := au.mutation.AddedDeletedBy(); ok {
		_spec.AddField(account.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := au.mutation.Account(); ok {
		_spec.SetField(account.FieldAccount, field.TypeString, value)
	}
	if value, ok := au.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if value, ok := au.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
	}
	if value, ok := au.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := au.mutation.AddedStatus(); ok {
		_spec.AddField(account.FieldStatus, field.TypeInt8, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AccountUpdateOne is the builder for updating a single Account entity.
type AccountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AccountMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AccountUpdateOne) SetUpdatedAt(i int64) *AccountUpdateOne {
	auo.mutation.ResetUpdatedAt()
	auo.mutation.SetUpdatedAt(i)
	return auo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (auo *AccountUpdateOne) AddUpdatedAt(i int64) *AccountUpdateOne {
	auo.mutation.AddUpdatedAt(i)
	return auo
}

// SetUpdatedBy sets the "updated_by" field.
func (auo *AccountUpdateOne) SetUpdatedBy(i int64) *AccountUpdateOne {
	auo.mutation.ResetUpdatedBy()
	auo.mutation.SetUpdatedBy(i)
	return auo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableUpdatedBy(i *int64) *AccountUpdateOne {
	if i != nil {
		auo.SetUpdatedBy(*i)
	}
	return auo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (auo *AccountUpdateOne) AddUpdatedBy(i int64) *AccountUpdateOne {
	auo.mutation.AddUpdatedBy(i)
	return auo
}

// SetDeletedAt sets the "deleted_at" field.
func (auo *AccountUpdateOne) SetDeletedAt(i int64) *AccountUpdateOne {
	auo.mutation.ResetDeletedAt()
	auo.mutation.SetDeletedAt(i)
	return auo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDeletedAt(i *int64) *AccountUpdateOne {
	if i != nil {
		auo.SetDeletedAt(*i)
	}
	return auo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (auo *AccountUpdateOne) AddDeletedAt(i int64) *AccountUpdateOne {
	auo.mutation.AddDeletedAt(i)
	return auo
}

// SetDeletedBy sets the "deleted_by" field.
func (auo *AccountUpdateOne) SetDeletedBy(i int64) *AccountUpdateOne {
	auo.mutation.ResetDeletedBy()
	auo.mutation.SetDeletedBy(i)
	return auo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableDeletedBy(i *int64) *AccountUpdateOne {
	if i != nil {
		auo.SetDeletedBy(*i)
	}
	return auo
}

// AddDeletedBy adds i to the "deleted_by" field.
func (auo *AccountUpdateOne) AddDeletedBy(i int64) *AccountUpdateOne {
	auo.mutation.AddDeletedBy(i)
	return auo
}

// SetAccount sets the "account" field.
func (auo *AccountUpdateOne) SetAccount(s string) *AccountUpdateOne {
	auo.mutation.SetAccount(s)
	return auo
}

// SetNillableAccount sets the "account" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableAccount(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetAccount(*s)
	}
	return auo
}

// SetPassword sets the "password" field.
func (auo *AccountUpdateOne) SetPassword(s string) *AccountUpdateOne {
	auo.mutation.SetPassword(s)
	return auo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillablePassword(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetPassword(*s)
	}
	return auo
}

// SetEmail sets the "email" field.
func (auo *AccountUpdateOne) SetEmail(s string) *AccountUpdateOne {
	auo.mutation.SetEmail(s)
	return auo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableEmail(s *string) *AccountUpdateOne {
	if s != nil {
		auo.SetEmail(*s)
	}
	return auo
}

// SetStatus sets the "status" field.
func (auo *AccountUpdateOne) SetStatus(i int8) *AccountUpdateOne {
	auo.mutation.ResetStatus()
	auo.mutation.SetStatus(i)
	return auo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (auo *AccountUpdateOne) SetNillableStatus(i *int8) *AccountUpdateOne {
	if i != nil {
		auo.SetStatus(*i)
	}
	return auo
}

// AddStatus adds i to the "status" field.
func (auo *AccountUpdateOne) AddStatus(i int8) *AccountUpdateOne {
	auo.mutation.AddStatus(i)
	return auo
}

// Mutation returns the AccountMutation object of the builder.
func (auo *AccountUpdateOne) Mutation() *AccountMutation {
	return auo.mutation
}

// Where appends a list predicates to the AccountUpdate builder.
func (auo *AccountUpdateOne) Where(ps ...predicate.Account) *AccountUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AccountUpdateOne) Select(field string, fields ...string) *AccountUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Account entity.
func (auo *AccountUpdateOne) Save(ctx context.Context) (*Account, error) {
	if err := auo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AccountUpdateOne) SaveX(ctx context.Context) *Account {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AccountUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AccountUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AccountUpdateOne) defaults() error {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		if account.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized account.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := account.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (auo *AccountUpdateOne) check() error {
	if v, ok := auo.mutation.Account(); ok {
		if err := account.AccountValidator(v); err != nil {
			return &ValidationError{Name: "account", err: fmt.Errorf(`ent: validator failed for field "Account.account": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	return nil
}

func (auo *AccountUpdateOne) sqlSave(ctx context.Context) (_node *Account, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(account.Table, account.Columns, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Account.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, account.FieldID)
		for _, f := range fields {
			if !account.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != account.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(account.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.UpdatedBy(); ok {
		_spec.SetField(account.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(account.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.DeletedAt(); ok {
		_spec.SetField(account.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(account.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.DeletedBy(); ok {
		_spec.SetField(account.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(account.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := auo.mutation.Account(); ok {
		_spec.SetField(account.FieldAccount, field.TypeString, value)
	}
	if value, ok := auo.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
	}
	if value, ok := auo.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
	}
	if value, ok := auo.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeInt8, value)
	}
	if value, ok := auo.mutation.AddedStatus(); ok {
		_spec.AddField(account.FieldStatus, field.TypeInt8, value)
	}
	_node = &Account{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{account.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}