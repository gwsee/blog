package server

import (
	"blog/api/account/v1"
	"blog/app/account/internal/conf"
	"blog/app/account/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
	"go.opentelemetry.io/otel/sdk/trace"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, tp *trace.TracerProvider, account *service.AccountService) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			//tracing.Server(
			//	tracing.WithTracerProvider(tp)),
			logging.Server(logger),
			ratelimit.Server(), //// 默认 bbr limiter
		),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
		)),
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
	srv := http.NewServer(opts...)
	v1.RegisterAccountHTTPServer(srv, account)
	return srv
}
