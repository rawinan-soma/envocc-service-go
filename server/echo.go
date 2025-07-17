package server

import (
	"envocc-service-go/config"
	"envocc-service-go/database"
	"envocc-service-go/src/auth"
	"envocc-service-go/src/users"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf config.Config
	mw   *auth.AuthMiddlewareContainer
}

func NewServer(conf *config.Config, db database.Database) Server {
	app := echo.New()
	app.Logger.SetLevel(log.DEBUG)
	sharedUserrepository := users.NewUserRepository(db)
	sharedAuthService := auth.NewAuthenService(sharedUserrepository, conf)
	mw := auth.NewMiddlewareConatiner(sharedAuthService)
	return &echoServer{
		app:  app,
		db:   db,
		conf: *conf,
		mw:   mw,
	}
}

func (e *echoServer) Start() {
	e.app.Use(middleware.Recover())
	e.app.Use(middleware.AddTrailingSlash())
	e.app.Use(middleware.Logger())

	v2 := e.app.Group("envocc-service/api")
	users.Wire(v2, e.db, *e.mw)
	auth.Wire(v2, e.db, *e.mw)

	serverUrl := fmt.Sprintf(":%d", e.conf.Server.Port)
	e.app.Logger.Fatal(e.app.Start(serverUrl))
}
