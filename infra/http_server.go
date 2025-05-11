package infra

import (
	"context"
	"errors"
	"net/http"
	"slices"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/healthcheck"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type HTTPServerConfig struct {
	Debug      bool
	Protocol   string
	Host       string
	Port       string
	ApiPrefix  string
	ApiVersion string
	LogEnable  bool
}

func NewHttpServer(txDB *TxDB) (*fiber.App, error) {
	f := fiber.New()
	f.Use(cors.New())
	f.Use(WithTransaction(txDB))

	f.Use(recover.New())
	f.Use(healthcheck.New())

	return f, nil
}

var failedStatuses = []int{http.StatusBadRequest, http.StatusInternalServerError}

func WithTransaction(txDB *TxDB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var err error

		txErr := txDB.RunInTransaction(c.Context(), func(ctx context.Context) bool {
			c.SetUserContext(ctx)
			err = c.Next()

			if err != nil {
				return false
			}
			if slices.Contains(failedStatuses, c.Response().StatusCode()) {
				return false
			}
			return true
		})

		return errors.Join(err, txErr)
	}
}
