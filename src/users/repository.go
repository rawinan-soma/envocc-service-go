package users

import (
	"envocc-service-go/database"
	"envocc-service-go/entities"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll() ([]entities.User, error)
	FindByUsernameOrEmail(username string, email string) (*entities.User, error)
	SaveUser(user *entities.User) error
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
	return users, r.db.Preload("Group").Preload("Position").Preload("PositionLevel").Find(&users).Error
}

func (r userRepository) FindByUsernameOrEmail(username string, email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).Or("email = ?", email).First(&user).Error

	return &user, err
}

func (r userRepository) SaveUser(user *entities.User) error {
	return r.db.Create(user).Error
}
