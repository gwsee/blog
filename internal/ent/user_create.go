// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/user"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (uc *UserCreate) SetCreatedAt(i int64) *UserCreate {
	uc.mutation.SetCreatedAt(i)
	return uc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedAt(i *int64) *UserCreate {
	if i != nil {
		uc.SetCreatedAt(*i)
	}
	return uc
}

// SetCreatedBy sets the "created_by" field.
func (uc *UserCreate) SetCreatedBy(i int64) *UserCreate {
	uc.mutation.SetCreatedBy(i)
	return uc
}

// SetNillableCreatedBy sets the "created_by" field if the given value is not nil.
func (uc *UserCreate) SetNillableCreatedBy(i *int64) *UserCreate {
	if i != nil {
		uc.SetCreatedBy(*i)
	}
	return uc
}

// SetUpdatedAt sets the "updated_at" field.
func (uc *UserCreate) SetUpdatedAt(i int64) *UserCreate {
	uc.mutation.SetUpdatedAt(i)
	return uc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedAt(i *int64) *UserCreate {
	if i != nil {
		uc.SetUpdatedAt(*i)
	}
	return uc
}

// SetUpdatedBy sets the "updated_by" field.
func (uc *UserCreate) SetUpdatedBy(i int64) *UserCreate {
	uc.mutation.SetUpdatedBy(i)
	return uc
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (uc *UserCreate) SetNillableUpdatedBy(i *int64) *UserCreate {
	if i != nil {
		uc.SetUpdatedBy(*i)
	}
	return uc
}

// SetDeletedAt sets the "deleted_at" field.
func (uc *UserCreate) SetDeletedAt(i int64) *UserCreate {
	uc.mutation.SetDeletedAt(i)
	return uc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedAt(i *int64) *UserCreate {
	if i != nil {
		uc.SetDeletedAt(*i)
	}
	return uc
}

// SetDeletedBy sets the "deleted_by" field.
func (uc *UserCreate) SetDeletedBy(i int64) *UserCreate {
	uc.mutation.SetDeletedBy(i)
	return uc
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (uc *UserCreate) SetNillableDeletedBy(i *int64) *UserCreate {
	if i != nil {
		uc.SetDeletedBy(*i)
	}
	return uc
}

// SetName sets the "name" field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.mutation.SetName(s)
	return uc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// SetAvatar sets the "avatar" field.
func (uc *UserCreate) SetAvatar(s string) *UserCreate {
	uc.mutation.SetAvatar(s)
	return uc
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uc *UserCreate) SetNillableAvatar(s *string) *UserCreate {
	if s != nil {
		uc.SetAvatar(*s)
	}
	return uc
}

// SetEmail sets the "email" field.
func (uc *UserCreate) SetEmail(s string) *UserCreate {
	uc.mutation.SetEmail(s)
	return uc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uc *UserCreate) SetNillableEmail(s *string) *UserCreate {
	if s != nil {
		uc.SetEmail(*s)
	}
	return uc
}

// SetProfessional sets the "professional" field.
func (uc *UserCreate) SetProfessional(s string) *UserCreate {
	uc.mutation.SetProfessional(s)
	return uc
}

// SetNillableProfessional sets the "professional" field if the given value is not nil.
func (uc *UserCreate) SetNillableProfessional(s *string) *UserCreate {
	if s != nil {
		uc.SetProfessional(*s)
	}
	return uc
}

// SetAddress sets the "address" field.
func (uc *UserCreate) SetAddress(s string) *UserCreate {
	uc.mutation.SetAddress(s)
	return uc
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uc *UserCreate) SetNillableAddress(s *string) *UserCreate {
	if s != nil {
		uc.SetAddress(*s)
	}
	return uc
}

// SetSkills sets the "skills" field.
func (uc *UserCreate) SetSkills(s []string) *UserCreate {
	uc.mutation.SetSkills(s)
	return uc
}

// SetDescription sets the "description" field.
func (uc *UserCreate) SetDescription(s string) *UserCreate {
	uc.mutation.SetDescription(s)
	return uc
}

// SetID sets the "id" field.
func (uc *UserCreate) SetID(i int) *UserCreate {
	uc.mutation.SetID(i)
	return uc
}

// Mutation returns the UserMutation object of the builder.
func (uc *UserCreate) Mutation() *UserMutation {
	return uc.mutation
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if err := uc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, uc.sqlSave, uc.mutation, uc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uc *UserCreate) Exec(ctx context.Context) error {
	_, err := uc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uc *UserCreate) ExecX(ctx context.Context) {
	if err := uc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uc *UserCreate) defaults() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		v := user.DefaultCreatedAt
		uc.mutation.SetCreatedAt(v)
	}
	if _, ok := uc.mutation.CreatedBy(); !ok {
		v := user.DefaultCreatedBy
		uc.mutation.SetCreatedBy(v)
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		v := user.DefaultUpdatedAt
		uc.mutation.SetUpdatedAt(v)
	}
	if _, ok := uc.mutation.UpdatedBy(); !ok {
		v := user.DefaultUpdatedBy
		uc.mutation.SetUpdatedBy(v)
	}
	if _, ok := uc.mutation.DeletedAt(); !ok {
		v := user.DefaultDeletedAt
		uc.mutation.SetDeletedAt(v)
	}
	if _, ok := uc.mutation.DeletedBy(); !ok {
		v := user.DefaultDeletedBy
		uc.mutation.SetDeletedBy(v)
	}
	if _, ok := uc.mutation.Name(); !ok {
		v := user.DefaultName
		uc.mutation.SetName(v)
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		v := user.DefaultAvatar
		uc.mutation.SetAvatar(v)
	}
	if _, ok := uc.mutation.Email(); !ok {
		v := user.DefaultEmail
		uc.mutation.SetEmail(v)
	}
	if _, ok := uc.mutation.Professional(); !ok {
		v := user.DefaultProfessional
		uc.mutation.SetProfessional(v)
	}
	if _, ok := uc.mutation.Address(); !ok {
		v := user.DefaultAddress
		uc.mutation.SetAddress(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (uc *UserCreate) check() error {
	if _, ok := uc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "User.created_at"`)}
	}
	if _, ok := uc.mutation.CreatedBy(); !ok {
		return &ValidationError{Name: "created_by", err: errors.New(`ent: missing required field "User.created_by"`)}
	}
	if _, ok := uc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "User.updated_at"`)}
	}
	if _, ok := uc.mutation.UpdatedBy(); !ok {
		return &ValidationError{Name: "updated_by", err: errors.New(`ent: missing required field "User.updated_by"`)}
	}
	if _, ok := uc.mutation.DeletedAt(); !ok {
		return &ValidationError{Name: "deleted_at", err: errors.New(`ent: missing required field "User.deleted_at"`)}
	}
	if _, ok := uc.mutation.DeletedBy(); !ok {
		return &ValidationError{Name: "deleted_by", err: errors.New(`ent: missing required field "User.deleted_by"`)}
	}
	if _, ok := uc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	if _, ok := uc.mutation.Avatar(); !ok {
		return &ValidationError{Name: "avatar", err: errors.New(`ent: missing required field "User.avatar"`)}
	}
	if _, ok := uc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "User.email"`)}
	}
	if _, ok := uc.mutation.Professional(); !ok {
		return &ValidationError{Name: "professional", err: errors.New(`ent: missing required field "User.professional"`)}
	}
	if _, ok := uc.mutation.Address(); !ok {
		return &ValidationError{Name: "address", err: errors.New(`ent: missing required field "User.address"`)}
	}
	if _, ok := uc.mutation.Skills(); !ok {
		return &ValidationError{Name: "skills", err: errors.New(`ent: missing required field "User.skills"`)}
	}
	if _, ok := uc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "User.description"`)}
	}
	if v, ok := uc.mutation.Description(); ok {
		if err := user.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "User.description": %w`, err)}
		}
	}
	if v, ok := uc.mutation.ID(); ok {
		if err := user.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "User.id": %w`, err)}
		}
	}
	return nil
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := uc.check(); err != nil {
		return nil, err
	}
	_node, _spec := uc.createSpec()
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	uc.mutation.id = &_node.ID
	uc.mutation.done = true
	return _node, nil
}

