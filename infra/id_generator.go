package infra

import (
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/service/idgen"
)

type IDGeneratorConfig struct {
	NodeID int64
}

func NewIDGenerator(cfg IDGeneratorConfig) inbound.IDGenerator {
	return idgen.NewIDGenerator(cfg.NodeID)
}
