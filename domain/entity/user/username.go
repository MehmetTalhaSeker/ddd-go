package entity

import (
	"errors"
	"strings"
)

var (
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrLongUsername  = errors.New("username cannot longer than 10 characters")
	ErrShortUsername = errors.New("username cannot shorter than 3 characters")
)

type Username string

func newUsername(username string) (Username, error) {
	username = strings.TrimSpace(username)
	if username == "" {
		return "", ErrEmptyUsername
	}

	if len(username) > 10 {
		return "", ErrLongUsername
	}

	if len(username) < 3 {
		return "", ErrShortUsername
	}

	return Username(username), nil
}
