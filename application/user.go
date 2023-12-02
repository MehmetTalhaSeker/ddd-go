package application

import (
	"github.com/MehmetTalhaSeker/ddd-go/domain/entity/user"
	"github.com/MehmetTalhaSeker/ddd-go/domain/repository"
)

type userApp struct {
	us repository.UserRepository
}

type UserApp interface {
	Create(*entity.UserCreateRequest) (*entity.UserResponse, []error)
}

func NewUserApp(us repository.UserRepository) UserApp {
	return &userApp{us: us}
}

func (ua *userApp) Create(req *entity.UserCreateRequest) (*entity.UserResponse, []error) {
	u, errors := entity.NewUser(req)
	if len(errors) > 0 {
		return nil, errors
	}

	cu, err := ua.us.Create(&u)
	if err != nil {
		return nil, append(errors, err)
	}

	return cu.ToUserDTO(), nil
}
