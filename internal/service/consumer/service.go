package consumer

import (
	"go-auth/internal/client/kafka"
	"go-auth/internal/repository"
	def "go-auth/internal/service"
)

var _ def.UserConsumerService = (*service)(nil)

type service struct {
	userRepository repository.UserRepository
	consumer       kafka.Consumer
}

// Create consumer service
func NewUserConsumerService(
	userRepository repository.UserRepository,
	consumer kafka.Consumer,
) *service {
	return &service{
		userRepository: userRepository,
		consumer:       consumer,
	}
}
