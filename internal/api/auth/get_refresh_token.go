package user

import (
	"context"

	"go-auth/internal/errors"
	desc "go-auth/pkg/auth_v1"
)

// GetRefreshToken Auth in desc layer
func (i *Implementation) GetRefreshToken(ctx context.Context, req *desc.GetRefreshTokenRequest) (*desc.GetRefreshTokenResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	refreshToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetRefreshTokenResponse{
		RefreshToken: refreshToken,
	}, nil
}
