package models

import "errors"

var (
	ErrUsernameTooLong  = errors.New("username too long")
	ErrUsernameTooShort = errors.New("username too short")

	ErrPasswordTooLong  = errors.New("password too long")
	ErrPasswordTooShort = errors.New("password too short")
)
