// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//go:build !wireinject
// +build !wireinject

package factory

import (
	"github.com/Justdanru/bhs-test/config"
	"github.com/Justdanru/bhs-test/internal/app"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/handler"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/router"
	"github.com/Justdanru/bhs-test/internal/controller/http/v1/server"
	"github.com/Justdanru/bhs-test/internal/infrastructure/repository/user"
	user2 "github.com/Justdanru/bhs-test/internal/infrastructure/service/user"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func StartApp() (*app.App, func(), error) {
	configConfig := config.NewConfig()
	errorsHandler := handler.NewErrorsHandler()
	db, cleanup, err := providePostgreSQLConnection(configConfig)
	if err != nil {
		return nil, nil, err
	}
	repositoryPostgreSQL := user.NewRepositoryPostgreSQL(db)
	service := user2.NewService(repositoryPostgreSQL)
	userHandler := handler.NewUserHandler(errorsHandler, service)
	checkUsernameHandler := handler.NewCheckUsernameHandler(errorsHandler, service)
	registerHandler := handler.NewRegisterHandler(errorsHandler, service)
	rootHandler := handler.NewRootHandler(userHandler, checkUsernameHandler, registerHandler)
	routerRouter := router.NewRouter(rootHandler)
	logger := provideLogger()
	httpServer := server.NewHTTPServer(configConfig, routerRouter, logger)
	appApp, cleanup2, err := startApp(httpServer)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	return appApp, func() {
		cleanup2()
		cleanup()
	}, nil
}
