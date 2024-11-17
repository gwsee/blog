package middleware

import (
	"blog/internal/constx"
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v5"
	"strings"
)

const authorizationKey = "Authorization"
const bearerWord = "Bearer"

func Auth(keyFunc jwt.Keyfunc) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			if header, ok := transport.FromServerContext(ctx); ok {
				var userInfo jwt.MapClaims
				auths := strings.SplitN(header.RequestHeader().Get(authorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
					return handler(ctx, req)
				}
				jwtToken := auths[1]
				var (
					tokenInfo *jwt.Token
				)
				tokenInfo, err = jwt.ParseWithClaims(jwtToken, &userInfo, keyFunc)
				if err != nil {
					return handler(ctx, req)
				}
				if !tokenInfo.Valid {
					return handler(ctx, req)
				}
				//if tokenInfo.Method != o.signingMethod {
				//	return nil, err
				//}
				vals, ex := (userInfo)[constx.UserInfo]
				if !ex {
					return handler(ctx, req)
				}
				valstr, ex := vals.(string)
				if !ex {
					return handler(ctx, req)
				}
				ctx = metadata.AppendToClientContext(ctx, constx.UserInfo, valstr)
				return handler(ctx, req)
			}
			return handler(ctx, req)
		}
	}
}
