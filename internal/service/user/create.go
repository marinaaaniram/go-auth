package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Create user_v1 in service layer
func (s *serv) Create(ctx context.Context, user *model.User) (*desc.User, error) {
	if user == nil {
		return nil, errors.ErrPointerIsNil("user")
	}

	userObj, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, errors.ErrPointerIsNil("userObj")
	}

	return converter.FromUserToDesc(userObj), nil
}