func (uc *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: uc.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	_spec.OnConflict = uc.conflict
	if id, ok := uc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := uc.mutation.CreatedAt(); ok {
		_spec.SetField(user.FieldCreatedAt, field.TypeInt64, value)
		_node.CreatedAt = value
	}
	if value, ok := uc.mutation.CreatedBy(); ok {
		_spec.SetField(user.FieldCreatedBy, field.TypeInt64, value)
		_node.CreatedBy = value
	}
	if value, ok := uc.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeInt64, value)
		_node.UpdatedAt = value
	}
	if value, ok := uc.mutation.UpdatedBy(); ok {
		_spec.SetField(user.FieldUpdatedBy, field.TypeInt64, value)
		_node.UpdatedBy = value
	}
	if value, ok := uc.mutation.DeletedAt(); ok {
		_spec.SetField(user.FieldDeletedAt, field.TypeInt64, value)
		_node.DeletedAt = value
	}
	if value, ok := uc.mutation.DeletedBy(); ok {
		_spec.SetField(user.FieldDeletedBy, field.TypeInt64, value)
		_node.DeletedBy = value
	}
	if value, ok := uc.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := uc.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
		_node.Avatar = value
	}
	if value, ok := uc.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := uc.mutation.Professional(); ok {
		_spec.SetField(user.FieldProfessional, field.TypeString, value)
		_node.Professional = value
	}
	if value, ok := uc.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
		_node.Address = value
	}
	if value, ok := uc.mutation.Skills(); ok {
		_spec.SetField(user.FieldSkills, field.TypeJSON, value)
		_node.Skills = value
	}
	if value, ok := uc.mutation.Description(); ok {
		_spec.SetField(user.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (uc *UserCreate) OnConflict(opts ...sql.ConflictOption) *UserUpsertOne {
	uc.conflict = opts
	return &UserUpsertOne{
		create: uc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uc *UserCreate) OnConflictColumns(columns ...string) *UserUpsertOne {
	uc.conflict = append(uc.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertOne{
		create: uc,
	}
}

type (
	// UserUpsertOne is the builder for "upsert"-ing
	//  one User node.
	UserUpsertOne struct {
		create *UserCreate
	}

	// UserUpsert is the "OnConflict" setter.
	UserUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsert) SetUpdatedAt(v int64) *UserUpsert {
	u.Set(user.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdatedAt() *UserUpsert {
	u.SetExcluded(user.FieldUpdatedAt)
	return u
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserUpsert) AddUpdatedAt(v int64) *UserUpsert {
	u.Add(user.FieldUpdatedAt, v)
	return u
}

// SetUpdatedBy sets the "updated_by" field.
func (u *UserUpsert) SetUpdatedBy(v int64) *UserUpsert {
	u.Set(user.FieldUpdatedBy, v)
	return u
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *UserUpsert) UpdateUpdatedBy() *UserUpsert {
	u.SetExcluded(user.FieldUpdatedBy)
	return u
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *UserUpsert) AddUpdatedBy(v int64) *UserUpsert {
	u.Add(user.FieldUpdatedBy, v)
	return u
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserUpsert) SetDeletedAt(v int64) *UserUpsert {
	u.Set(user.FieldDeletedAt, v)
	return u
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserUpsert) UpdateDeletedAt() *UserUpsert {
	u.SetExcluded(user.FieldDeletedAt)
	return u
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserUpsert) AddDeletedAt(v int64) *UserUpsert {
	u.Add(user.FieldDeletedAt, v)
	return u
}

// SetDeletedBy sets the "deleted_by" field.
func (u *UserUpsert) SetDeletedBy(v int64) *UserUpsert {
	u.Set(user.FieldDeletedBy, v)
	return u
}

// UpdateDeletedBy sets the "deleted_by" field to the value that was provided on create.
func (u *UserUpsert) UpdateDeletedBy() *UserUpsert {
	u.SetExcluded(user.FieldDeletedBy)
	return u
}

// AddDeletedBy adds v to the "deleted_by" field.
func (u *UserUpsert) AddDeletedBy(v int64) *UserUpsert {
	u.Add(user.FieldDeletedBy, v)
	return u
}

// SetName sets the "name" field.
func (u *UserUpsert) SetName(v string) *UserUpsert {
	u.Set(user.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsert) UpdateName() *UserUpsert {
	u.SetExcluded(user.FieldName)
	return u
}

// SetAvatar sets the "avatar" field.
func (u *UserUpsert) SetAvatar(v string) *UserUpsert {
	u.Set(user.FieldAvatar, v)
	return u
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *UserUpsert) UpdateAvatar() *UserUpsert {
	u.SetExcluded(user.FieldAvatar)
	return u
}

// SetEmail sets the "email" field.
func (u *UserUpsert) SetEmail(v string) *UserUpsert {
	u.Set(user.FieldEmail, v)
	return u
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsert) UpdateEmail() *UserUpsert {
	u.SetExcluded(user.FieldEmail)
	return u
}

// SetProfessional sets the "professional" field.
func (u *UserUpsert) SetProfessional(v string) *UserUpsert {
	u.Set(user.FieldProfessional, v)
	return u
}

// UpdateProfessional sets the "professional" field to the value that was provided on create.
func (u *UserUpsert) UpdateProfessional() *UserUpsert {
	u.SetExcluded(user.FieldProfessional)
	return u
}

// SetAddress sets the "address" field.
func (u *UserUpsert) SetAddress(v string) *UserUpsert {
	u.Set(user.FieldAddress, v)
	return u
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *UserUpsert) UpdateAddress() *UserUpsert {
	u.SetExcluded(user.FieldAddress)
	return u
}

// SetSkills sets the "skills" field.
func (u *UserUpsert) SetSkills(v []string) *UserUpsert {
	u.Set(user.FieldSkills, v)
	return u
}

// UpdateSkills sets the "skills" field to the value that was provided on create.
func (u *UserUpsert) UpdateSkills() *UserUpsert {
	u.SetExcluded(user.FieldSkills)
	return u
}

// SetDescription sets the "description" field.
func (u *UserUpsert) SetDescription(v string) *UserUpsert {
	u.Set(user.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *UserUpsert) UpdateDescription() *UserUpsert {
	u.SetExcluded(user.FieldDescription)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertOne) UpdateNewValues() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(user.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(user.FieldCreatedAt)
		}
		if _, exists := u.create.mutation.CreatedBy(); exists {
			s.SetIgnore(user.FieldCreatedBy)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserUpsertOne) Ignore() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertOne) DoNothing() *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreate.OnConflict
// documentation for more info.
func (u *UserUpsertOne) Update(set func(*UserUpsert)) *UserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertOne) SetUpdatedAt(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserUpsertOne) AddUpdatedAt(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdatedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *UserUpsertOne) SetUpdatedBy(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *UserUpsertOne) AddUpdatedBy(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateUpdatedBy() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserUpsertOne) SetDeletedAt(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserUpsertOne) AddDeletedAt(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDeletedAt() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeletedBy sets the "deleted_by" field.
func (u *UserUpsertOne) SetDeletedBy(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDeletedBy(v)
	})
}

// AddDeletedBy adds v to the "deleted_by" field.
func (u *UserUpsertOne) AddDeletedBy(v int64) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.AddDeletedBy(v)
	})
}

// UpdateDeletedBy sets the "deleted_by" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDeletedBy() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeletedBy()
	})
}

