package constant

import "time"

type UserRole string

const (
	UnknowUserRole UserRole = "UNKNOWN"
	UserUserRole   UserRole = "USER"
	AdminUserRole  UserRole = "ADMIN"

	AuthPrefix = "Bearer "

	RefreshTokenSecretKey = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g0="
	AccessTokenSecretKey  = "VqvguGiffXILza1f44TWXowDT4zwf03dtXmqWW4SYyE="

	RefreshTokenExpiration = 60 * time.Minute
	AccessTokenExpiration  = 5 * time.Minute
)
