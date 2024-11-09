package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/model"
)

// Update user_v1 in service layer
func (s *serv) Update(ctx context.Context, user *model.User) error {
	if user == nil {
		return status.Error(codes.Internal, "user is nil")
	}

	err := s.userRepository.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
