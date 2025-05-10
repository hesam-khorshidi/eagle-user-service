package inbound

import sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"

type JWTService interface {
	GenerateAccessToken(claims map[string]interface{}) (string, error)
	GenerateRefreshToken(claims map[string]interface{}) (sharedvo.ID, string, error)
}
