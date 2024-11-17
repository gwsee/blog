package constx

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/transport"
	jwt5 "github.com/golang-jwt/jwt/v5"
)

const UserInfo = "_user_info"
const EmptyUser = 0

var DefaultUser User

type User struct {
	Id      int64  `json:"id"`
	Account string `json:"account"`
	jwt5.MapClaims
}

func (User) Default(ctx context.Context) *User {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return &User{}
	}
	vals := tr.RequestHeader().Get(UserInfo)
	var user User
	_ = json.Unmarshal([]byte(vals), &user)
	return &user
}

func (User) User(ctx context.Context) (*User, error) {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, errors.New("请登陆")
	}
	vals := tr.RequestHeader().Get(UserInfo)
	var user User
	err := json.Unmarshal([]byte(vals), &user)
	if err != nil {
		return nil, errors.New("请登陆")
	}
	return &user, nil
}
