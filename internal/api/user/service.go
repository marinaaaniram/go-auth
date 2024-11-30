package user

import (
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"

	"github.com/marinaaaniram/go-auth/internal/service"
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
