package user

import (
	"context"
	"database/sql"
	"fmt"

	sharederr "github.com/hesam-khorshidi/eagle-user-service/internal/shared/errors"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/adapter/outbound/sql/model"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain"
	"github.com/hesam-khorshidi/eagle-user-service/internal/user/core/domain/valueobject"
	"github.com/pkg/errors"
)

func (r *Repository) FindBy(ctx context.Context, field valueobject.UserField, value any) (*domain.User, error) {
	db, dbErr := r.db.GetTX(ctx, nil)
	if dbErr != nil {
		return nil, errors.Wrap(dbErr, "error on user repository")
	}

	var user model.User
	err := db.NewSelect().
		Model(&user).
		Where(fmt.Sprintf("%s = ?", field), value).
		Scan(ctx)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sharederr.ErrEntityNotFound
		}
		return nil, errors.Wrap(err, "error on user repository")
	}

	domainUser := user.ToDomain()
	return &domainUser, nil
}
