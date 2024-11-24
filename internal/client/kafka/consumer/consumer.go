package consumer

import (
	"context"
	"log"
	"strings"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

type consumer struct {
	topicName            string
	consumerGroup        sarama.ConsumerGroup
	consumerGroupHandler *GroupHandler
}

// Create consumer
func NewConsumer(
	topicName string,
	consumerGroup sarama.ConsumerGroup,
	consumerGroupHandler *GroupHandler,
) *consumer {
	return &consumer{
		topicName:            topicName,
		consumerGroup:        consumerGroup,
		consumerGroupHandler: consumerGroupHandler,
	}
}

// Consume kafka
func (c *consumer) Consume(ctx context.Context, handler Handler) error {
	c.consumerGroupHandler.msgHandler = handler

	return c.consume(ctx)
}

// Close kafka
func (c *consumer) Close() error {
	return c.consumerGroup.Close()
}

// Consume logic
func (c *consumer) consume(ctx context.Context) error {
	for {
		err := c.consumerGroup.Consume(ctx, strings.Split(c.topicName, ","), c.consumerGroupHandler)
		if err != nil {
			if errors.Is(err, sarama.ErrClosedConsumerGroup) {
				return nil
			}

			return err
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}

		log.Printf("rebalancing...\n")
	}
}
