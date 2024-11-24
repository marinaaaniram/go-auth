package repository

import (
	"context"

	"go-auth/internal/model"
)

// Describe Auth pg repository interface
type AuthRepository interface {
	Login(ctx context.Context, auth *model.Auth) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, refreshToken string) (string, error)
}

// Describe Access pg repository interface
type AccessRepository interface {
	Check(ctx context.Context, endpointAddress string) error
}

// Describe User pg repository interface
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}

// Describe User redis repository interface
type UserRedisRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
