// Code generated by ent, DO NOT EDIT.

package travel

import (
	"blog/internal/ent/predicate"

	"entgo.io/ent/dialect/sql"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedBy applies equality check predicate on the "created_by" field. It's identical to CreatedByEQ.
func CreatedBy(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldCreatedBy, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedBy applies equality check predicate on the "updated_by" field. It's identical to UpdatedByEQ.
func UpdatedBy(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldUpdatedBy, v))
}

// DeletedAt applies equality check predicate on the "deleted_at" field. It's identical to DeletedAtEQ.
func DeletedAt(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedBy applies equality check predicate on the "deleted_by" field. It's identical to DeletedByEQ.
func DeletedBy(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldDeletedBy, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldDescription, v))
}

// Video applies equality check predicate on the "video" field. It's identical to VideoEQ.
func Video(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldVideo, v))
}

// BrowseNum applies equality check predicate on the "browse_num" field. It's identical to BrowseNumEQ.
func BrowseNum(v int) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldBrowseNum, v))
}

// ThumbNum applies equality check predicate on the "thumb_num" field. It's identical to ThumbNumEQ.
func ThumbNum(v int) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldThumbNum, v))
}

// CollectNum applies equality check predicate on the "collect_num" field. It's identical to CollectNumEQ.
func CollectNum(v int) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldCollectNum, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldCreatedAt, v))
}

// CreatedByEQ applies the EQ predicate on the "created_by" field.
func CreatedByEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldCreatedBy, v))
}

// CreatedByNEQ applies the NEQ predicate on the "created_by" field.
func CreatedByNEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldCreatedBy, v))
}

// CreatedByIn applies the In predicate on the "created_by" field.
func CreatedByIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldCreatedBy, vs...))
}

// CreatedByNotIn applies the NotIn predicate on the "created_by" field.
func CreatedByNotIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldCreatedBy, vs...))
}

// CreatedByGT applies the GT predicate on the "created_by" field.
func CreatedByGT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldCreatedBy, v))
}

// CreatedByGTE applies the GTE predicate on the "created_by" field.
func CreatedByGTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldCreatedBy, v))
}

// CreatedByLT applies the LT predicate on the "created_by" field.
func CreatedByLT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldCreatedBy, v))
}

// CreatedByLTE applies the LTE predicate on the "created_by" field.
func CreatedByLTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldCreatedBy, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldUpdatedAt, v))
}

// UpdatedByEQ applies the EQ predicate on the "updated_by" field.
func UpdatedByEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldUpdatedBy, v))
}

// UpdatedByNEQ applies the NEQ predicate on the "updated_by" field.
func UpdatedByNEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldUpdatedBy, v))
}

// UpdatedByIn applies the In predicate on the "updated_by" field.
func UpdatedByIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldUpdatedBy, vs...))
}

// UpdatedByNotIn applies the NotIn predicate on the "updated_by" field.
func UpdatedByNotIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldUpdatedBy, vs...))
}

// UpdatedByGT applies the GT predicate on the "updated_by" field.
func UpdatedByGT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldUpdatedBy, v))
}

// UpdatedByGTE applies the GTE predicate on the "updated_by" field.
func UpdatedByGTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldUpdatedBy, v))
}

// UpdatedByLT applies the LT predicate on the "updated_by" field.
func UpdatedByLT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldUpdatedBy, v))
}

// UpdatedByLTE applies the LTE predicate on the "updated_by" field.
func UpdatedByLTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldUpdatedBy, v))
}

// DeletedAtEQ applies the EQ predicate on the "deleted_at" field.
func DeletedAtEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldDeletedAt, v))
}

// DeletedAtNEQ applies the NEQ predicate on the "deleted_at" field.
func DeletedAtNEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldDeletedAt, v))
}

// DeletedAtIn applies the In predicate on the "deleted_at" field.
func DeletedAtIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldDeletedAt, vs...))
}

// DeletedAtNotIn applies the NotIn predicate on the "deleted_at" field.
func DeletedAtNotIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldDeletedAt, vs...))
}

