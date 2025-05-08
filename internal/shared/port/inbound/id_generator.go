package inbound

import "weasel/internal/shared/domain/valueobject"

type IDGenerator interface {
	ID() valueobject.ID
}
