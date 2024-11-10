package converter

import (
	"time"

	"github.com/marinaaaniram/go-auth/internal/model"
	modelRedis "github.com/marinaaaniram/go-auth/internal/repository/user/redis/model"
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
		Role:      model.UserRole(user.Role),
		CreatedAt: time.Unix(0, user.CreatedAtNs),
		UpdatedAt: &updatedAt,
	}
}
