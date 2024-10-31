package constx

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

const UserInfo = "_user_info"
const EmptyUser = 0

var DefaultUser User

type User struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
}

func (User) Default(ctx context.Context) *User {
	info, ex := jwt.FromContext(ctx)
	if !ex {
		return &User{}
	}
	data, ok := info.(jwt5.MapClaims)
	if !ok {
		return &User{}
	}
	val, ok := data[UserInfo]
	if !ok {
		return &User{}
	}
	vals, ok := val.(string)
	if !ok {
		return &User{}
	}
	var user User
	_ = json.Unmarshal([]byte(vals), &user)
	return &user
}
