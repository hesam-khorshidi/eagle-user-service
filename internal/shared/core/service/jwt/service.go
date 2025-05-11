package jwt

import (
	"time"

	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
	sharedinbound "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
)

var _ inbound.JWTService = (*Service)(nil)

type Config struct {
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
	JWTIssuer          string
	JWTAudience        string
}
type Service struct {
	config Config
	idGen  sharedinbound.IDGenerator
}

func New(cfg Config, idGen sharedinbound.IDGenerator) *Service {
	return &Service{
		config: cfg,
		idGen:  idGen,
	}
}
