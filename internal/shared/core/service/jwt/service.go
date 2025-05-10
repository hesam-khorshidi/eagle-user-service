package jwt

import (
	"github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
	sharedinbound "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/port/inbound"
	"time"
)

var _ inbound.JWTService = (*Service)(nil)

type Config struct {
	IdGeneratorNodeID  int64
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

func New(cfg Config) Service {
	return Service{
		config: cfg,
	}
}
