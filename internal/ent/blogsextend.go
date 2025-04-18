// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogsextend"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 博客关联信息
type BlogsExtend struct {
	config `json:"-"`
	// ID of the ent.
	// ID
	ID int `json:"id,omitempty"`
	// 博客ID
	BlogID int `json:"blog_id,omitempty"`
	// 账户ID
	AccountID int `json:"account_id,omitempty"`
	// 浏览量
	BrowseNum int `json:"browse_num,omitempty"`
	// 最后的浏览时间
	BrowseAt int64 `json:"browse_at,omitempty"`
	// 是否收藏
	Collect bool `json:"collect,omitempty"`
	// 收藏时间
	CollectAt int64 `json:"collect_at,omitempty"`
	// 是否点赞
	Love bool `json:"love,omitempty"`
	// 点赞时间
	LoveAt       int64 `json:"love_at,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BlogsExtend) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blogsextend.FieldCollect, blogsextend.FieldLove:
			values[i] = new(sql.NullBool)
		case blogsextend.FieldID, blogsextend.FieldBlogID, blogsextend.FieldAccountID, blogsextend.FieldBrowseNum, blogsextend.FieldBrowseAt, blogsextend.FieldCollectAt, blogsextend.FieldLoveAt:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BlogsExtend fields.
func (be *BlogsExtend) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blogsextend.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			be.ID = int(value.Int64)
		case blogsextend.FieldBlogID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field blog_id", values[i])
			} else if value.Valid {
				be.BlogID = int(value.Int64)
			}
		case blogsextend.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				be.AccountID = int(value.Int64)
			}
		case blogsextend.FieldBrowseNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field browse_num", values[i])
			} else if value.Valid {
				be.BrowseNum = int(value.Int64)
			}
		case blogsextend.FieldBrowseAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field browse_at", values[i])
			} else if value.Valid {
				be.BrowseAt = value.Int64
			}
		case blogsextend.FieldCollect:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field collect", values[i])
			} else if value.Valid {
				be.Collect = value.Bool
			}
		case blogsextend.FieldCollectAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field collect_at", values[i])
			} else if value.Valid {
				be.CollectAt = value.Int64
			}
		case blogsextend.FieldLove:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field love", values[i])
			} else if value.Valid {
				be.Love = value.Bool
			}
		case blogsextend.FieldLoveAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field love_at", values[i])
			} else if value.Valid {
				be.LoveAt = value.Int64
			}
		default:
			be.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the BlogsExtend.
// This includes values selected through modifiers, order, etc.
func (be *BlogsExtend) Value(name string) (ent.Value, error) {
	return be.selectValues.Get(name)
}

// Update returns a builder for updating this BlogsExtend.
// Note that you need to call BlogsExtend.Unwrap() before calling this method if this BlogsExtend
// was returned from a transaction, and the transaction was committed or rolled back.
func (be *BlogsExtend) Update() *BlogsExtendUpdateOne {
	return NewBlogsExtendClient(be.config).UpdateOne(be)
}

// Unwrap unwraps the BlogsExtend entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (be *BlogsExtend) Unwrap() *BlogsExtend {
	_tx, ok := be.config.driver.(*txDriver)
	if !ok {
		panic("ent: BlogsExtend is not a transactional entity")
	}
	be.config.driver = _tx.drv
	return be
}

// String implements the fmt.Stringer.
func (be *BlogsExtend) String() string {
	var builder strings.Builder
	builder.WriteString("BlogsExtend(")
	builder.WriteString(fmt.Sprintf("id=%v, ", be.ID))
	builder.WriteString("blog_id=")
	builder.WriteString(fmt.Sprintf("%v", be.BlogID))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", be.AccountID))
	builder.WriteString(", ")
	builder.WriteString("browse_num=")
	builder.WriteString(fmt.Sprintf("%v", be.BrowseNum))
	builder.WriteString(", ")
	builder.WriteString("browse_at=")
	builder.WriteString(fmt.Sprintf("%v", be.BrowseAt))
	builder.WriteString(", ")
	builder.WriteString("collect=")
	builder.WriteString(fmt.Sprintf("%v", be.Collect))
	builder.WriteString(", ")
	builder.WriteString("collect_at=")
	builder.WriteString(fmt.Sprintf("%v", be.CollectAt))
	builder.WriteString(", ")
	builder.WriteString("love=")
	builder.WriteString(fmt.Sprintf("%v", be.Love))
	builder.WriteString(", ")
	builder.WriteString("love_at=")
	builder.WriteString(fmt.Sprintf("%v", be.LoveAt))
	builder.WriteByte(')')
	return builder.String()
}

// BlogsExtends is a parsable slice of BlogsExtend.
type BlogsExtends []*BlogsExtend
