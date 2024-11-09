package repository

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/model"
)

// Describe User repository interface
type UserRepository interface {
	Create(ctx context.Context, user *model.User) (int64, error)
	Get(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}
