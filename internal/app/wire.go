//go:build wireinject
// +build wireinject

package app

import (
	"github.com/google/wire"
	"github.com/lcok/vault-sync/internal/config"
	"github.com/lcok/vault-sync/internal/handler"
	"github.com/lcok/vault-sync/internal/logger"
)

var BasicProviderSet = wire.NewSet(
	config.NewEnvConfig,
	logger.NewLogger,
)

func InitServer() (*handler.Server, error) {
	wire.Build(
		BasicProviderSet,
		handler.NewServer,
	)
	return nil, nil
}

func InitTask() (*handler.Task, error) {
	wire.Build(
		BasicProviderSet,
		handler.NewTask,
	)
	return nil, nil
}
