package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/model"
	converterRepo "github.com/marinaaaniram/go-auth/internal/repository/user/converter"
	modelRepo "github.com/marinaaaniram/go-auth/internal/repository/user/model"
)

// Create user_v1 in repository layer
func (r *repo) Create(ctx context.Context, user *model.User) (*model.User, error) {
	builderInsert := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Name, user.Password, user.Email, user.Role).
		Suffix(fmt.Sprintf("RETURNING %s, %s, %s, %s, %s, %s", idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn))

	query, args, err := builderInsert.ToSql()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.Create",
		QueryRaw: query,
	}

	var repoUser modelRepo.User
	err = r.db.DB().QueryRowContext(ctx, q, args...).Scan(&repoUser.ID, &repoUser.Name, &repoUser.Email, &repoUser.Role, &repoUser.CreatedAt, &repoUser.UpdatedAt)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to insert user: %v", err)
	}

	return converterRepo.FromRepoToUser(&repoUser), nil
}
