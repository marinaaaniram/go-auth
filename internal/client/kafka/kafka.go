package kafka

import (
	"context"

	"go-auth/internal/client/kafka/consumer"
)

type Consumer interface {
	Consume(ctx context.Context, handler consumer.Handler) (err error)
	Close() error
}

type Producer interface {
	SendMessage(ctx context.Context, data []byte) (err error)
	Close() error
}
