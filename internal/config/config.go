package config

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

type GRPCConfig interface {
	Address() string
}

type HTTPConfig interface {
	Address() string
}

type PGConfig interface {
	DSN() string
}

type RedisConfig interface {
	Address() string
	ConnectionTimeout() time.Duration
	MaxIdle() int
	IdleTimeout() time.Duration
}

type SwaggerConfig interface {
	Address() string
}

type StorageConfig interface {
	Mode() string
}

type KafkaConsumerConfig interface {
	Brokers() []string
	GroupID() string
	TopicName() string
	Config() *sarama.Config
}

type KafkaProducerConfig interface {
	Brokers() []string
	TopicName() string
	Config() *sarama.Config
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
