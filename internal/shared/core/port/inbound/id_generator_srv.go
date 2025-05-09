package inbound

import (
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
)

type IDGenerator interface {
	ID() valueobject.ID
}
