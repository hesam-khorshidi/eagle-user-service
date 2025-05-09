package errors

import (
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
)

var _ inbound.ErrorService = (*Service)(nil)

type Service struct{}

func New() Service {
	return Service{}
}
