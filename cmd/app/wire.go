//go:build wireinject
// +build wireinject

package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"weasel/config"
	"weasel/infra"
	shareddom "weasel/internal/shared/adapter/inbound/http"
)

var infraSet = wire.NewSet(
	infra.NewHttpServer,
	infra.NewIDGenerator,
	infra.NewDBWithTX,
	infra.NewRedisClient,
	provideHttpDependencies,
	provideRunInTransaction,
)

var configSet = wire.NewSet(
	provideIDGeneratorConfig,
	provideHttpServerConfig,
	provideDatabaseConfig,
)

var inboundSet = wire.NewSet()

var serviceSet = wire.NewSet()

var outboundSet = wire.NewSet()

func InitHttp(_ config.Config) (Http, func(), error) {
	wire.Build(infraSet, configSet, provideHttp)
	return Http{}, nil, nil
}

func provideHttpDependencies(f *fiber.App, cfg infra.HTTPServerConfig) shareddom.Dependencies {
	return shareddom.Dependencies{
		Fiber:  f,
		Prefix: cfg.ApiPrefix + cfg.ApiVersion,
		Debug:  cfg.Debug,
	}
}

func provideDatabaseConfig(cfg config.Config) infra.DatabaseConfig {
	return infra.DatabaseConfig{
		DatabasePort:       cfg.DatabasePort,
		DatabaseHost:       cfg.DatabaseHost,
		DatabaseName:       cfg.DatabaseName,
		DatabaseUsername:   cfg.DatabaseUsername,
		DatabasePassword:   cfg.DatabasePassword,
		DatabaseTimezone:   cfg.DatabaseTimezone,
		DatabaseSslMode:    cfg.DatabaseSSLMode,
		DatabaseLogEnabled: cfg.LoggingEnabled,
	}
}

func provideHttpServerConfig(cfg config.Config) infra.HTTPServerConfig {
	return infra.HTTPServerConfig{
		Debug:      cfg.AppEnvironment == "development",
		Protocol:   cfg.ServerProtocol,
		Host:       cfg.ServerHost,
		Port:       cfg.ServerPort,
		ApiPrefix:  cfg.ServerApiPrefix,
		ApiVersion: cfg.ServerApiVersion,
		LogEnable:  cfg.LoggingEnabled,
	}
}

func provideIDGeneratorConfig(cfg config.Config) infra.IDGeneratorConfig {
	return infra.IDGeneratorConfig{
		NodeID: cfg.IdGeneratorNodeID,
	}
}

func provideRunInTransaction(txDB *infra.TxDB) infra.IRunInTransaction {
	return txDB.RunInTransaction
}

func provideHttp(
	server *fiber.App,
) Http {
	return newHttp(server)
}
