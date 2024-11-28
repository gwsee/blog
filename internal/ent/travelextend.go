// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/travelextend"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 旅行关联关系
type TravelExtend struct {
	config `json:"-"`
	// ID of the ent.
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
	// 旅行的ID
	TravelID int `json:"travel_id,omitempty"`
	// 是否点赞
	IsThumb bool `json:"is_thumb,omitempty"`
	// 收藏量
	IsCollect    bool `json:"is_collect,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TravelExtend) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case travelextend.FieldIsThumb, travelextend.FieldIsCollect:
			values[i] = new(sql.NullBool)
		case travelextend.FieldID, travelextend.FieldCreatedAt, travelextend.FieldCreatedBy, travelextend.FieldUpdatedAt, travelextend.FieldUpdatedBy, travelextend.FieldDeletedAt, travelextend.FieldDeletedBy, travelextend.FieldAccountID, travelextend.FieldTravelID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TravelExtend fields.
func (te *TravelExtend) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case travelextend.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			te.ID = int(value.Int64)
		case travelextend.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				te.CreatedAt = value.Int64
			}
		case travelextend.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				te.CreatedBy = value.Int64
			}
		case travelextend.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				te.UpdatedAt = value.Int64
			}
		case travelextend.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				te.UpdatedBy = value.Int64
			}
		case travelextend.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				te.DeletedAt = value.Int64
			}
		case travelextend.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				te.DeletedBy = value.Int64
			}
		case travelextend.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				te.AccountID = int(value.Int64)
			}
		case travelextend.FieldTravelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field travel_id", values[i])
			} else if value.Valid {
				te.TravelID = int(value.Int64)
			}
		case travelextend.FieldIsThumb:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_thumb", values[i])
			} else if value.Valid {
				te.IsThumb = value.Bool
			}
		case travelextend.FieldIsCollect:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_collect", values[i])
			} else if value.Valid {
				te.IsCollect = value.Bool
			}
		default:
			te.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the TravelExtend.
// This includes values selected through modifiers, order, etc.
func (te *TravelExtend) Value(name string) (ent.Value, error) {
	return te.selectValues.Get(name)
}

// Update returns a builder for updating this TravelExtend.
// Note that you need to call TravelExtend.Unwrap() before calling this method if this TravelExtend
// was returned from a transaction, and the transaction was committed or rolled back.
func (te *TravelExtend) Update() *TravelExtendUpdateOne {
	return NewTravelExtendClient(te.config).UpdateOne(te)
}

// Unwrap unwraps the TravelExtend entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (te *TravelExtend) Unwrap() *TravelExtend {
	_tx, ok := te.config.driver.(*txDriver)
	if !ok {
		panic("ent: TravelExtend is not a transactional entity")
	}
	te.config.driver = _tx.drv
	return te
}

// String implements the fmt.Stringer.
func (te *TravelExtend) String() string {
	var builder strings.Builder
	builder.WriteString("TravelExtend(")
	builder.WriteString(fmt.Sprintf("id=%v, ", te.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", te.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", te.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", te.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", te.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", te.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", te.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", te.AccountID))
	builder.WriteString(", ")
	builder.WriteString("travel_id=")
	builder.WriteString(fmt.Sprintf("%v", te.TravelID))
	builder.WriteString(", ")
	builder.WriteString("is_thumb=")
	builder.WriteString(fmt.Sprintf("%v", te.IsThumb))
	builder.WriteString(", ")
	builder.WriteString("is_collect=")
	builder.WriteString(fmt.Sprintf("%v", te.IsCollect))
	builder.WriteByte(')')
	return builder.String()
}

// TravelExtends is a parsable slice of TravelExtend.
type TravelExtends []*TravelExtend