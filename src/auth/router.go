package auth

import (
	"envocc-service-go/config"
	"envocc-service-go/database"

	"github.com/labstack/echo/v4"
)

func Wire(group *echo.Group, db database.Database, mw AuthMiddlewareContainer) {
	repository := NewAuthenRepository(db)
	service := NewAuthenService(repository, config.GetConfig())
	controller := NewAuthenController(service)

	authGroup := group.Group("/auth")
	authGroup.POST("/login", controller.LoginHandler, mw.Local)
	authGroup.POST("/logout", controller.LogoutHandler, mw.JwtAccess)
	authGroup.POST("/refresh", controller.RefreshTokenHandler, mw.JwtRefresh)
	authGroup.POST("/register", controller.CreateUserHandler)
}
