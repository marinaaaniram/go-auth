package user

import (
	"go-auth/internal/repository"
	"go-auth/internal/service"
)

type serv struct {
	userRepository      repository.UserRepository
	userCacheService    service.UserCacheService
	userConsumerService service.UserConsumerService
	userProducerService service.UserProducerService
}

// Create User service
func NewUserService(
	userRepository repository.UserRepository,
	userCacheService service.UserCacheService,
	userConsumerService service.UserConsumerService,
	userProducerService service.UserProducerService,
) service.UserService {
	return &serv{
		userRepository:      userRepository,
		userCacheService:    userCacheService,
		userConsumerService: userConsumerService,
		userProducerService: userProducerService,
	}
}
