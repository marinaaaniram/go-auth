package model

import (
	"time"
)

// Internal User model
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt *time.Time
}
