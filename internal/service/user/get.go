package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// Get user_v1 in service layer
func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	if s.userCacheService != nil {
		userObj, err := s.userCacheService.Get(ctx, id)
		if err == nil && userObj != nil {
			return userObj, nil
		}
	}

	userObj, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if userObj == nil {
		return nil, errors.ErrPointerIsNil("userObj")
	}

	if s.userCacheService != nil {
		_, err = s.userCacheService.Create(ctx, userObj)
		if err != nil {
			return nil, err
		}
	}

	return userObj, nil
}
