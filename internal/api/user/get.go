package user

import (
	"context"

	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/errors"
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
