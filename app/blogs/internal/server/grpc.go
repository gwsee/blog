package server

import (
	v1 "blog/api/blogs/v1"
	"blog/app/blogs/internal/conf"
	"blog/app/blogs/internal/service"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, logger log.Logger, tp *trace.TracerProvider, blogs *service.BlogsService) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(
			//recovery.WithHandler(func(ctx context.Context, req, err interface{}) error {
			//	// do someting
			//	return nil
			//}),
			),
			//tracing.Server(
			//	tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			//selector.Server(
			//	jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
			//		return []byte(c.Auth.ApiKey), nil
			//	},
			//		jwt.WithSigningMethod(jwtv5.SigningMethodHS256), jwt.WithClaims(func() jwtv5.Claims {
			//			return &jwtv5.MapClaims{}
			//		})),
			//).
			//	Match(NewWhiteListMatcher()).
			//	Build(),
			metadata.Server(),
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
	v1.RegisterBlogsServer(srv, blogs)
	return srv
}
