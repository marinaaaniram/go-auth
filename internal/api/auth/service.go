package auth

import (
	desc "github.com/marinaaaniram/go-auth/pkg/auth_v1"

	"github.com/marinaaaniram/go-auth/internal/service"
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
