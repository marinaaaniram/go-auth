package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/model"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Create user_v1 in service layer
func (s *serv) Create(ctx context.Context, user *model.User) (*desc.User, error) {
	if user == nil {
		return nil, status.Error(codes.Internal, "user is nil")
	}

	userObj, err := s.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	if userObj == nil {
		return nil, status.Error(codes.Internal, "userObj is nil")
	}

	return converter.FromUserToDesc(userObj), nil
}
