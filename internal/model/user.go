package model

import (
	"time"

	"github.com/marinaaaniram/go-auth/internal/constant"
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
