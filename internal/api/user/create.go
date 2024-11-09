package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/errors"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Create User in desc layer
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	if err := validateName(req.Name); err != nil {
		return nil, err
	}

	if err := validateEmail(req.Email); err != nil {
		return nil, err
	}

	if err := validatePassword(req.Password, req.PasswordConfirm); err != nil {
		return nil, err
	}

	if err := validateRole(req.Role); err != nil {
		return nil, err
	}

	userId, err := i.userService.Create(ctx, converter.FromDescCreateToUser(req))
	if err != nil {
		return nil, err
	}

	return converter.FromUserIdToDescCreate(userId), nil
}
