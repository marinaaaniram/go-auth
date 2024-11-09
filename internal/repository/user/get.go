package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
	converterRepo "github.com/marinaaaniram/go-auth/internal/repository/user/converter"
	modelRepo "github.com/marinaaaniram/go-auth/internal/repository/user/model"
)

// Get User in repository layer
func (r *repo) Get(ctx context.Context, user *model.User) (*model.User, error) {
	if user == nil {
		return nil, errors.ErrPointerIsNil("user")
	}

	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, errors.ErrFailedToBuildQuery(err)
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var repoUser modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&repoUser.ID, &repoUser.Name, &repoUser.Email, &repoUser.Role, &repoUser.CreatedAt, &repoUser.UpdatedAt)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, errors.ErrObjectNotFount("user", user.ID)
		}
		return nil, errors.ErrFailedToSelectQuery(err)
	}

	return converterRepo.FromRepoToUserGet(&repoUser), nil
}
