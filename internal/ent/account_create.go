// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/account"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AccountCreate is the builder for creating a Account entity.
type AccountCreate struct {
	config
	mutation *AccountMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (ac *AccountCreate) SetCreatedAt(i int64) *AccountCreate {
	ac.mutation.SetCreatedAt(i)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedAt(i *int64) *AccountCreate {
	if i != nil {
		ac.SetCreatedAt(*i)
	}
	return ac
}

// SetCreatedBy sets the "created_by" field.
func (ac *AccountCreate) SetCreatedBy(s string) *AccountCreate {
	ac.mutation.SetCreatedBy(s)
	return ac
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (ac *AccountCreate) SetNillableCreatedBy(s *string) *AccountCreate {
	if s != nil {
		ac.SetCreatedBy(*s)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AccountCreate) SetUpdatedAt(i int64) *AccountCreate {
	ac.mutation.SetUpdatedAt(i)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedAt(i *int64) *AccountCreate {
	if i != nil {
		ac.SetUpdatedAt(*i)
	}
	return ac
}

// SetUpdatedBy sets the "updated_by" field.
func (ac *AccountCreate) SetUpdatedBy(s string) *AccountCreate {
	ac.mutation.SetUpdatedBy(s)
	return ac
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ac *AccountCreate) SetNillableUpdatedBy(s *string) *AccountCreate {
	if s != nil {
		ac.SetUpdatedBy(*s)
	}
	return ac
}

// SetIsDeleted sets the "is_deleted" field.
func (ac *AccountCreate) SetIsDeleted(u uint8) *AccountCreate {
	ac.mutation.SetIsDeleted(u)
	return ac
}

// SetNillableIsDeleted sets the "is_deleted" field if the given value is not nil.
func (ac *AccountCreate) SetNillableIsDeleted(u *uint8) *AccountCreate {
	if u != nil {
		ac.SetIsDeleted(*u)
	}
	return ac
}

// SetDeletedAt sets the "deleted_at" field.
func (ac *AccountCreate) SetDeletedAt(i int64) *AccountCreate {
	ac.mutation.SetDeletedAt(i)
	return ac
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ac *AccountCreate) SetNillableDeletedAt(i *int64) *AccountCreate {
	if i != nil {
		ac.SetDeletedAt(*i)
	}
	return ac
}

// SetDeletedBy sets the "deleted_by" field.
func (ac *AccountCreate) SetDeletedBy(s string) *AccountCreate {
	ac.mutation.SetDeletedBy(s)
	return ac
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ac *AccountCreate) SetNillableDeletedBy(s *string) *AccountCreate {
	if s != nil {
		ac.SetDeletedBy(*s)
	}
	return ac
}

// SetAccount sets the "account" field.
func (ac *AccountCreate) SetAccount(s string) *AccountCreate {
	ac.mutation.SetAccount(s)
	return ac
}

// SetPassword sets the "password" field.
func (ac *AccountCreate) SetPassword(s string) *AccountCreate {
	ac.mutation.SetPassword(s)
	return ac
}

// SetEmail sets the "email" field.
func (ac *AccountCreate) SetEmail(s string) *AccountCreate {
	ac.mutation.SetEmail(s)
	return ac
}

// SetStatus sets the "status" field.
func (ac *AccountCreate) SetStatus(i int8) *AccountCreate {
	ac.mutation.SetStatus(i)
	return ac
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ac *AccountCreate) SetNillableStatus(i *int8) *AccountCreate {
	if i != nil {
		ac.SetStatus(*i)
	}
	return ac
}

// SetID sets the "id" field.
func (ac *AccountCreate) SetID(i int32) *AccountCreate {
	ac.mutation.SetID(i)
	return ac
}

// Mutation returns the AccountMutation object of the builder.
func (ac *AccountCreate) Mutation() *AccountMutation {
	return ac.mutation
}

// Save creates the Account in the database.
func (ac *AccountCreate) Save(ctx context.Context) (*Account, error) {
	if err := ac.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AccountCreate) SaveX(ctx context.Context) *Account {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AccountCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AccountCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AccountCreate) defaults() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := account.DefaultCreatedAt
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.CreatedBy(); !ok {
		v := account.DefaultCreatedBy
		ac.mutation.SetCreatedBy(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := account.DefaultUpdatedAt
		ac.mutation.SetUpdatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedBy(); !ok {
		v := account.DefaultUpdatedBy
		ac.mutation.SetUpdatedBy(v)
	}
	if _, ok := ac.mutation.IsDeleted(); !ok {
		v := account.DefaultIsDeleted
		ac.mutation.SetIsDeleted(v)
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		v := account.DefaultDeletedAt
		ac.mutation.SetDeletedAt(v)
	}
	if _, ok := ac.mutation.DeletedBy(); !ok {
		v := account.DefaultDeletedBy
		ac.mutation.SetDeletedBy(v)
	}
	if _, ok := ac.mutation.Status(); !ok {
		v := account.DefaultStatus
		ac.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (ac *AccountCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Account.created_at"`)}
	}
	if _, ok := ac.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "Account.created_by"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Account.updated_at"`)}
	}
	if _, ok := ac.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "Account.updated_by"`)}
	}
	if _, ok := ac.mutation.IsDeleted(); !ok {
		return &ValidationError{Name: "is_deleted", err: errors.New(`ent: missing required field "Account.is_deleted"`)}
	}
	if _, ok := ac.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "Account.deleted_at"`)}
	}
	if _, ok := ac.mutation.DeletedBy(); !ok {
		return &ValidationError{Name: "deleted_by", err: errors.New(`ent: missing required field "Account.deleted_by"`)}
	}
	if _, ok := ac.mutation.Account(); !ok {
		return &ValidationError{Name: "account", err: errors.New(`ent: missing required field "Account.account"`)}
	}
	if _, ok := ac.mutation.Password(); !ok {
		return &ValidationError{Name: "password", err: errors.New(`ent: missing required field "Account.password"`)}
	}
	if v, ok := ac.mutation.Password(); ok {
		if err := account.PasswordValidator(v); err != nil {
			return &ValidationError{Name: "password", err: fmt.Errorf(`ent: validator failed for field "Account.password": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Account.email"`)}
	}
	if v, ok := ac.mutation.Email(); ok {
		if err := account.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Account.email": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Account.status"`)}
	}
	if v, ok := ac.mutation.ID(); ok {
		if err := account.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Account.id": %w`, err)}
		}
	}
	return nil
}

func (ac *AccountCreate) sqlSave(ctx context.Context) (*Account, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int32(id)
	}
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AccountCreate) createSpec() (*Account, *sqlgraph.CreateSpec) {
	var (
		_node = &Account{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(account.Table, sqlgraph.NewFieldSpec(account.FieldID, field.TypeInt32))
	)
	_spec.OnConflict = ac.conflict
	if id, ok := ac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(account.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.CreatedBy(); ok {
		_spec.SetField(account.FieldCreatedBy, field.TypeString, value)
		_node.CreatedBy = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(account.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.UpdatedBy(); ok {
		_spec.SetField(account.FieldUpdatedBy, field.TypeString, value)
		_node.UpdatedBy = value
	}
	if value, ok := ac.mutation.IsDeleted(); ok {
		_spec.SetField(account.FieldIsDeleted, field.TypeUint8, value)
		_node.IsDeleted = value
	}
	if value, ok := ac.mutation.DeletedAt(); ok {
		_spec.SetField(account.FieldDeletedAt, field.TypeInt64, value)
		_node.DeletedAt = value
	}
	if value, ok := ac.mutation.DeletedBy(); ok {
		_spec.SetField(account.FieldDeletedBy, field.TypeString, value)
		_node.DeletedBy = value
	}
	if value, ok := ac.mutation.Account(); ok {
		_spec.SetField(account.FieldAccount, field.TypeString, value)
		_node.Account = value
	}
	if value, ok := ac.mutation.Password(); ok {
		_spec.SetField(account.FieldPassword, field.TypeString, value)
		_node.Password = value
	}
	if value, ok := ac.mutation.Email(); ok {
		_spec.SetField(account.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := ac.mutation.Status(); ok {
		_spec.SetField(account.FieldStatus, field.TypeInt8, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ac *AccountCreate) OnConflict(opts ...sql.ConflictOption) *AccountUpsertOne {
	ac.conflict = opts
	return &AccountUpsertOne{
		create: ac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ac *AccountCreate) OnConflictColumns(columns ...string) *AccountUpsertOne {
	ac.conflict = append(ac.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertOne{
		create: ac,
	}
}

type (
	// AccountUpsertOne is the builder for "upsert"-ing
	//  one Account node.
	AccountUpsertOne struct {
		create *AccountCreate
	}

	// AccountUpsert is the "OnConflict" setter.
	AccountUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *AccountUpsert) SetUpdatedAt(v int64) *AccountUpsert {
	u.Set(account.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AccountUpsert) UpdateUpdatedAt() *AccountUpsert {
	u.SetExcluded(account.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AccountUpsert) AddUpdatedAt(v int64) *AccountUpsert {
	u.Add(account.FieldUpdatedAt, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AccountUpsert) SetUpdatedBy(v string) *AccountUpsert {
	u.Set(account.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AccountUpsert) UpdateUpdatedBy() *AccountUpsert {
	u.SetExcluded(account.FieldUpdatedBy)
	return u
}

// SetIsDeleted sets the "is_deleted" field.
func (u *AccountUpsert) SetIsDeleted(v uint8) *AccountUpsert {
	u.Set(account.FieldIsDeleted, v)
	return u
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *AccountUpsert) UpdateIsDeleted() *AccountUpsert {
	u.SetExcluded(account.FieldIsDeleted)
	return u
}

// AddIsDeleted adds v to the "is_deleted" field.
func (u *AccountUpsert) AddIsDeleted(v uint8) *AccountUpsert {
	u.Add(account.FieldIsDeleted, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AccountUpsert) SetDeletedAt(v int64) *AccountUpsert {
	u.Set(account.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AccountUpsert) UpdateDeletedAt() *AccountUpsert {
	u.SetExcluded(account.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AccountUpsert) AddDeletedAt(v int64) *AccountUpsert {
	u.Add(account.FieldDeletedAt, v)
	return u
}

// SetDeletedBy sets the "deleted_by" field.
func (u *AccountUpsert) SetDeletedBy(v string) *AccountUpsert {
	u.Set(account.FieldDeletedBy, v)
	return u
}

// UpdateDeletedBy sets the "deleted_by" field to the value that was provided on create.
func (u *AccountUpsert) UpdateDeletedBy() *AccountUpsert {
	u.SetExcluded(account.FieldDeletedBy)
	return u
}

// SetAccount sets the "account" field.
func (u *AccountUpsert) SetAccount(v string) *AccountUpsert {
	u.Set(account.FieldAccount, v)
	return u
}

// UpdateAccount sets the "account" field to the value that was provided on create.
func (u *AccountUpsert) UpdateAccount() *AccountUpsert {
	u.SetExcluded(account.FieldAccount)
	return u
}

// SetPassword sets the "password" field.
func (u *AccountUpsert) SetPassword(v string) *AccountUpsert {
	u.Set(account.FieldPassword, v)
	return u
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AccountUpsert) UpdatePassword() *AccountUpsert {
	u.SetExcluded(account.FieldPassword)
	return u
}

// SetEmail sets the "email" field.
func (u *AccountUpsert) SetEmail(v string) *AccountUpsert {
	u.Set(account.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AccountUpsert) UpdateEmail() *AccountUpsert {
	u.SetExcluded(account.FieldEmail)
	return u
}

// SetStatus sets the "status" field.
func (u *AccountUpsert) SetStatus(v int8) *AccountUpsert {
	u.Set(account.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AccountUpsert) UpdateStatus() *AccountUpsert {
	u.SetExcluded(account.FieldStatus)
	return u
}

// AddStatus adds v to the "status" field.
func (u *AccountUpsert) AddStatus(v int8) *AccountUpsert {
	u.Add(account.FieldStatus, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(account.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AccountUpsertOne) UpdateNewValues() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(account.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(account.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(account.FieldCreatedBy)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *AccountUpsertOne) Ignore() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertOne) DoNothing() *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreate.OnConflict
// documentation for more info.
func (u *AccountUpsertOne) Update(set func(*AccountUpsert)) *AccountUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AccountUpsertOne) SetUpdatedAt(v int64) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AccountUpsertOne) AddUpdatedAt(v int64) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateUpdatedAt() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AccountUpsertOne) SetUpdatedBy(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateUpdatedBy() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetIsDeleted sets the "is_deleted" field.
func (u *AccountUpsertOne) SetIsDeleted(v uint8) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetIsDeleted(v)
	})
}

// AddIsDeleted adds v to the "is_deleted" field.
func (u *AccountUpsertOne) AddIsDeleted(v uint8) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.AddIsDeleted(v)
	})
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateIsDeleted() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateIsDeleted()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AccountUpsertOne) SetDeletedAt(v int64) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AccountUpsertOne) AddDeletedAt(v int64) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateDeletedAt() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeletedBy sets the "deleted_by" field.
func (u *AccountUpsertOne) SetDeletedBy(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetDeletedBy(v)
	})
}

// UpdateDeletedBy sets the "deleted_by" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateDeletedBy() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDeletedBy()
	})
}

// SetAccount sets the "account" field.
func (u *AccountUpsertOne) SetAccount(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetAccount(v)
	})
}

// UpdateAccount sets the "account" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateAccount() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateAccount()
	})
}

// SetPassword sets the "password" field.
func (u *AccountUpsertOne) SetPassword(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdatePassword() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdatePassword()
	})
}

