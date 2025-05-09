package runners

import (
	"context"
	"log"

	"github.com/hesam-khorshidi/eagle-user-service/cmd/migrator"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var UpCmd = &cobra.Command{
	Use:   "up",
	Short: "Update database migration",
	Long:  "Update database migration",
	Run:   runUp,
}

func runUp(c *cobra.Command, _ []string) {
	if err := up(c.Context()); err != nil {
		log.Fatalf("migration up error: %v", err)
	}
	log.Println("migration up done ✅")
}

func up(ctx context.Context) error {
	m := migrator.GetMigrator()
	err := m.Lock(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer m.Unlock(ctx)

	group, err := m.Migrate(ctx)
	if err != nil {
		return errors.Wrap(err, "migrate error")
	}
	if group.IsZero() {
		return errors.New("migration group is zero")
	}
	return nil
}
