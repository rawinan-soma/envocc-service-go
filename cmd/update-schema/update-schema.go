package main

import (
	"envocc-service-go/config"
	"envocc-service-go/entities"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conf := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s", conf.Database.Host, conf.Database.User, conf.Database.Password, conf.Database.DBName, conf.Database.Port, conf.Database.TimeZone)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	if err := db.AutoMigrate(entities.GetAllModels()...); err != nil {
		panic("failed to migrate database")
	}

	fmt.Println("Database update complete")
}
