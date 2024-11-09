package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/converter"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Create User in desc layer
func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Internal, "req is nil")
	}

	if req.Password != req.PasswordConfirm {
		return nil, status.Error(codes.InvalidArgument, "'password' and 'password_confirm' do not match")
	}

	userDesc, err := i.userService.Create(ctx, converter.FromDescCreateToUser(req))
	if err != nil {
		return nil, err
	}
	if userDesc == nil {
		return nil, status.Error(codes.Internal, "userDesc is nil")
	}

	return &desc.CreateResponse{
		User: userDesc,
	}, nil
}
