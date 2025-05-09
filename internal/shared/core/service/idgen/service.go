package idgen

import (
	"log"

	"github.com/bwmarrin/snowflake"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
)

var _ inbound.IDGenerator = (*Service)(nil)

type Service struct {
	node *snowflake.Node
}

func NewIDGenerator(id int64) Service {
	node, err := snowflake.NewNode(id)
	if err != nil {
		log.Fatal(err)
	}
	return Service{node: node}
}

func (g Service) ID() valueobject.ID {
	return valueobject.ID(g.node.Generate().Int64())
}
