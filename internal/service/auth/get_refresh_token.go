package user

import (
	"context"
)

// GetRefreshToken Auth in service layer
func (s *serv) GetRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	refreshToken, err := s.authRepository.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}

	return refreshToken, nil
}
