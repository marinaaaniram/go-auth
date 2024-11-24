package producer

import (
	"go-auth/internal/client/kafka"
	"go-auth/internal/repository"
	def "go-auth/internal/service"
)

var _ def.UserProducerService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	producer       kafka.Producer
}

// Create producer service
func NewUserProducerService(
	userRepository repository.UserRepository,
	producer kafka.Producer,
) *service {
	return &service{
		userRepository: userRepository,
		producer:       producer,
	}
}
