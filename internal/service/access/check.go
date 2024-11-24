package user

import (
	"context"
	"go-auth/internal/model"
	"go-auth/internal/utils"
	"strings"
	"time"

	"google.golang.org/grpc/metadata"

	"go-auth/internal/errors"
)

const (
	grpcPort   = 50051
	authPrefix = "Bearer "

	refreshTokenSecretKey = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	accessTokenSecretKey  = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="

	refreshTokenExpiration = 60 * time.Minute
	accessTokenExpiration  = 5 * time.Minute
)

var accessibleRoles map[string]string

// Возвращает мапу с адресом эндпоинта и ролью, которая имеет доступ к нему
func (s *serverAccess) accessibleRoles(ctx context.Context) (map[string]string, error) {
	if accessibleRoles == nil {
		accessibleRoles = make(map[string]string)

		// Лезем в базу за данными о доступных ролях для каждого эндпоинта
		// Можно кэшировать данные, чтобы не лезть в базу каждый раз

		// Например, для эндпоинта /note_v1.NoteV1/Get доступна только роль admin
		accessibleRoles[model.ExamplePath] = "admin"
	}

	return accessibleRoles, nil
}

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

	if !strings.HasPrefix(authHeader[0], authPrefix) {
		return errors.ErrInvalidAuthHeaderFormat
	}

	accessToken := strings.TrimPrefix(authHeader[0], authPrefix)

	claims, err := utils.VerifyToken(accessToken, []byte(accessTokenSecretKey))
	if err != nil {
		return errors.ErrInvalidAccessToken
	}

	accessibleMap, err := s.accessibleRoles(ctx)
	if err != nil {
		return errors.ErrGetAccessibleRole
	}

	role, ok := accessibleMap[endpointAddress]
	if !ok {
		return nil
	}

	if role == claims.Role {
		return nil
	}

	return errors.ErrAccessDenied
}
