package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// Get user_v1 in service layer
func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	userObj, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, errors.ErrPointerIsNil("userObj")
	}

	return userObj, nil
}
