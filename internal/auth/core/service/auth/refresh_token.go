package auth

import (
	"context"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain/valueobject"
)

func (s Service) RefreshToken(ctx context.Context, refreshToken string) (*valueobject.Token, error) {
	//TODO implement me
	panic("implement me")
}
