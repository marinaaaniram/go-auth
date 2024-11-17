package service

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/model"
)

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
