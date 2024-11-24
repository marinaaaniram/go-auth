package model

// Internal Auth model
type Auth struct {
	ID           int64
	Email        string
	Password     string
	AccessToken  string
	RefreshToken string
}
