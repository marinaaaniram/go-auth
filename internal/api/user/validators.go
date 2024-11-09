package user

import "github.com/marinaaaniram/go-auth/internal/errors"

func validatePassword(password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.ErrPasswordsDoNotMatch
	}
	return nil
}
