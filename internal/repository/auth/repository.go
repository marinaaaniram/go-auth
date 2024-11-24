package user

import (
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"go-auth/internal/repository"
)

type repo struct {
	db db.Client
}

// Create Auth repository
func NewAuthRepository(db db.Client) repository.AuthRepository {
	return &repo{db: db}
}
