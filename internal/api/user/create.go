package user

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/converter"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	userDesc, err := i.userService.Create(ctx, converter.FromDescCreateToUser(req))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		User: userDesc,
	}, nil
}
