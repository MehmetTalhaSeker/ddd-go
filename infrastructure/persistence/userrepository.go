package persistence

import (
	"fmt"
	entity "github.com/MehmetTalhaSeker/ddd-go/domain/entity/user"
	"github.com/MehmetTalhaSeker/ddd-go/domain/repository"
	"github.com/google/uuid"
	"sync"
)

type InMemUserRepository struct {
	users map[uuid.UUID]entity.User
	sync.Mutex
}

func NewInMemUserRepository() repository.UserRepository {
	return &InMemUserRepository{
		users: make(map[uuid.UUID]entity.User),
	}
}

func (imr *InMemUserRepository) Create(u *entity.User) (*entity.User, error) {
	if imr.users == nil {
		imr.Lock()
		imr.users = make(map[uuid.UUID]entity.User)
		imr.Unlock()
	}

	if _, ok := imr.users[u.ID]; ok {
		return nil, fmt.Errorf("user already exists")
	}

	imr.Lock()
	imr.users[u.ID] = *u
	imr.Unlock()
	return u, nil
}
