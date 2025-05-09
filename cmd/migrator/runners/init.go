package runners

import (
	"log"

	"github.com/hesam-khorshidi/eagle-user-service/cmd/migrator"
	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "initialize database migrations",
	Long:  `Initialize database migration tables.`,
	Run:   runInit,
}

func runInit(c *cobra.Command, _ []string) {
	m := migrator.GetMigrator()
	err := m.Init(c.Context())
	if err != nil {
		log.Fatalf("migration init error: %v", err)
	}
	log.Println("migration up done ✅")
}
