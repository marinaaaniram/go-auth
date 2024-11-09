package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/marinaaaniram/go-auth/internal/converter"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Update User in desc layer
func (i *Implementation) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.Internal, "req is nil")
	}

	err := i.userService.Update(ctx, converter.FromDescUpdateToUser(req))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
