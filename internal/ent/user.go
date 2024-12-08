// Code generated by ent, DO NOT EDIT.

package ent

import (
	"blog/internal/ent/user"
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// 用户
type User struct {
	config `json:"-"`
	// ID of the ent.
	// 用户ID=账户ID
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
	// 姓名
	Name string `json:"name,omitempty"`
	// 头像
	Avatar string `json:"avatar,omitempty"`
	// 工作邮箱
	Email string `json:"email,omitempty"`
	// 职称
	Professional string `json:"professional,omitempty"`
	// 地址
	Address string `json:"address,omitempty"`
	// 技能
	Skills []string `json:"skills,omitempty"`
	// 个人简介
	Description string `json:"description,omitempty"`
	// 经历数
	Experience int `json:"experience,omitempty"`
	// 项目数
	Project      int `json:"project,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldSkills:
			values[i] = new([]byte)
		case user.FieldID, user.FieldCreatedAt, user.FieldCreatedBy, user.FieldUpdatedAt, user.FieldUpdatedBy, user.FieldDeletedAt, user.FieldDeletedBy, user.FieldExperience, user.FieldProject:
			values[i] = new(sql.NullInt64)
		case user.FieldName, user.FieldAvatar, user.FieldEmail, user.FieldProfessional, user.FieldAddress, user.FieldDescription:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Int64
			}
		case user.FieldCreatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field created_by", values[i])
			} else if value.Valid {
				u.CreatedBy = value.Int64
			}
		case user.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				u.UpdatedAt = value.Int64
			}
		case user.FieldUpdatedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field updated_by", values[i])
			} else if value.Valid {
				u.UpdatedBy = value.Int64
			}
		case user.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				u.DeletedAt = value.Int64
			}
		case user.FieldDeletedBy:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_by", values[i])
			} else if value.Valid {
				u.DeletedBy = value.Int64
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldAvatar:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field avatar", values[i])
			} else if value.Valid {
				u.Avatar = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldProfessional:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field professional", values[i])
			} else if value.Valid {
				u.Professional = value.String
			}
		case user.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				u.Address = value.String
			}
		case user.FieldSkills:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field skills", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Skills); err != nil {
					return fmt.Errorf("unmarshal field skills: %w", err)
				}
			}
		case user.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				u.Description = value.String
			}
		case user.FieldExperience:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field experience", values[i])
			} else if value.Valid {
				u.Experience = int(value.Int64)
			}
		case user.FieldProject:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field project", values[i])
			} else if value.Valid {
				u.Project = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the User.
// This includes values selected through modifiers, order, etc.
func (u *User) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return NewUserClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedAt))
	builder.WriteString(", ")
	builder.WriteString("created_by=")
	builder.WriteString(fmt.Sprintf("%v", u.CreatedBy))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedAt))
	builder.WriteString(", ")
	builder.WriteString("updated_by=")
	builder.WriteString(fmt.Sprintf("%v", u.UpdatedBy))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", u.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("deleted_by=")
	builder.WriteString(fmt.Sprintf("%v", u.DeletedBy))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("avatar=")
	builder.WriteString(u.Avatar)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("professional=")
	builder.WriteString(u.Professional)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(u.Address)
	builder.WriteString(", ")
	builder.WriteString("skills=")
	builder.WriteString(fmt.Sprintf("%v", u.Skills))
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(u.Description)
	builder.WriteString(", ")
	builder.WriteString("experience=")
	builder.WriteString(fmt.Sprintf("%v", u.Experience))
	builder.WriteString(", ")
	builder.WriteString("project=")
	builder.WriteString(fmt.Sprintf("%v", u.Project))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User
