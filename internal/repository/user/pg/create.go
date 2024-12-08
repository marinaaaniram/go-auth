package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
	"github.com/marinaaaniram/go-auth/internal/utils"
)

// Create User in repository layer
func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	if user == nil {
		return 0, errors.ErrPointerIsNil("user")
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return 0, errors.ErrFailedToHashPassword
	}

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Name, user.Email, hashedPassword, user.Role).
		Suffix(fmt.Sprintf("RETURNING %s", idColumn))

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var userId int64
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&userId)
	if err != nil {
		return 0, errors.ErrFailedToInsertQuery(err)
	}

	return userId, nil
}
