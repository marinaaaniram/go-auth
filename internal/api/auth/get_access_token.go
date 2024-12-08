package auth

import (
	"context"

	desc "github.com/marinaaaniram/go-auth/pkg/auth_v1"

	"github.com/marinaaaniram/go-auth/internal/errors"
)

// GetAccessToken Auth in desc layer
func (i *Implementation) GetAccessToken(ctx context.Context, req *desc.GetAccessTokenRequest) (*desc.GetAccessTokenResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	accessToken, err := i.authService.GetAccessToken(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &desc.GetAccessTokenResponse{
		AccessToken: accessToken,
	}, nil
}
