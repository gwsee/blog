package server

import (
	"blog/api/bff/v1"
	"blog/app/bff/internal/conf"
	"blog/app/bff/internal/server/middleware"
	"blog/app/bff/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, tp *trace.TracerProvider,
	account *service.AccountService, blogs *service.BlogsService, blogsComment *service.BlogsCommentService,
	tool *service.ToolsService, travel *service.TravelService, user *service.UserService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			metadata.Server(),
			recovery.Recovery(),
			middleware.Auth(func(token *jwtv5.Token) (interface{}, error) {
				return []byte("gwsee"), nil
			}),
			selector.Server(
				jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
					return []byte(c.Auth.ApiKey), nil
				},
					jwt.WithSigningMethod(jwtv5.SigningMethodHS256), jwt.WithClaims(func() jwtv5.Claims {
						return &jwtv5.MapClaims{}
					})),
			).
				Match(NewWhiteListMatcher()).
				Build(),
			//tracing.Server(
			//	tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			ratelimit.Server(), //// 默认 bbr limiter
		),
	}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterAccountServer(srv, account)
	v1.RegisterBlogsServer(srv, blogs)
	v1.RegisterBlogsCommentServer(srv, blogsComment)
	v1.RegisterToolsServer(srv, tool)
	v1.RegisterTravelsServer(srv, travel)
	v1.RegisterUserServer(srv, user)
	return srv
}
