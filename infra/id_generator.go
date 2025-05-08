package infra

import (
	"weasel/internal/shared/port/inbound"
	"weasel/internal/shared/service/idgen"
)

type IDGeneratorConfig struct {
	NodeID int64
}

func NewIDGenerator(cfg IDGeneratorConfig) inbound.IDGenerator {
	return idgen.NewIDGenerator(cfg.NodeID)
}
