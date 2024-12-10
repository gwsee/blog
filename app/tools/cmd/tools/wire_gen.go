// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"blog/app/tools/internal/biz"
	"blog/app/tools/internal/conf"
	"blog/app/tools/internal/data"
	"blog/app/tools/internal/server"
	"blog/app/tools/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(conf *conf.Bootstrap, tracerProvider *trace.TracerProvider,logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(conf.Data, logger)
	if err != nil {
		return nil, nil, err
	}
	toolsRepo := data.NewToolsRepo(dataData, logger)
	toolsUsecase := biz.NewToolsUsecase(toolsRepo, logger,conf.Oss)
	toolsService := service.NewToolsService(toolsUsecase)
	grpcServer := server.NewGRPCServer(conf.Server, toolsService, logger)
	httpServer := server.NewHTTPServer(conf.Server, toolsService, logger)
	register := server.NewEtcdRegistrar(conf.Etcd)
	app := newApp(logger, grpcServer, httpServer,register)
	return app, func() {
		cleanup()
	}, nil
}
