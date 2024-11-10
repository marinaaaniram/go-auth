package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/errors"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Get user_v1 in service layer
func (s *serv) Get(ctx context.Context, id int64) (*desc.User, error) {
	userObj, err := s.userRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, errors.ErrPointerIsNil("userObj")
	}

	return converter.FromUserToDesc(userObj), nil
}
