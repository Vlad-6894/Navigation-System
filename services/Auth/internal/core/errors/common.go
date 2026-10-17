package auth_errors

import "errors"

var (
	ErrInvalidLogin    = errors.New("login is invalid")
	ErrInvalidPassword = errors.New("password is invalid")
	ErrNegativeAge     = errors.New("age must not be negative")
	ErrNotFoundRole    = errors.New("this role is absent")
)
