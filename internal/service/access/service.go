package access

import (
	"github.com/marinaaaniram/go-auth/internal/repository"
	"github.com/marinaaaniram/go-auth/internal/service"
)

type serv struct {
	accessCacheService    service.AccessCacheService
	userRepository        repository.UserRepository
	accessRepository      repository.AccessRepository
	accessRedisRepository repository.AccessRedisRepository
}

// Create Access service
func NewAccessService(
	accessCacheService service.AccessCacheService,
	userRepository repository.UserRepository,
	accessRepository repository.AccessRepository,
	accessRedisRepository repository.AccessRedisRepository,
) service.AccessService {
	return &serv{
		accessCacheService:    accessCacheService,
		userRepository:        userRepository,
		accessRepository:      accessRepository,
		accessRedisRepository: accessRedisRepository,
	}
}
