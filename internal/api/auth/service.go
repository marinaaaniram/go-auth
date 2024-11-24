package user

import (
	"go-auth/internal/service"
	desc "go-auth/pkg/auth_v1"
)

type Implementation struct {
	desc.UnimplementedAuthV1Server
	authService service.AuthService
}

// Create Auth implementation
func NewAuthImplementation(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
