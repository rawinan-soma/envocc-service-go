package auth

import (
	"envocc-service-go/database"
	"envocc-service-go/entities"

	"gorm.io/gorm"
)

type AuthenRepository interface {
	FindByUsernameOrEmail(username string, email string) (*entities.User, error)
	SaveUser(user *entities.User) error
	FindByUsername(username string) (*entities.User, error)
	UpdateToken(token string, userID uint16) error
	FindByID(userID uint16) (*entities.User, error)
}

type authenRepository struct {
	db *gorm.DB
}

func NewAuthenRepository(db database.Database) AuthenRepository {
	return &authenRepository{
		db: db.GetDb(),
	}
}

func (r *authenRepository) FindByUsernameOrEmail(username string, email string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).Or("email = ?", email).First(&user).Error

	return &user, err
}

func (r *authenRepository) SaveUser(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *authenRepository) FindByUsername(username string) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("username = ?", username).First(&user).Error

	return &user, err

}

func (r *authenRepository) UpdateToken(token string, userID uint16) error {
	err := r.db.Model(&entities.User{}).Where("id = ?", userID).Update("hashedRefreshToken = ", token).Error

	return err
}

func (r *authenRepository) FindByID(userID uint16) (*entities.User, error) {
	var user entities.User
	err := r.db.Where("id = ?", userID).Find(&user).Error

	return &user, err
}
