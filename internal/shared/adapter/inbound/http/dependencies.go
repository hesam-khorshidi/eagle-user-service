package http

import (
	"github.com/gofiber/fiber/v2"
)

type Dependencies struct {
	Fiber  *fiber.App
	Prefix string
	Debug  bool
}
