package inbound

import (
	"context"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain/valueobject"
	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
)

type AuthorizationService interface {
	Login(ctx context.Context, email, password string) (*valueobject.Token, error)
	Logout(ctx context.Context, id sharedvo.ID) error
	Register(ctx context.Context, user domain.User) (*valueobject.Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (*valueobject.Token, error)
	InitializeResetPassword(ctx context.Context, id sharedvo.ID) error
	ResetPassword(ctx context.Context, resetToken string) error
}
