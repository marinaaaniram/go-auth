package user

import (
	"go-auth/internal/repository"
	"go-auth/internal/service"
)

type serv struct {
	accessRepository repository.AccessRepository
}

// Create Access service
func NewAccessService(
	accessRepository repository.AccessRepository,
) service.AccessService {
	return &serv{
		accessRepository: accessRepository,
	}
}
