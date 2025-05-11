package user

import (
	"context"

	"github.com/hesam-khorshidi/eagle-user-service/internal/user/adapter/outbound/sql/model"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
	"github.com/pkg/errors"
)

func (r *Repository) Create(ctx context.Context, user domain.User) error {
	db, dbErr := r.db.GetTX(ctx, nil)
	if dbErr != nil {
		return errors.Wrap(dbErr, "error on user repository")
	}

	modelUser := model.ToUser(user)
	_, err := db.NewInsert().
		Model(&modelUser).
		Exec(ctx)

	if err != nil {
		return errors.Wrap(err, "error on user repository")
	}
	return nil
}
