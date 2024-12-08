package cache

import (
	"context"
)

// Create AccessibleRoles in cache
func (s *serv) Create(ctx context.Context, accessibleRoles []string, endpointAddress string) (string, error) {
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		errTx = s.accessRedisRepository.Create(ctx, accessibleRoles, endpointAddress)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.accessRedisRepository.Get(ctx, endpointAddress)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return "", err
	}

	return endpointAddress, nil
}

// Get AccessibleRoles in cache
func (s *serv) Get(ctx context.Context, endpointAddress string) ([]string, error) {
	accessibleRoles, err := s.accessRedisRepository.Get(ctx, endpointAddress)
	if err != nil {
		return nil, err
	}

	return accessibleRoles, nil
}
