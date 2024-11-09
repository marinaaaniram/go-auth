package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/errors"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Get User in desc layer
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	userDesc, err := i.userService.Get(ctx, converter.FromDescGetToUser(req))
	if err != nil {
		return nil, err
	}
	if userDesc == nil {
		return nil, errors.ErrPointerIsNil("userDesc")
	}

	return &desc.GetResponse{
		User: userDesc,
	}, nil
}
