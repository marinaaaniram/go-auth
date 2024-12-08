package user

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"

	"github.com/marinaaaniram/go-auth/internal/errors"
)

// Delete User in desc layer
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	err := i.userService.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
