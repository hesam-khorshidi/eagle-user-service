package auth

import (
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/inbound"
	"github.com/hesam-khorshidi/eagle-user-service/internal/auth/core/port/outbound"
	sharedinbound "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
)

var _ inbound.AuthorizationService = (*Service)(nil)

type Service struct {
	userSrv    outbound.UserService
	tokenCache outbound.TokenCache
	errSrv     sharedinbound.ErrorService
	jwtService sharedinbound.JWTService
	idGen      sharedinbound.IDGenerator
}

func New(userSrv outbound.UserService,
	errSrv sharedinbound.ErrorService,
	tokenCache outbound.TokenCache,
	jwtSrv sharedinbound.JWTService,
	idGen sharedinbound.IDGenerator,
) *Service {
	return &Service{
		userSrv:    userSrv,
		errSrv:     errSrv,
		tokenCache: tokenCache,
		jwtService: jwtSrv,
		idGen:      idGen,
	}
}
