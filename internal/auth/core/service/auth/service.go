package auth

import (
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/inbound"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/outbound"
	errorsrv "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/service/errors"
)

var _ inbound.AuthorizationService = (*Service)(nil)

type Service struct {
	userSrv outbound.UserService
	errSrv  errorsrv.Service
}

func New(userSrv outbound.UserService, errSrv errorsrv.Service) Service {
	return Service{
		userSrv: userSrv,
		errSrv:  errSrv,
	}
}
