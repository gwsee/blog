// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/blogscomment"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 博客评论
type BlogsComment struct {
	config `json:"-"`
	// ID of the ent.
	// 评论ID
	ID int `json:"id,omitempty"`
	// 创建时间
	CreatedAt int64 `json:"created_at,omitempty"`
	// 创建人
	CreatedBy int64 `json:"created_by,omitempty"`
	// 更新时间
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// 更新人
	UpdatedBy int64 `json:"updated_by,omitempty"`
	// 软删除时间
	DeletedAt int64 `json:"deleted_at,omitempty"`
	// 删除人
	DeletedBy int64 `json:"deleted_by,omitempty"`
	// 账户ID
	AccountID int `json:"account_id,omitempty"`
	// 博客ID
	BlogID int `json:"blog_id,omitempty"`
	// 顶级ID
	TopID int `json:"top_id,omitempty"`
	// 父评论
	ParentID int `json:"parent_id,omitempty"`
	// 第几楼
	Level int `json:"level,omitempty"`
	// 当前级有多少数据;状态未0的
	Total int `json:"total,omitempty"`
	// 0显示,1隐藏
	Status int8 `json:"status,omitempty"`
	// 评论内容
	Content      string `json:"content,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BlogsComment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case blogscomment.FieldID, blogscomment.FieldCreatedAt, blogscomment.FieldCreatedBy, blogscomment.FieldUpdatedAt, blogscomment.FieldUpdatedBy, blogscomment.FieldDeletedAt, blogscomment.FieldDeletedBy, blogscomment.FieldAccountID, blogscomment.FieldBlogID, blogscomment.FieldTopID, blogscomment.FieldParentID, blogscomment.FieldLevel, blogscomment.FieldTotal, blogscomment.FieldStatus:
			values[i] = new(sql.NullInt64)
		case blogscomment.FieldContent:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BlogsComment fields.
func (bc *BlogsComment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case blogscomment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			bc.ID = int(value.Int64)
		case blogscomment.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bc.CreatedAt = value.Int64
			}
		case blogscomment.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				bc.CreatedBy = value.Int64
			}
		case blogscomment.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bc.UpdatedAt = value.Int64
			}
		case blogscomment.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				bc.UpdatedBy = value.Int64
			}
		case blogscomment.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				bc.DeletedAt = value.Int64
			}
		case blogscomment.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				bc.DeletedBy = value.Int64
			}
		case blogscomment.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				bc.AccountID = int(value.Int64)
			}
		case blogscomment.FieldBlogID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field blog_id", values[i])
			} else if value.Valid {
				bc.BlogID = int(value.Int64)
			}
		case blogscomment.FieldTopID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field top_id", values[i])
			} else if value.Valid {
				bc.TopID = int(value.Int64)
			}
		case blogscomment.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				bc.ParentID = int(value.Int64)
			}
		case blogscomment.FieldLevel:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field level", values[i])
			} else if value.Valid {
				bc.Level = int(value.Int64)
			}
		case blogscomment.FieldTotal:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field total", values[i])
			} else if value.Valid {
				bc.Total = int(value.Int64)
			}
		case blogscomment.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				bc.Status = int8(value.Int64)
			}
		case blogscomment.FieldContent:
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

// Value returns the ent.Value that was dynamically selected and assigned to the BlogsComment.
// This includes values selected through modifiers, order, etc.
func (bc *BlogsComment) Value(name string) (ent.Value, error) {
	return bc.selectValues.Get(name)
}

// Update returns a builder for updating this BlogsComment.
// Note that you need to call BlogsComment.Unwrap() before calling this method if this BlogsComment
// was returned from a transaction, and the transaction was committed or rolled back.
func (bc *BlogsComment) Update() *BlogsCommentUpdateOne {
	return NewBlogsCommentClient(bc.config).UpdateOne(bc)
}

// Unwrap unwraps the BlogsComment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bc *BlogsComment) Unwrap() *BlogsComment {
	_tx, ok := bc.config.driver.(*txDriver)
	if !ok {
		panic("ent: BlogsComment is not a transactional entity")
	}
	bc.config.driver = _tx.drv
	return bc
}

// String implements the fmt.Stringer.
func (bc *BlogsComment) String() string {
	var builder strings.Builder
	builder.WriteString("BlogsComment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", bc.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", bc.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", bc.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", bc.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", bc.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", bc.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", bc.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", bc.AccountID))
	builder.WriteString(", ")
	builder.WriteString("blog_id=")
	builder.WriteString(fmt.Sprintf("%v", bc.BlogID))
	builder.WriteString(", ")
	builder.WriteString("top_id=")
	builder.WriteString(fmt.Sprintf("%v", bc.TopID))
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", bc.ParentID))
	builder.WriteString(", ")
	builder.WriteString("level=")
	builder.WriteString(fmt.Sprintf("%v", bc.Level))
	builder.WriteString(", ")
	builder.WriteString("total=")
	builder.WriteString(fmt.Sprintf("%v", bc.Total))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", bc.Status))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(bc.Content)
	builder.WriteByte(')')
	return builder.String()
}

// BlogsComments is a parsable slice of BlogsComment.
type BlogsComments []*BlogsComment
