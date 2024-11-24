package consumer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/IBM/sarama"

	"go-auth/internal/model"
)

// Save User from consumer
func (s *service) UserSaveHandler(ctx context.Context, msg *sarama.ConsumerMessage) error {
	user := &model.User{}
	err := json.Unmarshal(msg.Value, user)
	if err != nil {
		return err
	}

	id, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}

	log.Printf("User with id %d created\n", id)

	return nil
}
