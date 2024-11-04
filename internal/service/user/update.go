package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/model"
)

func (s *serv) Update(ctx context.Context, user *model.User) error {
	err := s.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
