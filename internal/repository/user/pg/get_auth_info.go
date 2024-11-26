package user

import (
	"context"
	"go-auth/internal/errors"
	"go-auth/internal/model"
	converterRepo "go-auth/internal/repository/user/pg/converter"
	modelRepo "go-auth/internal/repository/user/pg/model"

	sq "github.com/Masterminds/squirrel"

	"github.com/jackc/pgx"
	"github.com/marinaaaniram/go-common-platform/pkg/db"
)

// Login Auth in repository layer
func (r *repo) GetAuthInfo(ctx context.Context, auth *model.Auth) (*model.User, error) {
	if auth == nil {
		return nil, errors.ErrPointerIsNil("auth")
	}

	builder := sq.Select(emailColumn, roleColumn, passwordColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{emailColumn: auth.Email})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "user_repository.GetAuthInfo",
		QueryRaw: query,
	}

	var repoUser modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&repoUser.Email, &repoUser.Role, &repoUser.Password)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, errors.ErrObjectContentNotFount("user", auth.Email)
		}
		return nil, errors.ErrFailedToSelectQuery(err)
	}

	user := converterRepo.FromRepoToUserGet(&repoUser)

	return user, nil
}
