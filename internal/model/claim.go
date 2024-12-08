package model

import "github.com/dgrijalva/jwt-go"

const (
	ExamplePath = "/user_v1.UserV1/Get"
)

type UserClaims struct {
	jwt.StandardClaims
	Email string `json:"email"`
	Role  string `json:"role"`
}
