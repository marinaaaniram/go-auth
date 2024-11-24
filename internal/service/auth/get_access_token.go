package user

import (
	"context"
)

// GetAccessToken Auth in service layer
func (s *serv) GetAccessToken(ctx context.Context, accessToken string) (string, error) {
	accessToken, err := s.authRepository.GetAccessToken(ctx, accessToken)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}
