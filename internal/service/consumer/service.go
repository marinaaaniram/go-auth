package consumer

import (
	"github.com/marinaaaniram/go-auth/internal/client/kafka"
	"github.com/marinaaaniram/go-auth/internal/repository"
	def "github.com/marinaaaniram/go-auth/internal/service"
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
