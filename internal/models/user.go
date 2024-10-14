package models

import (
	"golang.org/x/crypto/bcrypt"
	"unicode/utf8"
)

type User struct {
	id           uint64
	username     string
	passwordHash string
}

func NewUser(username string, password string) (*User, error) {
	user := &User{}

	if passwordHash, err := hashPassword(password); err == nil {
		user.SetPasswordHash(passwordHash)
	} else {
		return nil, err
	}

	if err := user.SetUsername(username); err != nil {
		return nil, err
	}

	return user, nil
}

func BuildUser(id uint64, username string, passwordHash string) *User {
	return &User{
		id:           id,
		username:     username,
		passwordHash: passwordHash,
	}
}

func (u *User) Id() uint64 {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) PasswordHash() string {
	return u.passwordHash
}

func (u *User) SetId(id uint64) {
	u.id = id
}

func (u *User) SetUsername(username string) error {
	if utf8.RuneCountInString(username) > maxUsernameLength {
		return ErrUsernameTooLong
	}

	if utf8.RuneCountInString(username) < minUsernameLength {
		return ErrUsernameTooShort
	}

	u.username = username

	return nil
}

func (u *User) SetPasswordHash(passwordHash string) {
	u.passwordHash = passwordHash
}

func hashPassword(password string) (string, error) {
	if utf8.RuneCountInString(password) < minPasswordLength {
		return "", ErrPasswordTooShort
	}

	bytes := []byte(password)

	if len(bytes) > 72 {
		return "", ErrPasswordTooLong
	}

	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
