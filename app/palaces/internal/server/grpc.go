package server

import (
	v1 "blog/api/palaces/v1"
	"blog/app/palaces/internal/conf"
	"blog/app/palaces/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Bootstrap, palaces *service.PalacesService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
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
	v1.RegisterPalacesServer(srv, palaces)
	return srv
}
