package user

import (
	"context"

	"go-auth/internal/errors"
	desc "go-auth/pkg/access_v1"
)

// Check Access in desc layer
func (i *Implementation) Check(ctx context.Context, req *desc.CheckRequest) error {
	if req == nil {
		return errors.ErrPointerIsNil("req")
	}

	return nil
}
