package domain

import (
	"time"

	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
)

type User struct {
	ID          sharedvo.ID
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Password    string
	Email       string
	DisplayName string
	AvatarURL   string
}
