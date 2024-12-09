package repository

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/model"
)

// Describe Access pg repository interface
type AccessRepository interface {
	GetAccessibleRoles(ctx context.Context, endpointAddress string) ([]string, error)
}

// Describe Access redis repository interface
type AccessRedisRepository interface {
	Create(ctx context.Context, accessibleRoles []string, endpointAddress string) error
	Get(ctx context.Context, endpointAddress string) ([]string, error)
}

// Describe User pg repository interface
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error

	GetAuthInfo(ctx context.Context, auth *model.AuthInput) (*model.User, error)
}

// Describe User redis repository interface
type UserRedisRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
