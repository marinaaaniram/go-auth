package model

// Internal Auth model
type AuthInput struct {
	Email    string
	Password string
}

// Internal Auth model
type AuthOutput struct {
	RefreshToken string
	AccessToken  string
}
