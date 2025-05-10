package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/adapter/inbound/http"
	"time"
)

type RegisterRequest struct {
	Password    string `json:"password"`
	Email       string `json:"email"`
	DisplayName string `json:"display_name"`
}

func (c Controller) Register(ctx *fiber.Ctx) error {
	var request RegisterRequest
	if err := ctx.BodyParser(&request); err != nil {
		return http.BadRequest(ctx, err)
	}
	token, err := c.authService.Register(ctx.UserContext(), request.ToDomain())
	if err != nil {
		return http.InternalError(ctx, err)
	}
	return http.Created(ctx, token)
}

func (r RegisterRequest) ToDomain() domain.User {
	now := time.Now()
	return domain.User{
		Email:       r.Email,
		Password:    r.Password,
		DisplayName: r.DisplayName,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}
