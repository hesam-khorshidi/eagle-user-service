//go:build wireinject
// +build wireinject

package migrator

import (
	"github.com/google/wire"
	"github.com/hesam-khorshidi/eagle-user-service/config"
	"github.com/hesam-khorshidi/eagle-user-service/infra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

var infraSet = wire.NewSet(
	infra.NewMigrationDB,
)

var configSet = wire.NewSet(
	provideDatabaseConfig,
	provideMigrators,
	provideMigrator,
)

func InitMigrator(_ config.Config) Migrator {
	wire.Build(infraSet, configSet)
	return Migrator{}
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

func provideMigrators() *migrate.Migrations {
	return InitMigrations()
}

func provideMigrator(db *bun.DB, migrations *migrate.Migrations) Migrator {
	return Migrator{
		DB:         db,
		Migrations: migrations,
	}
}
