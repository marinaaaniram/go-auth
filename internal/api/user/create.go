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

	err := validatePassword(req.Password, req.PasswordConfirm)
	if err != nil {
		return nil, err
	}

	userDesc, err := i.userService.Create(ctx, converter.FromDescCreateToUser(req))
	if err != nil {
		return nil, err
	}
	if userDesc == nil {
		return nil, errors.ErrPointerIsNil("userObj")
	}

	return &desc.CreateResponse{
		User: userDesc,
	}, nil
}
