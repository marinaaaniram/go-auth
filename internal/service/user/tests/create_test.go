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

func TestServiceUserCreate(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository
	type userCacheServiceMockFunc func(mc *minimock.Controller) service.UserCacheService

	type args struct {
		ctx context.Context
		req *model.User
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id       = gofakeit.Int64()
		name     = gofakeit.Name()
		email    = gofakeit.Email()
		password = gofakeit.Password(true, true, true, true, true, 10)
		role     = model.AdminUserRole

		repoErr = fmt.Errorf("Repo error")

		req = &model.User{
			Name:     name,
			Email:    email,
			Password: password,
			Role:     role,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name                 string
		args                 args
		want                 int64
		err                  error
		userRepositoryMock   userRepositoryMockFunc
		userCacheServiceMock userCacheServiceMockFunc
	}{
		{
			name: "Success case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: id,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(id, nil)
				return mock
			},
			userCacheServiceMock: func(mc *minimock.Controller) service.UserCacheService {
				return nil
			},
		},
		{
			name: "Api nil pointer",
			args: args{
				ctx: ctx,
				req: nil,
			},
			want: 0,
			err:  errors.ErrPointerIsNil("user"),
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				return nil
			},
			userCacheServiceMock: func(mc *minimock.Controller) service.UserCacheService {
				return nil
			},
		},
		{
			name: "Service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: 0,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.CreateMock.Expect(ctx, req).Return(0, repoErr)
				return mock
			},
			userCacheServiceMock: func(mc *minimock.Controller) service.UserCacheService {
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
			service := user.NewUserService(userRepoMock, userCacheMock)

			newID, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
