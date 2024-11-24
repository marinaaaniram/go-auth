package user

import (
	"context"
	"time"

	"go-auth/internal/errors"
	"go-auth/internal/model"
	"go-auth/internal/utils"
)

const (
	grpcPort   = 50051
	authPrefix = "Bearer "

	refreshTokenSecretKey = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	accessTokenSecretKey  = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="

	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 5 * time.Minute
)

// Login Auth in service layer
func (s *serv) Login(ctx context.Context, auth *model.Auth) (string, error) {
	if auth == nil {
		return "", errors.ErrPointerIsNil("auth")
	}

	// refreshToken, err := s.authRepository.Login(ctx, auth)
	// if err != nil {
	// 	return "", err
	// }

	refreshToken, err := utils.GenerateToken(model.UserInfo{
		Username: auth.Email,
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: "admin",
	},
		[]byte(refreshTokenSecretKey),
		refreshTokenExpiration,
	)
	if err != nil {
		return "", errors.ErrGenerateToken
	}

	return refreshToken, nil
}
