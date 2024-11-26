package access

import (
	"context"
)

var accessibleRoles map[string]string

// accessibleRoles Access in service layer
func (s *serv) getAccessibleRoles(ctx context.Context, endpointAddress string) ([]string, error) {
	var accessibleRoles []string

	accessibleRoles, err := s.accessRepository.GetAccessibleRoles(ctx, endpointAddress)
	if err != nil {
		return nil, err
	}

	return accessibleRoles, nil
}
