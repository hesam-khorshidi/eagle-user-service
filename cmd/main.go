package main

import (
	"log/slog"

	"github.com/hesam-khorshidi/eagle-user-service/cmd/app"
	"github.com/hesam-khorshidi/eagle-user-service/cmd/migrator"
	migrationrunners "github.com/hesam-khorshidi/eagle-user-service/cmd/migrator/runners"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "user-service",
		Short: "User service CLI",
		Long: `
$$\   $$\                                      $$$$$$\                                $$\                     
$$ |  $$ |                                    $$  __$$\                               \__|                    
$$ |  $$ | $$$$$$$\  $$$$$$\   $$$$$$\        $$ /  \__| $$$$$$\   $$$$$$\ $$\    $$\ $$\  $$$$$$$\  $$$$$$\  
$$ |  $$ |$$  _____|$$  __$$\ $$  __$$\       \$$$$$$\  $$  __$$\ $$  __$$\\$$\  $$  |$$ |$$  _____|$$  __$$\ 
$$ |  $$ |\$$$$$$\  $$$$$$$$ |$$ |  \__|       \____$$\ $$$$$$$$ |$$ |  \__|\$$\$$  / $$ |$$ /      $$$$$$$$ |
$$ |  $$ | \____$$\ $$   ____|$$ |            $$\   $$ |$$   ____|$$ |       \$$$  /  $$ |$$ |      $$   ____|
\$$$$$$  |$$$$$$$  |\$$$$$$$\ $$ |            \$$$$$$  |\$$$$$$$\ $$ |        \$  /   $$ |\$$$$$$$\ \$$$$$$$\ 
\______/ \_______/  \_______|\__|             \______/  \_______|\__|         \_/    \__| \_______| \_______|


`,
	}
	cmd.AddCommand(app.HttpCommand)

	migrator.MigrateCommand.AddCommand(migrationrunners.InitCmd)
	migrator.MigrateCommand.AddCommand(migrationrunners.UpCmd)
	migrator.MigrateCommand.AddCommand(migrationrunners.DownCmd)
	cmd.AddCommand(migrator.MigrateCommand)

	if err := cmd.Execute(); err != nil {
		slog.Error(err.Error())
	}
}
