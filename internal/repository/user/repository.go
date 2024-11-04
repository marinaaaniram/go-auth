package user

import (
	"context"
	"fmt"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/model"
	"github.com/marinaaaniram/go-auth/internal/repository"
	converterRepo "github.com/marinaaaniram/go-auth/internal/repository/user/converter"
	modelRepo "github.com/marinaaaniram/go-auth/internal/repository/user/model"
)

const (
	tableName = "auth_user"

	idColumn        = "id"
	nameColumn      = "name"
	emailColumn     = "email"
	passwordColumn  = "password"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

func NewRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}

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

func (r *repo) Update(ctx context.Context, user *model.User) error {
	builderUpdate := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(nameColumn, user.Name).
		Set(roleColumn, user.Role).
		Set(updatedAtColumn, time.Now()).
		Where(sq.Eq{idColumn: user.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		log.Fatalf("Failed to build query: %v", err)
	}

	q := db.Query{
		Name:     "user_repository.Update",
		QueryRaw: query,
	}

	res, err := r.db.DB().ExecContext(ctx, q, args...)
	if err != nil {
		log.Fatalf("Failed to update user: %v", err)
	}

	log.Printf("Updated %d rows", res.RowsAffected())

	return nil
}
