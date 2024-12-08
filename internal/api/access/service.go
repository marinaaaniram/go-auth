package access

import (
	desc "github.com/marinaaaniram/go-auth/pkg/access_v1"

	"github.com/marinaaaniram/go-auth/internal/service"
)

type Implementation struct {
	desc.UnimplementedAccessV1Server
	accessService service.AccessService
}

// Create Auth implementation
func NewAccessImplementation(accessService service.AccessService) *Implementation {
	return &Implementation{
		accessService: accessService,
	}
}
