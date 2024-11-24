package app

import (
	"context"
	"log"

	"github.com/IBM/sarama"
	redigo "github.com/gomodule/redigo/redis"
	"github.com/marinaaaniram/go-common-platform/pkg/closer"
	"github.com/marinaaaniram/go-common-platform/pkg/db"
	"github.com/marinaaaniram/go-common-platform/pkg/db/pg"
	"github.com/marinaaaniram/go-common-platform/pkg/db/transaction"

	"go-auth/internal/api/user"
	"go-auth/internal/client/cache"
	"go-auth/internal/client/cache/redis"
	"go-auth/internal/client/kafka"
	kafkaConsumer "go-auth/internal/client/kafka/consumer"
	kafkaProducer "go-auth/internal/client/kafka/producer"
	"go-auth/internal/config"
	"go-auth/internal/config/env"
	"go-auth/internal/repository"
	userRepository "go-auth/internal/repository/user/pg"
	userRedisRepository "go-auth/internal/repository/user/redis"
	"go-auth/internal/service"
	userConsumerService "go-auth/internal/service/consumer"
	userService "go-auth/internal/service/user"
	userCacheService "go-auth/internal/service/user/cache"
	userProduceService "go-auth/internal/service/user/producer"
)

type serviceProvider struct {
	pgConfig            config.PGConfig
	grpcConfig          config.GRPCConfig
	redisConfig         config.RedisConfig
	httpConfig          config.HTTPConfig
	swaggerConfig       config.SwaggerConfig
	kafkaConsumerConfig config.KafkaConsumerConfig
	kafkaProducerConfig config.KafkaProducerConfig

	dbClient  db.Client
	txManager db.TxManager

	redisPool   *redigo.Pool
	redisClient cache.RedisClient

	userRepository      repository.UserRepository
	userRedisRepository repository.UserRedisRepository

	userService         service.UserService
	userCacheService    service.UserCacheService
	userConsumerService service.UserConsumerService
	userProducerService service.UserProducerService

	consumer             kafka.Consumer
	consumerGroup        sarama.ConsumerGroup
	consumerGroupHandler *kafkaConsumer.GroupHandler

	producer     kafka.Producer
	producerSync sarama.SyncProducer

	userImpl *user.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// Get postgres config
func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

// Get GRPC config
func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

// Get redis config
func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := env.NewRedisConfig()
		if err != nil {
			log.Fatalf("failed to get redis config: %s", err.Error())
		}

		s.redisConfig = cfg
	}

	return s.redisConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) KafkaConsumerConfig() config.KafkaConsumerConfig {
	if s.kafkaConsumerConfig == nil {
		cfg, err := env.NewKafkaConsumerConfig()
		if err != nil {
			log.Fatalf("failed to get kafka consumer config: %s", err.Error())
		}

		s.kafkaConsumerConfig = cfg
	}

	return s.kafkaConsumerConfig
}

func (s *serviceProvider) KafkaProducerConfig() config.KafkaProducerConfig {
	if s.kafkaProducerConfig == nil {
		cfg, err := env.NewKafkaProducerConfig()
		if err != nil {
			log.Fatalf("failed to get kafka producer config: %s", err.Error())
		}

		s.kafkaProducerConfig = cfg
	}

	return s.kafkaProducerConfig
}

// Init db client
func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

// Init db transactions manager
func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

// Init redis pool
func (s *serviceProvider) RedisPool() *redigo.Pool {
	if s.redisPool == nil {
		s.redisPool = &redigo.Pool{
			MaxIdle:     s.RedisConfig().MaxIdle(),
			IdleTimeout: s.RedisConfig().IdleTimeout(),
			DialContext: func(ctx context.Context) (redigo.Conn, error) {
				return redigo.DialContext(ctx, "tcp", s.RedisConfig().Address())
			},
		}
	}

	return s.redisPool
}

// Init redis client
func (s *serviceProvider) RedisClient(ctx context.Context) cache.RedisClient {
	if s.redisClient == nil {
		s.redisClient = redis.NewClient(s.RedisPool(), s.RedisConfig())
	}

	return s.redisClient
}

