package user

import (
	"context"
	"go-auth/internal/errors"
	"go-auth/internal/model"
	"go-auth/internal/utils"
)

// GetAccessToken Auth in service layer
func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	accessToken, err := s.authRepository.GetAccessToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	claims, err := utils.VerifyToken(refreshToken, []byte(refreshTokenSecretKey))
	if err != nil {
		return "", errors.ErrInvalidRefreshToken
	}

	// Можем слазать в базу или в кэш за доп данными пользователя

	accessToken, err = utils.GenerateToken(model.UserInfo{
		Username: claims.Username,
		// Это пример, в реальности роль должна браться из базы или кэша
		Role: "admin",
	},
		[]byte(accessTokenSecretKey),
		accessTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
