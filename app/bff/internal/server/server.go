package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

var whiteList = map[string]struct{}{
	"/v1/blogs/get":          {},
	"/v1/blogs/list":         {},
	"/v1/blogs_comment/list": {},
	"/v1/blogs_comment/get":  {},

	"/v1/register":       {},
	"/v1/reset_password": {},
	"/v1/login":          {},
}

func NewWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
