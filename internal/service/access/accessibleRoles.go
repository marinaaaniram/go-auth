package access

import (
	"context"
)

// accessibleRoles Access in service layer
func (s *serv) getAccessibleRoles(ctx context.Context, endpointAddress string) ([]string, error) {
	if s.accessCacheService != nil {
		accessibleRoles, err := s.accessCacheService.Get(ctx, endpointAddress)
		if err == nil && accessibleRoles != nil {
			return accessibleRoles, nil
		}
	}

	var accessibleRoles []string
	accessibleRoles, err := s.accessRepository.GetAccessibleRoles(ctx, endpointAddress)
	if err != nil {
		return nil, err
	}

	if s.accessCacheService != nil {
		_, err = s.accessCacheService.Create(ctx, accessibleRoles, endpointAddress)
		if err != nil {
			return nil, err
		}
	}

	return accessibleRoles, nil
}
