package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Delete user_v1 in desc layer
func (i *Implementation) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	err := i.userService.Delete(ctx, converter.FromDescDeleteToUser(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
