// Code generated by ent, DO NOT EDIT.

package travels

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the travels type in the database.
	Label = "travels"
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
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldVideo holds the string denoting the video field in the database.
	FieldVideo = "video"
	// FieldIsHidden holds the string denoting the is_hidden field in the database.
	FieldIsHidden = "is_hidden"
	// FieldAccountID holds the string denoting the account_id field in the database.
	FieldAccountID = "account_id"
	// FieldPhotos holds the string denoting the photos field in the database.
	FieldPhotos = "photos"
	// FieldBrowseNum holds the string denoting the browse_num field in the database.
	FieldBrowseNum = "browse_num"
	// FieldThumbNum holds the string denoting the thumb_num field in the database.
	FieldThumbNum = "thumb_num"
	// FieldCollectNum holds the string denoting the collect_num field in the database.
	FieldCollectNum = "collect_num"
	// EdgeTravelExtends holds the string denoting the travel_extends edge name in mutations.
	EdgeTravelExtends = "travel_extends"
	// EdgeTravelAccount holds the string denoting the travel_account edge name in mutations.
	EdgeTravelAccount = "travel_account"
	// Table holds the table name of the travels in the database.
	Table = "travels"
	// TravelExtendsTable is the table that holds the travel_extends relation/edge.
	TravelExtendsTable = "travel_extends"
	// TravelExtendsInverseTable is the table name for the TravelExtends entity.
	// It exists in this package in order to avoid circular dependency with the "travelextends" package.
	TravelExtendsInverseTable = "travel_extends"
	// TravelExtendsColumn is the table column denoting the travel_extends relation/edge.
	TravelExtendsColumn = "travel_id"
	// TravelAccountTable is the table that holds the travel_account relation/edge.
	TravelAccountTable = "travels"
	// TravelAccountInverseTable is the table name for the Account entity.
	// It exists in this package in order to avoid circular dependency with the "account" package.
	TravelAccountInverseTable = "account"
	// TravelAccountColumn is the table column denoting the travel_account relation/edge.
	TravelAccountColumn = "account_id"
)

// Columns holds all SQL columns for travels fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldCreatedBy,
	FieldUpdatedAt,
	FieldUpdatedBy,
	FieldDeletedAt,
	FieldDeletedBy,
	FieldTitle,
	FieldDescription,
	FieldVideo,
	FieldIsHidden,
	FieldAccountID,
	FieldPhotos,
	FieldBrowseNum,
	FieldThumbNum,
	FieldCollectNum,
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
	Hooks [3]ent.Hook
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
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
	// DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	DescriptionValidator func(string) error
	// DefaultVideo holds the default value on creation for the "video" field.
	DefaultVideo string
	// DefaultIsHidden holds the default value on creation for the "is_hidden" field.
	DefaultIsHidden bool
	// DefaultBrowseNum holds the default value on creation for the "browse_num" field.
	DefaultBrowseNum int
	// DefaultThumbNum holds the default value on creation for the "thumb_num" field.
	DefaultThumbNum int
	// DefaultCollectNum holds the default value on creation for the "collect_num" field.
	DefaultCollectNum int
)

// OrderOption defines the ordering options for the Travels queries.
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

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByVideo orders the results by the video field.
func ByVideo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideo, opts...).ToFunc()
}

// ByIsHidden orders the results by the is_hidden field.
func ByIsHidden(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIsHidden, opts...).ToFunc()
}

// ByAccountID orders the results by the account_id field.
func ByAccountID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAccountID, opts...).ToFunc()
}

// ByBrowseNum orders the results by the browse_num field.
func ByBrowseNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBrowseNum, opts...).ToFunc()
}

// ByThumbNum orders the results by the thumb_num field.
func ByThumbNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldThumbNum, opts...).ToFunc()
}

// ByCollectNum orders the results by the collect_num field.
func ByCollectNum(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCollectNum, opts...).ToFunc()
}

// ByTravelExtendsCount orders the results by travel_extends count.
func ByTravelExtendsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTravelExtendsStep(), opts...)
	}
}

// ByTravelExtends orders the results by travel_extends terms.
func ByTravelExtends(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTravelExtendsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByTravelAccountField orders the results by travel_account field.
func ByTravelAccountField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTravelAccountStep(), sql.OrderByField(field, opts...))
	}
}
func newTravelExtendsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TravelExtendsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TravelExtendsTable, TravelExtendsColumn),
	)
}
func newTravelAccountStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TravelAccountInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TravelAccountTable, TravelAccountColumn),
	)
}
