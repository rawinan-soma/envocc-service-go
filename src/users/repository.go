package users

import (
	"envocc-service-go/database"
	"envocc-service-go/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db database.Database) UserRepository {
	return &userRepository{
		db: db.GetDb(),
	}
}

func (r userRepository) FindAll() ([]entities.User, error) {
	var users []entities.User
	return users, r.db.Find(&users).Error
}
