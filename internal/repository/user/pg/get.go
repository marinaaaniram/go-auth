package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
	converterRepo "github.com/marinaaaniram/go-auth/internal/repository/user/pg/converter"
	modelRepo "github.com/marinaaaniram/go-auth/internal/repository/user/pg/model"
)

// Get User in repository layer
func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	builder := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id})

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
			return nil, errors.ErrObjectNotFount("user", id)
		}
		return nil, errors.ErrFailedToSelectQuery(err)
	}

	user := converterRepo.FromRepoToUserGet(&repoUser)

	return user, nil
}
