package idgen

import (
	"github.com/bwmarrin/snowflake"
	"log"
	"weasel/internal/shared/domain/valueobject"
	"weasel/internal/shared/port/inbound"
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
