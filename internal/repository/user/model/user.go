package model

import (
	"database/sql"
	"time"

	"github.com/marinaaaniram/go-auth/internal/model"
)

// Repository User model
type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      model.UserRole
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

// Repository User Model with fields to update
type UserUpdate struct {
	ID   int64
	Name *string
	Role *model.UserRole
}
