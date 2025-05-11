package user

import (
	sharedinbound "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/port/inbound"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/port/outbound"
)

var _ inbound.UserService = (*Service)(nil)

type Service struct {
	errorSrv sharedinbound.ErrorService
	logSrv   sharedinbound.LogService
	userRepo outbound.UserRepository
	idGen    sharedinbound.IDGenerator
}

func New(
	userRepo outbound.UserRepository,
	errorSrv sharedinbound.ErrorService,
	logSrv sharedinbound.LogService,
	idGen sharedinbound.IDGenerator,
) *Service {
	return &Service{
		userRepo: userRepo,
		errorSrv: errorSrv,
		logSrv:   logSrv,
		idGen:    idGen,
	}
}
