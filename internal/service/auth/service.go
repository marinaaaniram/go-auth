package user

import (
	"go-auth/internal/repository"
	"go-auth/internal/service"
)

type serv struct {
	authRepository repository.AuthRepository
}

// Create Auth service
func NewAuthService(
	authRepository repository.AuthRepository,
) service.AuthService {
	return &serv{
		authRepository: authRepository,
	}
}
