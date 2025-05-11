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
	jwtsrv "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/service/jwt"
	logsrv "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/service/log"

	sharedhttp "github.com/hesam-khorshidi/eagle-user-service/internal/shared/adapter/inbound/http"
	sharedinboundprt "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"

	userrepo "github.com/hesam-khorshidi/eagle-user-service/internal/user/adapter/outbound/sql/user"
	userinboundprt "github.com/hesam-khorshidi/eagle-user-service/internal/user/core/port/inbound"
	useroutboundprt "github.com/hesam-khorshidi/eagle-user-service/internal/user/core/port/outbound"
	usersrv "github.com/hesam-khorshidi/eagle-user-service/internal/user/core/service/user"

	authhttp "github.com/hesam-khorshidi/eagle-user-service/internal/auth/adapter/inbound/http/auth"
	authcache "github.com/hesam-khorshidi/eagle-user-service/internal/auth/adapter/outbound/redis/tokencache"
	authinboundprt "github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/inbound"
	authoutboundprt "github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/outbound"
	authsrv "github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/service/auth"
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
	provideRedisConfig,
	provideJwtConfig,
)

var inboundSet = wire.NewSet(
	authhttp.Init,
)

var serviceSet = wire.NewSet(
	logsrv.New, wire.Bind(new(sharedinboundprt.LogService), new(*logsrv.Service)),
	errorsrv.New, wire.Bind(new(sharedinboundprt.ErrorService), new(*errorsrv.Service)),
	jwtsrv.New, wire.Bind(new(sharedinboundprt.JWTService), new(*jwtsrv.Service)),
	usersrv.New, wire.Bind(new(userinboundprt.UserService), new(*usersrv.Service)),
	wire.Bind(new(authoutboundprt.UserService), new(*usersrv.Service)),
	authsrv.New, wire.Bind(new(authinboundprt.AuthorizationService), new(*authsrv.Service)),
)

var outboundSet = wire.NewSet(
	userrepo.New, wire.Bind(new(useroutboundprt.UserRepository), new(*userrepo.Repository)),
	authcache.New, wire.Bind(new(authoutboundprt.TokenCache), new(*authcache.Repository)),
)

func InitHttp(_ config.Config) (Http, func(), error) {
	wire.Build(infraSet, configSet, provideHttp, serviceSet, inboundSet, outboundSet)
	return Http{}, nil, nil
}

func provideHttpDependencies(f *fiber.App, serverConfig infra.HTTPServerConfig, jwtConfig jwtsrv.Config) sharedhttp.Dependencies {
	return sharedhttp.Dependencies{
		Fiber:          f,
		Prefix:         serverConfig.ApiPrefix + serverConfig.ApiVersion,
		Debug:          serverConfig.Debug,
		AuthMiddleware: sharedhttp.AuthorizationMiddleware(jwtConfig),
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

func provideRedisConfig(cfg config.Config) infra.RedisConfig {
	return infra.RedisConfig{
		Password: cfg.RedisPassword,
		Host:     cfg.RedisHost,
		Port:     cfg.RedisPort,
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

func provideJwtConfig(cfg config.Config) jwtsrv.Config {
	return jwtsrv.Config{
		AccessTokenSecret:  cfg.AccessTokenSecret,
		RefreshTokenSecret: cfg.RefreshTokenSecret,
		AccessTokenExpiry:  cfg.AccessTokenExpiry,
		RefreshTokenExpiry: cfg.RefreshTokenExpiry,
		JWTIssuer:          cfg.JWTIssuer,
		JWTAudience:        cfg.JWTAudience,
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
	_ *authhttp.Controller,
) Http {
	return newHttp(server)
}
