package server

import (
	"blog/app/tools/internal/conf"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/registry"
	client "go.etcd.io/etcd/client/v3"
)

func NewEtcdRegistrar(conf *conf.Etcd) registry.Registrar {
	cli, err := client.New(client.Config{
		Endpoints: conf.Hosts,
	})
	if err != nil {
		panic(err)
	}
	return etcd.New(cli)
}
