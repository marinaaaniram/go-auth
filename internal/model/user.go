package model

import (
	"database/sql"
	"time"
)

// type UserRole string

// const (
// 	UnknowUserRole UserRole = "UNKNOWN"
// 	UserUserRole   UserRole = "USER"
// 	AdminUserRole  UserRole = "ADMIN"
// )

type User struct {
	ID        int64
	Name      string
	Email     string
	Password  string
	Role      string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}
