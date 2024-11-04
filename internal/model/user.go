package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
