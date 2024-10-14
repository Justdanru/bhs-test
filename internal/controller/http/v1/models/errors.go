package models

import "errors"

var (
	ErrUserIdNotPassed   = errors.New("user id not passed")
	ErrWrongUserIdFormat = errors.New("wrong user id format")
)
