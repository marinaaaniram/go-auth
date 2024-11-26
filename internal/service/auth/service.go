package auth

import (
	"go-auth/internal/repository"
	"go-auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

// Create Auth service
func NewAuthService(
	userRepository repository.UserRepository,
) service.AuthService {
	return &serv{
		userRepository: userRepository,
	}
}
