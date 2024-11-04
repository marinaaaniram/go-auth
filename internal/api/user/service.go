package user

import (
	"github.com/marinaaaniram/go-auth/internal/service"
	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

// Create User implementation
func NewImplementation(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