// Init User repository
func (s *serviceProvider) GetUserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewUserRepository(s.DBClient(ctx))
	}

	return s.userRepository
}

// Init User repository
func (s *serviceProvider) GetUserRedisRepository(ctx context.Context) repository.UserRedisRepository {
	if s.userRedisRepository == nil {
		s.userRedisRepository = userRedisRepository.NewRedisRepository(s.RedisClient(ctx))
	}

	return s.userRedisRepository
}

// Init User cache service
func (s *serviceProvider) GetUserCacheService(ctx context.Context) service.UserCacheService {
	if s.userCacheService == nil {
		s.userCacheService = userCacheService.NewUserCacheService(s.GetUserRedisRepository(ctx), s.txManager)
	}

	return s.userCacheService
}

// Init User producer service
func (s *serviceProvider) GetUserProducer(ctx context.Context) service.UserProducerService {
	if s.userProducerService == nil {
		s.userProducerService = userProduceService.NewUserProducerService(
			s.userRepository,
			s.Producer(),
		)
	}

	return s.userProducerService
}

// Init User producer service
func (s *serviceProvider) GetUserConsumer(ctx context.Context) service.UserConsumerService {
	if s.userConsumerService == nil {
		s.userConsumerService = userConsumerService.NewUserConsumerService(
			s.userRepository,
			s.Consumer(),
		)
	}

	return s.userConsumerService
}

// Init User service
func (s *serviceProvider) GetUserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewUserService(
			s.GetUserRepository(ctx),
			s.GetUserCacheService(ctx),
			s.GetUserConsumer(ctx),
			s.GetUserProducer(ctx),
		)
	}

	return s.userService
}

// Init User implementaion
func (s *serviceProvider) GetUserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewUserImplementation(s.GetUserService(ctx))
	}

	return s.userImpl
}

// Init Producer
func (s *serviceProvider) Producer() kafka.Producer {
	if s.producer == nil {
		s.producer = kafkaProducer.NewProducer(
			s.KafkaProducerConfig().TopicName(),
			s.ProducerSync(),
		)
		closer.Add(s.producer.Close)
	}

	return s.producer
}

// Init ProducerSync
func (s *serviceProvider) ProducerSync() sarama.SyncProducer {
	if s.producerSync == nil {
		producerSync, err := sarama.NewSyncProducer(
			s.KafkaProducerConfig().Brokers(),
			s.KafkaProducerConfig().Config(),
		)
		if err != nil {
			log.Fatalf("failed to create producer sync: %v", err)
		}

		s.producerSync = producerSync
	}

	return s.producerSync
}

// Init Consumer
func (s *serviceProvider) Consumer() kafka.Consumer {
	if s.consumer == nil {
		s.consumer = kafkaConsumer.NewConsumer(
			s.KafkaConsumerConfig().TopicName(),
			s.ConsumerGroup(),
			s.ConsumerGroupHandler(),
		)
		closer.Add(s.consumer.Close)
	}

	return s.consumer
}

// Init ConsumerGroup
func (s *serviceProvider) ConsumerGroup() sarama.ConsumerGroup {
	if s.consumerGroup == nil {
		consumerGroup, err := sarama.NewConsumerGroup(
			s.KafkaConsumerConfig().Brokers(),
			s.KafkaConsumerConfig().GroupID(),
			s.KafkaConsumerConfig().Config(),
		)
		if err != nil {
			log.Fatalf("failed to create consumer group: %v", err)
		}

		s.consumerGroup = consumerGroup
	}

	return s.consumerGroup
}

// Init ConsumerGroupHandler
func (s *serviceProvider) ConsumerGroupHandler() *kafkaConsumer.GroupHandler {
	if s.consumerGroupHandler == nil {
		s.consumerGroupHandler = kafkaConsumer.NewGroupHandler()
	}

	return s.consumerGroupHandler
}
