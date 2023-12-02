package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username Username
	Email    Email
	Password Password
}

type UserResponse struct {
	ID       uuid.UUID
	Username Username
	Email    Email
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewUser(r *UserCreateRequest) (User, []error) {
	var errs = make([]error, 0)

	un, err := newUsername(r.Username)
	if err != nil {
		errs = append(errs, err)
		err = nil
	}

	p, err := newPassword(r.Password)
	if err != nil {
		errs = append(errs, err)
		err = nil
	}

	e, err := newEmail(r.Email)
	if err != nil {
		errs = append(errs, err)
		err = nil
	}

	if len(errs) > 0 {
		return User{}, errs
	}

	return User{
		ID:       uuid.New(),
		Username: un,
		Password: p,
		Email:    e,
	}, nil
}

func (u *User) ToUserDTO() *UserResponse {
	return &UserResponse{
		ID:       u.ID,
		Email:    u.Email,
		Username: u.Username,
	}
}
