package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/marinaaaniram/go-auth/internal/errors"
	"github.com/marinaaaniram/go-auth/internal/model"
	"github.com/marinaaaniram/go-auth/internal/repository"
	repoMocks "github.com/marinaaaniram/go-auth/internal/repository/mocks"
	"github.com/marinaaaniram/go-auth/internal/service"
	"github.com/marinaaaniram/go-auth/internal/service/user"
)

func TestServiceUserGet(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type userCacheServiceMockFunc func(mc *minimock.Controller) service.UserCacheService
	type userConsumerServiceMockFunc func(mc *minimock.Controller) service.UserConsumerService
	type userProducerServiceMockFunc func(mc *minimock.Controller) service.UserProducerService

	type args struct {
		ctx context.Context
		req int64
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		password  = gofakeit.Password(true, true, true, true, true, 10)
		role      = model.AdminUserRole
		createdAt = gofakeit.Date()

		repoErr = fmt.Errorf("Repo error")

		res = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Password:  password,
			Role:      role,
			CreatedAt: createdAt,
			UpdatedAt: nil,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                    string
		args                    args
		want                    *model.User
		err                     error
		userRepositoryMock      userRepositoryMockFunc
		userCacheServiceMock    userCacheServiceMockFunc
		userConsumerServiceMock userConsumerServiceMockFunc
		userProducerServiceMock userProducerServiceMockFunc
	}{
		{
			name: "Success case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: res,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(res, nil)
				return mock
			},
			userCacheServiceMock: func(mc *minimock.Controller) service.UserCacheService {
				return nil
			},
			userConsumerServiceMock: func(mc *minimock.Controller) service.UserConsumerService {
				return nil
			},
			userProducerServiceMock: func(mc *minimock.Controller) service.UserProducerService {
				return nil
			},
		},
		{
			name: "Repo nil pointer",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  errors.ErrPointerIsNil("userObj"),
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, nil)
				return mock
			},
			userCacheServiceMock: func(mc *minimock.Controller) service.UserCacheService {
				return nil
			},
			userConsumerServiceMock: func(mc *minimock.Controller) service.UserConsumerService {
				return nil
			},
			userProducerServiceMock: func(mc *minimock.Controller) service.UserProducerService {
				return nil
			},
		},
		{
			name: "Service error case",
			args: args{
				ctx: ctx,
				req: id,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.GetMock.Expect(ctx, id).Return(nil, repoErr)
				return mock
			},
			userCacheServiceMock: func(mc *minimock.Controller) service.UserCacheService {
				return nil
			},
			userConsumerServiceMock: func(mc *minimock.Controller) service.UserConsumerService {
				return nil
			},
			userProducerServiceMock: func(mc *minimock.Controller) service.UserProducerService {
				return nil
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			userCacheMock := tt.userCacheServiceMock(mc)
			userConsumerMock := tt.userConsumerServiceMock(mc)
			userProducerMock := tt.userProducerServiceMock(mc)
			service := user.NewUserService(userRepoMock, userCacheMock, userConsumerMock, userProducerMock)

			userObj, err := service.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, userObj)
		})
	}
}