// SetEmail sets the "email" field.
func (u *AccountUpsertOne) SetEmail(v string) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateEmail() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateEmail()
	})
}

// SetStatus sets the "status" field.
func (u *AccountUpsertOne) SetStatus(v int8) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *AccountUpsertOne) AddStatus(v int8) *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AccountUpsertOne) UpdateStatus() *AccountUpsertOne {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *AccountUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AccountUpsertOne) ID(ctx context.Context) (id int32, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AccountUpsertOne) IDX(ctx context.Context) int32 {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AccountCreateBulk is the builder for creating many Account entities in bulk.
type AccountCreateBulk struct {
	config
	err      error
	builders []*AccountCreate
	conflict []sql.ConflictOption
}

// Save creates the Account entities in the database.
func (acb *AccountCreateBulk) Save(ctx context.Context) ([]*Account, error) {
	if acb.err != nil {
		return nil, acb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Account, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AccountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = acb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int32(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AccountCreateBulk) SaveX(ctx context.Context) []*Account {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AccountCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AccountCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Account.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AccountUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (acb *AccountCreateBulk) OnConflict(opts ...sql.ConflictOption) *AccountUpsertBulk {
	acb.conflict = opts
	return &AccountUpsertBulk{
		create: acb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (acb *AccountCreateBulk) OnConflictColumns(columns ...string) *AccountUpsertBulk {
	acb.conflict = append(acb.conflict, sql.ConflictColumns(columns...))
	return &AccountUpsertBulk{
		create: acb,
	}
}

// AccountUpsertBulk is the builder for "upsert"-ing
// a bulk of Account nodes.
type AccountUpsertBulk struct {
	create *AccountCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(account.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *AccountUpsertBulk) UpdateNewValues() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(account.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(account.FieldCreatedAt)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(account.FieldCreatedBy)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Account.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *AccountUpsertBulk) Ignore() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AccountUpsertBulk) DoNothing() *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AccountCreateBulk.OnConflict
// documentation for more info.
func (u *AccountUpsertBulk) Update(set func(*AccountUpsert)) *AccountUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AccountUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *AccountUpsertBulk) SetUpdatedAt(v int64) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *AccountUpsertBulk) AddUpdatedAt(v int64) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateUpdatedAt() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *AccountUpsertBulk) SetUpdatedBy(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateUpdatedBy() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetIsDeleted sets the "is_deleted" field.
func (u *AccountUpsertBulk) SetIsDeleted(v uint8) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetIsDeleted(v)
	})
}

// AddIsDeleted adds v to the "is_deleted" field.
func (u *AccountUpsertBulk) AddIsDeleted(v uint8) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.AddIsDeleted(v)
	})
}

