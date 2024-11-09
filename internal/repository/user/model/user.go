package model

import (
	"database/sql"
	"time"

	"github.com/marinaaaniram/go-auth/internal/model"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      model.UserRole
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserUpdate struct {
	ID   int64
	Name *string
	Role *model.UserRole
}
