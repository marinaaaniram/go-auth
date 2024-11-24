package env

import (
	"errors"
	"os"
	"strings"

	"github.com/IBM/sarama"
)

const (
	brokersProducerEnvName = "KAFKA_BROKERS"
	topicProducerEnvName   = "KAFKA_TOPIC_NAME"
)

type kafkaProducerConfig struct {
	brokers   []string
	topicName string
}

func NewKafkaProducerConfig() (*kafkaProducerConfig, error) {
	brokersStr := os.Getenv(brokersProducerEnvName)
	if len(brokersStr) == 0 {
		return nil, errors.New("kafka brokers address not found")
	}

	brokers := strings.Split(brokersStr, ",")

	topicName := os.Getenv(topicProducerEnvName)
	if len(topicName) == 0 {
		return nil, errors.New("kafka topic name not found")
	}

	return &kafkaProducerConfig{
		brokers:   brokers,
		topicName: topicName,
	}, nil
}

func (cfg *kafkaProducerConfig) Brokers() []string {
	return cfg.brokers
}

func (cfg *kafkaProducerConfig) TopicName() string {
	return cfg.topicName
}

func (cfg *kafkaProducerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Version = sarama.V2_6_0_0
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	return config
}
