package user

import (
	"context"

	"go-auth/internal/errors"
	"go-auth/internal/model"
)

// Update User in service layer
func (s *serv) Update(ctx context.Context, user *model.User) error {
	if user == nil {
		return errors.ErrPointerIsNil("user")
	}

	err := s.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
