package server

import (
	"envocc-service-go/config"
	"envocc-service-go/database"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf config.Config
}

func NewServer(conf *config.Config, db database.Database) Server {
	app := echo.New()
	app.Logger.SetLevel(log.DEBUG)
	return &echoServer{
		app:  app,
		db:   db,
		conf: *conf,
	}
}

func (e *echoServer) Start() {
	e.app.Use(middleware.Recover())
	e.app.Use(middleware.AddTrailingSlash())
	e.app.Use(middleware.Logger())

	serverUrl := fmt.Sprintf(":%d", e.conf.Server.Port)
	e.app.Logger.Fatal(e.app.Start(serverUrl))
}
