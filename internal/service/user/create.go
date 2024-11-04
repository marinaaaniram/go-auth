package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/model"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Create user_v1 in service layer
func (s *serv) Create(ctx context.Context, user *model.User) (*desc.User, error) {
	userObj, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return converter.FromUserToDesc(userObj), nil
}
