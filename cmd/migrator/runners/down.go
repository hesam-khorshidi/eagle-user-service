package runners

import (
	"context"
	"log"

	"github.com/hesam-khorshidi/eagle-user-service/cmd/migrator"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var DownCmd = &cobra.Command{
	Use:   "down",
	Short: "Roll back migrations",
	Long:  `Roll back migrations.`,
	Run:   runDown,
}

func runDown(c *cobra.Command, _ []string) {
	if err := down(c.Context()); err != nil {
		log.Fatalf("migration up error: %v", err)
	}
	log.Println("migration down done ✅")
}

func down(ctx context.Context) error {
	m := migrator.GetMigrator()
	err := m.Lock(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Unlock(ctx)

	group, err := m.Rollback(ctx)
	if err != nil {
		return errors.Wrap(err, "error on migration rollback")
	}
	if group.IsZero() {
		log.Println("migration rollback group is zero")
		return nil
	}

	log.Println("rolled back to ", group.ID)
	return nil
}
