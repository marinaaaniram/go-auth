package model

import (
	"time"

	"go-auth/internal/constant"
)

// Internal User model
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      constant.UserRole
	CreatedAt time.Time
	UpdatedAt *time.Time
}

type UserInfo struct {
	Username string `json:"username"`
	Role     string `json:"role"`
}
