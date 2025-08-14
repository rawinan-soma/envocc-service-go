package users_test

import (
	"envocc-service-go/entities"
	"envocc-service-go/src/users"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockUserRepository is a mock implementation of UserRepository interface
type MockUserRepository struct {
	mock.Mock
}

// FindAll mocks the FindAll method
func (m *MockUserRepository) FindAll() ([]entities.User, error) {
	args := m.Called()
	return args.Get(0).([]entities.User), args.Error(1)
}

// FindByUsernameOrEmail mocks the FindByUsernameOrEmail method
func (m *MockUserRepository) FindByUsernameOrEmail(username string, email string) (*entities.User, error) {
	args := m.Called(username, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

// FindByUsername mocks the FindByUsername method
func (m *MockUserRepository) FindByUsername(username string) (*entities.User, error) {
	args := m.Called(username)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

// FindByID mocks the FindByID method
func (m *MockUserRepository) FindByID(userID uint16) (*entities.User, error) {
	args := m.Called(userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

func (m *MockUserRepository) ResetRepository() {
	m.Calls = nil
	m.ExpectedCalls = nil
}

func TestRead(t *testing.T) {
	mockFullResponse := []entities.User{{ID: 1}, {ID: 2}}
	// var mockNullResponse []entities.User
	mockNullResponse := []entities.User{}
	mockErrorResponse := errors.New("something went wrong")
	repository := new(MockUserRepository)
	service := users.NewUserService(repository)

	t.Run("success", func(t *testing.T) {
		repository.ResetRepository()

		repository.On("FindAll").Return(mockFullResponse, nil)

		users, err := service.GetAllUsers()

		assert.NoError(t, err)
		assert.Equal(t, mockFullResponse, users)

		repository.AssertExpectations(t)
	})

	t.Run("null", func(t *testing.T) {
		repository.ResetRepository()

		repository.On("FindAll").Return(mockNullResponse, nil)

		users, err := service.GetAllUsers()

		assert.NoError(t, err)
		assert.Equal(t, mockNullResponse, users)

		repository.AssertExpectations(t)
	})

	t.Run("error", func(t *testing.T) {
		repository.ResetRepository()

		repository.On("FindAll").Return(mockNullResponse, mockErrorResponse)

		_, err := service.GetAllUsers()

		assert.Error(t, err)
		// assert.NotEqual(t, mockFullResponse, users)

		repository.AssertExpectations(t)
	})
}
