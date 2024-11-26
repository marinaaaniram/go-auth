package redis

import (
	"context"
	"fmt"

	redigo "github.com/gomodule/redigo/redis"

	"go-auth/internal/client/cache"
	"go-auth/internal/repository"
)

type repo struct {
	cl cache.RedisClient
}

// Create Redis repository
func NewAccessRedisRepository(cl cache.RedisClient) repository.AccessRedisRepository {
	return &repo{cl: cl}
}

// Create Access in redis
func (r *repo) Create(ctx context.Context, accessibleRoles []string, endpointAddress string) error {
	fmt.Printf("33333333\n")

	err := r.cl.HashSet(ctx, endpointAddress, accessibleRoles)
	if err != nil {
		return err
	}

	return nil
}

// Create Access from redis
func (r *repo) Get(ctx context.Context, endpointAddress string) ([]string, error) {
	values, err := r.cl.HGetAll(ctx, endpointAddress)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, nil
	}

	var accessibleRoles []string
	err = redigo.ScanStruct(values, &accessibleRoles)
	if err != nil {
		return nil, err
	}

	return accessibleRoles, nil
}
