package user

import (
	"github.com/marinaaaniram/go-auth/internal/client/db"
	"github.com/marinaaaniram/go-auth/internal/repository"
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
