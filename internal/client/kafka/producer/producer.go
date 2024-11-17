package producer

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

type producer struct {
	topicName    string
	producerSync sarama.SyncProducer
}

// Create producer
func NewProducer(
	topicName string,
	producerSync sarama.SyncProducer,
) *producer {
	return &producer{
		topicName:    topicName,
		producerSync: producerSync,
	}
}

// SendMessage kafka
func (p *producer) SendMessage(ctx context.Context, data []byte) error {
	defer func() {
		if err := p.Close(); err != nil {
			log.Fatalf("Failed to close producer: %v\n", err.Error())
		}
	}()

	msg := &sarama.ProducerMessage{
		Topic: p.topicName,
		Value: sarama.StringEncoder(data),
	}

	partition, offset, err := p.producerSync.SendMessage(msg)
	if err != nil {
		log.Printf("Failed to send message to Kafka: %v\n", err.Error())
		return err
	}

	log.Printf("Message sent to partition %d with offset %d\n", partition, offset)

	return nil
}

func (p *producer) Close() error {
	return p.producerSync.Close()
}
