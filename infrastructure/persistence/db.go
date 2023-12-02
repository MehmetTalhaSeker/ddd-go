package persistence

import (
	"github.com/MehmetTalhaSeker/ddd-go/domain/repository"
)

type Repositories struct {
	UserRepository repository.UserRepository
}

func NewRepositories() (*Repositories, error) {
	return &Repositories{
		UserRepository: NewInMemUserRepository(),
	}, nil
}
