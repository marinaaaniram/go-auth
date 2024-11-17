package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// Create User in service layer
func (s *serv) Create(ctx context.Context, user *model.User) (int64, error) {
	if user == nil {
		return 0, errors.ErrPointerIsNil("user")
	}

	if s.userProducerService != nil {
		err := s.userProducerService.SendUser(ctx, user)
		if err != nil {
			return 0, err
		}
	}

	userId, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
