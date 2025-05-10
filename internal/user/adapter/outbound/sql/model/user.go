package model

import (
	"time"

	sharedvo "github.com/hesam-khorshidi/eagle-user-service/internal/shared/core/domain/valueobject"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64     `bun:"id,pk"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
	Username      string    `bun:"username"`
	Password      string    `bun:"password"`
	Email         string    `bun:"email"`
	DisplayName   string    `bun:"display_name"`
	AvatarURL     string    `bun:"avatar_url"`
}

func (u User) ToDomain() domain.User {
	return domain.User{
		ID:          sharedvo.ID(u.ID),
		CreatedAt:   u.CreatedAt,
		UpdatedAt:   u.UpdatedAt,
		Password:    u.Password,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		AvatarURL:   u.AvatarURL,
	}
}

func ToUser(user domain.User) User {
	return User{
		ID:          int64(user.ID),
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
		Password:    user.Password,
		Email:       user.Email,
		DisplayName: user.DisplayName,
		AvatarURL:   user.AvatarURL,
	}
}
