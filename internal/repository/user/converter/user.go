package converter

import (
	"github.com/marinaaaniram/go-auth/internal/model"
	modelRepo "github.com/marinaaaniram/go-auth/internal/repository/user/model"
)

// Convert User model repo to internal model
func FromRepoToUser(user *modelRepo.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      model.UserRole(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
