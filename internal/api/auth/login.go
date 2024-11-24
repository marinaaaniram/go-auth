package user

import (
	"context"

	"go-auth/internal/converter"
	"go-auth/internal/errors"
	desc "go-auth/pkg/auth_v1"
)

// Login Auth in desc layer
func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	refreshToken, err := i.authService.Login(ctx, converter.FromDescLoginToAuth(req))
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: refreshToken,
	}, nil
}
