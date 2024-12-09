package converter

import (
	desc "github.com/marinaaaniram/go-auth/pkg/auth_v1"

	"github.com/marinaaaniram/go-auth/internal/model"
)

// Convert desc LoginRequest fields to internal Auth model
func FromDescLoginToAuth(req *desc.LoginRequest) *model.AuthInput {
	if req == nil {
		return nil
	}

	return &model.AuthInput{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}
