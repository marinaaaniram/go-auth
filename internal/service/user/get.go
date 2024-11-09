package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/converter"
	"github.com/marinaaaniram/go-auth/internal/model"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

// Get user_v1 in service layer
func (s *serv) Get(ctx context.Context, user *model.User) (*desc.User, error) {
	if user == nil {
		return nil, status.Error(codes.Internal, "user is nil")
	}

	userObj, err := s.userRepository.Get(ctx, user)
	if err != nil {
		return nil, err
	}

	return converter.FromUserToDesc(userObj), nil
}
