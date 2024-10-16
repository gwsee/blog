// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogcontent"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 博客内容
type BlogContent struct {
	config `json:"-"`
	// ID of the ent.
	// 博客ID
	ID int `json:"id,omitempty"`
	// 内容
	Content      string `json:"content,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BlogContent) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blogcontent.FieldID:
			values[i] = new(sql.NullInt64)
		case blogcontent.FieldContent:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BlogContent fields.
func (bc *BlogContent) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blogcontent.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bc.ID = int(value.Int64)
		case blogcontent.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				bc.Content = value.String
			}
		default:
			bc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BlogContent.
// This includes values selected through modifiers, order, etc.
func (bc *BlogContent) Value(name string) (ent.Value, error) {
	return bc.selectValues.Get(name)
}

// Update returns a builder for updating this BlogContent.
// Note that you need to call BlogContent.Unwrap() before calling this method if this BlogContent
// was returned from a transaction, and the transaction was committed or rolled back.
func (bc *BlogContent) Update() *BlogContentUpdateOne {
	return NewBlogContentClient(bc.config).UpdateOne(bc)
}

// Unwrap unwraps the BlogContent entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bc *BlogContent) Unwrap() *BlogContent {
	_tx, ok := bc.config.driver.(*txDriver)
	if !ok {
		panic("ent: BlogContent is not a transactional entity")
	}
	bc.config.driver = _tx.drv
	return bc
}

// String implements the fmt.Stringer.
func (bc *BlogContent) String() string {
	var builder strings.Builder
	builder.WriteString("BlogContent(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bc.ID))
	builder.WriteString("content=")
	builder.WriteString(bc.Content)
	builder.WriteByte(')')
	return builder.String()
}

// BlogContents is a parsable slice of BlogContent.
type BlogContents []*BlogContent
