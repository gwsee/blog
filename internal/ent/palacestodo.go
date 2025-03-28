// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/palacestodo"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 代办事项
type PalacesTodo struct {
	config `json:"-"`
	// ID of the ent.
	// ID
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
	// Theme holds the value of the "theme" field.
	Theme string `json:"theme,omitempty"`
	// 类型(1:按次,2：固定日期,2无固定日期)
	Type int8 `json:"type,omitempty"`
	// From holds the value of the "from" field.
	From int64 `json:"from,omitempty"`
	// To holds the value of the "to" field.
	To int64 `json:"to,omitempty"`
	// Num holds the value of the "num" field.
	Num int64 `json:"num,omitempty"`
	// Sort holds the value of the "sort" field.
	Sort int64 `json:"sort,omitempty"`
	// 内容
	Content string `json:"content,omitempty"`
	// 状态 0 代办,1 已做
	Status int8 `json:"status,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PalacesTodoQuery when eager-loading is set.
	Edges        PalacesTodoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PalacesTodoEdges holds the relations/edges for other nodes in the graph.
type PalacesTodoEdges struct {
	// Dones holds the value of the dones edge.
	Dones []*PalacesTodoDone `json:"dones,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// DonesOrErr returns the Dones value or an error if the edge
// was not loaded in eager-loading.
func (e PalacesTodoEdges) DonesOrErr() ([]*PalacesTodoDone, error) {
	if e.loadedTypes[0] {
		return e.Dones, nil
	}
	return nil, &NotLoadedError{edge: "dones"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PalacesTodo) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case palacestodo.FieldID, palacestodo.FieldCreatedAt, palacestodo.FieldCreatedBy, palacestodo.FieldUpdatedAt, palacestodo.FieldUpdatedBy, palacestodo.FieldDeletedAt, palacestodo.FieldDeletedBy, palacestodo.FieldAccountID, palacestodo.FieldType, palacestodo.FieldFrom, palacestodo.FieldTo, palacestodo.FieldNum, palacestodo.FieldSort, palacestodo.FieldStatus:
			values[i] = new(sql.NullInt64)
		case palacestodo.FieldTheme, palacestodo.FieldContent:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PalacesTodo fields.
func (pt *PalacesTodo) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case palacestodo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pt.ID = int(value.Int64)
		case palacestodo.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pt.CreatedAt = value.Int64
			}
		case palacestodo.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				pt.CreatedBy = value.Int64
			}
		case palacestodo.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				pt.UpdatedAt = value.Int64
			}
		case palacestodo.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				pt.UpdatedBy = value.Int64
			}
		case palacestodo.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				pt.DeletedAt = value.Int64
			}
		case palacestodo.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				pt.DeletedBy = value.Int64
			}
		case palacestodo.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				pt.AccountID = int(value.Int64)
			}
		case palacestodo.FieldTheme:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field theme", values[i])
			} else if value.Valid {
				pt.Theme = value.String
			}
		case palacestodo.FieldType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				pt.Type = int8(value.Int64)
			}
		case palacestodo.FieldFrom:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				pt.From = value.Int64
			}
		case palacestodo.FieldTo:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field to", values[i])
			} else if value.Valid {
				pt.To = value.Int64
			}
		case palacestodo.FieldNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field num", values[i])
			} else if value.Valid {
				pt.Num = value.Int64
			}
		case palacestodo.FieldSort:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field sort", values[i])
			} else if value.Valid {
				pt.Sort = value.Int64
			}
		case palacestodo.FieldContent:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value.Valid {
				pt.Content = value.String
			}
		case palacestodo.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				pt.Status = int8(value.Int64)
			}
		default:
			pt.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PalacesTodo.
// This includes values selected through modifiers, order, etc.
func (pt *PalacesTodo) Value(name string) (ent.Value, error) {
	return pt.selectValues.Get(name)
}

// QueryDones queries the "dones" edge of the PalacesTodo entity.
func (pt *PalacesTodo) QueryDones() *PalacesTodoDoneQuery {
	return NewPalacesTodoClient(pt.config).QueryDones(pt)
}

// Update returns a builder for updating this PalacesTodo.
// Note that you need to call PalacesTodo.Unwrap() before calling this method if this PalacesTodo
// was returned from a transaction, and the transaction was committed or rolled back.
func (pt *PalacesTodo) Update() *PalacesTodoUpdateOne {
	return NewPalacesTodoClient(pt.config).UpdateOne(pt)
}

// Unwrap unwraps the PalacesTodo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pt *PalacesTodo) Unwrap() *PalacesTodo {
	_tx, ok := pt.config.driver.(*txDriver)
	if !ok {
		panic("ent: PalacesTodo is not a transactional entity")
	}
	pt.config.driver = _tx.drv
	return pt
}

// String implements the fmt.Stringer.
func (pt *PalacesTodo) String() string {
	var builder strings.Builder
	builder.WriteString("PalacesTodo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pt.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", pt.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", pt.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", pt.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", pt.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", pt.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", pt.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", pt.AccountID))
	builder.WriteString(", ")
	builder.WriteString("theme=")
	builder.WriteString(pt.Theme)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(fmt.Sprintf("%v", pt.Type))
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(fmt.Sprintf("%v", pt.From))
	builder.WriteString(", ")
	builder.WriteString("to=")
	builder.WriteString(fmt.Sprintf("%v", pt.To))
	builder.WriteString(", ")
	builder.WriteString("num=")
	builder.WriteString(fmt.Sprintf("%v", pt.Num))
	builder.WriteString(", ")
	builder.WriteString("sort=")
	builder.WriteString(fmt.Sprintf("%v", pt.Sort))
	builder.WriteString(", ")
	builder.WriteString("content=")
	builder.WriteString(pt.Content)
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", pt.Status))
	builder.WriteByte(')')
	return builder.String()
}

// PalacesTodos is a parsable slice of PalacesTodo.
type PalacesTodos []*PalacesTodo
