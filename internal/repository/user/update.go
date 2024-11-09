package user

import (
	"context"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// Update User in repository layer
func (r *repo) Update(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.ErrPointerIsNil("user")
	}

	builderUpdate := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, user.Name).
		Set(roleColumn, user.Role).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return errors.ErrFailedToUpdateQuery(err)
	}

	log.Printf("Updated %d rows", res.RowsAffected())

	return nil
}
