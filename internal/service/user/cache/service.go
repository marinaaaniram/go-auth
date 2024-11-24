package cache

import (
	"go-auth/internal/repository"
	"go-auth/internal/service"

	"github.com/marinaaaniram/go-common-platform/pkg/db"
)

type serv struct {
	userRedisRepository repository.UserRedisRepository
	txManager           db.TxManager
}

// Create User cache service
func NewUserCacheService(userRedisRepository repository.UserRedisRepository, txManager db.TxManager) service.UserCacheService {
	return &serv{
		userRedisRepository: userRedisRepository,
		txManager:           txManager,
	}
}
