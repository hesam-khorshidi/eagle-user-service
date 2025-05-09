package inbound

import (
	"context"

	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain/valueobject"
)

type UserService interface {
	Create(ctx context.Context, user domain.User) error
	FindBy(ctx context.Context, field valueobject.UserField, value any) (*domain.User, error)
	Update(ctx context.Context, user domain.User, fields ...valueobject.UserField) error
	EmailExists(ctx context.Context, email string) (bool, error)
}
