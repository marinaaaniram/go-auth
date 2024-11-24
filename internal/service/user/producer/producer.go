package producer

import (
	"context"
	"encoding/json"
	"log"

	"github.com/marinaaaniram/go-auth/internal/model"
)

// Send User to producer
func (s *service) SendUser(ctx context.Context, user *model.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("failed to marshal data: %v\n", err.Error())
	}

	err = s.producer.SendMessage(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
