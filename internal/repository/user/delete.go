package user

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// Delete user_v1 in repository layer
func (r *repo) Delete(ctx context.Context, user *model.User) error {
	builderSelect := sq.Select("COUNT(*)").
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: user.ID})

	selectQuery, args, err := builderSelect.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to build select query: %v", err)
	}

	selectQ := db.Query{
		Name:     "user_repository.SelectId",
		QueryRaw: selectQuery,
	}

	var count int
	err = r.db.DB().QueryRowContext(ctx, selectQ, args...).Scan(&count)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to select user: %v", err)
	}

	if count == 0 {
		return status.Errorf(codes.NotFound, "User with id %d not found", user.ID)
	}

	builderDelete := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builderDelete.ToSql()
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to build delete query: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.Delete",
		QueryRaw: query,
	}

	_, err = r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		return status.Errorf(codes.Internal, "Failed to delete user: %v", err)
	}

	return nil
}
