package cache

import (
	"go-auth/internal/repository"
	"go-auth/internal/service"

	"github.com/marinaaaniram/go-common-platform/pkg/db"
)

type serv struct {
	accessRedisRepository repository.AccessRedisRepository
	txManager             db.TxManager
}

// Create Access cache service
func NewAccessCacheService(accessRedisRepository repository.AccessRedisRepository, txManager db.TxManager) service.AccessCacheService {
	return &serv{
		accessRedisRepository: accessRedisRepository,
		txManager:             txManager,
	}
}
