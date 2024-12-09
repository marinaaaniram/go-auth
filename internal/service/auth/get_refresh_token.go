package auth

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/constant"
	"github.com/marinaaaniram/go-auth/internal/model"
	"github.com/marinaaaniram/go-auth/internal/utils"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GetRefreshToken Auth in service layer
func (s *serv) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	claims, err := utils.VerifyToken(refreshToken, []byte(constant.RefreshTokenSecretKey))
	if err != nil {
		return "", status.Errorf(codes.Aborted, "Invalid refresh token")
	}

	var user *model.User
	user, err = s.userRepository.GetAuthInfo(ctx, &model.AuthInput{
		Email: claims.Email,
	})
	if err != nil {
		return "", err
	}

	refreshToken, err = utils.GenerateToken(model.UserAuthInfo{
		Email: user.Email,
		Role:  user.Role,
	},
		[]byte(constant.RefreshTokenSecretKey),
		constant.RefreshTokenExpiration,
	)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