// SetName sets the "name" field.
func (u *UserUpsertOne) SetName(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateName() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetAvatar sets the "avatar" field.
func (u *UserUpsertOne) SetAvatar(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateAvatar() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAvatar()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertOne) SetEmail(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateEmail() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetProfessional sets the "professional" field.
func (u *UserUpsertOne) SetProfessional(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetProfessional(v)
	})
}

// UpdateProfessional sets the "professional" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateProfessional() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateProfessional()
	})
}

// SetAddress sets the "address" field.
func (u *UserUpsertOne) SetAddress(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateAddress() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAddress()
	})
}

// SetSkills sets the "skills" field.
func (u *UserUpsertOne) SetSkills(v []string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetSkills(v)
	})
}

// UpdateSkills sets the "skills" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateSkills() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateSkills()
	})
}

// SetDescription sets the "description" field.
func (u *UserUpsertOne) SetDescription(v string) *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *UserUpsertOne) UpdateDescription() *UserUpsertOne {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *UserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
	conflict []sql.ConflictOption
}

// Save creates the User entities in the database.
func (ucb *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if ucb.err != nil {
		return nil, ucb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ucb.builders))
	nodes := make([]*User, len(ucb.builders))
	mutators := make([]Mutator, len(ucb.builders))
	for i := range ucb.builders {
		func(i int, root context.Context) {
			builder := ucb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					_, err = mutators[i+1].Mutate(root, ucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ucb.driver, spec); err != nil {
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
					nodes[i].ID = int(id)
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
		if _, err := mutators[0].Mutate(ctx, ucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ucb *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := ucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ucb *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := ucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucb *UserCreateBulk) ExecX(ctx context.Context) {
	if err := ucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.User.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserUpsertBulk {
	ucb.conflict = opts
	return &UserUpsertBulk{
		create: ucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (ucb *UserCreateBulk) OnConflictColumns(columns ...string) *UserUpsertBulk {
	ucb.conflict = append(ucb.conflict, sql.ConflictColumns(columns...))
	return &UserUpsertBulk{
		create: ucb,
	}
}

// UserUpsertBulk is the builder for "upsert"-ing
// a bulk of User nodes.
type UserUpsertBulk struct {
	create *UserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(user.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *UserUpsertBulk) UpdateNewValues() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(user.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(user.FieldCreatedAt)
			}
			if _, exists := b.mutation.CreatedBy(); exists {
				s.SetIgnore(user.FieldCreatedBy)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.User.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserUpsertBulk) Ignore() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserUpsertBulk) DoNothing() *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserCreateBulk.OnConflict
// documentation for more info.
func (u *UserUpsertBulk) Update(set func(*UserUpsert)) *UserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *UserUpsertBulk) SetUpdatedAt(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedAt(v)
	})
}

// AddUpdatedAt adds v to the "updated_at" field.
func (u *UserUpsertBulk) AddUpdatedAt(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdatedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedAt()
	})
}

// SetUpdatedBy sets the "updated_by" field.
func (u *UserUpsertBulk) SetUpdatedBy(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetUpdatedBy(v)
	})
}

