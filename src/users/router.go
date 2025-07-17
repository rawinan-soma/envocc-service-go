package users

import (
	"envocc-service-go/database"
	"envocc-service-go/src/auth"

	"github.com/labstack/echo/v4"
)

func Wire(group *echo.Group, db database.Database, mw auth.AuthMiddlewareContainer) {
	repository := NewUserRepository(db)
	service := NewUserService(repository)
	controller := NewUserController(service)

	userGroup := group.Group("/users")
	userGroup.GET("", controller.GetAllUserHandler)
	userGroup.GET("/protect", func(c echo.Context) error { return c.JSON(200, echo.Map{"msg": "protected"}) }, mw.JwtAccess)
}
