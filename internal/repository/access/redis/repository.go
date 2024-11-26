package redis

import (
	"context"
	"encoding/json"

	"go-auth/internal/client/cache"
	"go-auth/internal/errors"
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
	jsonData, err := json.Marshal(accessibleRoles)
	if err != nil {
		return errors.ErrFailedWithAccessCache(err)
	}

	// Сохраняем JSON в Redis
	err = r.cl.Set(ctx, endpointAddress, string(jsonData))
	if err != nil {
		return errors.ErrFailedWithAccessCache(err)
	}

	return nil
}

// Create Access from redis
func (r *repo) Get(ctx context.Context, endpointAddress string) ([]string, error) {
	jsonData, err := r.cl.Get(ctx, endpointAddress)
	if err != nil {
		return nil, errors.ErrFailedWithAccessCache(err)
	}

	if jsonData == nil {
		return nil, errors.ErrFailedWithAccessCache(err)
	}

	var rawString string
	switch v := jsonData.(type) {
	case string:
		rawString = v
	case []byte:
		rawString = string(v)
	default:
		return nil, errors.ErrFailedWithAccessCache(err)
	}

	var accessibleRoles []string
	err = json.Unmarshal([]byte(rawString), &accessibleRoles)
	if err != nil {
		return nil, errors.ErrFailedWithAccessCache(err)
	}

	return accessibleRoles, nil
}
