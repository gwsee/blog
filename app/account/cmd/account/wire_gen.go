// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"blog/app/account/internal/biz"
	_ "blog/app/account/internal/biz"
	"blog/app/account/internal/conf"
	"blog/app/account/internal/data"
	"blog/app/account/internal/server"
	"blog/app/account/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(cfg *conf.Bootstrap, tracerProvider *trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(cfg.Data, logger) //dataData
	if err != nil {
		return nil, nil, err
	}

	accountRepo := data.NewAccountRepo(dataData, logger)
	accountUserCase := biz.NewAccountUseCase(cfg.Auth, accountRepo, logger)

	accountService := service.NewAccountService(accountUserCase) //accountUsecase
	grpcServer := server.NewGRPCServer(cfg, logger, tracerProvider, accountService)
	httpServer := server.NewHTTPServer(cfg, logger, tracerProvider, accountService)
	register := server.NewEtcdRegistrar(cfg.Etcd)
	app := newApp(logger, grpcServer, httpServer, register)
	return app, func() {
		cleanup()
	}, nil
}
