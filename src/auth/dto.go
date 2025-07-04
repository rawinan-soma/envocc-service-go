package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID       int16  `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

type LoginRquest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenPayload struct {
	UserID int16
	jwt.RegisteredClaims
}

type UserResponse struct {
	Username string
	Role     string
}

type AuthenResponse struct {
	Username string
	ID       uint
	Role     string
}
