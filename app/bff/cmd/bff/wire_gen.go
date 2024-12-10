// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"blog/app/bff/internal/biz"
	"blog/app/bff/internal/conf"
	"blog/app/bff/internal/data"
	"blog/app/bff/internal/server"
	"blog/app/bff/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"go.opentelemetry.io/otel/sdk/trace"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(conf *conf.Bootstrap, tracerProvider *trace.TracerProvider, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(conf.Etcd,conf.Data.Redis, logger)
	if err != nil {
		return nil, nil, err
	}
	accountRepo := data.NewAccountRepo(dataData, logger)
	accountUseCase := biz.NewAccountUseCase(accountRepo, logger)
	accountService := service.NewAccountService(accountUseCase)

	blogsRepo := data.NewBlogsRepo(dataData, logger)
	blogsUseCase := biz.NewBlogsUseCase(blogsRepo, logger)
	blogsService := service.NewBlogsService(blogsUseCase)
	blogsCommentService := service.NewBlogsCommentService(blogsUseCase)

	travelRepo := data.NewTravelRepo(dataData, logger)
	travelUseCase :=biz.NewTravelUseCase(travelRepo, logger)
	travelService := service.NewTravelService(travelUseCase)

	toolRepo:=data.NewToolsRepo(dataData, logger)
	toolUseCase:=biz.NewToolUseCase(toolRepo, logger)
	toolService:=service.NewToolsService(toolUseCase)

	userRepo := data.NewUserRepo(dataData, logger)
	userUseCase := biz.NewUserUseCase(userRepo, logger)
	userService := service.NewUserService(userUseCase)

	grpcServer := server.NewGRPCServer(conf, logger, tracerProvider,
		accountService, blogsService, blogsCommentService,
		toolService,travelService,userService		)
	httpServer := server.NewHTTPServer(conf, logger, tracerProvider,
		accountService, blogsService, blogsCommentService,
		toolService,travelService,userService		)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
