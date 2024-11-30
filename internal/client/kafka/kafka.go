package kafka

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/client/kafka/consumer"
)

type Consumer interface {
	Consume(ctx context.Context, handler consumer.Handler) (err error)
	Close() error
}

type Producer interface {
	SendMessage(ctx context.Context, data []byte) (err error)
	Close() error
}
