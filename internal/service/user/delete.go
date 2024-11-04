package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/model"
)

func (s *serv) Delete(ctx context.Context, user *model.User) error {
	err := s.userRepository.Delete(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
