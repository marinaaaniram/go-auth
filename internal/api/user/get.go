package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Get User in desc layer
func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	userDesc, err := i.userService.Get(ctx, converter.FromDescGetToUser(req))
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: userDesc,
	}, nil
}
