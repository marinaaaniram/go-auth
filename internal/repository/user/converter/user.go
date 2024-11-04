package converter

import (
	"github.com/marinaaaniram/go-auth/internal/model"
	modelRepo "github.com/marinaaaniram/go-auth/internal/repository/user/model"
)

func FromRepoToUser(user *modelRepo.User) *model.User {
	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
