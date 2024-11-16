package app

import (
	"context"
	"log"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/marinaaaniram/go-common-platform/pkg/closer"
	"github.com/marinaaaniram/go-common-platform/pkg/db"
	"github.com/marinaaaniram/go-common-platform/pkg/db/pg"
	"github.com/marinaaaniram/go-common-platform/pkg/db/transaction"

	"github.com/marinaaaniram/go-auth/internal/api/user"
	"github.com/marinaaaniram/go-auth/internal/client/cache"
	"github.com/marinaaaniram/go-auth/internal/client/cache/redis"
	"github.com/marinaaaniram/go-auth/internal/config"
	"github.com/marinaaaniram/go-auth/internal/config/env"
	"github.com/marinaaaniram/go-auth/internal/repository"
	userRepository "github.com/marinaaaniram/go-auth/internal/repository/user/pg"
	userRedisRepository "github.com/marinaaaniram/go-auth/internal/repository/user/redis"
	"github.com/marinaaaniram/go-auth/internal/service"
	userService "github.com/marinaaaniram/go-auth/internal/service/user"
	userCacheService "github.com/marinaaaniram/go-auth/internal/service/user/cache"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	redisConfig   config.RedisConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig

	dbClient  db.Client
	txManager db.TxManager

	redisPool   *redigo.Pool
	redisClient cache.RedisClient

	userRepository      repository.UserRepository
	userRedisRepository repository.UserRedisRepository

	userService      service.UserService
	userCacheService service.UserCacheService

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

// Init User service
func (s *serviceProvider) GetUserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewUserService(s.GetUserRepository(ctx), s.GetUserCacheService(ctx))
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
