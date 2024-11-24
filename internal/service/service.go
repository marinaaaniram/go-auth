package service

import (
	"context"

	"go-auth/internal/model"
)

// Describe Auth service interface
type AuthService interface {
	Login(ctx context.Context, auth *model.Auth) (string, error)
	GetRefreshToken(ctx context.Context, refreshToken string) (string, error)
	GetAccessToken(ctx context.Context, accessToken string) (string, error)
}

// Describe Access service interface
type AccessService interface {
	Check(ctx context.Context, endpointAddress string) error
}

// Describe User service interface
type UserService interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
}

// Describe User cache service interface
type UserCacheService interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}

// Describe User consumer service interface
type UserConsumerService interface {
	RunConsumer(ctx context.Context) error
}

// Describe User producer service interface
type UserProducerService interface {
	SendUser(ctx context.Context, user *model.User) error
}
