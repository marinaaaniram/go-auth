package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/model"
	converterRepo "github.com/marinaaaniram/go-auth/internal/repository/user/converter"
	modelRepo "github.com/marinaaaniram/go-auth/internal/repository/user/model"
)

// Get user_v1 in repository layer
func (r *repo) Get(ctx context.Context, user *model.User) (*model.User, error) {
	builderSelect := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builderSelect.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.Get",
		QueryRaw: query,
	}

	var repoUser modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&repoUser.ID, &repoUser.Name, &repoUser.Email, &repoUser.Role, &repoUser.CreatedAt, &repoUser.UpdatedAt)
	if err != nil {
		if err.Error() == pgx.ErrNoRows.Error() {
			return nil, status.Errorf(codes.NotFound, "User with id %d not found", user.ID)
		}
		return nil, status.Errorf(codes.Internal, "Failed to query user: %v", err)
	}

	return converterRepo.FromRepoToUser(&repoUser), nil
}
