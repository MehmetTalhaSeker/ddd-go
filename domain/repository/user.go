package repository

import (
	"github.com/MehmetTalhaSeker/ddd-go/domain/entity/user"
)

type UserRepository interface {
	Create(*entity.User) (*entity.User, error)
}
