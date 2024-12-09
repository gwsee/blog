package data

import (
	account "blog/api/account/v1"
	blogs "blog/api/blogs/v1"
	tools "blog/api/tools/v1"
	travel "blog/api/travel/v1"
	user "blog/api/user/v1"

	"blog/app/bff/internal/conf"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAccountRepo, NewBlogsRepo, NewTravelRepo, NewUserRepo, NewToolsRepo)

// Data .
type Data struct {
	ac  account.AccountClient
	bc  blogs.BlogsClient
	bcm blogs.BlogsCommentClient
	uc  user.UserClient
	tc  travel.TravelClient
	t   tools.ToolsClient
	dis registry.Discovery
}

// NewData .
func NewData(c *conf.Etcd, logger log.Logger) (*Data, func(), error) {
	data := &Data{}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	client, err := clientv3.New(clientv3.Config{
		Endpoints: c.Hosts,
	})
	if err != nil {
		return nil, cleanup, err
	}
	data.dis = etcd.New(client)
	if err = data.NewAccountClient(); err != nil {
		return nil, cleanup, err
	}
	if err = data.NewBlogsClient(); err != nil {
		return nil, cleanup, err
	}
	if err = data.NewBlogsClientClient(); err != nil {
		return nil, cleanup, err
	}
	if err = data.NewUserClient(); err != nil {
		return nil, cleanup, err
	}
	if err = data.NewTravelClient(); err != nil {
		return nil, cleanup, err
	}
	if err = data.NewToolsClient(); err != nil {
		return nil, cleanup, err
	}
	return data, cleanup, nil
}

func (l *Data) NewAccountClient() error {
	endpoint := "discovery:///app-account"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis),
		grpc.WithMiddleware(recovery.Recovery(), metadata.Client(),
			jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
				return []byte("gwsee"), nil
			}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
				jwt.WithClaims(func() jwtv5.Claims {
					return &jwtv5.MapClaims{}
				}),
			),
		))
	if err != nil {
		return err
	}
	l.ac = account.NewAccountClient(conn)
	return nil
}
func (l *Data) NewBlogsClient() error {
	endpoint := "discovery:///app-blogs"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis), grpc.WithMiddleware(
		recovery.Recovery(),
		metadata.Client(),
		jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
			return []byte("gwsee"), nil
		}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
			jwt.WithClaims(func() jwtv5.Claims {
				return &jwtv5.MapClaims{}
			}),
		),
	))
	if err != nil {
		return err
	}
	l.bc = blogs.NewBlogsClient(conn)
	return nil
}
func (l *Data) NewBlogsClientClient() error {
	endpoint := "discovery:///app-blogs"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis),
		grpc.WithMiddleware(
			recovery.Recovery(),
			metadata.Client(),
			jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
				return []byte("gwsee"), nil
			}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
				jwt.WithClaims(func() jwtv5.Claims {
					return &jwtv5.MapClaims{}
				}),
			),
		))

	if err != nil {
		return err
	}
	l.bcm = blogs.NewBlogsCommentClient(conn)
	return nil
}
func (l *Data) NewUserClient() error {
	endpoint := "discovery:///app-user"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis), grpc.WithMiddleware(
		recovery.Recovery(),
		metadata.Client(),
		jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
			return []byte("gwsee"), nil
		}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
			jwt.WithClaims(func() jwtv5.Claims {
				return &jwtv5.MapClaims{}
			}),
		),
	))
	if err != nil {
		return err
	}
	l.uc = user.NewUserClient(conn)
	return nil
}
func (l *Data) NewTravelClient() error {
	endpoint := "discovery:///app-travel"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis), grpc.WithMiddleware(
		recovery.Recovery(),
		metadata.Client(),
		jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
			return []byte("gwsee"), nil
		}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
			jwt.WithClaims(func() jwtv5.Claims {
				return &jwtv5.MapClaims{}
			}),
		),
	))
	if err != nil {
		return err
	}
	l.tc = travel.NewTravelClient(conn)
	return nil
}
func (l *Data) NewToolsClient() error {
	endpoint := "discovery:///app-tools"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis), grpc.WithMiddleware(
		recovery.Recovery(),
		metadata.Client(),
		jwt.Client(func(token *jwtv5.Token) (interface{}, error) {
			return []byte("gwsee"), nil
		}, jwt.WithSigningMethod(jwtv5.SigningMethodHS256),
			jwt.WithClaims(func() jwtv5.Claims {
				return &jwtv5.MapClaims{}
			}),
		),
	))
	if err != nil {
		return err
	}
	l.t = tools.NewToolsClient(conn)
	return nil
}
