package user

import (
	"regexp"

	desc "github.com/marinaaaniram/go-auth/pkg/user_v1"

	"github.com/marinaaaniram/go-auth/internal/errors"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func validateUser(req *desc.CreateRequest) error {
	if err := validateName(req.Name); err != nil {
		return err
	}

	if err := validateEmail(req.Email); err != nil {
		return err
	}

	if err := validatePassword(req.Password, req.PasswordConfirm); err != nil {
		return err
	}

	if err := validateRole(req.Role); err != nil {
		return err
	}

	return nil
}

func validatePassword(password, confirmPassword string) error {
	if password != confirmPassword {
		return errors.ErrPasswordsDoNotMatch
	}
	return nil
}

func validateName(name string) error {
	if len(name) == 0 {
		return errors.ErrCanNotBeEmpty("name")
	}
	return nil
}

func validateEmail(email string) error {
	if len(email) == 0 {
		return errors.ErrCanNotBeEmpty("name")
	}
	if !emailRegex.MatchString(email) {
		return errors.ErrEmailIsNotValid
	}
	return nil
}

func validateRole(role desc.RoleEnum) error {
	switch role {
	case desc.RoleEnum_USER, desc.RoleEnum_ADMIN:
		return nil
	case desc.RoleEnum_UNKNOWN:
		return nil
	default:
		return errors.ErrRoleIsNotValid
	}
}
