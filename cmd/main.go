package main

import (
	"github.com/spf13/cobra"
	"log/slog"
	"weasel/cmd/app"
)

func main() {
	cmd := &cobra.Command{
		Use:   "weasel",
		Short: "Weasel CLI",
		Long: `
$$\      $$\                                         $$\ 
$$ | $\  $$ |                                        $$ |
$$ |$$$\ $$ | $$$$$$\   $$$$$$\   $$$$$$$\  $$$$$$\  $$ |
$$ $$ $$\$$ |$$  __$$\  \____$$\ $$  _____|$$  __$$\ $$ |
$$$$  _$$$$ |$$$$$$$$ | $$$$$$$ |\$$$$$$\  $$$$$$$$ |$$ |
$$$  / \$$$ |$$   ____|$$  __$$ | \____$$\ $$   ____|$$ |
$$  /   \$$ |\$$$$$$$\ \$$$$$$$ |$$$$$$$  |\$$$$$$$\ $$ |
\__/     \__| \_______| \_______|\_______/  \_______|\__|
`,
	}
	cmd.AddCommand(app.HttpCommand)

	if err := cmd.Execute(); err != nil {
		slog.Error(err.Error())
	}
}
