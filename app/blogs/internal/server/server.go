package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/google/wire"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewEtcdRegistrar)

var whiteList = map[string]struct{}{
	"/api.blogs.v1.Blogs/ListBlogs":               {},
	"/api.blogs.v1.Blogs/GetBlogs":                {},
	"/api.blogs.v1.BlogsComment/GetBlogsComment":  {},
	"/api.blogs.v1.BlogsComment/ListBlogsComment": {},
}

func NewWhiteListMatcher() selector.MatchFunc {
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
