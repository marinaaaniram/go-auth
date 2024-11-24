package converter

import (
	"time"

	"go-auth/internal/constant"
	"go-auth/internal/model"
	modelRedis "go-auth/internal/repository/user/redis/model"
)

// Convert User model redis to internal model
func FromRedisToModel(user *modelRedis.User) *model.User {
	var updatedAt time.Time
	if user.UpdatedAtNs != nil {
		updatedAt = time.Unix(0, *user.UpdatedAtNs)
	}

	return &model.User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		Role:      constant.UserRole(user.Role),
		CreatedAt: time.Unix(0, user.CreatedAtNs),
		UpdatedAt: &updatedAt,
	}
}
