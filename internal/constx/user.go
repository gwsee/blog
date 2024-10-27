package constx

import "context"

const UserInfo = "_user_info"
const EmptyUser = ""

var DefaultUser User

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

func (User) Default(ctx context.Context) *User {
	userInfo, ok := ctx.Value(UserInfo).(*User)
	if ok {
		return userInfo
	}
	return &User{}
}
