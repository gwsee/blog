package server

import (
	"blog/api/blogs/v1"
	"blog/app/blogs/internal/conf"
	"blog/app/blogs/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, blogs *service.BlogsService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.Middleware(
		jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			return []byte(c.Auth.ApiKey), nil
		}),
	))
	srv := http.NewServer(opts...)
	v1.RegisterBlogsHTTPServer(srv, blogs)
	return srv
}
