package user

import (
	"context"

	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/errors"
)

// Create User in desc layer
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	if err := validateUser(req); err != nil {
		return nil, err
	}

	userId, err := i.userService.Create(ctx, converter.FromDescCreateToUser(req))
	if err != nil {
		return nil, err
	}

	return converter.FromUserIdToDescCreate(userId), nil
}
