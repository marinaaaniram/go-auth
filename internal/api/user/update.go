package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go-auth/internal/converter"
	"go-auth/internal/errors"
	desc "go-auth/pkg/user_v1"
)

// Update User in desc layer
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	err := i.userService.Update(ctx, converter.FromDescUpdateToUser(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
