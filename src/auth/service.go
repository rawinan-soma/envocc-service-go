package auth

import (
	"envocc-service-go/config"
	"envocc-service-go/entities"

	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthenService interface {
	CreateUser(dto UserCreate) error
	GetAuthenticatedUser(username string, password string) (*entities.User, error)
	GetCookieWithAccessToken(userID uint16) (*http.Cookie, error)
	GetCookieWithRefreshToken(userID uint16) (*http.Cookie, string, error)
	SetCurrentRefreshToken(refreshToken string, userID uint16) error
	GetUserFromRefreshToken(refreshToken string, userID uint16) (*entities.User, error)
	GetLogoutCookie() []*http.Cookie
	RemoveRefreshToken(userID uint16) error
	ParseRefreshToken(tokenString string) (*TokenPayload, error)
	ParseAccessToken(tokenString string) (*TokenPayload, error)
	GetUserById(id uint16) (*entities.User, error)
}

type authenService struct {
	repository AuthenRepository
	config     *config.Config
}

func NewAuthenService(repository AuthenRepository, config *config.Config) AuthenService {
	return &authenService{
		repository: repository,
		config:     config,
	}
}

var (
	ErrExistingUser      = errors.New("username or email already exists")
	ErrInvalidCredential = errors.New("invalid credential")
	ErrInvalidToken      = errors.New("invalid token")
)

func (s *authenService) CreateUser(dto UserCreate) error {
	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), 10)
	if err != nil {
		return err
	}

	user, err := s.repository.FindByUsernameOrEmail(dto.Username, dto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if user != nil {
		return ErrExistingUser
	}

	dto.Password = string(hasedPassword)
	newUser := ToUserEntity(dto)

	return s.repository.SaveUser(newUser)
}

func (s *authenService) GetAuthenticatedUser(username string, password string) (*entities.User, error) {
	user, err := s.repository.FindByUsername(username)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInvalidCredential
	}

	if err != nil {
		return nil, err
	}

	if err := s.verifyPassword(password, user.Password); err != nil {
		return nil, err
	}

	user.Password = ""

	return user, nil

}

func (s *authenService) verifyPassword(password string, hashedPassword string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return ErrInvalidCredential
	}

	return nil
}

func (s *authenService) GetCookieWithAccessToken(userID uint16) (*http.Cookie, error) {
	claims := &TokenPayload{
		UserID: int16(userID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.config.Jwt.Access_exp) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.Jwt.Access))
	if err != nil {
		return nil, err
	}

	cookie := &http.Cookie{
		Name:     "Authentication",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int(s.config.Jwt.Access_exp),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}

	return cookie, nil
}

func (s *authenService) GetCookieWithRefreshToken(userID uint16) (*http.Cookie, string, error) {
	claims := &TokenPayload{
		UserID: int16(userID),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.config.Jwt.Refresh_exp) * time.Second)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.config.Jwt.Refresh))

	if err != nil {
		return nil, "", err
	}

	cookie := &http.Cookie{
		Name:     "Refresh",
		Value:    tokenString,
		Path:     "/",
		MaxAge:   int(s.config.Jwt.Refresh_exp),
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   false,
	}

	return cookie, tokenString, nil
}

func (s *authenService) SetCurrentRefreshToken(refreshToken string, userID uint16) error {
	hashedRefreshToken, err := bcrypt.GenerateFromPassword([]byte(refreshToken), 10)
	if err != nil {
		return err
	}

	if err := s.repository.UpdateToken(string(hashedRefreshToken), userID); err != nil {
		return err
	}

	return nil
}

func (s *authenService) GetUserFromRefreshToken(refreshToken string, userID uint16) (*entities.User, error) {
	user, err := s.repository.FindByID(userID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrInvalidCredential
	}

	if user.HashedRefreshToken == nil {
		return nil, ErrInvalidToken
	}

	if err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(*user.HashedRefreshToken), []byte(refreshToken)); err != nil {
		return nil, ErrInvalidToken
	}

	user.Password = ""
	user.HashedRefreshToken = nil

	return user, nil
}

func (s *authenService) GetLogoutCookie() []*http.Cookie {
	return []*http.Cookie{
		{
			Name:     "Authentication",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		},
		{
			Name:     "Refresh",
			Value:    "",
			Path:     "/",
			MaxAge:   -1,
			HttpOnly: true,
		},
	}
}

func (s *authenService) RemoveRefreshToken(userID uint16) error {
	if err := s.repository.UpdateToken("", userID); err != nil {
		return err
	}
	return nil
}

func (s *authenService) GetUserById(userID uint16) (*entities.User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *authenService) ParseAccessToken(tokenString string) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenPayload{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.Jwt.Access), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenPayload); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}

// Helper function to parse refresh token
func (s *authenService) ParseRefreshToken(tokenString string) (*TokenPayload, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenPayload{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(s.config.Jwt.Refresh), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenPayload); ok && token.Valid {
		return claims, nil
	}

	return nil, ErrInvalidToken
}
