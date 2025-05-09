package user

import (
	"context"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/adapter/outbound/sql/model"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain/valueobject"
	"github.com/hesam-khorshidi/eagle-user-service/pkg/slices"
	"github.com/pkg/errors"
)

func (r Repository) Update(ctx context.Context, user domain.User, fields ...valueobject.UserField) error {
	db, dbErr := r.db.GetTX(ctx, nil)
	if dbErr != nil {
		return errors.Wrap(dbErr, "error on user repository")
	}

	modelUser := model.ToUser(user)
	result, err := db.NewUpdate().
		Model(&modelUser).
		WherePK().
		Column(slices.ToString(fields)...).
		Exec(ctx)

	if err != nil {
		return errors.Wrap(err, "error on user repository")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error on user repository")
	}

	if rowsAffected == 0 {
		return sharederr.ErrEntityNotFound
	}
	return nil
}
