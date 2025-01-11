// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/predicate"
	"blog/internal/ent/userexperience"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
)

// UserExperienceUpdate is the builder for updating UserExperience entities.
type UserExperienceUpdate struct {
	config
	hooks    []Hook
	mutation *UserExperienceMutation
}

// Where appends a list predicates to the UserExperienceUpdate builder.
func (ueu *UserExperienceUpdate) Where(ps ...predicate.UserExperience) *UserExperienceUpdate {
	ueu.mutation.Where(ps...)
	return ueu
}

// SetUpdatedAt sets the "updated_at" field.
func (ueu *UserExperienceUpdate) SetUpdatedAt(i int64) *UserExperienceUpdate {
	ueu.mutation.ResetUpdatedAt()
	ueu.mutation.SetUpdatedAt(i)
	return ueu
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ueu *UserExperienceUpdate) AddUpdatedAt(i int64) *UserExperienceUpdate {
	ueu.mutation.AddUpdatedAt(i)
	return ueu
}

// SetUpdatedBy sets the "updated_by" field.
func (ueu *UserExperienceUpdate) SetUpdatedBy(i int64) *UserExperienceUpdate {
	ueu.mutation.ResetUpdatedBy()
	ueu.mutation.SetUpdatedBy(i)
	return ueu
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableUpdatedBy(i *int64) *UserExperienceUpdate {
	if i != nil {
		ueu.SetUpdatedBy(*i)
	}
	return ueu
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ueu *UserExperienceUpdate) AddUpdatedBy(i int64) *UserExperienceUpdate {
	ueu.mutation.AddUpdatedBy(i)
	return ueu
}

// SetDeletedAt sets the "deleted_at" field.
func (ueu *UserExperienceUpdate) SetDeletedAt(i int64) *UserExperienceUpdate {
	ueu.mutation.ResetDeletedAt()
	ueu.mutation.SetDeletedAt(i)
	return ueu
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableDeletedAt(i *int64) *UserExperienceUpdate {
	if i != nil {
		ueu.SetDeletedAt(*i)
	}
	return ueu
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ueu *UserExperienceUpdate) AddDeletedAt(i int64) *UserExperienceUpdate {
	ueu.mutation.AddDeletedAt(i)
	return ueu
}

// SetDeletedBy sets the "deleted_by" field.
func (ueu *UserExperienceUpdate) SetDeletedBy(i int64) *UserExperienceUpdate {
	ueu.mutation.ResetDeletedBy()
	ueu.mutation.SetDeletedBy(i)
	return ueu
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableDeletedBy(i *int64) *UserExperienceUpdate {
	if i != nil {
		ueu.SetDeletedBy(*i)
	}
	return ueu
}

// AddDeletedBy adds i to the "deleted_by" field.
func (ueu *UserExperienceUpdate) AddDeletedBy(i int64) *UserExperienceUpdate {
	ueu.mutation.AddDeletedBy(i)
	return ueu
}

// SetUserID sets the "user_id" field.
func (ueu *UserExperienceUpdate) SetUserID(i int) *UserExperienceUpdate {
	ueu.mutation.ResetUserID()
	ueu.mutation.SetUserID(i)
	return ueu
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableUserID(i *int) *UserExperienceUpdate {
	if i != nil {
		ueu.SetUserID(*i)
	}
	return ueu
}

// AddUserID adds i to the "user_id" field.
func (ueu *UserExperienceUpdate) AddUserID(i int) *UserExperienceUpdate {
	ueu.mutation.AddUserID(i)
	return ueu
}

// SetCompany sets the "company" field.
func (ueu *UserExperienceUpdate) SetCompany(s string) *UserExperienceUpdate {
	ueu.mutation.SetCompany(s)
	return ueu
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableCompany(s *string) *UserExperienceUpdate {
	if s != nil {
		ueu.SetCompany(*s)
	}
	return ueu
}

// SetRole sets the "role" field.
func (ueu *UserExperienceUpdate) SetRole(s string) *UserExperienceUpdate {
	ueu.mutation.SetRole(s)
	return ueu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableRole(s *string) *UserExperienceUpdate {
	if s != nil {
		ueu.SetRole(*s)
	}
	return ueu
}

// SetLocation sets the "location" field.
func (ueu *UserExperienceUpdate) SetLocation(s string) *UserExperienceUpdate {
	ueu.mutation.SetLocation(s)
	return ueu
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableLocation(s *string) *UserExperienceUpdate {
	if s != nil {
		ueu.SetLocation(*s)
	}
	return ueu
}

// SetStart sets the "start" field.
func (ueu *UserExperienceUpdate) SetStart(i int64) *UserExperienceUpdate {
	ueu.mutation.ResetStart()
	ueu.mutation.SetStart(i)
	return ueu
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableStart(i *int64) *UserExperienceUpdate {
	if i != nil {
		ueu.SetStart(*i)
	}
	return ueu
}

// AddStart adds i to the "start" field.
func (ueu *UserExperienceUpdate) AddStart(i int64) *UserExperienceUpdate {
	ueu.mutation.AddStart(i)
	return ueu
}

// SetEnd sets the "end" field.
func (ueu *UserExperienceUpdate) SetEnd(i int64) *UserExperienceUpdate {
	ueu.mutation.ResetEnd()
	ueu.mutation.SetEnd(i)
	return ueu
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableEnd(i *int64) *UserExperienceUpdate {
	if i != nil {
		ueu.SetEnd(*i)
	}
	return ueu
}

// AddEnd adds i to the "end" field.
func (ueu *UserExperienceUpdate) AddEnd(i int64) *UserExperienceUpdate {
	ueu.mutation.AddEnd(i)
	return ueu
}

// SetDescription sets the "description" field.
func (ueu *UserExperienceUpdate) SetDescription(s string) *UserExperienceUpdate {
	ueu.mutation.SetDescription(s)
	return ueu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableDescription(s *string) *UserExperienceUpdate {
	if s != nil {
		ueu.SetDescription(*s)
	}
	return ueu
}

// SetResponsibilities sets the "responsibilities" field.
func (ueu *UserExperienceUpdate) SetResponsibilities(s string) *UserExperienceUpdate {
	ueu.mutation.SetResponsibilities(s)
	return ueu
}

// SetNillableResponsibilities sets the "responsibilities" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableResponsibilities(s *string) *UserExperienceUpdate {
	if s != nil {
		ueu.SetResponsibilities(*s)
	}
	return ueu
}

// SetAchievements sets the "achievements" field.
func (ueu *UserExperienceUpdate) SetAchievements(s string) *UserExperienceUpdate {
	ueu.mutation.SetAchievements(s)
	return ueu
}

// SetNillableAchievements sets the "achievements" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableAchievements(s *string) *UserExperienceUpdate {
	if s != nil {
		ueu.SetAchievements(*s)
	}
	return ueu
}

// SetSkills sets the "skills" field.
func (ueu *UserExperienceUpdate) SetSkills(s []string) *UserExperienceUpdate {
	ueu.mutation.SetSkills(s)
	return ueu
}

// AppendSkills appends s to the "skills" field.
func (ueu *UserExperienceUpdate) AppendSkills(s []string) *UserExperienceUpdate {
	ueu.mutation.AppendSkills(s)
	return ueu
}

// SetProject sets the "project" field.
func (ueu *UserExperienceUpdate) SetProject(i int) *UserExperienceUpdate {
	ueu.mutation.ResetProject()
	ueu.mutation.SetProject(i)
	return ueu
}

// SetNillableProject sets the "project" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableProject(i *int) *UserExperienceUpdate {
	if i != nil {
		ueu.SetProject(*i)
	}
	return ueu
}

// AddProject adds i to the "project" field.
func (ueu *UserExperienceUpdate) AddProject(i int) *UserExperienceUpdate {
	ueu.mutation.AddProject(i)
	return ueu
}

// SetImage sets the "image" field.
func (ueu *UserExperienceUpdate) SetImage(s string) *UserExperienceUpdate {
	ueu.mutation.SetImage(s)
	return ueu
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (ueu *UserExperienceUpdate) SetNillableImage(s *string) *UserExperienceUpdate {
	if s != nil {
		ueu.SetImage(*s)
	}
	return ueu
}

// Mutation returns the UserExperienceMutation object of the builder.
func (ueu *UserExperienceUpdate) Mutation() *UserExperienceMutation {
	return ueu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ueu *UserExperienceUpdate) Save(ctx context.Context) (int, error) {
	if err := ueu.defaults(); err != nil {
		return 0, err
	}
	return withHooks(ctx, ueu.sqlSave, ueu.mutation, ueu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ueu *UserExperienceUpdate) SaveX(ctx context.Context) int {
	affected, err := ueu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ueu *UserExperienceUpdate) Exec(ctx context.Context) error {
	_, err := ueu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ueu *UserExperienceUpdate) ExecX(ctx context.Context) {
	if err := ueu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ueu *UserExperienceUpdate) defaults() error {
	if _, ok := ueu.mutation.UpdatedAt(); !ok {
		if userexperience.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userexperience.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userexperience.UpdateDefaultUpdatedAt()
		ueu.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ueu *UserExperienceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(userexperience.Table, userexperience.Columns, sqlgraph.NewFieldSpec(userexperience.FieldID, field.TypeInt))
	if ps := ueu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ueu.mutation.UpdatedAt(); ok {
		_spec.SetField(userexperience.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(userexperience.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.UpdatedBy(); ok {
		_spec.SetField(userexperience.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userexperience.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.DeletedAt(); ok {
		_spec.SetField(userexperience.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.AddedDeletedAt(); ok {
		_spec.AddField(userexperience.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.DeletedBy(); ok {
		_spec.SetField(userexperience.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.AddedDeletedBy(); ok {
		_spec.AddField(userexperience.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.UserID(); ok {
		_spec.SetField(userexperience.FieldUserID, field.TypeInt, value)
	}
	if value, ok := ueu.mutation.AddedUserID(); ok {
		_spec.AddField(userexperience.FieldUserID, field.TypeInt, value)
	}
	if value, ok := ueu.mutation.Company(); ok {
		_spec.SetField(userexperience.FieldCompany, field.TypeString, value)
	}
	if value, ok := ueu.mutation.Role(); ok {
		_spec.SetField(userexperience.FieldRole, field.TypeString, value)
	}
	if value, ok := ueu.mutation.Location(); ok {
		_spec.SetField(userexperience.FieldLocation, field.TypeString, value)
	}
	if value, ok := ueu.mutation.Start(); ok {
		_spec.SetField(userexperience.FieldStart, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.AddedStart(); ok {
		_spec.AddField(userexperience.FieldStart, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.End(); ok {
		_spec.SetField(userexperience.FieldEnd, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.AddedEnd(); ok {
		_spec.AddField(userexperience.FieldEnd, field.TypeInt64, value)
	}
	if value, ok := ueu.mutation.Description(); ok {
		_spec.SetField(userexperience.FieldDescription, field.TypeString, value)
	}
	if value, ok := ueu.mutation.Responsibilities(); ok {
		_spec.SetField(userexperience.FieldResponsibilities, field.TypeString, value)
	}
	if value, ok := ueu.mutation.Achievements(); ok {
		_spec.SetField(userexperience.FieldAchievements, field.TypeString, value)
	}
	if value, ok := ueu.mutation.Skills(); ok {
		_spec.SetField(userexperience.FieldSkills, field.TypeJSON, value)
	}
	if value, ok := ueu.mutation.AppendedSkills(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, userexperience.FieldSkills, value)
		})
	}
	if value, ok := ueu.mutation.Project(); ok {
		_spec.SetField(userexperience.FieldProject, field.TypeInt, value)
	}
	if value, ok := ueu.mutation.AddedProject(); ok {
		_spec.AddField(userexperience.FieldProject, field.TypeInt, value)
	}
	if value, ok := ueu.mutation.Image(); ok {
		_spec.SetField(userexperience.FieldImage, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ueu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userexperience.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ueu.mutation.done = true
	return n, nil
}

// UserExperienceUpdateOne is the builder for updating a single UserExperience entity.
type UserExperienceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserExperienceMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ueuo *UserExperienceUpdateOne) SetUpdatedAt(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.ResetUpdatedAt()
	ueuo.mutation.SetUpdatedAt(i)
	return ueuo
}

// AddUpdatedAt adds i to the "updated_at" field.
func (ueuo *UserExperienceUpdateOne) AddUpdatedAt(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.AddUpdatedAt(i)
	return ueuo
}

// SetUpdatedBy sets the "updated_by" field.
func (ueuo *UserExperienceUpdateOne) SetUpdatedBy(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.ResetUpdatedBy()
	ueuo.mutation.SetUpdatedBy(i)
	return ueuo
}

// SetNillableUpdatedBy sets the "updated_by" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableUpdatedBy(i *int64) *UserExperienceUpdateOne {
	if i != nil {
		ueuo.SetUpdatedBy(*i)
	}
	return ueuo
}

// AddUpdatedBy adds i to the "updated_by" field.
func (ueuo *UserExperienceUpdateOne) AddUpdatedBy(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.AddUpdatedBy(i)
	return ueuo
}

// SetDeletedAt sets the "deleted_at" field.
func (ueuo *UserExperienceUpdateOne) SetDeletedAt(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.ResetDeletedAt()
	ueuo.mutation.SetDeletedAt(i)
	return ueuo
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableDeletedAt(i *int64) *UserExperienceUpdateOne {
	if i != nil {
		ueuo.SetDeletedAt(*i)
	}
	return ueuo
}

// AddDeletedAt adds i to the "deleted_at" field.
func (ueuo *UserExperienceUpdateOne) AddDeletedAt(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.AddDeletedAt(i)
	return ueuo
}

// SetDeletedBy sets the "deleted_by" field.
func (ueuo *UserExperienceUpdateOne) SetDeletedBy(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.ResetDeletedBy()
	ueuo.mutation.SetDeletedBy(i)
	return ueuo
}

// SetNillableDeletedBy sets the "deleted_by" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableDeletedBy(i *int64) *UserExperienceUpdateOne {
	if i != nil {
		ueuo.SetDeletedBy(*i)
	}
	return ueuo
}

// AddDeletedBy adds i to the "deleted_by" field.
func (ueuo *UserExperienceUpdateOne) AddDeletedBy(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.AddDeletedBy(i)
	return ueuo
}

// SetUserID sets the "user_id" field.
func (ueuo *UserExperienceUpdateOne) SetUserID(i int) *UserExperienceUpdateOne {
	ueuo.mutation.ResetUserID()
	ueuo.mutation.SetUserID(i)
	return ueuo
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableUserID(i *int) *UserExperienceUpdateOne {
	if i != nil {
		ueuo.SetUserID(*i)
	}
	return ueuo
}

// AddUserID adds i to the "user_id" field.
func (ueuo *UserExperienceUpdateOne) AddUserID(i int) *UserExperienceUpdateOne {
	ueuo.mutation.AddUserID(i)
	return ueuo
}

// SetCompany sets the "company" field.
func (ueuo *UserExperienceUpdateOne) SetCompany(s string) *UserExperienceUpdateOne {
	ueuo.mutation.SetCompany(s)
	return ueuo
}

// SetNillableCompany sets the "company" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableCompany(s *string) *UserExperienceUpdateOne {
	if s != nil {
		ueuo.SetCompany(*s)
	}
	return ueuo
}

// SetRole sets the "role" field.
func (ueuo *UserExperienceUpdateOne) SetRole(s string) *UserExperienceUpdateOne {
	ueuo.mutation.SetRole(s)
	return ueuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableRole(s *string) *UserExperienceUpdateOne {
	if s != nil {
		ueuo.SetRole(*s)
	}
	return ueuo
}

// SetLocation sets the "location" field.
func (ueuo *UserExperienceUpdateOne) SetLocation(s string) *UserExperienceUpdateOne {
	ueuo.mutation.SetLocation(s)
	return ueuo
}

// SetNillableLocation sets the "location" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableLocation(s *string) *UserExperienceUpdateOne {
	if s != nil {
		ueuo.SetLocation(*s)
	}
	return ueuo
}

// SetStart sets the "start" field.
func (ueuo *UserExperienceUpdateOne) SetStart(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.ResetStart()
	ueuo.mutation.SetStart(i)
	return ueuo
}

// SetNillableStart sets the "start" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableStart(i *int64) *UserExperienceUpdateOne {
	if i != nil {
		ueuo.SetStart(*i)
	}
	return ueuo
}

// AddStart adds i to the "start" field.
func (ueuo *UserExperienceUpdateOne) AddStart(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.AddStart(i)
	return ueuo
}

// SetEnd sets the "end" field.
func (ueuo *UserExperienceUpdateOne) SetEnd(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.ResetEnd()
	ueuo.mutation.SetEnd(i)
	return ueuo
}

// SetNillableEnd sets the "end" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableEnd(i *int64) *UserExperienceUpdateOne {
	if i != nil {
		ueuo.SetEnd(*i)
	}
	return ueuo
}

// AddEnd adds i to the "end" field.
func (ueuo *UserExperienceUpdateOne) AddEnd(i int64) *UserExperienceUpdateOne {
	ueuo.mutation.AddEnd(i)
	return ueuo
}

// SetDescription sets the "description" field.
func (ueuo *UserExperienceUpdateOne) SetDescription(s string) *UserExperienceUpdateOne {
	ueuo.mutation.SetDescription(s)
	return ueuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableDescription(s *string) *UserExperienceUpdateOne {
	if s != nil {
		ueuo.SetDescription(*s)
	}
	return ueuo
}

// SetResponsibilities sets the "responsibilities" field.
func (ueuo *UserExperienceUpdateOne) SetResponsibilities(s string) *UserExperienceUpdateOne {
	ueuo.mutation.SetResponsibilities(s)
	return ueuo
}

// SetNillableResponsibilities sets the "responsibilities" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableResponsibilities(s *string) *UserExperienceUpdateOne {
	if s != nil {
		ueuo.SetResponsibilities(*s)
	}
	return ueuo
}

// SetAchievements sets the "achievements" field.
func (ueuo *UserExperienceUpdateOne) SetAchievements(s string) *UserExperienceUpdateOne {
	ueuo.mutation.SetAchievements(s)
	return ueuo
}

// SetNillableAchievements sets the "achievements" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableAchievements(s *string) *UserExperienceUpdateOne {
	if s != nil {
		ueuo.SetAchievements(*s)
	}
	return ueuo
}

// SetSkills sets the "skills" field.
func (ueuo *UserExperienceUpdateOne) SetSkills(s []string) *UserExperienceUpdateOne {
	ueuo.mutation.SetSkills(s)
	return ueuo
}

// AppendSkills appends s to the "skills" field.
func (ueuo *UserExperienceUpdateOne) AppendSkills(s []string) *UserExperienceUpdateOne {
	ueuo.mutation.AppendSkills(s)
	return ueuo
}

// SetProject sets the "project" field.
func (ueuo *UserExperienceUpdateOne) SetProject(i int) *UserExperienceUpdateOne {
	ueuo.mutation.ResetProject()
	ueuo.mutation.SetProject(i)
	return ueuo
}

// SetNillableProject sets the "project" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableProject(i *int) *UserExperienceUpdateOne {
	if i != nil {
		ueuo.SetProject(*i)
	}
	return ueuo
}

// AddProject adds i to the "project" field.
func (ueuo *UserExperienceUpdateOne) AddProject(i int) *UserExperienceUpdateOne {
	ueuo.mutation.AddProject(i)
	return ueuo
}

// SetImage sets the "image" field.
func (ueuo *UserExperienceUpdateOne) SetImage(s string) *UserExperienceUpdateOne {
	ueuo.mutation.SetImage(s)
	return ueuo
}

// SetNillableImage sets the "image" field if the given value is not nil.
func (ueuo *UserExperienceUpdateOne) SetNillableImage(s *string) *UserExperienceUpdateOne {
	if s != nil {
		ueuo.SetImage(*s)
	}
	return ueuo
}

// Mutation returns the UserExperienceMutation object of the builder.
func (ueuo *UserExperienceUpdateOne) Mutation() *UserExperienceMutation {
	return ueuo.mutation
}

// Where appends a list predicates to the UserExperienceUpdate builder.
func (ueuo *UserExperienceUpdateOne) Where(ps ...predicate.UserExperience) *UserExperienceUpdateOne {
	ueuo.mutation.Where(ps...)
	return ueuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ueuo *UserExperienceUpdateOne) Select(field string, fields ...string) *UserExperienceUpdateOne {
	ueuo.fields = append([]string{field}, fields...)
	return ueuo
}

// Save executes the query and returns the updated UserExperience entity.
func (ueuo *UserExperienceUpdateOne) Save(ctx context.Context) (*UserExperience, error) {
	if err := ueuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, ueuo.sqlSave, ueuo.mutation, ueuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ueuo *UserExperienceUpdateOne) SaveX(ctx context.Context) *UserExperience {
	node, err := ueuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ueuo *UserExperienceUpdateOne) Exec(ctx context.Context) error {
	_, err := ueuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ueuo *UserExperienceUpdateOne) ExecX(ctx context.Context) {
	if err := ueuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ueuo *UserExperienceUpdateOne) defaults() error {
	if _, ok := ueuo.mutation.UpdatedAt(); !ok {
		if userexperience.UpdateDefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized userexperience.UpdateDefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := userexperience.UpdateDefaultUpdatedAt()
		ueuo.mutation.SetUpdatedAt(v)
	}
	return nil
}

func (ueuo *UserExperienceUpdateOne) sqlSave(ctx context.Context) (_node *UserExperience, err error) {
	_spec := sqlgraph.NewUpdateSpec(userexperience.Table, userexperience.Columns, sqlgraph.NewFieldSpec(userexperience.FieldID, field.TypeInt))
	id, ok := ueuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserExperience.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ueuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userexperience.FieldID)
		for _, f := range fields {
			if !userexperience.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != userexperience.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ueuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ueuo.mutation.UpdatedAt(); ok {
		_spec.SetField(userexperience.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.AddedUpdatedAt(); ok {
		_spec.AddField(userexperience.FieldUpdatedAt, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.UpdatedBy(); ok {
		_spec.SetField(userexperience.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.AddedUpdatedBy(); ok {
		_spec.AddField(userexperience.FieldUpdatedBy, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.DeletedAt(); ok {
		_spec.SetField(userexperience.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.AddedDeletedAt(); ok {
		_spec.AddField(userexperience.FieldDeletedAt, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.DeletedBy(); ok {
		_spec.SetField(userexperience.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.AddedDeletedBy(); ok {
		_spec.AddField(userexperience.FieldDeletedBy, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.UserID(); ok {
		_spec.SetField(userexperience.FieldUserID, field.TypeInt, value)
	}
	if value, ok := ueuo.mutation.AddedUserID(); ok {
		_spec.AddField(userexperience.FieldUserID, field.TypeInt, value)
	}
	if value, ok := ueuo.mutation.Company(); ok {
		_spec.SetField(userexperience.FieldCompany, field.TypeString, value)
	}
	if value, ok := ueuo.mutation.Role(); ok {
		_spec.SetField(userexperience.FieldRole, field.TypeString, value)
	}
	if value, ok := ueuo.mutation.Location(); ok {
		_spec.SetField(userexperience.FieldLocation, field.TypeString, value)
	}
	if value, ok := ueuo.mutation.Start(); ok {
		_spec.SetField(userexperience.FieldStart, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.AddedStart(); ok {
		_spec.AddField(userexperience.FieldStart, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.End(); ok {
		_spec.SetField(userexperience.FieldEnd, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.AddedEnd(); ok {
		_spec.AddField(userexperience.FieldEnd, field.TypeInt64, value)
	}
	if value, ok := ueuo.mutation.Description(); ok {
		_spec.SetField(userexperience.FieldDescription, field.TypeString, value)
	}
	if value, ok := ueuo.mutation.Responsibilities(); ok {
		_spec.SetField(userexperience.FieldResponsibilities, field.TypeString, value)
	}
	if value, ok := ueuo.mutation.Achievements(); ok {
		_spec.SetField(userexperience.FieldAchievements, field.TypeString, value)
	}
	if value, ok := ueuo.mutation.Skills(); ok {
		_spec.SetField(userexperience.FieldSkills, field.TypeJSON, value)
	}
	if value, ok := ueuo.mutation.AppendedSkills(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, userexperience.FieldSkills, value)
		})
	}
	if value, ok := ueuo.mutation.Project(); ok {
		_spec.SetField(userexperience.FieldProject, field.TypeInt, value)
	}
	if value, ok := ueuo.mutation.AddedProject(); ok {
		_spec.AddField(userexperience.FieldProject, field.TypeInt, value)
	}
	if value, ok := ueuo.mutation.Image(); ok {
		_spec.SetField(userexperience.FieldImage, field.TypeString, value)
	}
	_node = &UserExperience{config: ueuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ueuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userexperience.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ueuo.mutation.done = true
	return _node, nil
}
