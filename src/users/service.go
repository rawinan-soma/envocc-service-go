package users

import "envocc-service-go/entities"

type UserService interface {
	GetAllUsers() ([]entities.User, error)
}

type userService struct {
	repository UserRepository
}

func NewUserService(repository UserRepository) UserService {
	return &userService{
		repository: repository,
	}
}

func (s *userService) GetAllUsers() ([]entities.User, error) {
	var users []entities.User

	users, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return users, err
}
