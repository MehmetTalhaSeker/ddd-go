package entity

import (
	"errors"
)

var (
	ErrEmptyPassword = errors.New("password cannot be empty")
	ErrLongPassword  = errors.New("password cannot be longer than 55 characters")
	ErrShortPassword = errors.New("password cannot be shorter than 8 characters")
)

type Password string

func newPassword(password string) (Password, error) {
	if len(password) == 0 {
		return "", ErrEmptyPassword
	}

	if len(password) > 55 {
		return "", ErrLongPassword
	}

	if len(password) < 8 {
		return "", ErrShortPassword
	}

	return Password(password), nil
}
