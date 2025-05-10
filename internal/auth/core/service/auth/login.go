package auth

import (
	"context"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/domain/valueobject"
)

func (s Service) Login(ctx context.Context, email, password string) (*valueobject.Token, error) {
	//TODO implement me
	panic("implement me")
}
