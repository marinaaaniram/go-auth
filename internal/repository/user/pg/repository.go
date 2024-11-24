package user

import (
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"go-auth/internal/repository"
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

// Create User repository
func NewUserRepository(db db.Client) repository.UserRepository {
	return &repo{db: db}
}
