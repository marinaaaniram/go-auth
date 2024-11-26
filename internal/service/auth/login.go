package auth

import (
	"context"

	"go-auth/internal/constant"
	"go-auth/internal/errors"
	"go-auth/internal/model"
	"go-auth/internal/utils"
)

// Login Auth in service layer
func (s *serv) Login(ctx context.Context, auth *model.Auth) (string, error) {
	if auth == nil {
		return "", errors.ErrPointerIsNil("auth")
	}

	var user *model.User
	user, err := s.userRepository.GetAuthInfo(ctx, auth)
	if err != nil {
		return "", err
	}

	if !utils.VerifyPassword(user.Password, auth.Password) {
		return "", errors.ErrIncorrectPassword
	}

	refreshToken, err := utils.GenerateToken(model.UserAuthInfo{
		Email: user.Email,
		Role:  user.Role,
	},
		[]byte(constant.RefreshTokenSecretKey),
		constant.RefreshTokenExpiration,
	)
	if err != nil {
		return "", errors.ErrGenerateToken
	}

	return refreshToken, nil
}
