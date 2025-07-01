package auth

import (
	"envocc-service-go/src/users"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenService interface {
	Register(dto users.UserCreate) error
}

type authenService struct {
	repository users.UserRepository
}

func NewAuthenService(repository users.UserRepository) AuthenService {
	return &authenService{
		repository: repository,
	}
}

func (s authenService) Register(dto users.UserCreate) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		return err
	}

	user, err := s.repository.FindByUsernameOrEmail(dto.Username, dto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if user != nil {
		return gorm.ErrDuplicatedKey
	}

	dto.Password = string(hasedPassword)
	newUser := users.ToUserEntity(dto)

	return s.repository.SaveUser(newUser)
}
