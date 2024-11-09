package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/marinaaaniram/go-auth/internal/model"
)

// Delete user_v1 in service layer
func (s *serv) Delete(ctx context.Context, user *model.User) error {
	if user == nil {
		return status.Error(codes.Internal, "user is nil")
	}

	err := s.userRepository.Delete(ctx, user)
	if err != nil {
		return err
	}

	return nil
}