// DeletedAtGT applies the GT predicate on the "deleted_at" field.
func DeletedAtGT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldDeletedAt, v))
}

// DeletedAtGTE applies the GTE predicate on the "deleted_at" field.
func DeletedAtGTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldDeletedAt, v))
}

// DeletedAtLT applies the LT predicate on the "deleted_at" field.
func DeletedAtLT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldDeletedAt, v))
}

// DeletedAtLTE applies the LTE predicate on the "deleted_at" field.
func DeletedAtLTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldDeletedAt, v))
}

// DeletedByEQ applies the EQ predicate on the "deleted_by" field.
func DeletedByEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldDeletedBy, v))
}

// DeletedByNEQ applies the NEQ predicate on the "deleted_by" field.
func DeletedByNEQ(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldDeletedBy, v))
}

// DeletedByIn applies the In predicate on the "deleted_by" field.
func DeletedByIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldDeletedBy, vs...))
}

// DeletedByNotIn applies the NotIn predicate on the "deleted_by" field.
func DeletedByNotIn(vs ...int64) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldDeletedBy, vs...))
}

// DeletedByGT applies the GT predicate on the "deleted_by" field.
func DeletedByGT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldDeletedBy, v))
}

// DeletedByGTE applies the GTE predicate on the "deleted_by" field.
func DeletedByGTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldDeletedBy, v))
}

// DeletedByLT applies the LT predicate on the "deleted_by" field.
func DeletedByLT(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldDeletedBy, v))
}

// DeletedByLTE applies the LTE predicate on the "deleted_by" field.
func DeletedByLTE(v int64) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldDeletedBy, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Travel {
	return predicate.Travel(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Travel {
	return predicate.Travel(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Travel {
	return predicate.Travel(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Travel {
	return predicate.Travel(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Travel {
	return predicate.Travel(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Travel {
	return predicate.Travel(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Travel {
	return predicate.Travel(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Travel {
	return predicate.Travel(sql.FieldContainsFold(FieldDescription, v))
}

// VideoEQ applies the EQ predicate on the "video" field.
func VideoEQ(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldVideo, v))
}

// VideoNEQ applies the NEQ predicate on the "video" field.
func VideoNEQ(v string) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldVideo, v))
}

// VideoIn applies the In predicate on the "video" field.
func VideoIn(vs ...string) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldVideo, vs...))
}

// VideoNotIn applies the NotIn predicate on the "video" field.
func VideoNotIn(vs ...string) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldVideo, vs...))
}

// VideoGT applies the GT predicate on the "video" field.
func VideoGT(v string) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldVideo, v))
}

// VideoGTE applies the GTE predicate on the "video" field.
func VideoGTE(v string) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldVideo, v))
}

// VideoLT applies the LT predicate on the "video" field.
func VideoLT(v string) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldVideo, v))
}

// VideoLTE applies the LTE predicate on the "video" field.
func VideoLTE(v string) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldVideo, v))
}

// VideoContains applies the Contains predicate on the "video" field.
func VideoContains(v string) predicate.Travel {
	return predicate.Travel(sql.FieldContains(FieldVideo, v))
}

// VideoHasPrefix applies the HasPrefix predicate on the "video" field.
func VideoHasPrefix(v string) predicate.Travel {
	return predicate.Travel(sql.FieldHasPrefix(FieldVideo, v))
}

// VideoHasSuffix applies the HasSuffix predicate on the "video" field.
func VideoHasSuffix(v string) predicate.Travel {
	return predicate.Travel(sql.FieldHasSuffix(FieldVideo, v))
}

// VideoEqualFold applies the EqualFold predicate on the "video" field.
func VideoEqualFold(v string) predicate.Travel {
	return predicate.Travel(sql.FieldEqualFold(FieldVideo, v))
}

// VideoContainsFold applies the ContainsFold predicate on the "video" field.
func VideoContainsFold(v string) predicate.Travel {
	return predicate.Travel(sql.FieldContainsFold(FieldVideo, v))
}

// BrowseNumEQ applies the EQ predicate on the "browse_num" field.
func BrowseNumEQ(v int) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldBrowseNum, v))
}

