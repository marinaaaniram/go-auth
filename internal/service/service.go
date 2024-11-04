package service

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/model"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Describe User service interface
type UserService interface {
	Create(ctx context.Context, user *model.User) (*desc.User, error)
	Get(ctx context.Context, user *model.User) (*desc.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, user *model.User) error
}
