package model

type UserRole string

const (
	UnknowUserRole UserRole = "UNKNOWN"
	UserUserRole   UserRole = "USER"
	AdminUserRole  UserRole = "ADMIN"
)
