package main

import (
	"envocc-service-go/config"
	"envocc-service-go/database"
	"envocc-service-go/server"
)

func main() {
	conf := config.GetConfig()
	db := database.NewPostgresDatabase(conf)
	server.NewServer(conf, db).Start()
}
