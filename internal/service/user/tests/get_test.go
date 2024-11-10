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
	"github.com/marinaaaniram/go-auth/internal/repository/user/model"
	"github.com/marinaaaniram/go-auth/internal/service/user"
)

func TestServiceUserGet(t *testing.T) {
	t.Parallel()
	type userRepositoryMockFunc func(mc *minimock.Controller) repository.UserRepository

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
		role      = model.AdminUserRole
		createdAt = gofakeit.Date()
		// updatedAt = gofakeit.Date()

		repoErr = fmt.Errorf("Repo error")

		res = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      role,
			CreatedAt: createdAt,
			// UpdatedAt: sql.NullTime{
			// 	Time:  updatedAt,
			// 	Valid: true,
			// },
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name               string
		args               args
		want               int64
		err                error
		userRepositoryMock userRepositoryMockFunc
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
				mock.GetMock.Expect(ctx, id).Return(nil)
				return mock
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
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userRepoMock := tt.userRepositoryMock(mc)
			service := user.NewUserService(userRepoMock)

			err := service.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
