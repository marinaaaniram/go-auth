package user

import (
	"github.com/marinaaaniram/go-auth/internal/repository"
	"github.com/marinaaaniram/go-auth/internal/service"
)

type serv struct {
	userRepository   repository.UserRepository
	userCacheService service.UserCacheService
}

// Create User service
func NewUserService(userRepository repository.UserRepository, userCacheService service.UserCacheService) service.UserService {
	return &serv{
		userRepository:   userRepository,
		userCacheService: userCacheService,
	}
}
