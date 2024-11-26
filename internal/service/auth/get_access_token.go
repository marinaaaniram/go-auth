package auth

import (
	"context"
	"go-auth/internal/constant"
	"go-auth/internal/model"
	"go-auth/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetAccessToken Auth in service layer
func (s *serv) GetAccessToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(constant.RefreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "Invalid refresh token")
	}

	var user *model.User
	user, err = s.userRepository.GetAuthInfo(ctx, &model.Auth{
		Email: claims.Email,
	})
	if err != nil {
		return "", err
	}

	accessToken, err := utils.GenerateToken(model.UserAuthInfo{
		Email: user.Email,
		Role:  user.Role,
	},
		[]byte(constant.AccessTokenSecretKey),
		constant.AccessTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
