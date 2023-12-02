package interfaces

import (
	"errors"
	"github.com/MehmetTalhaSeker/ddd-go/application"
	"github.com/MehmetTalhaSeker/ddd-go/domain/entity/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Users struct {
	ua application.UserApp
}

func NewUsersHandler(ua application.UserApp) *Users {
	return &Users{
		ua: ua,
	}
}

func (s *Users) SaveUser(c *gin.Context) {
	var req entity.UserCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}

	ur, errs := s.ua.Create(&req)
	if len(errs) > 0 {
		c.JSON(http.StatusBadRequest, errors.Join(errs...).Error())
		return
	}

	c.JSON(http.StatusCreated, ur)
}
