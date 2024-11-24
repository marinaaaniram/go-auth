package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/wrapperspb"

	"go-auth/internal/api/user"
	"go-auth/internal/constant"
	"go-auth/internal/errors"
	"go-auth/internal/model"
	"go-auth/internal/service"
	serviceMocks "go-auth/internal/service/mocks"
	desc "go-auth/pkg/user_v1"
)

func TestApiUserUpdate(t *testing.T) {
	t.Parallel()
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.UpdateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id   = gofakeit.Int64()
		name = gofakeit.Name()

		serviceErr = fmt.Errorf("Service error")

		serviceReq = &model.User{
			ID:   id,
			Name: name,
			Role: constant.AdminUserRole,
		}

		serviceReq_2 = &model.User{
			ID:   id,
			Name: name,
			Role: constant.UnknowUserRole,
		}

		serviceReq_3 = &model.User{
			ID:   id,
			Role: constant.AdminUserRole,
		}

		req = &desc.UpdateRequest{
			Id:   id,
			Name: wrapperspb.String(name),
			Role: desc.RoleEnum_ADMIN,
		}

		req_2 = &desc.UpdateRequest{
			Id:   id,
			Name: wrapperspb.String(name),
		}

		req_3 = &desc.UpdateRequest{
			Id:   id,
			Role: desc.RoleEnum_ADMIN,
		}

		res = &emptypb.Empty{}
	)

	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *emptypb.Empty
		err             error
		userServiceMock userServiceMockFunc
	}{
		{
			name: "Success case 1",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, serviceReq).Return(nil)
				return mock
			},
		},
		{
			name: "Success case 2",
			args: args{
				ctx: ctx,
				req: req_2,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, serviceReq_2).Return(nil)
				return mock
			},
		},
		{
			name: "Success case 3",
			args: args{
				ctx: ctx,
				req: req_3,
			},
			want: res,
			err:  nil,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, serviceReq_3).Return(nil)
				return mock
			},
		},
		{
			name: "Api nil pointer",
			args: args{
				ctx: ctx,
				req: nil,
			},
			want: nil,
			err:  errors.ErrPointerIsNil("req"),
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				return nil
			},
		},
		{
			name: "Service error case",
			args: args{
				ctx: ctx,
				req: req,
			},
			want: nil,
			err:  serviceErr,
			userServiceMock: func(mc *minimock.Controller) service.UserService {
				mock := serviceMocks.NewUserServiceMock(mc)
				mock.UpdateMock.Expect(ctx, serviceReq).Return(serviceErr)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			userServiceMock := tt.userServiceMock(mc)
			api := user.NewUserImplementation(userServiceMock)

			_, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
		})
	}
}