// BrowseNumNEQ applies the NEQ predicate on the "browse_num" field.
func BrowseNumNEQ(v int) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldBrowseNum, v))
}

// BrowseNumIn applies the In predicate on the "browse_num" field.
func BrowseNumIn(vs ...int) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldBrowseNum, vs...))
}

// BrowseNumNotIn applies the NotIn predicate on the "browse_num" field.
func BrowseNumNotIn(vs ...int) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldBrowseNum, vs...))
}

// BrowseNumGT applies the GT predicate on the "browse_num" field.
func BrowseNumGT(v int) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldBrowseNum, v))
}

// BrowseNumGTE applies the GTE predicate on the "browse_num" field.
func BrowseNumGTE(v int) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldBrowseNum, v))
}

// BrowseNumLT applies the LT predicate on the "browse_num" field.
func BrowseNumLT(v int) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldBrowseNum, v))
}

// BrowseNumLTE applies the LTE predicate on the "browse_num" field.
func BrowseNumLTE(v int) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldBrowseNum, v))
}

// ThumbNumEQ applies the EQ predicate on the "thumb_num" field.
func ThumbNumEQ(v int) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldThumbNum, v))
}

// ThumbNumNEQ applies the NEQ predicate on the "thumb_num" field.
func ThumbNumNEQ(v int) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldThumbNum, v))
}

// ThumbNumIn applies the In predicate on the "thumb_num" field.
func ThumbNumIn(vs ...int) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldThumbNum, vs...))
}

// ThumbNumNotIn applies the NotIn predicate on the "thumb_num" field.
func ThumbNumNotIn(vs ...int) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldThumbNum, vs...))
}

// ThumbNumGT applies the GT predicate on the "thumb_num" field.
func ThumbNumGT(v int) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldThumbNum, v))
}

// ThumbNumGTE applies the GTE predicate on the "thumb_num" field.
func ThumbNumGTE(v int) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldThumbNum, v))
}

// ThumbNumLT applies the LT predicate on the "thumb_num" field.
func ThumbNumLT(v int) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldThumbNum, v))
}

// ThumbNumLTE applies the LTE predicate on the "thumb_num" field.
func ThumbNumLTE(v int) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldThumbNum, v))
}

// CollectNumEQ applies the EQ predicate on the "collect_num" field.
func CollectNumEQ(v int) predicate.Travel {
	return predicate.Travel(sql.FieldEQ(FieldCollectNum, v))
}

// CollectNumNEQ applies the NEQ predicate on the "collect_num" field.
func CollectNumNEQ(v int) predicate.Travel {
	return predicate.Travel(sql.FieldNEQ(FieldCollectNum, v))
}

// CollectNumIn applies the In predicate on the "collect_num" field.
func CollectNumIn(vs ...int) predicate.Travel {
	return predicate.Travel(sql.FieldIn(FieldCollectNum, vs...))
}

// CollectNumNotIn applies the NotIn predicate on the "collect_num" field.
func CollectNumNotIn(vs ...int) predicate.Travel {
	return predicate.Travel(sql.FieldNotIn(FieldCollectNum, vs...))
}

// CollectNumGT applies the GT predicate on the "collect_num" field.
func CollectNumGT(v int) predicate.Travel {
	return predicate.Travel(sql.FieldGT(FieldCollectNum, v))
}

// CollectNumGTE applies the GTE predicate on the "collect_num" field.
func CollectNumGTE(v int) predicate.Travel {
	return predicate.Travel(sql.FieldGTE(FieldCollectNum, v))
}

// CollectNumLT applies the LT predicate on the "collect_num" field.
func CollectNumLT(v int) predicate.Travel {
	return predicate.Travel(sql.FieldLT(FieldCollectNum, v))
}

// CollectNumLTE applies the LTE predicate on the "collect_num" field.
func CollectNumLTE(v int) predicate.Travel {
	return predicate.Travel(sql.FieldLTE(FieldCollectNum, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Travel) predicate.Travel {
	return predicate.Travel(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Travel) predicate.Travel {
	return predicate.Travel(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Travel) predicate.Travel {
	return predicate.Travel(sql.NotPredicates(p))
}