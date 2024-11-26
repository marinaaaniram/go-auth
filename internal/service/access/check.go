package access

import (
	"context"
	"strings"

	"go-auth/internal/constant"
	"go-auth/internal/errors"
	"go-auth/internal/utils"

	"google.golang.org/grpc/metadata"
)

// Check Access in service layer
func (s *serv) Check(ctx context.Context, endpointAddress string) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return errors.ErrMetedataNotProvided
	}

	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return errors.ErrAuthHeaderNotProvided
	}

	if !strings.HasPrefix(authHeader[0], constant.AuthPrefix) {
		return errors.ErrInvalidAuthHeaderFormat
	}

	accessToken := strings.TrimPrefix(authHeader[0], constant.AuthPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(constant.AccessTokenSecretKey))
	if err != nil {
		return errors.ErrInvalidAccessToken
	}

	accessibleRoles, err := s.getAccessibleRoles(ctx, endpointAddress)
	if err != nil {
		return errors.ErrGetAccessibleRole(err)
	}

	for _, i := range accessibleRoles {
		if i == claims.Role {
			return nil
		}
	}

	return errors.ErrAccessDenied
}
