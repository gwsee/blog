// Code generated by ent, DO NOT EDIT.

package travelextends

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the travelextends type in the database.
	Label = "travel_extends"
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
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldTravelID holds the string denoting the travel_id field in the database.
	FieldTravelID = "travel_id"
	// FieldIsThumb holds the string denoting the is_thumb field in the database.
	FieldIsThumb = "is_thumb"
	// FieldIsCollect holds the string denoting the is_collect field in the database.
	FieldIsCollect = "is_collect"
	// EdgeExtends holds the string denoting the extends edge name in mutations.
	EdgeExtends = "extends"
	// Table holds the table name of the travelextends in the database.
	Table = "travels_extend"
	// ExtendsTable is the table that holds the extends relation/edge.
	ExtendsTable = "travels_extend"
	// ExtendsInverseTable is the table name for the Travels entity.
	// It exists in this package in order to avoid circular dependency with the "travels" package.
	ExtendsInverseTable = "travels"
	// ExtendsColumn is the table column denoting the extends relation/edge.
	ExtendsColumn = "travel_id"
)

// Columns holds all SQL columns for travelextends fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldAccountID,
	FieldTravelID,
	FieldIsThumb,
	FieldIsCollect,
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
	// AccountIDValidator is a validator for the "account_id" field. It is called by the builders before save.
	AccountIDValidator func(int) error
	// DefaultIsThumb holds the default value on creation for the "is_thumb" field.
	DefaultIsThumb bool
	// DefaultIsCollect holds the default value on creation for the "is_collect" field.
	DefaultIsCollect bool
)

// OrderOption defines the ordering options for the TravelExtends queries.
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

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByTravelID orders the results by the travel_id field.
func ByTravelID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTravelID, opts...).ToFunc()
}

// ByIsThumb orders the results by the is_thumb field.
func ByIsThumb(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsThumb, opts...).ToFunc()
}

// ByIsCollect orders the results by the is_collect field.
func ByIsCollect(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsCollect, opts...).ToFunc()
}

// ByExtendsField orders the results by extends field.
func ByExtendsField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newExtendsStep(), sql.OrderByField(field, opts...))
	}
}
func newExtendsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ExtendsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ExtendsTable, ExtendsColumn),
	)
}
