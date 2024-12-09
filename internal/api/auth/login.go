package auth

import (
	"context"

	desc "github.com/marinaaaniram/go-auth/pkg/auth_v1"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
)

// Login Auth in desc layer
func (i *Implementation) Login(ctx context.Context, req *desc.LoginRequest) (*desc.LoginResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	var authOutput *model.AuthOutput
	authOutput, err := i.authService.Login(ctx, converter.FromDescLoginToAuth(req))
	if err != nil {
		return nil, err
	}

	return &desc.LoginResponse{
		RefreshToken: authOutput.RefreshToken,
		AccessToken:  authOutput.AccessToken,
	}, nil
}