// AddUpdatedBy adds v to the "updated_by" field.
func (u *UserUpsertBulk) AddUpdatedBy(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddUpdatedBy(v)
	})
}

// UpdateUpdatedBy sets the "updated_by" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateUpdatedBy() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateUpdatedBy()
	})
}

// SetDeletedAt sets the "deleted_at" field.
func (u *UserUpsertBulk) SetDeletedAt(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDeletedAt(v)
	})
}

// AddDeletedAt adds v to the "deleted_at" field.
func (u *UserUpsertBulk) AddDeletedAt(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddDeletedAt(v)
	})
}

// UpdateDeletedAt sets the "deleted_at" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDeletedAt() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeletedAt()
	})
}

// SetDeletedBy sets the "deleted_by" field.
func (u *UserUpsertBulk) SetDeletedBy(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDeletedBy(v)
	})
}

// AddDeletedBy adds v to the "deleted_by" field.
func (u *UserUpsertBulk) AddDeletedBy(v int64) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.AddDeletedBy(v)
	})
}

// UpdateDeletedBy sets the "deleted_by" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDeletedBy() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDeletedBy()
	})
}

// SetName sets the "name" field.
func (u *UserUpsertBulk) SetName(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateName() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateName()
	})
}

// SetAvatar sets the "avatar" field.
func (u *UserUpsertBulk) SetAvatar(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetAvatar(v)
	})
}

// UpdateAvatar sets the "avatar" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateAvatar() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAvatar()
	})
}

// SetEmail sets the "email" field.
func (u *UserUpsertBulk) SetEmail(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetEmail(v)
	})
}

// UpdateEmail sets the "email" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateEmail() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateEmail()
	})
}

// SetProfessional sets the "professional" field.
func (u *UserUpsertBulk) SetProfessional(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetProfessional(v)
	})
}

// UpdateProfessional sets the "professional" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateProfessional() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateProfessional()
	})
}

// SetAddress sets the "address" field.
func (u *UserUpsertBulk) SetAddress(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetAddress(v)
	})
}

// UpdateAddress sets the "address" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateAddress() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateAddress()
	})
}

// SetSkills sets the "skills" field.
func (u *UserUpsertBulk) SetSkills(v []string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetSkills(v)
	})
}

// UpdateSkills sets the "skills" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateSkills() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateSkills()
	})
}

// SetDescription sets the "description" field.
func (u *UserUpsertBulk) SetDescription(v string) *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *UserUpsertBulk) UpdateDescription() *UserUpsertBulk {
	return u.Update(func(s *UserUpsert) {
		s.UpdateDescription()
	})
}

// Exec executes the query.
func (u *UserUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
