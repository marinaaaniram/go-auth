package user

import (
	"context"
)

// GetAccessToken Auth in repository layer
func (r *repo) GetAccessToken(ctx context.Context, accessToken string) (string, error) {
	return "", nil
}
