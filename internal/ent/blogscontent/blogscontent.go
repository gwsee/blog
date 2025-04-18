// Code generated by ent, DO NOT EDIT.

package blogscontent

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the blogscontent type in the database.
	Label = "blogs_content"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldContent holds the string denoting the content field in the database.
	FieldContent = "content"
	// FieldFiles holds the string denoting the files field in the database.
	FieldFiles = "files"
	// Table holds the table name of the blogscontent in the database.
	Table = "blogs_content"
)

// Columns holds all SQL columns for blogscontent fields.
var Columns = []string{
	FieldID,
	FieldContent,
	FieldFiles,
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

var (
	// ContentValidator is a validator for the "content" field. It is called by the builders before save.
	ContentValidator func(string) error
)

// OrderOption defines the ordering options for the BlogsContent queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByContent orders the results by the content field.
func ByContent(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldContent, opts...).ToFunc()
}
