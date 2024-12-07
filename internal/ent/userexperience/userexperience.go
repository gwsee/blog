// Code generated by ent, DO NOT EDIT.

package userexperience

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the userexperience type in the database.
	Label = "user_experience"
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
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldLocation holds the string denoting the location field in the database.
	FieldLocation = "location"
	// FieldStart holds the string denoting the start field in the database.
	FieldStart = "start"
	// FieldEnd holds the string denoting the end field in the database.
	FieldEnd = "end"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldResponsibilities holds the string denoting the responsibilities field in the database.
	FieldResponsibilities = "responsibilities"
	// FieldAchievements holds the string denoting the achievements field in the database.
	FieldAchievements = "achievements"
	// FieldSkills holds the string denoting the skills field in the database.
	FieldSkills = "skills"
	// Table holds the table name of the userexperience in the database.
	Table = "user_experience"
)

// Columns holds all SQL columns for userexperience fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldUserID,
	FieldCompany,
	FieldRole,
	FieldLocation,
	FieldStart,
	FieldEnd,
	FieldDescription,
	FieldResponsibilities,
	FieldAchievements,
	FieldSkills,
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

// OrderOption defines the ordering options for the UserExperience queries.
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

// ByCompany orders the results by the company field.
func ByCompany(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCompany, opts...).ToFunc()
}

// ByRole orders the results by the role field.
func ByRole(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRole, opts...).ToFunc()
}

// ByLocation orders the results by the location field.
func ByLocation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLocation, opts...).ToFunc()
}

// ByStart orders the results by the start field.
func ByStart(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStart, opts...).ToFunc()
}

// ByEnd orders the results by the end field.
func ByEnd(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEnd, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByResponsibilities orders the results by the responsibilities field.
func ByResponsibilities(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResponsibilities, opts...).ToFunc()
}

// ByAchievements orders the results by the achievements field.
func ByAchievements(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAchievements, opts...).ToFunc()
}