package http

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data,omitempty" swaggerignore:"true"`
	Message string `json:"message,omitempty" swaggerignore:"true"`
	Meta    any    `json:"meta,omitempty" swaggerignore:"true"`
}

type ListMeta struct {
	TotalCount int `json:"total_count"`
}

func InternalError(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusInternalServerError).JSON(
		Response{
			Success: false,
			Message: err.Error(),
		})
}

func BadRequest(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusBadRequest).JSON(
		Response{
			Success: false,
			Message: err.Error(),
		})
}

func PaymentRequired(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusPaymentRequired).JSON(Response{
		Success: false,
		Message: err.Error(),
	})
}

func Notfound(ctx *fiber.Ctx, err error) error {
	return ctx.Status(http.StatusNotFound).
		JSON(Response{
			Success: false,
			Message: err.Error(),
		})
}

func Success(ctx *fiber.Ctx, data any) error {
	return ctx.Status(http.StatusOK).JSON(Response{
		Success: true,
		Data:    data,
	})
}

func SuccessWithMeta(ctx *fiber.Ctx, data, meta any) error {
	return ctx.Status(http.StatusOK).JSON(Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	})
}

func SuccessWithListMeta(ctx *fiber.Ctx, data any, totalCount int) error {
	return SuccessWithMeta(ctx, data, ListMeta{TotalCount: totalCount})
}

func NoContent(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusNoContent)
}

func Forbidden(ctx *fiber.Ctx, _ error) error {
	return ctx.Status(http.StatusForbidden).JSON(Response{
		Success: false,
		Message: "access denied!",
	})
}

func TooManyRequest(ctx *fiber.Ctx) error {
	return ctx.SendStatus(http.StatusTooManyRequests)
}

func Unauthorized(ctx *fiber.Ctx, _ error) error {
	return ctx.Status(http.StatusUnauthorized).JSON(Response{
		Success: false,
		Message: "unauthorized!",
	})
}
