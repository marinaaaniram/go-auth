package user

import (
	"context"
)

// GetRefreshToken Auth in repository layer
func (r *repo) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}
