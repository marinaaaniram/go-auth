package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Update User in desc layer
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, converter.FromDescUpdateToUser(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}