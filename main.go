package main

import (
	"github.com/MehmetTalhaSeker/ddd-go/application"
	"github.com/MehmetTalhaSeker/ddd-go/infrastructure/persistence"
	"github.com/MehmetTalhaSeker/ddd-go/interfaces"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	repositories, err := persistence.NewRepositories()
	if err != nil {
		panic(err)
	}

	ua := application.NewUserApp(repositories.UserRepository)
	uh := interfaces.NewUsersHandler(ua)

	r := gin.Default()

	r.POST("/users", uh.SaveUser)

	log.Fatal(r.Run(":8080"))
}
