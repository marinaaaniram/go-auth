package user

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"golang.org/x/crypto/bcrypt"

	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
)

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Create User in repository layer
func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	if user == nil {
		return 0, errors.ErrPointerIsNil("user")
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return 0, errors.ErrFailedToHashPassword
	}

	builder := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, roleColumn).
		Values(user.Name, hashedPassword, user.Email, user.Role).
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