// UpdateIsDeleted sets the "is_deleted" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateIsDeleted() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateIsDeleted()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *AccountUpsertBulk) SetDeletedAt(v int64) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *AccountUpsertBulk) AddDeletedAt(v int64) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateDeletedAt() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeletedBy sets the "deleted_by" field.
func (u *AccountUpsertBulk) SetDeletedBy(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetDeletedBy(v)
	})
}

// UpdateDeletedBy sets the "deleted_by" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateDeletedBy() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateDeletedBy()
	})
}

// SetAccount sets the "account" field.
func (u *AccountUpsertBulk) SetAccount(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetAccount(v)
	})
}

// UpdateAccount sets the "account" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateAccount() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateAccount()
	})
}

// SetPassword sets the "password" field.
func (u *AccountUpsertBulk) SetPassword(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetPassword(v)
	})
}

// UpdatePassword sets the "password" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdatePassword() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdatePassword()
	})
}

// SetEmail sets the "email" field.
func (u *AccountUpsertBulk) SetEmail(v string) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateEmail() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateEmail()
	})
}

// SetStatus sets the "status" field.
func (u *AccountUpsertBulk) SetStatus(v int8) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.SetStatus(v)
	})
}

// AddStatus adds v to the "status" field.
func (u *AccountUpsertBulk) AddStatus(v int8) *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.AddStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *AccountUpsertBulk) UpdateStatus() *AccountUpsertBulk {
	return u.Update(func(s *AccountUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *AccountUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AccountCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AccountCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AccountUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
