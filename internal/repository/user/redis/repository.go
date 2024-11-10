package pg

import (
	"context"
	"strconv"

	redigo "github.com/gomodule/redigo/redis"

	"github.com/marinaaaniram/go-auth/internal/client/cache"
	"github.com/marinaaaniram/go-auth/internal/model"
	"github.com/marinaaaniram/go-auth/internal/repository"
	"github.com/marinaaaniram/go-auth/internal/repository/user/redis/converter"
	modelRedis "github.com/marinaaaniram/go-auth/internal/repository/user/redis/model"
)

type repo struct {
	cl cache.RedisClient
}

func NewRedisRepository(cl cache.RedisClient) repository.UserRedisRepository {
	return &repo{cl: cl}
}

func (r *repo) Create(ctx context.Context, user *model.User) (int64, error) {
	id := int64(1)

	idStr := strconv.FormatInt(id, 10)
	err := r.cl.HashSet(ctx, idStr, user)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	idStr := strconv.FormatInt(id, 10)
	values, err := r.cl.HGetAll(ctx, idStr)
	if err != nil {
		return nil, err
	}

	if len(values) == 0 {
		return nil, nil
	}

	var user modelRedis.User
	err = redigo.ScanStruct(values, &user)
	if err != nil {
		return nil, err
	}

	return converter.FromRedisToModel(&user), nil
}
