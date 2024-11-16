package cache

import (
	"github.com/marinaaaniram/go-auth/internal/repository"
	"github.com/marinaaaniram/go-auth/internal/service"
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
