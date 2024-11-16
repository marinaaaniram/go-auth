package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
	"github.com/marinaaaniram/go-auth/internal/repository/user/converter"
)

// Update User in repository layer
func (r *repo) Update(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.ErrPointerIsNil("user")
	}

	repoUserUpdate := converter.FromUserToRepoUpdate(user)

	builder := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: user.ID}).
		Set(updatedAtColumn, time.Now())

	if repoUserUpdate.Name != nil {
		builder = builder.Set(nameColumn, *repoUserUpdate.Name)
	}
	if repoUserUpdate.Role != nil {
		builder = builder.Set(roleColumn, *repoUserUpdate.Role)
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errors.ErrFailedToUpdateQuery(err)
	}

	return nil
}
