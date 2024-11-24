package user

import (
	"context"

	"go-auth/internal/errors"
	"go-auth/internal/model"
)

// Login Auth in service layer
func (s *serv) Login(ctx context.Context, auth *model.Auth) (string, error) {
	if auth == nil {
		return "", errors.ErrPointerIsNil("auth")
	}

	refreshToken, err := s.authRepository.Login(ctx, auth)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
