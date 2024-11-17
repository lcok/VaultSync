// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/google/wire"
	"github.com/lcok/vault-sync/internal/config"
	"github.com/lcok/vault-sync/internal/handler"
	"github.com/lcok/vault-sync/internal/logger"
)

// Injectors from wire.go:

func InitServer() (*handler.Server, error) {
	envConfig, err := config.NewEnvConfig()
	if err != nil {
		return nil, err
	}
	loggerLogger, err := logger.NewLogger(envConfig)
	if err != nil {
		return nil, err
	}
	server := handler.NewServer(envConfig, loggerLogger)
	return server, nil
}

func InitTask() (*handler.Task, error) {
	envConfig, err := config.NewEnvConfig()
	if err != nil {
		return nil, err
	}
	loggerLogger, err := logger.NewLogger(envConfig)
	if err != nil {
		return nil, err
	}
	task := handler.NewTask(loggerLogger)
	return task, nil
}

// wire.go:

var BasicProviderSet = wire.NewSet(config.NewEnvConfig, logger.NewLogger)