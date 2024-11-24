package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"go-auth/internal/repository"
	repoMocks "go-auth/internal/repository/mocks"
	"go-auth/internal/service"
	"go-auth/internal/service/user"
)

func TestServiceUserDelete(t *testing.T) {
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

		id = gofakeit.Int64()

		repoErr = fmt.Errorf("Repo error")
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                    string
		args                    args
		want                    int64
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
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(nil)
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
			want: 0,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(repoErr)
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

			err := service.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
