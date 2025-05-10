package domain

import (
	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	userdomain "github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
	"time"
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

func (u User) ToUserDomain() userdomain.User {
	return userdomain.User{
		ID:          u.ID,
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		Password:    u.Password,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		AvatarURL:   u.AvatarURL,
	}
}
