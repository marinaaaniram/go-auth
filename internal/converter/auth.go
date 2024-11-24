package converter

import (
	"go-auth/internal/model"
	desc "go-auth/pkg/auth_v1"
)

// Convert desc LoginRequest fields to internal Auth model
func FromDescLoginToAuth(req *desc.LoginRequest) *model.Auth {
	if req == nil {
		return nil
	}

	return &model.Auth{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}
}
