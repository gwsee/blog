// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/travelextends"
	"blog/internal/ent/travels"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 旅行关联关系
type TravelExtends struct {
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
	IsCollect bool `json:"is_collect,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TravelExtendsQuery when eager-loading is set.
	Edges        TravelExtendsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TravelExtendsEdges holds the relations/edges for other nodes in the graph.
type TravelExtendsEdges struct {
	// Extends holds the value of the extends edge.
	Extends *Travels `json:"extends,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ExtendsOrErr returns the Extends value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TravelExtendsEdges) ExtendsOrErr() (*Travels, error) {
	if e.Extends != nil {
		return e.Extends, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: travels.Label}
	}
	return nil, &NotLoadedError{edge: "extends"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*TravelExtends) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case travelextends.FieldIsThumb, travelextends.FieldIsCollect:
			values[i] = new(sql.NullBool)
		case travelextends.FieldID, travelextends.FieldCreatedAt, travelextends.FieldCreatedBy, travelextends.FieldUpdatedAt, travelextends.FieldUpdatedBy, travelextends.FieldDeletedAt, travelextends.FieldDeletedBy, travelextends.FieldAccountID, travelextends.FieldTravelID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the TravelExtends fields.
func (te *TravelExtends) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case travelextends.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			te.ID = int(value.Int64)
		case travelextends.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				te.CreatedAt = value.Int64
			}
		case travelextends.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				te.CreatedBy = value.Int64
			}
		case travelextends.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				te.UpdatedAt = value.Int64
			}
		case travelextends.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				te.UpdatedBy = value.Int64
			}
		case travelextends.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				te.DeletedAt = value.Int64
			}
		case travelextends.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				te.DeletedBy = value.Int64
			}
		case travelextends.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				te.AccountID = int(value.Int64)
			}
		case travelextends.FieldTravelID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field travel_id", values[i])
			} else if value.Valid {
				te.TravelID = int(value.Int64)
			}
		case travelextends.FieldIsThumb:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_thumb", values[i])
			} else if value.Valid {
				te.IsThumb = value.Bool
			}
		case travelextends.FieldIsCollect:
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

// Value returns the ent.Value that was dynamically selected and assigned to the TravelExtends.
// This includes values selected through modifiers, order, etc.
func (te *TravelExtends) Value(name string) (ent.Value, error) {
	return te.selectValues.Get(name)
}

// QueryExtends queries the "extends" edge of the TravelExtends entity.
func (te *TravelExtends) QueryExtends() *TravelsQuery {
	return NewTravelExtendsClient(te.config).QueryExtends(te)
}

// Update returns a builder for updating this TravelExtends.
// Note that you need to call TravelExtends.Unwrap() before calling this method if this TravelExtends
// was returned from a transaction, and the transaction was committed or rolled back.
func (te *TravelExtends) Update() *TravelExtendsUpdateOne {
	return NewTravelExtendsClient(te.config).UpdateOne(te)
}

// Unwrap unwraps the TravelExtends entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (te *TravelExtends) Unwrap() *TravelExtends {
	_tx, ok := te.config.driver.(*txDriver)
	if !ok {
		panic("ent: TravelExtends is not a transactional entity")
	}
	te.config.driver = _tx.drv
	return te
}

// String implements the fmt.Stringer.
func (te *TravelExtends) String() string {
	var builder strings.Builder
	builder.WriteString("TravelExtends(")
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

// TravelExtendsSlice is a parsable slice of TravelExtends.
type TravelExtendsSlice []*TravelExtends
