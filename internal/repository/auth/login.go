package user

import (
	"context"
	"go-auth/internal/errors"
	"go-auth/internal/model"
)

// Login Auth in repository layer
func (r *repo) Login(ctx context.Context, auth *model.Auth) (string, error) {
	if auth == nil {
		return "", errors.ErrPointerIsNil("auth")
	}

	return "", nil
}
