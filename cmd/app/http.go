package app

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/hesam-khorshidi/eagle-user-service/config"
	"github.com/spf13/cobra"
)

type Http struct {
	httpServer *fiber.App
}

var HttpCommand = &cobra.Command{
	Use:   "http",
	Short: "Starts the HTTP server",
	Run:   RunHTTP,
}

func newHttp(httpServer *fiber.App) Http {
	return Http{
		httpServer: httpServer,
	}
}

func RunHTTP(_ *cobra.Command, _ []string) {
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	app, _, err := InitHttp(*cfg)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	msg := make(chan error)
	go func() {
		msg <- app.httpServer.Listen(cfg.ServerHost + ":" + cfg.ServerPort)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		msg <- fmt.Errorf("%s", <-c)
	}()
	if err := <-msg; err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
