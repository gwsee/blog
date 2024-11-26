package server

import (
	v1 "blog/api/bff/v1"
	"blog/app/bff/internal/conf"
	"blog/app/bff/internal/server/middleware"
	"blog/app/bff/internal/service"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/handlers"
	"go.opentelemetry.io/otel/sdk/trace"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Bootstrap, logger log.Logger, tp *trace.TracerProvider,
	account *service.AccountService, blogs *service.BlogsService,
	blogsComment *service.BlogsCommentService) *http.Server {
	opts := []http.ServerOption{
		//http.ResponseEncoder(JsonResponse),
		http.Middleware(
			recovery.Recovery(),
			metadata.Server(),
			//tracing.Server(
			//	tracing.WithTracerProvider(tp)),
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
	v1.RegisterBlogsHTTPServer(srv, blogs)
	v1.RegisterBlogsCommentHTTPServer(srv, blogsComment)
	return srv
}

// MarshalOptions is a configurable JSON format marshaller.
var MarshalOptions = protojson.MarshalOptions{
	UseProtoNames:   true, // 使用原生的 proto 字段名
	UseEnumNumbers:  false,
	EmitUnpopulated: true,
}

// JsonResponse 对于数字字符串问题 先不处理；应该是bff层的输出在定义上不使用pb要稍微友好一点
func JsonResponse(w http.ResponseWriter, r *http.Request, v interface{}) (err error) {
	var data []byte
	switch m := v.(type) {
	case json.Marshaler:
		data, err = m.MarshalJSON()
	case proto.Message:
		data, err = MarshalOptions.Marshal(m)
	default:
		data, err = json.Marshal(m)
	}
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	_, err = w.Write(data)
	if err != nil {
		return err
	}
	return nil
}
