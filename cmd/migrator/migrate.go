package migrator

import (
	"embed"
	"log"

	"github.com/hesam-khorshidi/eagle-user-service/config"
	"github.com/spf13/cobra"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"
)

//go:embed migration/*.sql
var sqlMigrations embed.FS

var MigrateCommand = &cobra.Command{
	Use:   "migrate",
	Short: "manage database migrations",
}

type Migrator struct {
	DB         *bun.DB
	Migrations *migrate.Migrations
}

func InitMigrations() *migrate.Migrations {
	migrator := migrate.NewMigrations()
	if err := migrator.DiscoverCaller(); err != nil {
		log.Fatal(err)
	}
	if err := migrator.Discover(sqlMigrations); err != nil {
		log.Fatal(err)
	}
	return migrator
}

func GetMigrator() *migrate.Migrator {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatal(err)
	}
	m := InitMigrator(*cfg)
	return migrate.NewMigrator(m.DB, m.Migrations, migrate.WithMarkAppliedOnSuccess(true))
}
