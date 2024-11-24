package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrPasswordsDoNotMatch  = status.Errorf(codes.InvalidArgument, "'password' and 'password_confirm' do not match")
	ErrFailedToHashPassword = status.Errorf(codes.InvalidArgument, "Failed to hash password")
	ErrEmailIsNotValid      = status.Errorf(codes.InvalidArgument, "Email format is invalid")
	ErrRoleIsNotValid       = status.Errorf(codes.InvalidArgument, "Role format is invalid")
)

func ErrCanNotBeEmpty(argumentName string) error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("%s cannot be empty", argumentName))
}

func ErrPointerIsNil(argumentName string) error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("%s is nil", argumentName))
}

func ErrFailedToBuildQuery(argumentName error) error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to build query: %v", argumentName))
}

func ErrFailedToSelectQuery(argumentName error) error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to select query: %v", argumentName))
}

func ErrFailedToInsertQuery(argumentName error) error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to insert query: %v", argumentName))
}

func ErrFailedToUpdateQuery(argumentName error) error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to update query: %v", argumentName))
}

func ErrFailedToDeleteQuery(argumentName error) error {
	return status.Errorf(codes.InvalidArgument, fmt.Sprintf("Failed to delete query: %v", argumentName))
}

func ErrObjectNotFount(objectName string, objectId int64) error {
	return status.Errorf(codes.NotFound, fmt.Sprintf("%s with id %d not found", objectName, objectId))
}
