package user

import (
	"context"

	"go-auth/internal/converter"
	"go-auth/internal/errors"
	desc "go-auth/pkg/user_v1"
)

// Get User in desc layer
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	userObj, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, errors.ErrPointerIsNil("userDesc")
	}

	return &desc.GetResponse{
		User: converter.FromUserToDesc(userObj),
	}, nil
}
