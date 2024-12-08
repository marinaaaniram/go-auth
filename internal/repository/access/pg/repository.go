package access

import (
	"github.com/marinaaaniram/go-common-platform/pkg/db"

	"github.com/marinaaaniram/go-auth/internal/repository"
)

const (
	tableName = "access_endpoint"

	idColumn        = "id"
	endpointColumn  = "endpoint"
	roleColumn      = "role"
	createdAtColumn = "created_at"
	updatedAtColumn = "updated_at"
)

type repo struct {
	db db.Client
}

// Create User repository
func NewAccessRepository(db db.Client) repository.AccessRepository {
	return &repo{db: db}
}
