package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"

	"github.com/marinaaaniram/go-auth/internal/model"
	"github.com/marinaaaniram/go-auth/internal/repository"
	repoMocks "github.com/marinaaaniram/go-auth/internal/repository/mocks"
	"github.com/marinaaaniram/go-auth/internal/service/user"
)

func TestServiceUserUpdate(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

	type args struct {
		ctx context.Context
		req *model.User
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id   = gofakeit.Int64()
		name = gofakeit.Name()
		role = model.AdminUserRole

		repoErr = fmt.Errorf("Repo error")

		req = &model.User{
			ID:   id,
			Name: name,
			Role: role,
		}

		req_2 = &model.User{
			ID:   id,
			Name: name,
		}

		req_3 = &model.User{
			ID:   id,
			Role: model.AdminUserRole,
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               *model.User
		err                error
		userRepositoryMock userRepositoryMockFunc
	}{
		{
			name: "Success case 1",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req).Return(nil)
				return mock
			},
		},
		{
			name: "Success case 2",
			args: args{
				ctx: ctx,
				req: req_2,
			},
			want: nil,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req_2).Return(nil)
				return mock
			},
		},
		{
			name: "Success case 3",
			args: args{
				ctx: ctx,
				req: req_3,
			},
			want: nil,
			err:  nil,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req_3).Return(nil)
				return mock
			},
		},
		{
			name: "Service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  repoErr,
			userRepositoryMock: func(mc *minimock.Controller) repository.UserRepository {
				mock := repoMocks.NewUserRepositoryMock(mc)
				mock.UpdateMock.Expect(ctx, req).Return(repoErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewUserService(userRepoMock)

			err := service.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
