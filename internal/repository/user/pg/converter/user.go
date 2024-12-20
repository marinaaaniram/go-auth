package converter

import (
	"database/sql"
	"time"

	"github.com/marinaaaniram/go-auth/internal/constant"
	"github.com/marinaaaniram/go-auth/internal/model"
	modelRepo "github.com/marinaaaniram/go-auth/internal/repository/user/pg/model"
)

// Convert sql.NullTime to *time.Time
func repoNullTimeToTime(userUpdatedAt sql.NullTime) *time.Time {
	var updatedAt *time.Time
	if userUpdatedAt.Valid {
		updatedAt = &userUpdatedAt.Time
	} else {
		updatedAt = nil
	}
	return updatedAt
}

// Convert User model repo to internal model
func FromRepoToUserGet(user *modelRepo.User) *model.User {
	if user == nil {
		return nil
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      constant.UserRole(user.Role),
		CreatedAt: user.CreatedAt,
		UpdatedAt: repoNullTimeToTime(user.UpdatedAt),
	}
}

// Convert update params of User model to internal model
func FromUserToRepoUpdate(user *model.User) *modelRepo.UserUpdate {
	userRoleStr := string(user.Role)
	return &modelRepo.UserUpdate{
		ID:   user.ID,
		Name: &user.Name,
		Role: &userRoleStr,
	}
}
