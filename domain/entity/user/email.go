package entity

import (
	"errors"
	"net/mail"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Email string

func newEmail(email string) (Email, error) {
	e, err := mail.ParseAddress(email)
	if err != nil {
		return "", ErrInvalidEmail
	}

	return Email(e.Address), nil
}
