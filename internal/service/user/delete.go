package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// Delete user_v1 in service layer
func (s *serv) Delete(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.ErrPointerIsNil("user")
	}

	err := s.userRepository.Delete(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
