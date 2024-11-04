package user

import (
	"github.com/marinaaaniram/go-auth/internal/repository"
	"github.com/marinaaaniram/go-auth/internal/service"
)

type serv struct {
	userRepository repository.UserRepository
}

// Create User service
func NewService(userRepository repository.UserRepository) service.UserService {
	return &serv{
		userRepository: userRepository,
	}
}
