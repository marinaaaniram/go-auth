package user

import (
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"go-auth/internal/repository"
)

type repo struct {
	db db.Client
}

// Create Access repository
func NewAccessRepository(db db.Client) repository.AccessRepository {
	return &repo{db: db}
}
