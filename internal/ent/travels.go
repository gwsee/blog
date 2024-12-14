// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/account"
	"blog/internal/ent/travels"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 旅行
type Travels struct {
	config `json:"-"`
	// ID of the ent.
	// 旅行记录的ID
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
	// 标题
	Title string `json:"title,omitempty"`
	// 旅行简介
	Description string `json:"description,omitempty"`
	// 旅行视频
	Video string `json:"video,omitempty"`
	// 是否隐藏:0否,1是
	IsHidden bool `json:"is_hidden,omitempty"`
	// 账户ID
	AccountID int `json:"account_id,omitempty"`
	// 旅行的照片
	Photos []string `json:"photos,omitempty"`
	// 浏览量
	BrowseNum int `json:"browse_num,omitempty"`
	// 点赞量
	ThumbNum int `json:"thumb_num,omitempty"`
	// 收藏量
	CollectNum int `json:"collect_num,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the TravelsQuery when eager-loading is set.
	Edges        TravelsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// TravelsEdges holds the relations/edges for other nodes in the graph.
type TravelsEdges struct {
	// TravelExtends holds the value of the travel_extends edge.
	TravelExtends []*TravelExtends `json:"travel_extends,omitempty"`
	// TravelAccount holds the value of the travel_account edge.
	TravelAccount *Account `json:"travel_account,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TravelExtendsOrErr returns the TravelExtends value or an error if the edge
// was not loaded in eager-loading.
func (e TravelsEdges) TravelExtendsOrErr() ([]*TravelExtends, error) {
	if e.loadedTypes[0] {
		return e.TravelExtends, nil
	}
	return nil, &NotLoadedError{edge: "travel_extends"}
}

// TravelAccountOrErr returns the TravelAccount value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e TravelsEdges) TravelAccountOrErr() (*Account, error) {
	if e.TravelAccount != nil {
		return e.TravelAccount, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: account.Label}
	}
	return nil, &NotLoadedError{edge: "travel_account"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Travels) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case travels.FieldPhotos:
			values[i] = new([]byte)
		case travels.FieldIsHidden:
			values[i] = new(sql.NullBool)
		case travels.FieldID, travels.FieldCreatedAt, travels.FieldCreatedBy, travels.FieldUpdatedAt, travels.FieldUpdatedBy, travels.FieldDeletedAt, travels.FieldDeletedBy, travels.FieldAccountID, travels.FieldBrowseNum, travels.FieldThumbNum, travels.FieldCollectNum:
			values[i] = new(sql.NullInt64)
		case travels.FieldTitle, travels.FieldDescription, travels.FieldVideo:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Travels fields.
func (t *Travels) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case travels.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case travels.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Int64
			}
		case travels.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				t.CreatedBy = value.Int64
			}
		case travels.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Int64
			}
		case travels.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				t.UpdatedBy = value.Int64
			}
		case travels.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = value.Int64
			}
		case travels.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				t.DeletedBy = value.Int64
			}
		case travels.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case travels.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case travels.FieldVideo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video", values[i])
			} else if value.Valid {
				t.Video = value.String
			}
		case travels.FieldIsHidden:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_hidden", values[i])
			} else if value.Valid {
				t.IsHidden = value.Bool
			}
		case travels.FieldAccountID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field account_id", values[i])
			} else if value.Valid {
				t.AccountID = int(value.Int64)
			}
		case travels.FieldPhotos:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field photos", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Photos); err != nil {
					return fmt.Errorf("unmarshal field photos: %w", err)
				}
			}
		case travels.FieldBrowseNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field browse_num", values[i])
			} else if value.Valid {
				t.BrowseNum = int(value.Int64)
			}
		case travels.FieldThumbNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field thumb_num", values[i])
			} else if value.Valid {
				t.ThumbNum = int(value.Int64)
			}
		case travels.FieldCollectNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field collect_num", values[i])
			} else if value.Valid {
				t.CollectNum = int(value.Int64)
			}
		default:
			t.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Travels.
// This includes values selected through modifiers, order, etc.
func (t *Travels) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// QueryTravelExtends queries the "travel_extends" edge of the Travels entity.
func (t *Travels) QueryTravelExtends() *TravelExtendsQuery {
	return NewTravelsClient(t.config).QueryTravelExtends(t)
}

// QueryTravelAccount queries the "travel_account" edge of the Travels entity.
func (t *Travels) QueryTravelAccount() *AccountQuery {
	return NewTravelsClient(t.config).QueryTravelAccount(t)
}

// Update returns a builder for updating this Travels.
// Note that you need to call Travels.Unwrap() before calling this method if this Travels
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Travels) Update() *TravelsUpdateOne {
	return NewTravelsClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Travels entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Travels) Unwrap() *Travels {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Travels is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Travels) String() string {
	var builder strings.Builder
	builder.WriteString("Travels(")
	builder.WriteString(fmt.Sprintf("id=%v, ", t.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", t.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", t.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", t.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", t.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(t.Title)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(t.Description)
	builder.WriteString(", ")
	builder.WriteString("video=")
	builder.WriteString(t.Video)
	builder.WriteString(", ")
	builder.WriteString("is_hidden=")
	builder.WriteString(fmt.Sprintf("%v", t.IsHidden))
	builder.WriteString(", ")
	builder.WriteString("account_id=")
	builder.WriteString(fmt.Sprintf("%v", t.AccountID))
	builder.WriteString(", ")
	builder.WriteString("photos=")
	builder.WriteString(fmt.Sprintf("%v", t.Photos))
	builder.WriteString(", ")
	builder.WriteString("browse_num=")
	builder.WriteString(fmt.Sprintf("%v", t.BrowseNum))
	builder.WriteString(", ")
	builder.WriteString("thumb_num=")
	builder.WriteString(fmt.Sprintf("%v", t.ThumbNum))
	builder.WriteString(", ")
	builder.WriteString("collect_num=")
	builder.WriteString(fmt.Sprintf("%v", t.CollectNum))
	builder.WriteByte(')')
	return builder.String()
}

// TravelsSlice is a parsable slice of Travels.
type TravelsSlice []*Travels
