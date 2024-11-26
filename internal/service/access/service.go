package access

import (
	"go-auth/internal/repository"
	"go-auth/internal/service"
)

type serv struct {
	userRepository        repository.UserRepository
	accessRepository      repository.AccessRepository
	accessRedisRepository repository.AccessRedisRepository
}

// Create Access service
func NewAccessService(
	userRepository repository.UserRepository,
	accessRepository repository.AccessRepository,
	accessRedisRepository repository.AccessRedisRepository,
) service.AccessService {
	return &serv{
		userRepository:        userRepository,
		accessRepository:      accessRepository,
		accessRedisRepository: accessRedisRepository,
	}
}
