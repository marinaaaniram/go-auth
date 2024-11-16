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

type UserRedis struct {
	ID          int64  `redis:"id"`
	Name        string `redis:"name"`
	Email       string `redis:"email"`
	Password    string `redis:"password"`
	Role        string `redis:"role"`
	CreatedAtNs int64  `redis:"created_at"`
	UpdatedAtNs *int64 `redis:"updated_at"`
}
