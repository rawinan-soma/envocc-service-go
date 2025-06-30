package users

import (
	"envocc-service-go/database"

	"github.com/labstack/echo/v4"
)

func Wire(group *echo.Group, db database.Database) {
	repository := NewUserRepository(db)
	service := NewUserService(repository)
	controller := NewUserController(service)

	userGroup := group.Group("/users")
	userGroup.GET("", controller.GetAllUserHandler)
}
