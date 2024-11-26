package access

import (
	"context"

	"go-auth/internal/errors"
	desc "go-auth/pkg/access_v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Check Access in desc layer
func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, errors.ErrPointerIsNil("req")
	}

	err := i.accessService.Check(ctx, req.GetEndpointAddress())
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
