// Code generated by ent, DO NOT EDIT.

package userproject

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the userproject type in the database.
	Label = "user_project"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldCreatedBy holds the string denoting the created_by field in the database.
	FieldCreatedBy = "created_by"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldUpdatedBy holds the string denoting the updated_by field in the database.
	FieldUpdatedBy = "updated_by"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldDeletedBy holds the string denoting the deleted_by field in the database.
	FieldDeletedBy = "deleted_by"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldExperienceID holds the string denoting the experience_id field in the database.
	FieldExperienceID = "experience_id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldSkills holds the string denoting the skills field in the database.
	FieldSkills = "skills"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// FieldLink holds the string denoting the link field in the database.
	FieldLink = "link"
	// FieldPhotos holds the string denoting the photos field in the database.
	FieldPhotos = "photos"
	// Table holds the table name of the userproject in the database.
	Table = "users_project"
)

// Columns holds all SQL columns for userproject fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldUserID,
	FieldExperienceID,
	FieldTitle,
	FieldDescription,
	FieldSkills,
	FieldStart,
	FieldEnd,
	FieldLink,
	FieldPhotos,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "blog/internal/ent/runtime"
var (
	Hooks        [3]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt int64
	// UpdateDefaultCreatedAt holds the default value on update for the "created_at" field.
	UpdateDefaultCreatedAt func() int64
	// DefaultCreatedBy holds the default value on creation for the "created_by" field.
	DefaultCreatedBy int64
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt int64
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() int64
	// DefaultUpdatedBy holds the default value on creation for the "updated_by" field.
	DefaultUpdatedBy int64
	// DefaultDeletedAt holds the default value on creation for the "deleted_at" field.
	DefaultDeletedAt int64
	// DefaultDeletedBy holds the default value on creation for the "deleted_by" field.
	DefaultDeletedBy int64
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(int) error
)

// OrderOption defines the ordering options for the UserProject queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByCreatedBy orders the results by the created_by field.
func ByCreatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedBy, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByUpdatedBy orders the results by the updated_by field.
func ByUpdatedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedBy, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByDeletedBy orders the results by the deleted_by field.
func ByDeletedBy(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedBy, opts...).ToFunc()
}

// ByUserID orders the results by the user_id field.
func ByUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserID, opts...).ToFunc()
}

// ByExperienceID orders the results by the experience_id field.
func ByExperienceID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExperienceID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByStart orders the results by the start field.
func ByStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStart, opts...).ToFunc()
}

// ByEnd orders the results by the end field.
func ByEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnd, opts...).ToFunc()
}

// ByLink orders the results by the link field.
func ByLink(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLink, opts...).ToFunc()
}
