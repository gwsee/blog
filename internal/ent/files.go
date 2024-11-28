// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/files"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 文件
type Files struct {
	config `json:"-"`
	// ID of the ent.
	// 文件的ID(内容MD5)
	ID string `json:"id,omitempty"`
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
	// 文件类型
	Type string `json:"type,omitempty"`
	// 文件大小
	Size int64 `json:"size,omitempty"`
	// 文件名称
	Name string `json:"name,omitempty"`
	// 文件真实存放地址
	Path         string `json:"path,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Files) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case files.FieldCreatedAt, files.FieldCreatedBy, files.FieldUpdatedAt, files.FieldUpdatedBy, files.FieldDeletedAt, files.FieldDeletedBy, files.FieldSize:
			values[i] = new(sql.NullInt64)
		case files.FieldID, files.FieldType, files.FieldName, files.FieldPath:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Files fields.
func (f *Files) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case files.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				f.ID = value.String
			}
		case files.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				f.CreatedAt = value.Int64
			}
		case files.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				f.CreatedBy = value.Int64
			}
		case files.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				f.UpdatedAt = value.Int64
			}
		case files.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				f.UpdatedBy = value.Int64
			}
		case files.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				f.DeletedAt = value.Int64
			}
		case files.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				f.DeletedBy = value.Int64
			}
		case files.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				f.Type = value.String
			}
		case files.FieldSize:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field size", values[i])
			} else if value.Valid {
				f.Size = value.Int64
			}
		case files.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				f.Name = value.String
			}
		case files.FieldPath:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field path", values[i])
			} else if value.Valid {
				f.Path = value.String
			}
		default:
			f.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Files.
// This includes values selected through modifiers, order, etc.
func (f *Files) Value(name string) (ent.Value, error) {
	return f.selectValues.Get(name)
}

// Update returns a builder for updating this Files.
// Note that you need to call Files.Unwrap() before calling this method if this Files
// was returned from a transaction, and the transaction was committed or rolled back.
func (f *Files) Update() *FilesUpdateOne {
	return NewFilesClient(f.config).UpdateOne(f)
}

// Unwrap unwraps the Files entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (f *Files) Unwrap() *Files {
	_tx, ok := f.config.driver.(*txDriver)
	if !ok {
		panic("ent: Files is not a transactional entity")
	}
	f.config.driver = _tx.drv
	return f
}

// String implements the fmt.Stringer.
func (f *Files) String() string {
	var builder strings.Builder
	builder.WriteString("Files(")
	builder.WriteString(fmt.Sprintf("id=%v, ", f.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", f.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", f.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", f.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", f.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", f.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", f.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(f.Type)
	builder.WriteString(", ")
	builder.WriteString("size=")
	builder.WriteString(fmt.Sprintf("%v", f.Size))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(f.Name)
	builder.WriteString(", ")
	builder.WriteString("path=")
	builder.WriteString(f.Path)
	builder.WriteByte(')')
	return builder.String()
}

// FilesSlice is a parsable slice of Files.
type FilesSlice []*Files