package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/converter"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Get User in desc layer
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Internal, "req is nil")
	}

	userDesc, err := i.userService.Get(ctx, converter.FromDescGetToUser(req))
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: userDesc,
	}, nil
}
