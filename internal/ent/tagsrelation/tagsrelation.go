// Code generated by ent, DO NOT EDIT.

package tagsrelation

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the tagsrelation type in the database.
	Label = "tags_relation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTagID holds the string denoting the tag_id field in the database.
	FieldTagID = "tag_id"
	// FieldRelation holds the string denoting the relation field in the database.
	FieldRelation = "relation"
	// FieldRelationID holds the string denoting the relation_id field in the database.
	FieldRelationID = "relation_id"
	// EdgeBlog holds the string denoting the blog edge name in mutations.
	EdgeBlog = "blog"
	// EdgeTag holds the string denoting the tag edge name in mutations.
	EdgeTag = "tag"
	// Table holds the table name of the tagsrelation in the database.
	Table = "tags_relation"
	// BlogTable is the table that holds the blog relation/edge.
	BlogTable = "tags_relation"
	// BlogInverseTable is the table name for the Blogs entity.
	// It exists in this package in order to avoid circular dependency with the "blogs" package.
	BlogInverseTable = "blogs"
	// BlogColumn is the table column denoting the blog relation/edge.
	BlogColumn = "relation_id"
	// TagTable is the table that holds the tag relation/edge.
	TagTable = "tags_relation"
	// TagInverseTable is the table name for the Tags entity.
	// It exists in this package in order to avoid circular dependency with the "tags" package.
	TagInverseTable = "tags"
	// TagColumn is the table column denoting the tag relation/edge.
	TagColumn = "tag_id"
)

// Columns holds all SQL columns for tagsrelation fields.
var Columns = []string{
	FieldID,
	FieldTagID,
	FieldRelation,
	FieldRelationID,
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

// OrderOption defines the ordering options for the TagsRelation queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTagID orders the results by the tag_id field.
func ByTagID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTagID, opts...).ToFunc()
}

// ByRelation orders the results by the relation field.
func ByRelation(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelation, opts...).ToFunc()
}

// ByRelationID orders the results by the relation_id field.
func ByRelationID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRelationID, opts...).ToFunc()
}

// ByBlogField orders the results by blog field.
func ByBlogField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newBlogStep(), sql.OrderByField(field, opts...))
	}
}

// ByTagField orders the results by tag field.
func ByTagField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTagStep(), sql.OrderByField(field, opts...))
	}
}
func newBlogStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(BlogInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, BlogTable, BlogColumn),
	)
}
func newTagStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TagInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
	)
}
