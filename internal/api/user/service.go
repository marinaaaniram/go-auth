package user

import (
	"go-auth/internal/service"
	desc "go-auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// Create User implementation
func NewUserImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
