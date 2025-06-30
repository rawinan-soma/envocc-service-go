package database

import (
	"envocc-service-go/config"
	"fmt"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgressDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgressDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d TimeZone=%s", conf.Database.Host, conf.Database.User, conf.Database.Password, conf.Database.DBName, conf.Database.Port, conf.Database.TimeZone)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		dbInstance = &postgressDatabase{Db: db}
	})

	return dbInstance
}

func (p *postgressDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
