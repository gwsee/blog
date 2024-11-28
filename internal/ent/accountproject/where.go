// Code generated by ent, DO NOT EDIT.

package accountproject

import (
	"blog/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldDeletedBy, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldUserID, v))
}

// ExperienceID applies equality check predicate on the "experience_id" field. It's identical to ExperienceIDEQ.
func ExperienceID(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldExperienceID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldDescription, v))
}

// Start applies equality check predicate on the "start" field. It's identical to StartEQ.
func Start(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldStart, v))
}

// End applies equality check predicate on the "end" field. It's identical to EndEQ.
func End(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldEnd, v))
}

// Link applies equality check predicate on the "link" field. It's identical to LinkEQ.
func Link(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldLink, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldDeletedBy, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldUserID, v))
}

// ExperienceIDEQ applies the EQ predicate on the "experience_id" field.
func ExperienceIDEQ(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldExperienceID, v))
}

// ExperienceIDNEQ applies the NEQ predicate on the "experience_id" field.
func ExperienceIDNEQ(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldExperienceID, v))
}

// ExperienceIDIn applies the In predicate on the "experience_id" field.
func ExperienceIDIn(vs ...int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldExperienceID, vs...))
}

// ExperienceIDNotIn applies the NotIn predicate on the "experience_id" field.
func ExperienceIDNotIn(vs ...int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldExperienceID, vs...))
}

// ExperienceIDGT applies the GT predicate on the "experience_id" field.
func ExperienceIDGT(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldExperienceID, v))
}

// ExperienceIDGTE applies the GTE predicate on the "experience_id" field.
func ExperienceIDGTE(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldExperienceID, v))
}

// ExperienceIDLT applies the LT predicate on the "experience_id" field.
func ExperienceIDLT(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldExperienceID, v))
}

// ExperienceIDLTE applies the LTE predicate on the "experience_id" field.
func ExperienceIDLTE(v int) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldExperienceID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldContainsFold(FieldDescription, v))
}

// StartEQ applies the EQ predicate on the "start" field.
func StartEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldStart, v))
}

// StartNEQ applies the NEQ predicate on the "start" field.
func StartNEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldStart, v))
}

// StartIn applies the In predicate on the "start" field.
func StartIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldStart, vs...))
}

// StartNotIn applies the NotIn predicate on the "start" field.
func StartNotIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldStart, vs...))
}

// StartGT applies the GT predicate on the "start" field.
func StartGT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldStart, v))
}

// StartGTE applies the GTE predicate on the "start" field.
func StartGTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldStart, v))
}

// StartLT applies the LT predicate on the "start" field.
func StartLT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldStart, v))
}

// StartLTE applies the LTE predicate on the "start" field.
func StartLTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldStart, v))
}

// EndEQ applies the EQ predicate on the "end" field.
func EndEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldEnd, v))
}

// EndNEQ applies the NEQ predicate on the "end" field.
func EndNEQ(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldEnd, v))
}

// EndIn applies the In predicate on the "end" field.
func EndIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldEnd, vs...))
}

// EndNotIn applies the NotIn predicate on the "end" field.
func EndNotIn(vs ...int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldEnd, vs...))
}

// EndGT applies the GT predicate on the "end" field.
func EndGT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldEnd, v))
}

// EndGTE applies the GTE predicate on the "end" field.
func EndGTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldEnd, v))
}

// EndLT applies the LT predicate on the "end" field.
func EndLT(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldEnd, v))
}

// EndLTE applies the LTE predicate on the "end" field.
func EndLTE(v int64) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldEnd, v))
}

// LinkEQ applies the EQ predicate on the "link" field.
func LinkEQ(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEQ(FieldLink, v))
}

// LinkNEQ applies the NEQ predicate on the "link" field.
func LinkNEQ(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNEQ(FieldLink, v))
}

// LinkIn applies the In predicate on the "link" field.
func LinkIn(vs ...string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldIn(FieldLink, vs...))
}

// LinkNotIn applies the NotIn predicate on the "link" field.
func LinkNotIn(vs ...string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldNotIn(FieldLink, vs...))
}

// LinkGT applies the GT predicate on the "link" field.
func LinkGT(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGT(FieldLink, v))
}

// LinkGTE applies the GTE predicate on the "link" field.
func LinkGTE(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldGTE(FieldLink, v))
}

// LinkLT applies the LT predicate on the "link" field.
func LinkLT(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLT(FieldLink, v))
}

// LinkLTE applies the LTE predicate on the "link" field.
func LinkLTE(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldLTE(FieldLink, v))
}

// LinkContains applies the Contains predicate on the "link" field.
func LinkContains(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldContains(FieldLink, v))
}

// LinkHasPrefix applies the HasPrefix predicate on the "link" field.
func LinkHasPrefix(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldHasPrefix(FieldLink, v))
}

// LinkHasSuffix applies the HasSuffix predicate on the "link" field.
func LinkHasSuffix(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldHasSuffix(FieldLink, v))
}

// LinkEqualFold applies the EqualFold predicate on the "link" field.
func LinkEqualFold(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldEqualFold(FieldLink, v))
}

// LinkContainsFold applies the ContainsFold predicate on the "link" field.
func LinkContainsFold(v string) predicate.AccountProject {
	return predicate.AccountProject(sql.FieldContainsFold(FieldLink, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.AccountProject) predicate.AccountProject {
	return predicate.AccountProject(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.AccountProject) predicate.AccountProject {
	return predicate.AccountProject(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.AccountProject) predicate.AccountProject {
	return predicate.AccountProject(sql.NotPredicates(p))
}