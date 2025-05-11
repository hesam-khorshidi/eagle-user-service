package user

import (
	"github.com/hesam-khorshidi/eagle-user-service/infra"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/port/outbound"
)

var _ outbound.UserRepository = (*Repository)(nil)

type Repository struct {
	db *infra.TxDB
}

func New(db *infra.TxDB) *Repository {
	return &Repository{db: db}
}
