package cache

import (
	"context"

	"github.com/marinaaaniram/go-auth/internal/model"
)

func (s *serv) Get(ctx context.Context, id int64) (*model.User, error) {
	userObj, err := s.userRedisRepository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return userObj, nil
}

func (s *serv) Create(ctx context.Context, info *model.User) (int64, error) {
	var id int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		var errTx error
		id, errTx = s.userRedisRepository.Create(ctx, info)
		if errTx != nil {
			return errTx
		}

		_, errTx = s.userRedisRepository.Get(ctx, id)
		if errTx != nil {
			return errTx
		}

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
