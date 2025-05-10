//go:build wireinject
// +build wireinject

package app

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"github.com/hesam-khorshidi/eagle-user-service/config"
	"github.com/hesam-khorshidi/eagle-user-service/infra"

	errorsrv "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/service/errors"
	logsrv "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/service/log"

	shareddom "github.com/hesam-khorshidi/eagle-user-service/internal/shared/adapter/inbound/http"
	sharedinbound "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"

	userrepo "github.com/hesam-khorshidi/eagle-user-service/internal/user/adapter/outbound/sql/user"
	userinbound "github.com/hesam-khorshidi/eagle-user-service/internal/user/core/port/inbound"
	useroutound "github.com/hesam-khorshidi/eagle-user-service/internal/user/core/port/outbound"
	usersrv "github.com/hesam-khorshidi/eagle-user-service/internal/user/core/service/user"

	authoutbound "github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/outbound"
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
	provideLoggingLevel,
)

var inboundSet = wire.NewSet()

var serviceSet = wire.NewSet(
	logsrv.New, wire.Bind(new(sharedinbound.LogService), new(logsrv.Service)),
	errorsrv.New, wire.Bind(new(sharedinbound.ErrorService), new(errorsrv.Service)),
	usersrv.New, wire.Bind(new(userinbound.UserService), new(usersrv.Service)),
	wire.Bind(new(authoutbound.UserService), new(usersrv.Service)),
)

var outboundSet = wire.NewSet(
	userrepo.New, wire.Bind(new(useroutound.UserRepository), new(userrepo.Repository)),
)

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

func provideLoggingLevel(cfg config.Config) slog.Level {
	switch cfg.LoggingLevel {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn", "warning":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
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
