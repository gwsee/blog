package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer)

var whiteList = map[string]struct{}{
	"/api.bff.v1.Blogs/GetBlogs":                {},
	"/api.bff.v1.Blogs/ListBlogs":               {},
	"/api.bff.v1.BlogsComment/ListBlogsComment": {},
	"/api.bff.v1.BlogsComment/GetBlogsComment":  {},

	"/api.bff.v1.Account/CreateAccount":  {},
	"/api.bff.v1.Account/ResetPassword":  {},
	"/api.bff.v1.Account/LoginByAccount": {},

	"/api.bff.v1.Tools/UploadFile": {},
	"/api.bff.v1.Tools/Files":      {},
}

func NewWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
