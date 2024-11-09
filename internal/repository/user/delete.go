package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// Delete User in repository layer
func (r *repo) Delete(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.ErrPointerIsNil("user")
	}

	builderSelect := sq.Select("COUNT(*)").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: user.ID})

	selectQuery, args, err := builderSelect.ToSql()
	if err != nil {
		return errors.ErrFailedToBuildQuery(err)
	}

	selectQ := db.Query{
		Name:     "user_repository.SelectId",
		QueryRaw: selectQuery,
	}

	var count int
	err = r.db.DB().QueryRowContext(ctx, selectQ, args...).Scan(&count)
	if err != nil {
		return errors.ErrFailedToSelectQuery(err)
	}

	if count == 0 {
		return errors.ErrObjectNotFount("user", user.ID)
	}

	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errors.ErrFailedToDeleteQuery(err)
	}

	return nil
}
