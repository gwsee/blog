package data

import (
	account "blog/api/account/v1"
	blogs "blog/api/blogs/v1"
	"blog/app/bff/internal/conf"
	"context"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewAccountRepo, NewBlogsRepo)

// Data .
type Data struct {
	ac  account.AccountClient
	bc  blogs.BlogsClient
	bcm blogs.BlogsCommentClient
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
	return data, cleanup, nil
}

func (l *Data) NewAccountClient() error {
	endpoint := "discovery:///app-account"
	conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis))
	if err != nil {
		return err
	}
	l.ac = account.NewAccountClient(conn)
	return nil
}
func (l *Data) NewBlogsClient() error {
	endpoint := "discovery:///app-blogs"
	conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis))
	if err != nil {
		return err
	}
	l.bc = blogs.NewBlogsClient(conn)
	return nil
}
func (l *Data) NewBlogsClientClient() error {
	endpoint := "discovery:///app-blogs"
	conn, err := grpc.Dial(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(l.dis))
	if err != nil {
		return err
	}
	l.bcm = blogs.NewBlogsCommentClient(conn)
	return nil
}
