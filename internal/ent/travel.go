// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/travel"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 旅行
type Travel struct {
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
	// 旅行的照片
	Photos []string `json:"photos,omitempty"`
	// 浏览量
	BrowseNum int `json:"browse_num,omitempty"`
	// 点赞量
	ThumbNum int `json:"thumb_num,omitempty"`
	// 收藏量
	CollectNum   int `json:"collect_num,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Travel) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case travel.FieldPhotos:
			values[i] = new([]byte)
		case travel.FieldID, travel.FieldCreatedAt, travel.FieldCreatedBy, travel.FieldUpdatedAt, travel.FieldUpdatedBy, travel.FieldDeletedAt, travel.FieldDeletedBy, travel.FieldBrowseNum, travel.FieldThumbNum, travel.FieldCollectNum:
			values[i] = new(sql.NullInt64)
		case travel.FieldTitle, travel.FieldDescription, travel.FieldVideo:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Travel fields.
func (t *Travel) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case travel.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			t.ID = int(value.Int64)
		case travel.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				t.CreatedAt = value.Int64
			}
		case travel.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				t.CreatedBy = value.Int64
			}
		case travel.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				t.UpdatedAt = value.Int64
			}
		case travel.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				t.UpdatedBy = value.Int64
			}
		case travel.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				t.DeletedAt = value.Int64
			}
		case travel.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				t.DeletedBy = value.Int64
			}
		case travel.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				t.Title = value.String
			}
		case travel.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				t.Description = value.String
			}
		case travel.FieldVideo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video", values[i])
			} else if value.Valid {
				t.Video = value.String
			}
		case travel.FieldPhotos:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field photos", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &t.Photos); err != nil {
					return fmt.Errorf("unmarshal field photos: %w", err)
				}
			}
		case travel.FieldBrowseNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field browse_num", values[i])
			} else if value.Valid {
				t.BrowseNum = int(value.Int64)
			}
		case travel.FieldThumbNum:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field thumb_num", values[i])
			} else if value.Valid {
				t.ThumbNum = int(value.Int64)
			}
		case travel.FieldCollectNum:
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

// Value returns the ent.Value that was dynamically selected and assigned to the Travel.
// This includes values selected through modifiers, order, etc.
func (t *Travel) Value(name string) (ent.Value, error) {
	return t.selectValues.Get(name)
}

// Update returns a builder for updating this Travel.
// Note that you need to call Travel.Unwrap() before calling this method if this Travel
// was returned from a transaction, and the transaction was committed or rolled back.
func (t *Travel) Update() *TravelUpdateOne {
	return NewTravelClient(t.config).UpdateOne(t)
}

// Unwrap unwraps the Travel entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (t *Travel) Unwrap() *Travel {
	_tx, ok := t.config.driver.(*txDriver)
	if !ok {
		panic("ent: Travel is not a transactional entity")
	}
	t.config.driver = _tx.drv
	return t
}

// String implements the fmt.Stringer.
func (t *Travel) String() string {
	var builder strings.Builder
	builder.WriteString("Travel(")
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

// Travels is a parsable slice of Travel.
type Travels []*Travel